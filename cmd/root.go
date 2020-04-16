package cmd

import (
	"caixa-falso/console"
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "fmreqs",
	Short: "Envia várias requisições simultâneas para um determinado URL.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute executa o rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(fmt.Sprintf("Error in root cmd: %q\n", err))
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Flags //
	// =========================

	// Config
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "configuration file (default to \"./config.json\"")

	// Proxy
	rootCmd.PersistentFlags().StringSliceP("proxy", "p", []string{}, "proxy list to use in request (default to empty)")
	viper.BindPFlag("request.proxy.list", rootCmd.PersistentFlags().Lookup("proxy"))

	// Random proxy
	rootCmd.PersistentFlags().Bool("randomProxy", true, "use random proxy from proxy list (default to true)")
	viper.BindPFlag("request.proxy.random", rootCmd.PersistentFlags().Lookup("randomProxy"))

	// Shell Width
	rootCmd.PersistentFlags().Int("shellWidth", 100, "sets the shell width (default to 100)")
	viper.BindPFlag("app.shell.width", rootCmd.PersistentFlags().Lookup("shellWidth"))

	// Shell Left Offset
	rootCmd.PersistentFlags().Int("logOffset", 10, "sets the shell log left offset (default to 10)")
	viper.BindPFlag("app.shell.offset", rootCmd.PersistentFlags().Lookup("logOffset"))
}

func initConfig() {
	// Load config file
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Error loading config file")
	}

	// Init console config
	console.Init()
}
