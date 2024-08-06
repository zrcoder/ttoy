package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zrcoder/ttoy/pkg/formatter"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:     "fmt",
	Aliases: []string{"format"},
	GroupID: "main",
	Short:   "format json, yaml, toml and xml",
	Run: func(cmd *cobra.Command, args []string) {
		formatters := map[string]Action{
			"json": formatter.JsonCli,
			"yaml": formatter.YamlCli,
			"yml":  formatter.YamlCli,
			"toml": formatter.TomlCli,
			"xml":  formatter.HtmlCli,
			"html": formatter.HtmlCli,
		}
		action := formatters[InputFormat]
		do(action)
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
