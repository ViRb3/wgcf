package update

import (
	"log"
	"slices"

	"github.com/ViRb3/wgcf/v2/cloudflare"
	. "github.com/ViRb3/wgcf/v2/cmd/shared"
	"github.com/ViRb3/wgcf/v2/config"
	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deviceName string
var licenseKey string
var removeDevices []string
var deactivateDevices []string
var activateDevices []string

var shortMsg = "Updates the current Cloudflare Warp account, preparing it for connection"

var Cmd = &cobra.Command{
	Use:   "update",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
If a new/different license key is provided, the current device will be bound to the new key and its parent account. 
Please note that there is a maximum limit of 5 active devices linked to the same account at a given time.`),
	Run: func(cmd *cobra.Command, args []string) {
		RunCommandFatal(updateAccount)
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Update this device's name")
	Cmd.PersistentFlags().StringVarP(&licenseKey, "license-key", "l", "", "Update this device's license key")
	Cmd.PersistentFlags().StringSliceVar(&removeDevices, "remove", []string{}, "Remove devices by their ID (can be used multiple times)")
	Cmd.PersistentFlags().StringSliceVar(&deactivateDevices, "deactivate", []string{}, "Deactivate devices by their ID (can be used multiple times)")
	Cmd.PersistentFlags().StringSliceVar(&activateDevices, "activate", []string{}, "Activate devices by their ID (can be used multiple times)")
}

func updateAccount() error {
	if err := EnsureConfigValidAccount(); err != nil {
		return errors.WithStack(err)
	}

	ctx := CreateContext()
	if licenseKey != "" {
		ctx.LicenseKey = licenseKey
	}

	account, err := cloudflare.GetAccount(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	if account.License != ctx.LicenseKey {
		log.Println("Updated license key detected, re-binding device to new account")
		if _, err := cloudflare.UpdateLicenseKey(ctx); err != nil {
			return errors.WithStack(err)
		}
		viper.Set(config.LicenseKey, ctx.LicenseKey)
		if err := viper.WriteConfig(); err != nil {
			return errors.WithStack(err)
		}
	}

	boundDevice, err := cloudflare.GetSourceBoundDevice(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	if boundDevice.Name == nil || (deviceName != "" && deviceName != *boundDevice.Name) {
		log.Println("Setting device name")
		if _, err := SetDeviceName(ctx, deviceName); err != nil {
			return errors.WithStack(err)
		}
	}

	deviceActions := map[string]func() error{}
	for _, deviceId := range removeDevices {
		deviceActions[deviceId] = func() error {
			log.Printf("Deleting device: %s\n", deviceId)
			err := cloudflare.DeleteBoundDevice(ctx, deviceId)
			return errors.WithStack(err)
		}
	}
	for _, deviceId := range deactivateDevices {
		deviceActions[deviceId] = func() error {
			log.Printf("Deactivating device: %s\n", deviceId)
			if _, err := cloudflare.UpdateSourceBoundDeviceActive(ctx, deviceId, false); err != nil {
				return errors.WithStack(err)
			}
			return errors.WithStack(err)
		}
	}
	for _, deviceId := range activateDevices {
		deviceActions[deviceId] = func() error {
			log.Printf("Activating device: %s\n", deviceId)
			_, err := cloudflare.UpdateSourceBoundDeviceActive(ctx, deviceId, true)
			return errors.WithStack(err)
		}
	}

	boundDevices, err := cloudflare.GetBoundDevices(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	for deviceId, action := range deviceActions {
		deviceFound := slices.ContainsFunc(boundDevices, func(device cloudflare.BoundDevice) bool {
			return device.Id == deviceId
		})
		if !deviceFound {
			return errors.Newf("device not found: %s", deviceId)
		}
		if err := action(); err != nil {
			return errors.WithStack(err)
		}
	}

	// refresh in case e.g. account type changed
	account, err = cloudflare.GetAccount(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	boundDevices, err = cloudflare.GetBoundDevices(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	PrintAccountDetails(account, boundDevices)

	log.Println("Successfully updated Cloudflare Warp account")
	return nil
}
