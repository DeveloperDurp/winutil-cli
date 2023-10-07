package cmd

import (
	"os"

	"winutil-cli/cmd/app"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Winutil CLI",
	Short: "CLI Utility for WinUtil",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(app.Appcmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
