package cloudflare

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestDefaultTransportTLSSignatureSchemes(t *testing.T) {
	signatureSchemes := make(chan []tls.SignatureScheme, 1)
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()

	serverDone := make(chan error, 1)
	go func() {
		defer serverConn.Close()
		server := tls.Server(serverConn, &tls.Config{
			GetConfigForClient: func(hello *tls.ClientHelloInfo) (*tls.Config, error) {
				signatureSchemes <- append([]tls.SignatureScheme(nil), hello.SignatureSchemes...)
				return nil, errors.New("stop after capturing ClientHello")
			},
		})
		serverDone <- server.Handshake()
	}()

	transport := DefaultTransport.Clone()
	transport.Proxy = nil
	transport.DialContext = func(context.Context, string, string) (net.Conn, error) {
		return clientConn, nil
	}
	client := &http.Client{Transport: transport}
	defer client.CloseIdleConnections()

	requestContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(requestContext, http.MethodGet, "https://example.test", nil)
	if err != nil {
		t.Fatalf("create request: %v", err)
	}
	if _, err := client.Do(request); err == nil {
		t.Fatal("request unexpectedly succeeded")
	}

	select {
	case err := <-serverDone:
		if err == nil {
			t.Fatal("server handshake unexpectedly succeeded")
		}
	case <-requestContext.Done():
		t.Fatal("server did not receive ClientHello")
	}

	got := <-signatureSchemes
	want := []tls.SignatureScheme{
		tls.PSSWithSHA256,
		tls.ECDSAWithP256AndSHA256,
		tls.Ed25519,
		tls.PSSWithSHA384,
		tls.PSSWithSHA512,
		tls.PKCS1WithSHA256,
		tls.PKCS1WithSHA384,
		tls.PKCS1WithSHA512,
		tls.ECDSAWithP384AndSHA384,
		tls.ECDSAWithP521AndSHA512,
		tls.PKCS1WithSHA1,
		tls.ECDSAWithSHA1,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(
			"TLS signature schemes changed:\n got: %s\nwant: %s",
			signatureSchemeNames(got),
			signatureSchemeNames(want),
		)
	}
}

func signatureSchemeNames(schemes []tls.SignatureScheme) string {
	names := make([]string, len(schemes))
	for i, scheme := range schemes {
		names[i] = scheme.String()
	}
	return fmt.Sprint(names)
}
