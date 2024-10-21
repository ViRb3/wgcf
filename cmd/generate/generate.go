package generate

import (
	"log"

	"github.com/ViRb3/wgcf/v2/cloudflare"
	. "github.com/ViRb3/wgcf/v2/cmd/shared"
	"github.com/ViRb3/wgcf/v2/config"
	"github.com/ViRb3/wgcf/v2/util"
	"github.com/ViRb3/wgcf/v2/wireguard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var profileFile string
var shortMsg = "Generates a WireGuard profile from the current Cloudflare Warp account"

var Cmd = &cobra.Command{
	Use:   "generate",
	Short: shortMsg,
	Long:  FormatMessage(shortMsg, ``),
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
	if !IsConfigValidAccount() {
		return errors.New("no account detected")
	}

	ctx := CreateContext()
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return err
	}
	boundDevice, err := cloudflare.GetSourceBoundDevice(ctx)
	if err != nil {
		return err
	}

	profile, err := wireguard.NewProfile(&wireguard.ProfileData{
		PrivateKey: viper.GetString(config.PrivateKey),
		Address1:   thisDevice.Config.Interface.Addresses.V4,
		Address2:   thisDevice.Config.Interface.Addresses.V6,
		PublicKey:  thisDevice.Config.Peers[0].PublicKey,
		Endpoint:   thisDevice.Config.Peers[0].Endpoint.Host,
	})
	if err != nil {
		return err
	}
	if err := profile.Save(profileFile); err != nil {
		return err
	}

	PrintDeviceData(thisDevice, boundDevice)
	log.Println("Successfully generated WireGuard profile:", profileFile)
	return nil
}
