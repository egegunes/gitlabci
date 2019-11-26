package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
