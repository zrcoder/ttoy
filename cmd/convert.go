package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zrcoder/ttoy/converter"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:     "conv",
	Aliases: []string{"convert"},
	GroupID: "main",
	Short:   "convert between json, yaml and toml",
	Run: func(cmd *cobra.Command, args []string) {
		converters := map[string]Action{
			"json-yaml": converter.Json2Yaml,
			"json-toml": converter.Json2Toml,
			"yaml-json": converter.Yaml2Json,
			"yaml-toml": converter.Yaml2Toml,
			"toml-json": converter.Toml2Json,
			"toml-yaml": converter.Toml2Yaml,
		}
		key := InputFormat + "-" + OutputFormat
		action := converters[key]
		do(action)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
