package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	pipelineCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [project] [ref]",
	Short: "Create pipeline for project",
	Long:  "Create pipeline for project",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		ref := args[1]

		opts := &gitlab.CreatePipelineOptions{Ref: gitlab.String(ref)}
		pipeline, _, err := git.Pipelines.CreatePipeline(pid, opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't create pipeline: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Pipeline %d is running for %s\n", pipeline.ID, pipeline.Ref)
		os.Exit(0)
	},
}
