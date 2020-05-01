package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("token", "t", "SECRET", "Your Gitlab API Access Token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

func initConfig() {
	viper.SetConfigName(".gitlab")
	viper.AddConfigPath("$HOME/.config/gitlabci")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "couldn't read config file: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gitlabci [command]",
	Short: "gitlabci lets you manage and track Gitlab CI pipelines",
	Long:  "gitlabci lets you manage and track Gitlab CI pipelines",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
