package generate

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"wgcf/cloudflare/api"
	. "wgcf/cmd/util"
	"wgcf/config"
	"wgcf/util"
	"wgcf/wireguard"
)

var profileFile string
var shortMsg = "Generates a WireGuard profile from the current Cloudflare Warp account"

var Cmd = &cobra.Command{
	Use:   "generate",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
Will change various account settings to ensure WireGuard connection will succeed.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := generateProfile(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&profileFile, "profile", "p", "wgcf-profile.conf", "WireGuard profile file")
}

func generateProfile() error {
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

	if err := EnsureWarpEnabled(ctx, deviceData); err != nil {
		return err
	}
	if err := EnsureDeviceActive(ctx, boundDevice); err != nil {
		return err
	}

	profile, err := wireguard.NewProfile(&wireguard.ProfileData{
		PrivateKey: viper.GetString(config.PrivateKey),
		Address1:   deviceData.Config.Interface.Addresses.V4,
		Address2:   deviceData.Config.Interface.Addresses.V6,
		PublicKey:  deviceData.Config.Peers[0].PublicKey,
		Endpoint:   deviceData.Config.Peers[0].Endpoint.Host,
	})
	if err != nil {
		return err
	}
	if err := profile.Save(profileFile); err != nil {
		return err
	}

	PrintDeviceData(deviceData, boundDevice)
	log.Println("Successfully generated WireGuard profile:", profileFile)
	return nil
}
