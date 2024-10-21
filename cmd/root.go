package cmd

import (
	"errors"
	"log"

	"github.com/ViRb3/wgcf/v2/cmd/generate"
	"github.com/ViRb3/wgcf/v2/cmd/register"
	. "github.com/ViRb3/wgcf/v2/cmd/shared"
	"github.com/ViRb3/wgcf/v2/cmd/status"
	"github.com/ViRb3/wgcf/v2/cmd/trace"
	"github.com/ViRb3/wgcf/v2/cmd/update"
	"github.com/ViRb3/wgcf/v2/config"
	"github.com/ViRb3/wgcf/v2/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "wgcf",
	Short: "WireGuard Cloudflare Warp utility",
	Long: FormatMessage("", `
wgcf is a utility for Cloudflare Warp that allows you to create and
manage accounts, assign license keys, and generate WireGuard profiles.
Made by Victor (@ViRb3). Project website: https://github.com/ViRb3/wgcf`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "wgcf-account.toml", "Configuration file")
	RootCmd.AddCommand(register.Cmd)
	RootCmd.AddCommand(update.Cmd)
	RootCmd.AddCommand(generate.Cmd)
	RootCmd.AddCommand(status.Cmd)
	RootCmd.AddCommand(trace.Cmd)
}

var unsupportedConfigError viper.UnsupportedConfigError

func initConfig() {
	initConfigDefaults()
	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix("WGCF")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); errors.As(err, &unsupportedConfigError) {
		log.Fatal(err)
	} else {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initConfigDefaults() {
	viper.SetDefault(config.DeviceId, "")
	viper.SetDefault(config.AccessToken, "")
	viper.SetDefault(config.PrivateKey, "")
	viper.SetDefault(config.LicenseKey, "")
}
