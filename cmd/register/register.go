package register

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"wgcf/cloudflare"
	. "wgcf/cmd/shared"
	"wgcf/config"
	"wgcf/util"
	"wgcf/wireguard"
)

var deviceName string
var deviceModel string
var acceptedTOS = false
var shortMsg = "Registers a new Cloudflare Warp device and creates a new account, preparing it for connection"

var Cmd = &cobra.Command{
	Use:   "register",
	Short: shortMsg,
	Long:  FormatMessage(shortMsg, ``),
	Run: func(cmd *cobra.Command, args []string) {
		if err := registerAccount(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Device name displayed under the 1.1.1.1 app (defaults to random)")
	Cmd.PersistentFlags().StringVarP(&deviceModel, "model", "m", "PC", "Device model displayed under the 1.1.1.1 app")
	Cmd.PersistentFlags().BoolVar(&acceptedTOS, "accept-tos", false, "Accept Cloudflare's Terms of Service non-interactively")
}

func registerAccount() error {
	if IsConfigValidAccount() {
		return errors.New("existing account detected")
	}
	if accepted, err := checkTOS(); err != nil || !accepted {
		return err
	}

	privateKey, err := wireguard.NewPrivateKey()
	if err != nil {
		return err
	}
	device, err := cloudflare.Register(privateKey.Public(), deviceModel)
	if err != nil {
		return err
	}

	viper.Set(config.PrivateKey, privateKey.String())
	viper.Set(config.DeviceId, device.Id)
	viper.Set(config.AccessToken, device.Token)
	viper.Set(config.LicenseKey, device.Account.License)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	ctx := CreateContext()
	_, err = SetDeviceName(ctx, deviceName)
	if err != nil {
		return err
	}
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return err
	}

	boundDevice, err := cloudflare.UpdateSourceBoundDeviceActive(ctx, true)
	if err != nil {
		return err
	}
	if !boundDevice.Active {
		return errors.New("failed to activate device")
	}

	PrintDeviceData(thisDevice, boundDevice)
	log.Println("Successfully created Cloudflare Warp account")
	return nil
}

func checkTOS() (bool, error) {
	if !acceptedTOS {
		fmt.Println("This project is in no way affiliated with Cloudflare")
		fmt.Println("Cloudflare's Terms of Service: https://www.cloudflare.com/application/terms/")
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
