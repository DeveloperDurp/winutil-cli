package app

import (
	"github.com/spf13/cobra"
)

var Appcmd = &cobra.Command{
	Use:   "app",
	Short: "Install/uninstall Applications",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
