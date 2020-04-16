package util

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"strings"
	"wgcf/cloudflare/api"
	"wgcf/cloudflare/structs/request"
	"wgcf/cloudflare/structs/resp"
	"wgcf/config"
	"wgcf/util"
)

func FormatMessage(shortMessage string, longMessage string) string {
	if longMessage != "" {
		if strings.HasPrefix(longMessage, "\n") {
			longMessage = longMessage[1:]
		}
		longMessage = strings.Replace(longMessage, "\n", " ", -1)
	}
	if shortMessage != "" && longMessage != "" {
		return shortMessage + ". " + longMessage
	} else if shortMessage != "" {
		return shortMessage
	} else {
		return longMessage
	}
}

func ValidAccount() bool {
	return viper.GetString(config.DeviceId) != "" &&
		viper.GetString(config.AccessToken) != "" &&
		viper.GetString(config.PrivateKey) != ""
}

func CreateContext() *config.Context {
	ctx := config.Context{
		DeviceId:    viper.GetString(config.DeviceId),
		AccessToken: viper.GetString(config.AccessToken),
		PrivateKey:  viper.GetString(config.PrivateKey),
		LicenseKey:  viper.GetString(config.LicenseKey),
	}
	return &ctx
}

func PrintDeviceData(data *resp.DeviceData, boundDevice *resp.BoundDevice) {
	log.Println("=======================================")
	log.Println("Device name:", boundDevice.Name)
	log.Println("Device model:", data.Model)
	log.Println("Device active:", boundDevice.Active)
	log.Println("Account type:", data.Account.AccountType)
	log.Println("Warp+ enabled:", data.Account.WarpPlus)
	log.Println("Premium data:", data.Account.PremiumData)
	log.Println("Quota:", data.Account.Quota)
	log.Println("=======================================")
}

// modifies device pointer if updated
func EnsureWarpEnabled(ctx *config.Context, deviceData *resp.DeviceData) error {
	if !deviceData.WarpEnabled {
		log.Println("Enabling Warp")
		success, err := api.EnableWarp(ctx)
		if err != nil {
			return err
		}
		if !success {
			return errors.New("could not enable warp")
		}
		newDeviceData, err := api.GetDeviceData(ctx)
		if err != nil {
			return err
		}
		*deviceData = *newDeviceData
	}
	return nil
}

// modifies device pointer if updated
func EnsureDeviceActive(ctx *config.Context, device *resp.BoundDevice) error {
	if !device.Active {
		log.Println("Device not active, activating it")
		newDevice, err := api.SetBoundDeviceActive(ctx, &request.DeviceActiveData{Active: true})
		if err != nil {
			return err
		}
		if !newDevice.Active {
			return errors.New("could not activate device")
		}
		*device = *newDevice
	}
	return nil
}

// changing the bound account (e.g. changing license key) will reset the device name
func SetDeviceName(ctx *config.Context, deviceName string) (*resp.BoundDevice, error) {
	if deviceName == "" {
		deviceName += util.RandomHexString(3)
	}
	device, err := api.SetBoundDeviceName(ctx, &request.DeviceNameData{Name: deviceName})
	if err != nil {
		return nil, err
	}
	if device.Name != deviceName {
		return nil, errors.New("could not update device name")
	}
	return device, nil
}
