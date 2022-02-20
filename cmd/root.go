package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	token   string
	addr    string
	rootCmd = &cobra.Command{
		Use:   "sutol",
		Short: "A simple CLI to test deal proposals to Filecoin miner.",
		Long: `Sutol is a CLI library that can be used to send self-forged proposals to your Filecoin miner (WIP) and replay already sent proposals.
	`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is golang:os.UseConfigDir()/sutol/.sutol.yaml, i.e in Linux $HOME/.config/sutol/.sutol.yaml)")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Token to connect to lotus-daemon (default to env var SUTOL_TOKEN else value in conf file else empty string)")
	rootCmd.PersistentFlags().StringVarP(&addr, "addr", "u", "", "Socket address (IP/hostname:port) of the lotus-daemon to connect to (default to env var SUTOL_ADDR else value in conf file else 'localhost:1234')")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
	viper.SetDefault("token", "")
	viper.SetDefault("addr", "localhost:1234")
}

func initConfig() {
	viper.SetEnvPrefix("sutol")
	viper.BindEnv("token")
	viper.BindEnv("addr")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := os.UserConfigDir()
		configDir = configDir + "/sutol"
		cobra.CheckErr(err)

		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sutol")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	token = viper.GetString("token")
	addr = viper.GetString("addr")

}
