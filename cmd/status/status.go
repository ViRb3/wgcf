package status

import (
	"log"

	"github.com/ViRb3/wgcf/v2/cloudflare"
	. "github.com/ViRb3/wgcf/v2/cmd/shared"
	"github.com/ViRb3/wgcf/v2/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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
	thisDevice, err := cloudflare.GetSourceDevice(ctx)
	if err != nil {
		return err
	}
	boundDevice, err := cloudflare.GetSourceBoundDevice(ctx)
	if err != nil {
		return err
	}

	PrintDeviceData(thisDevice, boundDevice)
	return nil
}
