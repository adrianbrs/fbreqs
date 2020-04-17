package cmd

import (
	"fbreqs/client"
	"fbreqs/console"
	"fbreqs/sender"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia as requisições",
	Run: func(cmd *cobra.Command, args []string) {
		console.Header().SetLogger(console.PrimaryLogger).SetValues(
			strings.ToUpper(viper.GetString("app.name")),
			viper.GetString("app.version"),
		).Render()
		console.Separator()

		// Create new client and run request sender
		client := client.New()
		sender.Run(client)
	},
}
