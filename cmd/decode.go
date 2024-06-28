/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:     "dec",
	Aliases: []string{"decode"},
	GroupID: "code",
	Short:   "decode url or html",
	RunE: func(cmd *cobra.Command, args []string) error {
		return encodeOrDecode(false)
	},
}

func init() {
	decodeCmd.Flags().StringVarP(&url, "url", "u", "", "url")
	rootCmd.AddCommand(decodeCmd)
}
