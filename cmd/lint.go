package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	rootCmd.AddCommand(lintCmd)
}

var lintCmd = &cobra.Command{
	Use:   "lint [.gitlab-ci.yml]",
	Short: "Lint .gitlab-ci.yml",
	Long:  "Lint .gitlab-ci.yml",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		yamlContent, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't read %s: %v\n", args[0], err)
			os.Exit(1)
		}

		jsonContent, err := yaml.YAMLToJSON(yamlContent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't convert %s to json: %v\n", args[0], err)
			os.Exit(1)
		}

		content := string(jsonContent)

		result, _, err := git.Validate.Lint(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't validate %s: %v\n", args[0], err)
			os.Exit(1)
		}

		if result.Status == "valid" {
			fmt.Fprintf(os.Stdout, "%s: %s is valid\n", GREEN("OK"), args[0])
			os.Exit(0)
		}

		for _, e := range result.Errors {
			fmt.Fprintf(os.Stdout, "%s: %s\n", RED("ERROR"), e)
		}

		os.Exit(1)
	},
}
