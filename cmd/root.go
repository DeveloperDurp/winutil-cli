package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func setDefaults() {
}

func init() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	cobra.OnInitialize(initConfig)

	setDefaults()
	cfgFile := home + "/.winutil.yaml"
	err = viper.WriteConfigAs(cfgFile)
	if err != nil {
		fmt.Println(err)
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.winutil.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".winutil")

	viper.AutomaticEnv()
	viper.ReadInConfig()

}
