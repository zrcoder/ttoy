package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zrcoder/ttoy/generator"
	"github.com/zrcoder/ttoy/util"
)

var generateCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"generate"},
	GroupID: "main",
	Short:   "generate hash, uuid, svg graph or go struct from json",
}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "generate hash",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generator.Hash(util.Input)
	},
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "generate uuid",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generator.Uuid(util.Input)
	},
}

var svgCmd = &cobra.Command{
	Use:   "svg",
	Short: "generate svg graph for json",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generator.Json2Svg(util.Input)
	},
}

var structCmd = &cobra.Command{
	Use:   "struct",
	Short: "generate go struct for json",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generator.Json2Struct(util.Input)
	},
}

var d2Cmd = &cobra.Command{
	Use:   "d2",
	Short: "generate d2 svg graph from d2 scripts",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generator.D2(util.Input)
	},
}

var ndorCmd = &cobra.Command{
	Use:   "ndor",
	Short: "generate ndor png from source code",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generator.Ndor(util.Input)
	},
}

func init() {
	generateCmd.AddCommand(hashCmd)
	generateCmd.AddCommand(uuidCmd)
	generateCmd.AddCommand(svgCmd)
	generateCmd.AddCommand(structCmd)
	generateCmd.AddCommand(d2Cmd)
	generateCmd.AddCommand(ndorCmd)
	rootCmd.AddCommand(generateCmd)
}
