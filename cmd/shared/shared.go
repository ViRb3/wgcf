package shared

import (
	errors2 "errors"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/ViRb3/wgcf/v2/cloudflare"
	"github.com/ViRb3/wgcf/v2/config"
	"github.com/ViRb3/wgcf/v2/util"
	"github.com/dustin/go-humanize"

	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"
)

var ErrExistingAccount = errors2.New("existing account detected, refusing to overwrite")
var ErrNoAccount = errors2.New("no account detected, register one first")
var ErrTOSNotAccepted = errors2.New("TOS not accepted")

func RunCommandFatal(cmd func() error) {
	if err := cmd(); err != nil {
		expectedErrs := []error{ErrNoAccount, ErrExistingAccount, ErrTOSNotAccepted}
		if slices.ContainsFunc(expectedErrs, func(e error) bool { return errors.Is(err, e) }) {
			log.Fatalln(err)
		} else {
			log.Fatalf("%+v\n", err)
		}
	}
}

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

func isConfigValidAccount() bool {
	return viper.GetString(config.DeviceId) != "" &&
		viper.GetString(config.AccessToken) != "" &&
		viper.GetString(config.PrivateKey) != ""
}

func EnsureConfigValidAccount() error {
	if isConfigValidAccount() {
		return nil
	} else {
		return ErrNoAccount
	}
}

func EnsureNoExistingAccount() error {
	if isConfigValidAccount() {
		return ErrExistingAccount
	} else {
		return nil
	}
}

func CreateContext() *config.Context {
	ctx := config.Context{
		DeviceId:    viper.GetString(config.DeviceId),
		AccessToken: viper.GetString(config.AccessToken),
		LicenseKey:  viper.GetString(config.LicenseKey),
	}
	return &ctx
}

func PrintAccountDetails(account *cloudflare.Account, boundDevices []cloudflare.BoundDevice) {
	log.Println("Printing account details:")
	fmt.Println()
	fmt.Println("================================================================")
	fmt.Println("Account")
	fmt.Println("================================================================")
	fmt.Printf("%-12s : %s\n", "Id", account.Id)
	fmt.Printf("%-12s : %s\n", "Account type", account.AccountType)
	fmt.Printf("%-12s : %s\n", "Created", account.Created)
	fmt.Printf("%-12s : %s\n", "Updated", account.Updated)
	fmt.Printf("%-12s : %s\n", "Premium data", humanize.Bytes(uint64(account.PremiumData)))
	fmt.Printf("%-12s : %s\n", "Quota", humanize.Bytes(uint64(account.Quota)))
	fmt.Printf("%-12s : %s\n", "Role", account.Role)
	fmt.Println()
	fmt.Println("================================================================")
	fmt.Println("Devices")
	fmt.Println("================================================================")
	for _, device := range boundDevices {
		name := "N/A"
		if device.Name != nil {
			name = *device.Name
		}
		id := device.Id
		if device.Id == viper.GetString(config.DeviceId) {
			id += " (current)"
		}
		fmt.Printf("%-9s : %s\n", "Id", id)
		fmt.Printf("%-9s : %s\n", "Type", device.Type)
		fmt.Printf("%-9s : %s\n", "Model", device.Model)
		fmt.Printf("%-9s : %s\n", "Name", name)
		fmt.Printf("%-9s : %t\n", "Active", device.Active)
		fmt.Printf("%-9s : %s\n", "Created", device.Created)
		fmt.Printf("%-9s : %s\n", "Activated", device.Activated)
		fmt.Printf("%-9s : %s\n", "Role", device.Role)
		fmt.Println()
	}
}

func SetDeviceName(ctx *config.Context, deviceName string) (*cloudflare.BoundDevice, error) {
	if deviceName == "" {
		deviceName += util.RandomHexString(3)
	}
	device, err := cloudflare.UpdateSourceBoundDeviceName(ctx, ctx.DeviceId, deviceName)
	if err == nil {
		if device.Name == nil || *device.Name != deviceName {
			return nil, errors.New("could not update device name")
		}
	} else if util.IsHttp500Error(err) {
		// server-side issue, but the operation still works
	} else {
		return nil, errors.WithStack(err)
	}

	return device, nil
}
