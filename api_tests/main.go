package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	opticgo "github.com/ViRb3/optic-go"
	"github.com/ViRb3/sling/v2"
	"github.com/ViRb3/wgcf/v2/cloudflare"
	"github.com/ViRb3/wgcf/v2/util"
	"github.com/ViRb3/wgcf/v2/wireguard"
)

var defaultHeaders = map[string]string{}
var defaultTransport = cloudflare.DefaultTransport

func init() {
	for key, val := range cloudflare.DefaultHeaders {
		defaultHeaders[key] = val
	}
}

type CustomTransport struct{}

func (CustomTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for key, val := range defaultHeaders {
		r.Header.Set(key, val)
	}
	response, err := defaultTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("bad code: %d", response.StatusCode))
	}
	return response, nil
}

func main() {
	if err := testAPI(); err != nil {
		log.Fatalln(err)
	}
}

func generateKeyPair() (*wireguard.Key, *wireguard.Key) {
	privateKey, err := wireguard.NewPrivateKey()
	if err != nil {
		panic(err)
	}
	return privateKey, privateKey.Public()
}

func testAPI() error {
	testConfig := opticgo.Config{
		ApiUrl:          opticgo.MustUrl(cloudflare.ApiUrl).ResolveReference(opticgo.MustUrl(cloudflare.ApiVersion)),
		OpticUrl:        opticgo.MustUrl("http://localhost:8889"),
		ProxyListenAddr: "localhost",
		DebugPrint:      true,
		TripFunc: func(tripper http.RoundTripper) http.RoundTripper {
			return &CustomTransport{}
		},
		InternetCheckTimeout: 10 * time.Second,
	}
	tester, err := opticgo.NewTester(testConfig)
	if err != nil {
		return err
	}
	proxyErrChan, err := tester.StartProxy()
	if err != nil {
		return err
	}
	go func() {
		for err := range proxyErrChan {
			log.Fatalln(err)
		}
	}()

	client := sling.New().Client(&http.Client{Transport: CustomTransport{}}).Path(testConfig.OpticUrl.String())

	_, publicKey := generateKeyPair()
	_, publicKey2 := generateKeyPair()
	regData := struct {
		PublicKey string `json:"key"`
		InstallID string `json:"install_id"`
		FcmToken  string `json:"fcm_token"`
		Tos       string `json:"tos"`
		Model     string `json:"model"`
		Type      string `json:"type"`
		Locale    string `json:"locale"`
	}{
		publicKey.String(),
		"", // not empty on actual client
		"", // not empty on actual client
		util.GetTimestamp(),
		"PC",
		"Android",
		"en_US",
	}
	var regResp map[string]interface{}
	regPath := path.Join(testConfig.ApiUrl.Path, "reg")
	if _, err := client.New().Post(regPath).BodyJSON(regData).ReceiveSuccess(&regResp); err != nil {
		return err
	}

	deviceId := regResp["id"].(string)
	accessToken := regResp["token"].(string)
	initialLicenseKey := regResp["account"].(map[string]interface{})["license"].(string)

	defaultHeaders["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)

	var tests = []opticgo.TestDefinition{
		{
			"get device",
			nil,
			fmt.Sprintf("reg/%s", deviceId),
			"GET",
		},
		{
			"get account",
			nil,
			fmt.Sprintf("reg/%s/account", deviceId),
			"GET",
		},
		{
			"get account devices",
			nil,
			fmt.Sprintf("reg/%s/account/devices", deviceId),
			"GET",
		},
		{
			"set device active",
			struct {
				Active bool `json:"active"`
			}{
				true,
			},
			fmt.Sprintf("reg/%s/account/reg/%s", deviceId, deviceId),
			"PATCH",
		},
		{
			"set device name",
			struct {
				Name string `json:"name"`
			}{
				"TEST",
			},
			fmt.Sprintf("reg/%s/account/reg/%s", deviceId, deviceId),
			"PATCH",
		},
		{
			"get account devices, this time with the name set",
			nil,
			fmt.Sprintf("reg/%s/account/devices", deviceId),
			"GET",
		},
		{
			"get client config",
			nil,
			fmt.Sprintf("client_config"),
			"GET",
		},
		{
			"set license key",
			struct {
				LicenseKey string `json:"license"`
			}{
				initialLicenseKey, // TODO: don't use same key
			},
			fmt.Sprintf("reg/%s/account", deviceId),
			"PUT",
		},
		{
			"set public key",
			struct {
				PublicKey string `json:"key"`
			}{
				publicKey2.String(),
			},
			fmt.Sprintf("reg/%s", deviceId),
			"PATCH",
		},
		{
			"recreate license key",
			nil,
			fmt.Sprintf("reg/%s/account/license", deviceId),
			"POST",
		},
	}

	errChan, _, err := tester.StartAll(tests)
	if err != nil {
		return err
	}
	errText := ""
	for err := range errChan {
		errText += err.Error() + "\n"
	}
	if errText != "" {
		return errors.New(errText)
	}
	return nil
}
