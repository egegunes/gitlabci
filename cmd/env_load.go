package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	envCmd.AddCommand(envLoadCmd)
}

var envLoadCmd = &cobra.Command{
	Use:   "load [project] [file]",
	Short: "Loads CI variables to Gitlab",
	Long:  "Loads project level CI variables to Gitlab. Works with both project ID and NAMESPACE/PROJECTNAME.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		inputFile := args[1]

		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't read input file: %v\n", err)
			os.Exit(1)
		}

		var variables []gitlab.ProjectVariable
		if err = json.Unmarshal(content, &variables); err != nil {
			fmt.Fprintf(os.Stderr, "couldn't load file contents: %v\n", err)
			os.Exit(1)
		}

		for _, variable := range variables {
			fmt.Fprintf(os.Stderr, "updating %s... ", variable.Key)
			_, _, err := git.ProjectVariables.UpdateVariable(
				pid,
				variable.Key,
				&gitlab.UpdateProjectVariableOptions{
					Value:            &variable.Value,
					Masked:           &variable.Masked,
					Protected:        &variable.Protected,
					EnvironmentScope: &variable.EnvironmentScope,
				},
				nil,
			)

			if err != nil {
				fmt.Fprintf(os.Stderr, "\n%s not found, creating... ", variable.Key)
				_, _, err := git.ProjectVariables.CreateVariable(
					pid,
					&gitlab.CreateProjectVariableOptions{
						Key:              &variable.Key,
						Value:            &variable.Value,
						Masked:           &variable.Masked,
						Protected:        &variable.Protected,
						EnvironmentScope: &variable.EnvironmentScope,
					},
					nil,
				)

				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", RED("error"))
					fmt.Fprintf(os.Stderr, "couldn't create variable %s: %v\n", variable.Key, err)
					os.Exit(1)
				}
			}

			fmt.Fprintf(os.Stderr, "%s\n", GREEN("done"))
		}

		os.Exit(0)
	},
}
