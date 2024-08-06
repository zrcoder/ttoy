package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zrcoder/ttoy/pkg/converter"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:     "conv",
	Aliases: []string{"convert"},
	GroupID: "main",
	Short:   "convert between json, yaml and toml",
	Run: func(cmd *cobra.Command, args []string) {
		converters := map[string]Action{
			"json-yaml": converter.Json2YamlCli,
			"json-toml": converter.Json2TomlCli,
			"yaml-json": converter.Yaml2JsonCli,
			"yaml-toml": converter.Yaml2TomlCli,
			"toml-json": converter.Toml2JsonCli,
			"toml-yaml": converter.Toml2YamlCli,
		}
		key := InputFormat + "-" + OutputFormat
		action := converters[key]
		do(action)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
