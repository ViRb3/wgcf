package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"wgcf/cmd/generate"
	"wgcf/cmd/register"
	"wgcf/cmd/status"
	"wgcf/cmd/trace"
	"wgcf/cmd/update"
	"wgcf/cmd/util"
	"wgcf/config"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "wgcf",
	Short: "WireGuard Cloudflare Warp utility",
	Long: util.FormatMessage("",`
wgcf is a utility for Cloudflare Warp that allows you to create and
manage accounts, assign license keys, and generate WireGuard profiles.
Made by Victor (@ViRb3). Project website: https://github.com/ViRb3/wgcf`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatal(err)
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

func initConfig() {
	initConfigDefaults()
	viper.SetConfigFile(cfgFile)
	viper.SetEnvPrefix("WGCF")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initConfigDefaults() {
	viper.SetDefault(config.DeviceId, "")
	viper.SetDefault(config.AccessToken, "")
	viper.SetDefault(config.PrivateKey, "")
	viper.SetDefault(config.LicenseKey, "")
}
