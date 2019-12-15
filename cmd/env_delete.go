package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	envCmd.AddCommand(envDeleteCmd)
}

var envDeleteCmd = &cobra.Command{
	Use:   "delete [project] [key]",
	Short: "delete CI variable",
	Long:  "delete CI variable. Works with both project id and NAMESPACE/PROJECTNAME",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		key := args[1]

		_, err := git.ProjectVariables.RemoveVariable(pid, key, nil)
		fmt.Fprintf(os.Stderr, "deleting %s...", key)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", RED("error"))
			fmt.Fprintf(os.Stderr, "couldn't delete variable: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "%s\n", GREEN("done"))

		os.Exit(0)
	},
}
