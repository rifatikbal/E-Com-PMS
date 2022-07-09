package cmd

import (
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
	_ "github.com/spf13/viper/remote"
)

const (
	remoteConfigProvider = "consul"
	envPrefix            = "pms"
	remoteConfigEndpoint = "consul_url"
	remoteConfigPath     = "consul_path"
	remoteConfigType     = "yaml"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pms",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	endpoint := viper.GetString(remoteConfigEndpoint)
	path := viper.GetString(remoteConfigPath)

	cobra.CheckErr(viper.AddRemoteProvider(remoteConfigProvider, endpoint, path))
	viper.SetConfigType(remoteConfigType)
	cobra.CheckErr(viper.ReadRemoteConfig())
}
