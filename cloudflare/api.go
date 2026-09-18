package cloudflare

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"io"
	"mime"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ViRb3/wgcf/v2/config"
	"github.com/ViRb3/wgcf/v2/openapi"
	"github.com/ViRb3/wgcf/v2/util"
	"github.com/ViRb3/wgcf/v2/wireguard"
	"github.com/cockroachdb/errors"
	utls "github.com/refraction-networking/utls"
)

const (
	ApiUrl                  = "https://api.cloudflareclient.com"
	ApiVersion              = "v0a5641"
	warpTLSHandshakeTimeout = 10 * time.Second
)

var (
	DefaultHeaders = map[string]string{
		"User-Agent":        "1.1.1.1/6.38.9-5641 (Android 16.0.0)",
		"CF-Client-Version": "a-6.38.9-5641",
	}
	warpSessionCache = utls.NewLRUClientSessionCache(64)
	warpDialer       = &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	warpTLSConfig = &tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS12,
	}
	DefaultTransport = &http.Transport{
		// TLSClientConfig supplies verification settings to dialWarpTLS. The
		// ClientHello itself is assembled by uTLS to match the Android app.
		TLSClientConfig:   warpTLSConfig,
		DialTLSContext:    dialWarpTLS,
		ForceAttemptHTTP2: false,
		// A CONNECT proxy makes net/http bypass DialTLSContext for the target
		// handshake. Stay direct so every API connection uses the WARP hello.
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   warpTLSHandshakeTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}
)

type cloudflareResponseTransport struct {
	base http.RoundTripper
}

func (t cloudflareResponseTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	// Match the OkHttp request headers emitted by the Android client.
	request.Header.Del("Accept")
	request.Header.Set("Connection", "Keep-Alive")
	requestContentType, _, requestContentTypeErr := mime.ParseMediaType(request.Header.Get("Content-Type"))
	if requestContentTypeErr == nil && requestContentType == "application/json" {
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	response, err := t.base.RoundTrip(request)
	if err != nil || response == nil || !strings.HasSuffix(request.URL.Path, "/client_config") {
		return response, err
	}

	contentType, _, parseErr := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if parseErr == nil && contentType == "text/plain" {
		// The Android endpoint returns JSON while labelling it as text/plain.
		// The generated OpenAPI decoder dispatches exclusively on this header.
		response.Header.Set("Content-Type", "application/json")
	}
	return response, nil
}

func dialWarpTLS(ctx context.Context, network, address string) (net.Conn, error) {
	rawConn, err := warpDialer.DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}

	serverName, _, err := net.SplitHostPort(address)
	if err != nil {
		rawConn.Close()
		return nil, errors.Wrap(err, "split TLS address")
	}

	handshakeCtx, cancel := context.WithTimeout(ctx, warpTLSHandshakeTimeout)
	defer cancel()
	tlsConn, err := handshakeWarpTLS(handshakeCtx, rawConn, serverName, warpTLSConfig)
	if err != nil {
		rawConn.Close()
		return nil, err
	}
	return tlsConn, nil
}

func handshakeWarpTLS(ctx context.Context, rawConn net.Conn, serverName string, config *tls.Config) (*utls.UConn, error) {
	if config.ServerName != "" {
		serverName = config.ServerName
	}
	uConfig := &utls.Config{
		ServerName:            serverName,
		RootCAs:               config.RootCAs,
		InsecureSkipVerify:    config.InsecureSkipVerify,
		VerifyPeerCertificate: config.VerifyPeerCertificate,
		ClientSessionCache:    warpSessionCache,
		MinVersion:            utls.VersionTLS12,
		MaxVersion:            utls.VersionTLS12,
		NextProtos:            []string{"http/1.1"},
		KeyLogWriter:          config.KeyLogWriter,
	}
	uConn := utls.UClient(rawConn, uConfig, utls.HelloCustom)
	if err := uConn.ApplyPreset(warpClientHelloSpec()); err != nil {
		return nil, errors.Wrap(err, "apply WARP TLS fingerprint")
	}

	// Conscrypt leaves the session ID empty on a fresh TLS 1.2 connection.
	// uTLS normally generates one for every custom ClientHello, so clear it
	// before session lookup and restore it only when a cached ticket is used.
	uConn.HandshakeState.Hello.SessionId = nil
	if err := uConn.BuildHandshakeState(); err != nil {
		return nil, errors.Wrap(err, "build WARP TLS ClientHello")
	}
	if warpSessionTicketPresent(uConn) {
		sessionID := make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, sessionID); err != nil {
			return nil, errors.Wrap(err, "generate TLS session ID")
		}
		uConn.HandshakeState.Hello.SessionId = sessionID
		if err := uConn.MarshalClientHello(); err != nil {
			return nil, errors.Wrap(err, "marshal resumed WARP TLS ClientHello")
		}
	}

	if err := uConn.HandshakeContext(ctx); err != nil {
		return nil, errors.Wrap(err, "WARP TLS handshake")
	}
	return uConn, nil
}

func warpSessionTicketPresent(conn *utls.UConn) bool {
	for _, extension := range conn.Extensions {
		if ticket, ok := extension.(*utls.SessionTicketExtension); ok {
			return len(ticket.Ticket) != 0
		}
	}
	return false
}

func warpClientHelloSpec() *utls.ClientHelloSpec {
	return &utls.ClientHelloSpec{
		TLSVersMin: utls.VersionTLS12,
		TLSVersMax: utls.VersionTLS12,
		CipherSuites: []uint16{
			utls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			utls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
		CompressionMethods: []uint8{0},
		Extensions: []utls.TLSExtension{
			&utls.SNIExtension{},
			&utls.ExtendedMasterSecretExtension{},
			&utls.RenegotiationInfoExtension{Renegotiation: utls.RenegotiateNever},
			&utls.SupportedCurvesExtension{Curves: []utls.CurveID{
				utls.X25519,
				utls.CurveP256,
				utls.CurveP384,
			}},
			&utls.SupportedPointsExtension{SupportedPoints: []uint8{0}},
			&utls.SessionTicketExtension{},
			&utls.ALPNExtension{AlpnProtocols: []string{"http/1.1"}},
			&utls.StatusRequestExtension{},
			&utls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []utls.SignatureScheme{
				utls.ECDSAWithP256AndSHA256,
				utls.PSSWithSHA256,
				utls.PKCS1WithSHA256,
				utls.ECDSAWithP384AndSHA384,
				utls.PSSWithSHA384,
				utls.PKCS1WithSHA384,
				utls.PSSWithSHA512,
				utls.PKCS1WithSHA512,
				utls.PKCS1WithSHA1,
			}},
			&utls.UtlsPaddingExtension{GetPaddingLen: utls.BoringPaddingStyle},
		},
	}
}

var apiClient = MakeApiClient(nil)
var apiClientAuth *openapi.APIClient

func MakeApiClient(authToken *string) *openapi.APIClient {
	httpClient := http.Client{Transport: cloudflareResponseTransport{base: DefaultTransport}}
	defaultHeaders := make(map[string]string, len(DefaultHeaders))
	for name, value := range DefaultHeaders {
		// UserAgent is emitted separately by the generated client. Copying the
		// remaining headers also prevents per-client auth from mutating globals.
		if !strings.EqualFold(name, "User-Agent") {
			defaultHeaders[name] = value
		}
	}
	apiClient := openapi.NewAPIClient(&openapi.Configuration{
		DefaultHeader: defaultHeaders,
		UserAgent:     DefaultHeaders["User-Agent"],
		Debug:         false,
		Servers: []openapi.ServerConfiguration{
			{URL: ApiUrl},
		},
		HTTPClient: &httpClient,
	})
	if authToken != nil {
		apiClient.GetConfig().DefaultHeader["Authorization"] = "Bearer " + *authToken
	}
	return apiClient
}

func Register(publicKey *wireguard.Key, deviceModel string) (*openapi.Register200Response, error) {
	request := registrationRequest(publicKey.String(), deviceModel, util.GetTimestamp())
	result, _, err := apiClient.DefaultAPI.
		Register(nil, ApiVersion).
		RegisterRequest(request).
		Execute()
	return result, errors.WithStack(err)
}

func registrationRequest(publicKey, deviceModel, timestamp string) openapi.RegisterRequest {
	// The app normalizes Android's release to a three-component version.
	osVersion := "16.0.0"
	keyType := "curve25519"
	tunnelType := "wireguard"
	return openapi.RegisterRequest{
		FcmToken:     "", // populated by Firebase in the Android app
		InstallId:    "", // populated by Firebase in the Android app
		Key:          publicKey,
		Locale:       "en_US",
		Model:        &deviceModel,
		Tos:          &timestamp,
		SerialNumber: "",
		OsVersion:    &osVersion,
		KeyType:      &keyType,
		TunnelType:   &tunnelType,
	}
}

func WireGuardEndpoint(endpoint openapi.Endpoint) (string, error) {
	if endpoint.Host != nil && *endpoint.Host != "" {
		return *endpoint.Host, nil
	}
	host := endpoint.V4
	if host == "" {
		host = endpoint.V6
	}
	if host == "" {
		return "", errors.New("Cloudflare response did not contain a WireGuard endpoint")
	}
	if parsedHost, _, err := net.SplitHostPort(host); err == nil {
		host = parsedHost
	}
	port := int32(2408)
	if len(endpoint.Ports) != 0 {
		port = endpoint.Ports[0]
	}
	return net.JoinHostPort(host, strconv.Itoa(int(port))), nil
}

type SourceDevice openapi.GetSourceDevice200Response

func GetSourceDevice(ctx *config.Context) (*SourceDevice, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultAPI.
		GetSourceDevice(nil, ApiVersion, ctx.DeviceId).
		Execute()
	return (*SourceDevice)(result), errors.WithStack(err)
}

func globalClientAuth(authToken string) *openapi.APIClient {
	if apiClientAuth == nil {
		apiClientAuth = MakeApiClient(&authToken)
	}
	return apiClientAuth
}

type Account openapi.Account

func GetAccount(ctx *config.Context) (*Account, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultAPI.
		GetAccount(nil, ctx.DeviceId, ApiVersion).
		Execute()
	castResult := (*Account)(result)
	return castResult, errors.WithStack(err)
}

func UpdateLicenseKey(ctx *config.Context) error {
	_, _, err := globalClientAuth(ctx.AccessToken).DefaultAPI.
		UpdateAccount(nil, ctx.DeviceId, ApiVersion).
		UpdateAccountRequest(openapi.UpdateAccountRequest{License: ctx.LicenseKey}).
		Execute()
	return errors.WithStack(err)
}

type BoundDevice openapi.BoundDevice

func GetBoundDevices(ctx *config.Context) ([]BoundDevice, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultAPI.
		GetBoundDevices(nil, ctx.DeviceId, ApiVersion).
		Execute()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var castResult []BoundDevice
	for _, device := range result {
		castResult = append(castResult, BoundDevice(device))
	}
	return castResult, nil
}

func GetSourceBoundDevice(ctx *config.Context) (*BoundDevice, error) {
	result, err := GetBoundDevices(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return FindDevice(result, ctx.DeviceId)
}

func UpdateSourceBoundDeviceName(ctx *config.Context, targetDeviceId string, newName string) (*BoundDevice, error) {
	return updateSourceBoundDevice(ctx, targetDeviceId, openapi.UpdateBoundDeviceRequest{
		Name: &newName,
	})
}

func UpdateSourceBoundDeviceActive(ctx *config.Context, targetDeviceId string, active bool) (*BoundDevice, error) {
	return updateSourceBoundDevice(ctx, targetDeviceId, openapi.UpdateBoundDeviceRequest{
		Active: &active,
	})
}

func updateSourceBoundDevice(ctx *config.Context, targetDeviceId string, data openapi.UpdateBoundDeviceRequest) (*BoundDevice, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultAPI.
		UpdateBoundDevice(nil, ctx.DeviceId, ApiVersion, targetDeviceId).
		UpdateBoundDeviceRequest(data).
		Execute()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var castResult []BoundDevice
	for _, device := range result {
		castResult = append(castResult, BoundDevice(device))
	}
	return FindDevice(castResult, ctx.DeviceId)
}

func DeleteBoundDevice(ctx *config.Context, targetDeviceId string) error {
	if _, _, err := globalClientAuth(ctx.AccessToken).DefaultAPI.
		DeleteBoundDevice(nil, ctx.DeviceId, ApiVersion, targetDeviceId).
		Execute(); err != nil {
		return err
	}
	return nil
}
