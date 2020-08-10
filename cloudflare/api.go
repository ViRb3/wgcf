package cloudflare

import (
	"crypto/tls"
	"github.com/antihax/optional"
	"net/http"
	"time"
	"wgcf/config"
	"wgcf/openapi"
	"wgcf/util"
	"wgcf/wireguard"
)

var apiClient = MakeApiClient(nil)
var apiClientAuth *openapi.APIClient

func MakeApiClient(authToken *string) *openapi.APIClient {
	httpClient := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// Match app's TLS config or API will reject us with code 403 error 1020
				MinVersion: tls.VersionTLS10,
				MaxVersion: tls.VersionTLS12},
		}}
	apiClient := openapi.NewAPIClient(&openapi.Configuration{
		BasePath:      "https://api.cloudflareclient.com",
		DefaultHeader: map[string]string{},
		UserAgent:     "okhttp/3.12.1",
		Debug:         false,
		Servers:       nil,
		HTTPClient:    &httpClient,
	})
	if authToken != nil {
		apiClient.GetConfig().DefaultHeader["Authorization"] = "Bearer " + *authToken
	}
	return apiClient
}

func Register(publicKey *wireguard.Key, deviceModel string) (*openapi.Register200Response, error) {
	timestamp := util.GetTimestamp()
	result, _, err := apiClient.DefaultApi.Register(nil,
		&openapi.RegisterOpts{optional.NewInterface(openapi.RegisterRequest{
			Key:       publicKey.String(),
			InstallId: "", // not empty on actual client
			FcmToken:  "", // not empty on actual client
			Tos:       timestamp,
			Model:     deviceModel,
			Type:      "Android",
			Locale:    "en_US",
		})})
	return &result, err
}

type Device openapi.GetSourceDevice200Response

func GetSourceDevice(ctx *config.Context) (*Device, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultApi.GetSourceDevice(nil, ctx.DeviceId)
	castResult := Device(result)
	return &castResult, err
}

func globalClientAuth(authToken string) *openapi.APIClient {
	if apiClientAuth == nil {
		apiClientAuth = MakeApiClient(&authToken)
	}
	return apiClientAuth
}

type Account openapi.GetAccount200Response

func GetAccount(ctx *config.Context) (*Account, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultApi.GetAccount(nil, ctx.DeviceId)
	castResult := Account(result)
	return &castResult, err
}

func UpdateLicenseKey(ctx *config.Context, newPublicKey string) (*openapi.UpdateAccount200Response, *Device, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultApi.UpdateAccount(nil, ctx.DeviceId,
		&openapi.UpdateAccountOpts{optional.NewInterface(openapi.UpdateAccountRequest{License: ctx.LicenseKey})})
	if err != nil {
		return nil, nil, err
	}
	// change public key as per official client
	result2, _, err := globalClientAuth(ctx.AccessToken).DefaultApi.UpdateSourceDevice(nil, ctx.DeviceId,
		&openapi.UpdateSourceDeviceOpts{optional.NewInterface(openapi.UpdateSourceDeviceRequest{Key: newPublicKey})})
	castResult := Device(result2)
	if err != nil {
		return nil, nil, err
	}
	return &result, &castResult, nil
}

type BoundDevice openapi.GetBoundDevices200Response

func GetBoundDevices(ctx *config.Context) ([]BoundDevice, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultApi.GetBoundDevices(nil, ctx.DeviceId)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return FindDevice(result, ctx.DeviceId)
}

func UpdateSourceBoundDeviceName(ctx *config.Context, newName string) (*BoundDevice, error) {
	return UpdateSourceBoundDevice(ctx,
		&openapi.UpdateBoundDeviceOpts{optional.NewInterface(openapi.UpdateBoundDeviceRequest{
			Name: newName,
		})})
}

func UpdateSourceBoundDeviceActive(ctx *config.Context, active bool) (*BoundDevice, error) {
	return UpdateSourceBoundDevice(ctx,
		&openapi.UpdateBoundDeviceOpts{optional.NewInterface(openapi.UpdateBoundDeviceRequest{
			Active: active,
		})})
}

func UpdateSourceBoundDevice(ctx *config.Context, data *openapi.UpdateBoundDeviceOpts) (*BoundDevice, error) {
	result, _, err := globalClientAuth(ctx.AccessToken).DefaultApi.UpdateBoundDevice(nil, ctx.DeviceId, ctx.DeviceId, data)
	if err != nil {
		return nil, err
	}
	var castResult []BoundDevice
	for _, device := range result {
		castResult = append(castResult, BoundDevice(device))
	}
	return FindDevice(castResult, ctx.DeviceId)
}
