package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zrcoder/ttoy/util"
)

var rootCmd = &cobra.Command{
	Use:   "ttoy",
	Short: "Dev toys on the terminal",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		util.ShowFatal(err)
	}
}

func init() {
	// hide help flags and command
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.PersistentFlags().Lookup("help").Hidden = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	// hide completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().StringVarP(&util.InputFile, "input", "i", "", "input file")
	rootCmd.PersistentFlags().StringVarP(&util.OutputFile, "output", "o", "", "output file")
	rootCmd.PersistentFlags().StringVarP(&util.InputFormat, "input-format", "", "", "input format")
	rootCmd.PersistentFlags().StringVarP(&util.OutputFormat, "output-format", "", "", "output format")
	rootCmd.AddGroup(&cobra.Group{ID: "main", Title: ""})
	rootCmd.AddGroup(&cobra.Group{ID: "code", Title: ""})
	cobra.OnInitialize(util.Init)
}
