package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	envCmd.AddCommand(envDumpCmd)
}

var envDumpCmd = &cobra.Command{
	Use:   "dump [project]",
	Short: "Dumps CI variables to stdout",
	Long:  "Dumps CI variables to stdout. Works with both project ID and NAMESPACE/PROJECTNAME.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		opts := &gitlab.ListProjectVariablesOptions{PerPage: 100, Page: 1}
		variables, _, err := git.ProjectVariables.ListVariables(pid, opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't get project variables: %v\n", err)
			os.Exit(1)
		}

		dump, err := json.MarshalIndent(variables, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't dump project variables: %v", err)
			os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, string(dump))

		os.Exit(0)
	},
}
