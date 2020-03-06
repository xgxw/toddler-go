package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xgxw/foundation-go/database"
	"github.com/xgxw/foundation-go/log"
)

var cfgFile = ""

var rootCmd = &cobra.Command{
	Use:   "toddler",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".toddler-go.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

type (
	IpdbOptions struct {
		Path string `yaml:"path" mapstructure:"path"`
	}
	Options struct {
		Logging log.Options      `mapstructure:"logger"`
		DB      database.Options `mapstructure:"db" yaml:"db"`
		Demo    DemoOption       `mapstructure:"demo" yaml:"demo"`
		Server  ServerOption     `mapstructure:"server" yaml:"server"`
	}
)

func loadOptions() (*Options, error) {
	o := new(Options)
	err := viper.Unmarshal(o)
	return o, err
}
