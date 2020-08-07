package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dghubble/sling"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"
	"wgcf/util"
	"wgcf/wireguard"
)

var globalClient *sling.Sling
var defaultHeaders = map[string]string{"User-Agent": "okhttp/3.12.1"}
var customTransport = &CustomTransport{}
var parsedApiEndpoint *url.URL
var wg sync.WaitGroup

const (
	apiEndpoint = "https://api.cloudflareclient.com/"
	OpticPort   = "8888"
	DebugPrint  = false
)

var fixedTransport = &http.Transport{
	TLSClientConfig: &tls.Config{
		// Match app's TLS config or API will reject us with code 403 error 1020
		MinVersion: tls.VersionTLS10,
		MaxVersion: tls.VersionTLS12},
}

type CustomTransport struct{}

func (CustomTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Println("Requesting: ", r.URL.String())
	r.Header.Del("X-Forwarded-For") // inserted by httputil.NewSingleHostReverseProxy
	if DebugPrint {
		b, err := httputil.DumpRequestOut(r, false)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(b))
	}
	return fixedTransport.RoundTrip(r)
}

func init() {
	parsed, err := url.Parse(apiEndpoint)
	if err != nil {
		log.Fatal(err)
		return
	}
	parsedApiEndpoint = parsed
	doer, _ := util.NewBaseDoer()
	globalClient = sling.New().Doer(doer).SetMany(defaultHeaders).Path("http://localhost:" + OpticPort)
}

func serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(parsedApiEndpoint)
		proxy.Director = func(r *http.Request) {
			r.Host = parsedApiEndpoint.Host
			r.URL = parsedApiEndpoint.ResolveReference(r.URL)
		}
		proxy.Transport = customTransport
		proxy.ServeHTTP(w, r)
		log.Println("Returned: ", r.URL.String())
		wg.Done()
	})
	return http.ListenAndServe(":"+os.Getenv("OPTIC_API_PORT"), nil)
}

func main() {
	waitInternetAccess()
	go func() { log.Fatal(serve()) }()
	if err := testAPI(); err != nil {
		log.Println(err)
	}
	wg.Wait()
	// wait for last request to be returned to Optic
	time.Sleep(1 * time.Second)
}

// e.g. get firewall permission
func waitInternetAccess() {
	for {
		_, err := http.Get("https://google.com/")
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func generateKeyPair() (*wireguard.Key, *wireguard.Key) {
	privateKey, err := wireguard.NewPrivateKey()
	if err != nil {
		panic(err)
	}
	return privateKey, privateKey.Public()
}

type TestConfig struct {
	name       string
	data       interface{}
	requestUrl string
	method     string
}

func runTest(config *TestConfig) error {
	log.Println("Running test: ", config.requestUrl)
	request := globalClient.New()
	switch config.method {
	case "GET":
		request.Get(config.requestUrl)
	case "POST":
		request.Post(config.requestUrl)
	case "PUT":
		request.Put(config.requestUrl)
	case "PATCH":
		request.Patch(config.requestUrl)
	default:
		return errors.New("unknown request method")
	}
	if config.data != nil {
		request.BodyJSON(config.data)
	}
	wg.Add(1)
	if _, err := request.ReceiveSuccess(nil); err != nil {
		return err
	}
	return nil
}

func testAPI() error {
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
	wg.Add(1)
	if _, err := globalClient.New().Post("/v0a977/reg").BodyJSON(regData).ReceiveSuccess(&regResp); err != nil {
		return err
	}

	deviceId := regResp["id"].(string)
	accessToken := regResp["token"].(string)
	initialLicenseKey := regResp["account"].(map[string]interface{})["license"].(string)

	globalClient.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	var tests = []TestConfig{
		{
			"get device",
			nil,
			fmt.Sprintf("/v0a977/reg/%s", deviceId),
			"GET",
		},
		{
			"get account",
			nil,
			fmt.Sprintf("/v0a977/reg/%s/account", deviceId),
			"GET",
		},
		{
			"get account devices",
			nil,
			fmt.Sprintf("/v0a977/reg/%s/account/devices", deviceId),
			"GET",
		},
		{
			"set device active",
			struct {
				Active bool `json:"active"`
			}{
				true,
			},
			fmt.Sprintf("/v0a977/reg/%s/account/reg/%s", deviceId, deviceId),
			"PATCH",
		},
		{
			"set device name",
			struct {
				Name string `json:"name"`
			}{
				"TEST",
			},
			fmt.Sprintf("/v0a977/reg/%s/account/reg/%s", deviceId, deviceId),
			"PATCH",
		},
		{
			"get account devices, this time with the name set",
			nil,
			fmt.Sprintf("/v0a977/reg/%s/account/devices", deviceId),
			"GET",
		},
		{
			"get client config",
			nil,
			fmt.Sprintf("/v0a977/client_config"),
			"GET",
		},
		{
			"set license key",
			struct {
				LicenseKey string `json:"license"`
			}{
				initialLicenseKey, // TODO: don't use same key
			},
			fmt.Sprintf("/v0a977/reg/%s/account", deviceId),
			"PUT",
		},
		{
			"set public key",
			struct {
				PublicKey string `json:"key"`
			}{
				publicKey2.String(),
			},
			fmt.Sprintf("/v0a977/reg/%s", deviceId),
			"PATCH",
		},
		{
			"recreate license key",
			nil,
			fmt.Sprintf("/v0a977/reg/%s/account/license", deviceId),
			"POST",
		},
	}

	for _, test := range tests {
		if err := runTest(&test); err != nil {
			return err
		}
	}
	return nil
}
