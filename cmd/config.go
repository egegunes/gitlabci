package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".gitlab")
	viper.AddConfigPath("$HOME/.config/gitlabenv")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "couldn't read config file: %v", err)
		os.Exit(1)
	}
}
