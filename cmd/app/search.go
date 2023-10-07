package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"winutil-cli/objects"

	"github.com/spf13/cobra"
)

var app string

var searchcmd = &cobra.Command{
	Use:   "search",
	Short: "Generates a config file using current config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !strings.HasPrefix(app, "WPFInstall") {
			app = "WPFInstall" + app
		}
		url := "https://raw.githubusercontent.com/ChrisTitusTech/winutil/main/config/applications.json"

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer response.Body.Close()

		var data objects.Applications
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if installInfo, ok := data[app]; ok {
			fmt.Println("Key:", app)
			fmt.Println("Winget:", installInfo.Winget)
			fmt.Println("Choco:", installInfo.Choco)
		} else {
			fmt.Println("App not found:", app)
		}
	},
}

func init() {
	searchcmd.Flags().StringVarP(&app, "app", "a", "", "Application to find")
	Appcmd.AddCommand(searchcmd)
}
