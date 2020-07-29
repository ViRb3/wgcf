package shared

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"strings"
	"wgcf/cloudflare"
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

func IsConfigValidAccount() bool {
	return viper.GetString(config.DeviceId) != "" &&
		viper.GetString(config.AccessToken) != "" &&
		viper.GetString(config.PrivateKey) != ""
}

func CreateContext() *config.Context {
	ctx := config.Context{
		DeviceId:    viper.GetString(config.DeviceId),
		AccessToken: viper.GetString(config.AccessToken),
		LicenseKey:  viper.GetString(config.LicenseKey),
	}
	return &ctx
}

func PrintDeviceData(thisDevice *cloudflare.Device, boundDevice *cloudflare.Device) {
	log.Println("=======================================")
	log.Printf("%-13s : %s\n", "Device name", *boundDevice.Name)
	log.Printf("%-13s : %s\n", "Device model", thisDevice.Model)
	log.Printf("%-13s : %t\n", "Device active", *boundDevice.Active)
	log.Printf("%-13s : %s\n", "Account type", thisDevice.Account.AccountType)
	log.Printf("%-13s : %s\n", "Role", thisDevice.Account.Role)
	log.Printf("%-13s : %d\n", "Premium data", thisDevice.Account.PremiumData)
	log.Printf("%-13s : %d\n", "Quota", thisDevice.Account.Quota)
	log.Println("=======================================")
}

// changing the bound account (e.g. changing license key) will reset the device name
func SetDeviceName(ctx *config.Context, deviceName string) (*cloudflare.Device, error) {
	if deviceName == "" {
		deviceName += util.RandomHexString(3)
	}
	device, err := cloudflare.SetThisBoundDeviceName(ctx, cloudflare.SetBoundDeviceNameRequest{Name: deviceName})
	if err != nil {
		return nil, err
	}
	if device.Name == nil || *device.Name != deviceName {
		return nil, errors.New("could not update device name")
	}
	return device, nil
}
