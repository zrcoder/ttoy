package cmd

import (
	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:     "enc",
	Aliases: []string{"encode"},
	GroupID: "code",
	Short:   "encode url or html",
	Run: func(cmd *cobra.Command, args []string) {
		encodeOrDecode(true)
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&url, "url", "u", "", "url")
	rootCmd.AddCommand(encodeCmd)
}
