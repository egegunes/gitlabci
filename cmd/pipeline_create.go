package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

var env []string

func init() {
	pipelineCmd.AddCommand(createCmd)
	createCmd.Flags().StringSliceVarP(&env, "variable", "e", []string{}, "Override environment variables. KEY=VALUE")
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

		var variables []*gitlab.PipelineVariable

		for _, variable := range env {
			v := strings.Split(variable, "=")

			variables = append(variables, &gitlab.PipelineVariable{
				Key:          v[0],
				Value:        v[1],
				VariableType: "env_var",
			})
		}

		opts := &gitlab.CreatePipelineOptions{Ref: gitlab.String(ref), Variables: variables}
		pipeline, _, err := git.Pipelines.CreatePipeline(pid, opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't create pipeline: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Pipeline %d is running for %s\n", pipeline.ID, pipeline.Ref)
		os.Exit(0)
	},
}
