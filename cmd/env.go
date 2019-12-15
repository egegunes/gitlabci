package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(envCmd)
}

var envCmd = &cobra.Command{
	Use:   "env [subcommand]",
	Short: "Manage CI environment variables",
	Long:  "Manage CI environment variables",
}
