package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	envCmd.AddCommand(envAddCmd)
}

var envAddCmd = &cobra.Command{
	Use:   "add [project] [key] [value]",
	Short: "add CI variable",
	Long:  "add CI variable. Works with both project id and NAMESPACE/PROJECTNAME",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		key := args[1]
		value := args[2]
		masked := false
		protected := false
		scope := "*"

		variableOptions := &gitlab.CreateProjectVariableOptions{
			Key:              &key,
			Value:            &value,
			Masked:           &masked,
			Protected:        &protected,
			EnvironmentScope: &scope,
		}

		variable, _, err := git.ProjectVariables.CreateVariable(pid, variableOptions, nil)
		fmt.Fprintf(os.Stderr, "creating %s...", variable.Key)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", RED("error"))
			fmt.Fprintf(os.Stderr, "couldn't add variable: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "%s\n", GREEN("done"))

		os.Exit(0)
	},
}
