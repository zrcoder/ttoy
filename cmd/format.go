package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zrcoder/ttoy/formatter"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:     "fmt",
	Aliases: []string{"format"},
	GroupID: "main",
	Short:   "format json, yaml, toml and xml",
	Run: func(cmd *cobra.Command, args []string) {
		formatters := map[string]Action{
			"json": formatter.Json,
			"yaml": formatter.Yaml,
			"yml":  formatter.Yaml,
			"toml": formatter.Toml,
			"xml":  formatter.Html,
			"html": formatter.Html,
		}
		action := formatters[InputFormat]
		do(action)
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
