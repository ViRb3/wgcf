package generate

import (
	"log"
	"net/netip"
	"strings"

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
var dnsServers []string
var noDNS bool
var shortMsg = "Generates a WireGuard profile from the current Cloudflare Warp account"

const maxKeepalive = 65535

var defaultDNS = []string{
	"1.1.1.1",
	"1.0.0.1",
	"2606:4700:4700::1111",
	"2606:4700:4700::1001",
}

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
	Cmd.PersistentFlags().StringSliceVarP(&dnsServers, "dns", "d", defaultDNS, "DNS server IP addresses (comma-separated or repeatable)")
	Cmd.PersistentFlags().BoolVar(&noDNS, "no-dns", false, "Omit the DNS field from the generated profile")
	Cmd.MarkFlagsMutuallyExclusive("dns", "no-dns")
}

func generateProfile() error {
	if err := validateDNS(dnsServers); err != nil {
		return err
	}

	if err := validateKeepalive(keepalive); err != nil {
		return err
	}

	if err := EnsureConfigValidAccount(); err != nil {
		return errors.WithStack(err)
	}

	var profileDNS string
	if !noDNS {
		profileDNS = strings.Join(dnsServers, ", ")
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
		DNS:        profileDNS,
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

func validateKeepalive(interval int) error {
	if interval < 0 || interval > maxKeepalive {
		return errors.Errorf("persistent keepalive interval must be between 0 and %d seconds, got %d", maxKeepalive, interval)
	}
	return nil
}

func validateDNS(servers []string) error {
	if len(servers) == 0 {
		return errors.Errorf("at least one DNS server must be specified")
	}
	for _, server := range servers {
		addr, err := netip.ParseAddr(server)
		if err != nil || addr.Zone() != "" {
			return errors.Errorf("invalid DNS server IP address: %q", server)
		}
	}
	return nil
}
