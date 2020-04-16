package update

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"wgcf/cloudflare/api"
	"wgcf/cloudflare/structs/resp"
	. "wgcf/cmd/util"
	"wgcf/config"
	"wgcf/util"
)

var deviceName string

var shortMsg = "Updates the current Cloudflare Warp account, preparing it for connection"

var Cmd = &cobra.Command{
	Use:   "update",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
If a new/different license key is provided, the current device will be bound to the new key and its parent account. 
Please note that there is a maximum limit of 5 active devices linked to the same account at a given time.
Will change various account settings to ensure WireGuard connection will succeed.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateAccount(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Device name displayed under the 1.1.1.1 app")
}

func updateAccount() error {
	if !ValidAccount() {
		return errors.New("no account detected")
	}

	ctx := CreateContext()
	deviceData, err := api.GetDeviceData(ctx)
	if err != nil {
		return err
	}
	boundDevice, err := api.GetBoundDevice(ctx)
	if err != nil {
		return err
	}

	if err := ensureLicenseKeyUpdated(ctx, deviceData, boundDevice); err != nil {
		return err
	}
	if err := EnsureDeviceActive(ctx, boundDevice); err != nil {
		return err
	}

	PrintDeviceData(deviceData, boundDevice)
	log.Println("Successfully updated Cloudflare Warp account")
	return nil
}

// modifies device pointers if updated
func ensureLicenseKeyUpdated(ctx *config.Context, deviceData *resp.DeviceData, device *resp.BoundDevice) error {
	if deviceData.Account.LicenseKey != ctx.LicenseKey {
		log.Println("Updated license key detected, re-binding device to new account")
		if _, err := updateLicenseKey(ctx); err != nil {
			return err
		}
		if _, err := SetDeviceName(ctx, deviceName); err != nil {
			return err
		}
		// make sure new license key is synced with config file if it came from other stream (e.g. env variable)
		if err := viper.WriteConfig(); err != nil {
			return err
		}

		newDeviceData, err := api.GetDeviceData(ctx)
		if err != nil {
			return err
		}
		newBoundDevice, err := api.GetBoundDevice(ctx)
		if err != nil {
			return err
		}
		*deviceData = *newDeviceData
		*device = *newBoundDevice
	}
	return nil
}

func updateLicenseKey(ctx *config.Context) (*resp.UpdateLicenseData, error) {
	accountData, err := api.UpdateLicenseKey(ctx)
	if err != nil {
		return nil, err
	}
	//TODO: Check if key actually updated
	return accountData, nil
}
