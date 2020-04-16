package register

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"wgcf/cloudflare/api"
	"wgcf/cloudflare/url"
	. "wgcf/cmd/util"
	"wgcf/config"
	"wgcf/util"
	"wgcf/wireguard"
)

var deviceName string
var deviceModel string
var existingKey = false
var acceptedTOS = false
var shortMsg = "Registers a new Cloudflare Warp device and creates a new account, preparing it for connection"

var Cmd = &cobra.Command{
	Use:   "register",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
Will change various account settings to ensure WireGuard connection will succeed.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := registerAccount(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Device name displayed under the 1.1.1.1 app (defaults to random)")
	Cmd.PersistentFlags().StringVarP(&deviceModel, "model", "m", "PC", "Device model displayed under the 1.1.1.1 app")
	Cmd.PersistentFlags().BoolVar(&existingKey, "existing-key", false, "Use existing private key for the new account")
	Cmd.PersistentFlags().BoolVar(&acceptedTOS, "accept-tos", false, "Accept Cloudflare's Terms of Service non-interactively")
}

func registerAccount() error {
	if ValidAccount() {
		return errors.New("existing account detected")
	}

	ctx, err := createNewAccount()
	if err != nil {
		return err
	}
	deviceData, err := api.GetDeviceData(ctx)
	if err != nil {
		return err
	}
	boundDevice, err := api.GetBoundDevice(ctx)
	if err != nil {
		return err
	}

	if err := EnsureWarpEnabled(ctx, deviceData); err != nil {
		return err
	}
	if err := EnsureDeviceActive(ctx, boundDevice); err != nil {
		return err
	}

	PrintDeviceData(deviceData, boundDevice)
	log.Println("Successfully created Cloudflare Warp account")
	return nil
}

func createNewAccount() (*config.Context, error) {
	if err := ensurePrivateKey(); err != nil {
		return nil, err
	}
	if accepted, err := checkTOS(); err != nil || !accepted {
		return nil, err
	}

	registrationData, err := api.Register(viper.GetString(config.PrivateKey), deviceModel)
	if err != nil {
		return nil, err
	}

	viper.Set(config.DeviceId, registrationData.ID)
	viper.Set(config.AccessToken, registrationData.AccessToken)
	viper.Set(config.LicenseKey, registrationData.LicenseKey)
	if err := viper.WriteConfig(); err != nil {
		return nil, err
	}

	ctx := CreateContext()
	if _, err := SetDeviceName(ctx, deviceName); err != nil {
		return ctx, err
	}
	return ctx, nil
}

func ensurePrivateKey() error {
	if viper.GetString(config.PrivateKey) == "" || !existingKey {
		key, err := wireguard.NewPrivateKey()
		if err != nil {
			return err
		}
		viper.Set(config.PrivateKey, key.String())
	}
	return nil
}

func checkTOS() (bool, error) {
	if !acceptedTOS {
		fmt.Println("This project is in no way affiliated with Cloudflare")
		fmt.Println("Cloudflare's Terms of Service:", url.TermsUrl)
		prompt := promptui.Select{
			Label: "Do you agree?",
			Items: []string{"Yes", "No"},
		}
		if _, result, err := prompt.Run(); err != nil || result != "Yes" {
			return false, err
		}
	}
	return true, nil
}
