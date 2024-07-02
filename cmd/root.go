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

	rootCmd.PersistentFlags().StringVarP(&util.InputFile, "in", "i", "", "input file")
	rootCmd.PersistentFlags().StringVarP(&util.OutputFile, "out", "o", "", "output file")
	rootCmd.PersistentFlags().StringVarP(&util.InputFormat, "in-fmt", "", "", "input format")
	rootCmd.PersistentFlags().StringVarP(&util.OutputFormat, "out-fmt", "", "", "output format")
	rootCmd.AddGroup(&cobra.Group{ID: "main", Title: "Converter/Fomatter/Generator:"})
	rootCmd.AddGroup(&cobra.Group{ID: "code", Title: "Encoder/Decoder:"})

	cobra.OnInitialize(util.Init)
}
