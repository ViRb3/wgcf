package update

import (
	"log"

	"github.com/ViRb3/wgcf/v2/cloudflare"
	. "github.com/ViRb3/wgcf/v2/cmd/shared"
	"github.com/ViRb3/wgcf/v2/config"
	"github.com/ViRb3/wgcf/v2/util"
	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
)

var deviceName string

var shortMsg = "Updates the current Cloudflare Warp account, preparing it for connection"

var Cmd = &cobra.Command{
	Use:   "update",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
If a new/different license key is provided, the current device will be bound to the new key and its parent account. 
Please note that there is a maximum limit of 5 active devices linked to the same account at a given time.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateAccount(); err != nil {
			log.Fatalf("%+v\n", err)
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Device name displayed under the 1.1.1.1 app")
}

func updateAccount() error {
	if !IsConfigValidAccount() {
		return errors.New("no account detected")
	}

	ctx := CreateContext()
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	_, thisDevice, err = ensureLicenseKeyUpToDate(ctx, thisDevice)
	if err != nil {
		return errors.WithStack(err)
	}

	boundDevice, err := cloudflare.GetSourceBoundDevice(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	if boundDevice.Name == nil || (deviceName != "" && deviceName != *boundDevice.Name) {
		log.Println("Setting device name")
		if _, err := SetDeviceName(ctx, deviceName); util.IsHttp500Error(err) {
			// server-side issue, but the operation still works
		} else if err != nil {
			return errors.WithStack(err)
		}
	}

	if _, err = cloudflare.UpdateSourceBoundDeviceActive(ctx, true); util.IsHttp500Error(err) {
		// server-side issue, but the operation still works
	} else if err != nil {
		return errors.WithStack(err)
	}

	boundDevice, err = cloudflare.GetSourceBoundDevice(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	if !boundDevice.Active {
		return errors.New("failed activating device")
	}

	PrintDeviceData(thisDevice, boundDevice)
	log.Println("Successfully updated Cloudflare Warp account")
	return nil
}

func ensureLicenseKeyUpToDate(ctx *config.Context, thisDevice *cloudflare.SourceDevice) (*cloudflare.Account, *cloudflare.SourceDevice, error) {
	if thisDevice.Account.License != ctx.LicenseKey {
		log.Println("Updated license key detected, re-binding device to new account")
		return updateLicenseKey(ctx)
	}
	return nil, thisDevice, nil
}

func updateLicenseKey(ctx *config.Context) (*cloudflare.Account, *cloudflare.SourceDevice, error) {

	if _, err := cloudflare.UpdateLicenseKey(ctx); err != nil {
		return nil, nil, errors.WithStack(err)
	}

	account, err := cloudflare.GetAccount(ctx)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if account.License != ctx.LicenseKey {
		return nil, nil, errors.New("failed to update license key")
	}

	return account, thisDevice, nil
}
