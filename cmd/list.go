package cmd

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

var Group bool
var Status string
var IncludeJobs bool
var Number int

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&Status, "status", "s", "", "pipeline status to filter")
	listCmd.Flags().BoolVarP(&IncludeJobs, "jobs", "j", false, "include jobs to output")
	listCmd.Flags().BoolVarP(&Group, "group", "g", false, "get all pipelines in group")
	listCmd.Flags().IntVarP(&Number, "number", "n", 1, "number of pipelines to list")
}

type Pipeline struct {
	*gitlab.PipelineInfo
	Project string
}

func (p Pipeline) String() string {
	status := p.Status
	if status == "success" {
		status = GREEN(status)
	} else if status == "failed" {
		status = RED(status)
	}

	createdAt := p.CreatedAt.Format(time.RFC822Z)

	return fmt.Sprintf("%-40s %-30s %-10d %22s %-9s", p.Project, p.Ref, p.ID, createdAt, status)
}

type Job struct {
	*gitlab.Job
}

func (j Job) String() string {
	status := j.Status
	if status == "success" {
		status = GREEN(status)
	} else if status == "failed" {
		status = RED(status)
	}

	return fmt.Sprintf("%-12d %-10s %-20s %-9s %6.2f seconds", j.ID, j.Stage, j.Name, status, j.Duration)
}

var listCmd = &cobra.Command{
	Use:   "list [project|group]",
	Short: "List pipelines",
	Long: `List pipelines in a project or all projects in a group.
If you want to list for a group, you have to use -g flag.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		var allPipelines []Pipeline
		var projects []string

		if Group {
			groupProjects, _, err := git.Groups.ListGroupProjects(args[0], nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "couldn't get group projects: %v\n", err)
				os.Exit(1)
			}
			for _, project := range groupProjects {
				projects = append(projects, project.PathWithNamespace)
			}
		} else {
			projects = append(projects, args[0])
		}

		for _, pid := range projects {
			opts := &gitlab.ListProjectPipelinesOptions{OrderBy: gitlab.String("id")}
			opts.PerPage = Number
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
				allPipelines = append(allPipelines, Pipeline{pipeline, pid})
			}
		}

		sort.Slice(allPipelines, func(i, j int) bool {
			return allPipelines[i].ID > allPipelines[j].ID
		})

		for _, pipeline := range allPipelines {
			fmt.Fprintf(os.Stdout, "%s\n", pipeline)
			if IncludeJobs {
				jobs, _, err := git.Jobs.ListPipelineJobs(pipeline.Project, pipeline.ID, nil)
				if err != nil {
					fmt.Fprintf(os.Stderr, "couldn't get pipeline jobs: %v\n", err)
					os.Exit(1)
				}
				for _, job := range jobs {
					fmt.Fprintf(os.Stdout, "\t%s\n", Job{job})
				}
			}
		}

		os.Exit(0)
	},
}
