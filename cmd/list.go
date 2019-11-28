package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

var Status string
var IncludeJobs bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&Status, "status", "s", "", "pipeline status to filter")
	listCmd.Flags().BoolVarP(&IncludeJobs, "jobs", "j", false, "include jobs to output")
}

var listCmd = &cobra.Command{
	Use:   "list [project]",
	Short: "List pipelines in project",
	Long:  "List pipelines in project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]

		opts := &gitlab.ListProjectPipelinesOptions{OrderBy: gitlab.String("id")}
		if Status != "" {
			status := gitlab.BuildState(gitlab.BuildStateValue(Status))
			opts.Status = status
		}
		pipelines, _, err := git.Pipelines.ListProjectPipelines(pid, opts)

		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't get project pipelines: %v\n", err)
			os.Exit(1)
		}

		for _, pipeline := range pipelines {
			fmt.Fprintf(os.Stdout, "%d %s %s\n", pipeline.ID, pipeline.Status, pipeline.Ref)
			if IncludeJobs {
				jobs, _, err := git.Jobs.ListPipelineJobs(pid, pipeline.ID, nil)
				if err != nil {
					fmt.Fprintf(os.Stderr, "couldn't get pipeline jobs: %v\n", err)
					os.Exit(1)
				}
				for _, job := range jobs {
					fmt.Fprintf(os.Stdout, "    %d %s %s %f %s\n",
						job.ID, job.Name, job.Stage, job.Duration, job.Status,
					)
				}
			}
		}

		os.Exit(0)
	},
}
