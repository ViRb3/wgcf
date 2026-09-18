package cloudflare

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/ViRb3/wgcf/v2/openapi"
)

func TestLatestAPIIdentityAndRegistrationPayload(t *testing.T) {
	if ApiVersion != "v0a5641" {
		t.Fatalf("API version = %q", ApiVersion)
	}
	if got := DefaultHeaders["CF-Client-Version"]; got != "a-6.38.9-5641" {
		t.Fatalf("CF-Client-Version = %q", got)
	}
	if got := DefaultHeaders["User-Agent"]; got != "1.1.1.1/6.38.9-5641 (Android 16.0.0)" {
		t.Fatalf("User-Agent = %q", got)
	}

	payload, err := json.Marshal(registrationRequest("public-key", "PC", "tos-time"))
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]any
	if err := json.Unmarshal(payload, &fields); err != nil {
		t.Fatal(err)
	}
	want := map[string]any{
		"key":           "public-key",
		"install_id":    "",
		"fcm_token":     "",
		"tos":           "tos-time",
		"model":         "PC",
		"serial_number": "",
		"os_version":    "16.0.0",
		"key_type":      "curve25519",
		"tunnel_type":   "wireguard",
		"locale":        "en_US",
	}
	if !reflect.DeepEqual(fields, want) {
		t.Fatalf("registration payload = %#v, want %#v", fields, want)
	}
}

func TestWireGuardEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		endpoint openapi.Endpoint
		want     string
	}{
		{"IPv4", openapi.Endpoint{V4: "192.0.2.10", Ports: []int32{2408, 500}}, "192.0.2.10:2408"},
		{"IPv6", openapi.Endpoint{V6: "2001:db8::1", Ports: []int32{500}}, "[2001:db8::1]:500"},
		{"IPv4 placeholder port", openapi.Endpoint{V4: "198.51.100.4:0", Ports: []int32{2408, 500}}, "198.51.100.4:2408"},
		{"IPv6 placeholder port", openapi.Endpoint{V6: "[2001:db8::4]:0", Ports: []int32{2408, 500}}, "[2001:db8::4]:2408"},
		{"default port", openapi.Endpoint{V4: "192.0.2.10"}, "192.0.2.10:2408"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := WireGuardEndpoint(test.endpoint)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Fatalf("endpoint = %q, want %q", got, test.want)
			}
		})
	}
}

func TestGetClientConfigAcceptsTextPlainJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if values := request.Header.Values("User-Agent"); len(values) != 1 || values[0] != DefaultHeaders["User-Agent"] {
			t.Errorf("User-Agent values = %q", values)
		}
		if request.Header.Get("CF-Client-Version") != DefaultHeaders["CF-Client-Version"] {
			t.Errorf("CF-Client-Version = %q", request.Header.Get("CF-Client-Version"))
		}
		if request.Header.Get("Accept") != "" {
			t.Errorf("Accept = %q", request.Header.Get("Accept"))
		}
		if request.Header.Get("Connection") != "Keep-Alive" {
			t.Errorf("Connection = %q", request.Header.Get("Connection"))
		}
		writer.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		_, _ = io.WriteString(writer, `{"errors":[],"messages":[],"result":{"captive_portal":[{"name":"example_captive","networks":[{"address":"192.0.2.0/24"}]}],"denylist":[],"premium_data_bytes":0,"referral_reward_bytes":1000000000},"success":true}`)
	}))
	defer server.Close()

	client := MakeApiClient(nil)
	client.GetConfig().Servers = openapi.ServerConfigurations{{URL: server.URL}}
	result, response, err := client.DefaultAPI.GetClientConfig(context.Background(), ApiVersion).Execute()
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if !result.Success || len(result.Result.CaptivePortal) != 1 {
		t.Fatalf("client config = %#v", result)
	}
}

func TestMakeApiClientDoesNotShareAuthorizationHeaders(t *testing.T) {
	token := "token"
	authenticated := MakeApiClient(&token)
	if got := authenticated.GetConfig().DefaultHeader["Authorization"]; got != "Bearer token" {
		t.Fatalf("authenticated Authorization = %q", got)
	}
	if _, exists := DefaultHeaders["Authorization"]; exists {
		t.Fatal("authenticated client mutated DefaultHeaders")
	}
	if _, exists := MakeApiClient(nil).GetConfig().DefaultHeader["Authorization"]; exists {
		t.Fatal("unauthenticated client inherited Authorization")
	}
}

func TestUpdateAccountDecodes200Response(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPut {
			t.Errorf("method = %q", request.Method)
		}
		if request.Header.Get("Content-Type") != "application/json; charset=UTF-8" {
			t.Errorf("Content-Type = %q", request.Header.Get("Content-Type"))
		}
		if request.Header.Get("Accept") != "" {
			t.Errorf("Accept = %q", request.Header.Get("Accept"))
		}
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = io.WriteString(writer, `{"created":"0001-01-01T00:00:00Z","id":"account-id","premium_data":0,"quota":0,"referral_count":0,"referral_renewal_countdown":0,"role":"child","updated":"2026-09-18T16:15:36.549493834Z","warp_plus":true}`)
	}))
	defer server.Close()

	client := MakeApiClient(nil)
	client.GetConfig().Servers = openapi.ServerConfigurations{{URL: server.URL}}
	result, response, err := client.DefaultAPI.
		UpdateAccount(context.Background(), "device-id", ApiVersion).
		UpdateAccountRequest(openapi.UpdateAccountRequest{License: "license-key"}).
		Execute()
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if result.Id != "account-id" || result.Role != "child" || !result.WarpPlus {
		t.Fatalf("updated account = %#v", result)
	}
}

func TestWarpClientHelloMatchesAndroidApp(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer serverConn.Close()

	handshakeDone := make(chan error, 1)
	go func() {
		_, err := handshakeWarpTLS(context.Background(), clientConn, "api.cloudflareclient.com", &tls.Config{
			InsecureSkipVerify: true,
		})
		handshakeDone <- err
	}()

	if err := serverConn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	recordHeader := make([]byte, 5)
	if _, err := io.ReadFull(serverConn, recordHeader); err != nil {
		t.Fatalf("read TLS record header: %v", err)
	}
	if got, want := recordHeader[:3], []byte{22, 3, 1}; !reflect.DeepEqual(got, want) {
		t.Fatalf("TLS record header = %v, want %v", got, want)
	}
	record := make([]byte, int(binary.BigEndian.Uint16(recordHeader[3:])))
	if _, err := io.ReadFull(serverConn, record); err != nil {
		t.Fatalf("read TLS record: %v", err)
	}
	if got, want := len(record), 161; got != want {
		t.Errorf("ClientHello TLS record length = %d, want %d", got, want)
	}
	actualRecord := append(append([]byte(nil), recordHeader...), record...)
	zeroClientHelloRandom(actualRecord)
	expectedRecord, err := hex.DecodeString(warpAndroidInitialClientHelloGolden)
	if err != nil {
		t.Fatalf("decode Android ClientHello golden record: %v", err)
	}
	if !bytes.Equal(actualRecord, expectedRecord) {
		t.Errorf("generated ClientHello differs from the Android golden record")
	}
	serverConn.Close()
	<-handshakeDone

	hello := parseClientHello(t, record)
	if got, want := hello.version, uint16(tls.VersionTLS12); got != want {
		t.Errorf("ClientHello version = %#x, want %#x", got, want)
	}
	if len(hello.sessionID) != 0 {
		t.Errorf("initial session ID length = %d, want 0", len(hello.sessionID))
	}
	if got, want := hello.cipherSuites, []uint16{0xc02c, 0xc030}; !reflect.DeepEqual(got, want) {
		t.Errorf("cipher suites = %#v, want %#v", got, want)
	}
	if got, want := hello.compressionMethods, []byte{0}; !reflect.DeepEqual(got, want) {
		t.Errorf("compression methods = %v, want %v", got, want)
	}
	if got, want := hello.extensionIDs, []uint16{0, 23, 65281, 10, 11, 35, 16, 5, 13}; !reflect.DeepEqual(got, want) {
		t.Errorf("extension order = %v, want %v", got, want)
	}
	if got, want := uint16List(t, hello.extensions[10][2:]), []uint16{29, 23, 24}; !reflect.DeepEqual(got, want) {
		t.Errorf("supported groups = %v, want %v", got, want)
	}
	if got, want := hello.extensions[11], []byte{1, 0}; !reflect.DeepEqual(got, want) {
		t.Errorf("point formats = %v, want %v", got, want)
	}
	if got, want := hello.extensions[16], []byte{0, 9, 8, 'h', 't', 't', 'p', '/', '1', '.', '1'}; !reflect.DeepEqual(got, want) {
		t.Errorf("ALPN = %v, want %v", got, want)
	}
	if got, want := uint16List(t, hello.extensions[13][2:]), []uint16{
		0x0403, 0x0804, 0x0401, 0x0503, 0x0805, 0x0501, 0x0806, 0x0601, 0x0201,
	}; !reflect.DeepEqual(got, want) {
		t.Errorf("signature algorithms = %#v, want %#v", got, want)
	}
}

// Golden TLS record for Android WARP 6.38.9 (build 5641). The 32-byte
// ClientHello random field is zeroed so the fixture contains no session data.
const warpAndroidInitialClientHelloGolden = "16030100a10100009d03030000000000000000000000000000000000000000000000000000000000000000000004c02cc030010000700000001d001b0000186170692e636c6f7564666c617265636c69656e742e636f6d00170000ff01000100000a00080006001d00170018000b00020100002300000010000b000908687474702f312e31000500050100000000000d00140012040308040401050308050501080606010201"

func zeroClientHelloRandom(record []byte) {
	const (
		randomOffset = 5 + 4 + 2
		randomLength = 32
	)
	for i := randomOffset; i < randomOffset+randomLength; i++ {
		record[i] = 0
	}
}

func TestWarpClientHelloResumptionShape(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	server.TLS = &tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS12,
	}
	server.StartTLS()
	defer server.Close()

	dial := func() parsedClientHello {
		rawConn, err := net.Dial("tcp", server.Listener.Addr().String())
		if err != nil {
			t.Fatalf("dial test server: %v", err)
		}
		recorded := &recordingConn{Conn: rawConn}
		tlsConn, err := handshakeWarpTLS(context.Background(), recorded, "warp-session.test", &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			t.Fatalf("TLS handshake: %v", err)
		}
		tlsConn.Close()

		data := recorded.writes.Bytes()
		if len(data) < 5 {
			t.Fatal("client did not write a TLS record")
		}
		recordLength := int(binary.BigEndian.Uint16(data[3:5]))
		if len(data) < 5+recordLength {
			t.Fatalf("TLS record length = %d, available %d bytes", recordLength, len(data)-5)
		}
		return parseClientHello(t, data[5:5+recordLength])
	}

	initial := dial()
	if len(initial.sessionID) != 0 || len(initial.extensions[35]) != 0 {
		t.Fatalf("initial handshake unexpectedly contains session state")
	}

	resumed := dial()
	if got, want := len(resumed.sessionID), 32; got != want {
		t.Errorf("resumed session ID length = %d, want %d", got, want)
	}
	if len(resumed.extensions[35]) == 0 {
		t.Error("resumed handshake has an empty session ticket")
	}
	wantPrefix := []uint16{0, 23, 65281, 10, 11, 35, 16, 5, 13}
	if got := resumed.extensionIDs[:len(wantPrefix)]; !reflect.DeepEqual(got, wantPrefix) {
		t.Errorf("resumed extension order = %v, want prefix %v", resumed.extensionIDs, wantPrefix)
	}
	if resumed.extensionIDs[len(resumed.extensionIDs)-1] != 21 {
		t.Errorf("last resumed extension = %d, want padding (21)", resumed.extensionIDs[len(resumed.extensionIDs)-1])
	}
}

type recordingConn struct {
	net.Conn
	writes bytes.Buffer
}

func (conn *recordingConn) Write(data []byte) (int, error) {
	conn.writes.Write(data)
	return conn.Conn.Write(data)
}

type parsedClientHello struct {
	version            uint16
	sessionID          []byte
	cipherSuites       []uint16
	compressionMethods []byte
	extensionIDs       []uint16
	extensions         map[uint16][]byte
}

func parseClientHello(t *testing.T, record []byte) parsedClientHello {
	t.Helper()
	if len(record) < 4 || record[0] != 1 {
		t.Fatalf("record does not contain a ClientHello")
	}
	bodyLength := int(record[1])<<16 | int(record[2])<<8 | int(record[3])
	if bodyLength != len(record)-4 {
		t.Fatalf("ClientHello length = %d, record contains %d", bodyLength, len(record)-4)
	}
	body := record[4:]
	if len(body) < 35 {
		t.Fatal("truncated ClientHello")
	}
	hello := parsedClientHello{
		version:    binary.BigEndian.Uint16(body),
		extensions: make(map[uint16][]byte),
	}
	offset := 34
	sessionIDLength := int(body[offset])
	offset++
	hello.sessionID = append([]byte(nil), body[offset:offset+sessionIDLength]...)
	offset += sessionIDLength
	cipherLength := int(binary.BigEndian.Uint16(body[offset:]))
	offset += 2
	hello.cipherSuites = uint16List(t, body[offset:offset+cipherLength])
	offset += cipherLength
	compressionLength := int(body[offset])
	offset++
	hello.compressionMethods = append([]byte(nil), body[offset:offset+compressionLength]...)
	offset += compressionLength
	extensionsLength := int(binary.BigEndian.Uint16(body[offset:]))
	offset += 2
	end := offset + extensionsLength
	if end != len(body) {
		t.Fatalf("extensions end at %d, ClientHello ends at %d", end, len(body))
	}
	for offset < end {
		id := binary.BigEndian.Uint16(body[offset:])
		length := int(binary.BigEndian.Uint16(body[offset+2:]))
		offset += 4
		hello.extensionIDs = append(hello.extensionIDs, id)
		hello.extensions[id] = append([]byte(nil), body[offset:offset+length]...)
		offset += length
	}
	return hello
}

func uint16List(t *testing.T, data []byte) []uint16 {
	t.Helper()
	if len(data)%2 != 0 {
		t.Fatalf("odd uint16 list length: %d", len(data))
	}
	result := make([]uint16, 0, len(data)/2)
	for len(data) != 0 {
		result = append(result, binary.BigEndian.Uint16(data))
		data = data[2:]
	}
	return result
}
