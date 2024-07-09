package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zrcoder/ttoy/encoder"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:     "dec",
	Aliases: []string{"decode"},
	GroupID: "code",
	Short:   "decode url, html or qrcode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return encodeOrDecode(false)
	},
}

var decqrCmd = &cobra.Command{
	Use:     "qr",
	Aliases: []string{"qrcode"},
	Short:   "decode qrcode",
	RunE: func(cmd *cobra.Command, args []string) error {
		return encoder.DecodeQr(Input)
	},
}

func init() {
	decodeCmd.Flags().StringVarP(&url, "url", "u", "", "url")
	decodeCmd.AddCommand(decqrCmd)
	rootCmd.AddCommand(decodeCmd)
}
