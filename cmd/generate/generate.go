package generate

import (
	"log"

	"github.com/ViRb3/wgcf/v2/cloudflare"
	. "github.com/ViRb3/wgcf/v2/cmd/shared"
	"github.com/ViRb3/wgcf/v2/config"
	"github.com/ViRb3/wgcf/v2/wireguard"
	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var profileFile string
var keepalive int
var shortMsg = "Generates a WireGuard profile from the current Cloudflare Warp account"

var Cmd = &cobra.Command{
	Use:   "generate",
	Short: shortMsg,
	Long:  FormatMessage(shortMsg, ``),
	Run: func(cmd *cobra.Command, args []string) {
		RunCommandFatal(generateProfile)
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&profileFile, "profile", "p", "wgcf-profile.conf", "WireGuard profile file")
	Cmd.PersistentFlags().IntVarP(&keepalive, "keepalive", "k", 0, "Persistent keepalive interval in seconds (default 25)")
	Cmd.PersistentFlags().Lookup("keepalive").NoOptDefVal = "25"
}

func generateProfile() error {
	if err := EnsureConfigValidAccount(); err != nil {
		return errors.WithStack(err)
	}

	ctx := CreateContext()
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	if thisDevice.Config == nil || len(thisDevice.Config.Peers) == 0 {
		return errors.New("Cloudflare response did not contain a WireGuard peer")
	}
	peer := thisDevice.Config.Peers[0]
	endpoint, err := cloudflare.WireGuardEndpoint(peer.Endpoint)
	if err != nil {
		return errors.WithStack(err)
	}

	profile, err := wireguard.NewProfile(&wireguard.ProfileData{
		PrivateKey: viper.GetString(config.PrivateKey),
		Address1:   thisDevice.Config.Interface.Addresses.V4,
		Address2:   thisDevice.Config.Interface.Addresses.V6,
		PublicKey:  peer.PublicKey,
		Endpoint:   endpoint,
		Keepalive:  keepalive,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	if err := profile.Save(profileFile); err != nil {
		return errors.WithStack(err)
	}

	log.Println("Successfully generated WireGuard profile:", profileFile)
	return nil
}
