package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pipelineCmd)
}

var pipelineCmd = &cobra.Command{
	Use:   "pipeline [subcommand]",
	Short: "Manage CI pipelines",
	Long:  "Manage CI pipelines",
}
