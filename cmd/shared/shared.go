package shared

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/ViRb3/wgcf/v2/cloudflare"
	"github.com/ViRb3/wgcf/v2/config"
	"github.com/ViRb3/wgcf/v2/util"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
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

func F32ToHumanReadable(number float32) string {
	for i := 8; i >= 0; i-- {
		humanReadable := number / float32(math.Pow(1024, float64(i)))
		if humanReadable >= 1 && humanReadable < 1024 {
			return fmt.Sprintf("%.2f %ciB", humanReadable, "KMGTPEZY"[i-1])
		}
	}
	return fmt.Sprintf("%.2f B", number)
}

func PrintDeviceData(thisDevice *cloudflare.Device, boundDevice *cloudflare.BoundDevice) {
	log.Println("=======================================")
	log.Printf("%-13s : %s\n", "Device name", *boundDevice.Name)
	log.Printf("%-13s : %s\n", "Device model", thisDevice.Model)
	log.Printf("%-13s : %t\n", "Device active", boundDevice.Active)
	log.Printf("%-13s : %s\n", "Account type", thisDevice.Account.AccountType)
	log.Printf("%-13s : %s\n", "Role", thisDevice.Account.Role)
	log.Printf("%-13s : %s\n", "Premium data", F32ToHumanReadable(thisDevice.Account.PremiumData))
	log.Printf("%-13s : %s\n", "Quota", F32ToHumanReadable(thisDevice.Account.Quota))
	log.Println("=======================================")
}

// changing the bound account (e.g. changing license key) will reset the device name
func SetDeviceName(ctx *config.Context, deviceName string) (*cloudflare.BoundDevice, error) {
	if deviceName == "" {
		deviceName += util.RandomHexString(3)
	}
	device, err := cloudflare.UpdateSourceBoundDeviceName(ctx, deviceName)
	if err != nil {
		return nil, err
	}
	if device.Name == nil || *device.Name != deviceName {
		return nil, errors.New("could not update device name")
	}
	return device, nil
}
