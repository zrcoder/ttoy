package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zrcoder/ttoy/pkg/generator"
)

var generateCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"generate"},
	GroupID: "main",
	Short:   "generate hash, uuid, qrcode, svg graph or go struct from json",
}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "generate hash",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Hash(Input)
	},
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "generate uuid",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Uuid(Input)
	},
}

var svgCmd = &cobra.Command{
	Use:   "svg",
	Short: "generate svg graph for json",
	Long:  "generate svg graph for json, see https://github.com/zrcoder/cdor",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Json2SvgCli(Input)
	},
}

var structCmd = &cobra.Command{
	Use:   "struct",
	Short: "generate go struct for json",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Json2Struct(Input)
	},
}

var d2Cmd = &cobra.Command{
	Use:   "d2",
	Short: "generate d2 svg graph from d2 scripts",
	Long:  "generate d2 svg graph from d2 scripts, see https://d2lang.com",
	Run: func(cmd *cobra.Command, args []string) {
		generator.D2(Input)
	},
}

var ndorCmd = &cobra.Command{
	Use:   "ndor",
	Short: "generate ndor png from source code",
	Long:  "generate ndor png from source code, see https://github.com/zrcoder/ndor/wiki",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Ndor(Input)
	},
}

var genqrCmd = &cobra.Command{
	Use:     "qr",
	Aliases: []string{"qrcode"},
	Short:   "encode drcode",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Qrcode(Input)
	},
}

func init() {
	structCmd.Flags().StringVarP(&generator.StructOption.Name, "name", "n", "Ttoy", "struct name")
	structCmd.Flags().StringVarP(&generator.StructOption.Pkg, "pkg", "p", "ttoy", "pkg name")
	structCmd.Flags().StringArrayVarP(&generator.StructOption.Tags, "tags", "t", []string{"json"}, "tags to add")
	structCmd.Flags().BoolVarP(&generator.StructOption.ConvertFloats, "float", "f", true, "if convert floats")
	structCmd.Flags().BoolVarP(&generator.StructOption.SubStruct, "sub", "s", false, "if generate sub struct")
	generateCmd.AddCommand(hashCmd)
	generateCmd.AddCommand(uuidCmd)
	generateCmd.AddCommand(svgCmd)
	generateCmd.AddCommand(structCmd)
	generateCmd.AddCommand(d2Cmd)
	generateCmd.AddCommand(ndorCmd)
	generateCmd.AddCommand(genqrCmd)
	rootCmd.AddCommand(generateCmd)
}
