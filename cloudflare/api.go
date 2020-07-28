package cloudflare

import (
	"crypto/tls"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"wgcf/config"
	"wgcf/util"
	"wgcf/wireguard"
)

var globalClient *sling.Sling
var defaultHeaders = map[string]string{"User-Agent": "okhttp/3.12.1"}

func init() {
	doer, client := util.NewBaseDoer()
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			// Match app's TLS config or API will reject us with code 403 error 1020
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS12},
	}
	globalClient = sling.New().Doer(doer)
	globalClient.SetMany(defaultHeaders)
}

func globalClientAuth(token string) *sling.Sling {
	return globalClient.New().Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func Register(publicKey *wireguard.Key, deviceModel string) (*Device, error) {
	timestamp := GetTimestamp()
	data := RegRequest{
		PublicKey: publicKey.String(),
		InstallID: "", // not empty on actual client
		FcmToken:  "", // not empty on actual client
		Tos:       timestamp,
		Model:     deviceModel,
		Type:      "Android",
		Locale:    "en_US",
	}
	var result Device
	if _, err := globalClient.New().
		Post(regUrl).
		BodyJSON(data).
		ReceiveSuccess(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetThisDevice(ctx *config.Context) (*Device, error) {
	var result Device
	if _, err := globalClientAuth(ctx.AccessToken).
		Get(getThisDeviceUrl(ctx.DeviceId)).
		ReceiveSuccess(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetAccount(ctx *config.Context) (*Account, error) {
	var result Account
	if _, err := globalClientAuth(ctx.AccessToken).
		Get(getAccountUrl(ctx.DeviceId)).
		ReceiveSuccess(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdateLicenseKey(ctx *config.Context, data2 UpdateLicenseKeyRequest2) (*Account, *Device, error) {
	data := UpdateLicenseKeyRequest{LicenseKey: ctx.LicenseKey}
	var result Account
	if _, err := globalClientAuth(ctx.AccessToken).
		Put(getAccountUrl(ctx.DeviceId)).
		BodyJSON(data).
		ReceiveSuccess(&result); err != nil {
		return nil, nil, err
	}
	// change public key as per official client
	var result2 Device
	if _, err := globalClientAuth(ctx.AccessToken).
		Patch(getThisDeviceUrl(ctx.DeviceId)).
		BodyJSON(data2).
		ReceiveSuccess(&result2); err != nil {
		return nil, nil, err
	}
	return &result, &result2, nil
}

func GetBoundDevices(ctx *config.Context) ([]Device, error) {
	var result []Device
	if _, err := globalClientAuth(ctx.AccessToken).
		Get(getBoundDevicesUrl(ctx.DeviceId)).
		ReceiveSuccess(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetThisBoundDevice(ctx *config.Context) (*Device, error) {
	devices, err := GetBoundDevices(ctx)
	if err != nil {
		return nil, err
	}
	return FindDevice(devices, ctx.DeviceId)
}

func SetThisBoundDeviceName(ctx *config.Context, data SetBoundDeviceNameRequest) (*Device, error) {
	return setThisBoundDeviceData(ctx, data)
}

func SetThisBoundDeviceActive(ctx *config.Context, data SetBoundDeviceActiveRequest) (*Device, error) {
	return setThisBoundDeviceData(ctx, data)
}

func setThisBoundDeviceData(ctx *config.Context, data interface{}) (*Device, error) {
	var result []Device
	if _, err := globalClientAuth(ctx.AccessToken).
		Patch(getBoundDeviceUrl(ctx.DeviceId, ctx.DeviceId)).
		BodyJSON(data).
		ReceiveSuccess(&result); err != nil {
		return nil, err
	}
	device, err := FindDevice(result, ctx.DeviceId)
	if err != nil {
		return nil, err
	}
	return device, nil
}
