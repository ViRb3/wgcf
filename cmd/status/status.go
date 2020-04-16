package status

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"wgcf/cloudflare/api"
	. "wgcf/cmd/util"
	"wgcf/util"
)

var shortMsg = "Prints the status of the current Cloudflare Warp device"

var Cmd = &cobra.Command{
	Use:   "status",
	Short: shortMsg,
	Long:  FormatMessage(shortMsg, ""),
	Run: func(cmd *cobra.Command, args []string) {
		if err := status(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
}

func status() error {
	if !ValidAccount() {
		return errors.New("no valid account detected")
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

	PrintDeviceData(deviceData, boundDevice)
	return nil
}
