package status

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"wgcf/cloudflare"
	. "wgcf/cmd/shared"
	"wgcf/util"
)

var shortMsg = "Prints the status of the current Cloudflare Warp device"

var Cmd = &cobra.Command{
	Use:   "status",
	Short: shortMsg,
	Long:  FormatMessage(shortMsg, ``),
	Run: func(cmd *cobra.Command, args []string) {
		if err := status(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
}

func status() error {
	if !IsConfigValidAccount() {
		return errors.New("no valid account detected")
	}

	ctx := CreateContext()
	thisDevice, err := cloudflare.GetThisDevice(ctx)
	if err != nil {
		return err
	}
	boundDevice, err := cloudflare.GetThisBoundDevice(ctx)
	if err != nil {
		return err
	}

	PrintDeviceData(thisDevice, boundDevice)
	return nil
}
