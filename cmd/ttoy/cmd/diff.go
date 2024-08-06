package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/zrcoder/ttoy/pkg/diff"
	"github.com/zrcoder/ttoy/pkg/util"
)

var diffInline bool

var diffCmd = &cobra.Command{
	Use:     "diff",
	GroupID: "main",
	Short:   "show differences between two files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			util.ShowFatal(errors.New("need source file and dest file"))
		}
		diff.Show(args[0], args[1], diffInline)
	},
}

func init() {
	diffCmd.Flags().BoolVar(&diffInline, "inline", false, "if show diff inline")
	diffCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Printf("Usage:\n  %s\n\n", cmd.UseLine())
		fmt.Println(cmd.Short)
		fmt.Println("\nFlags:")
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			switch f.Name {
			case "i", "in", "in-fmt", "o", "out", "out-fmt":
			default:
				fmt.Printf("  --%s\t%s\n", f.Name, f.Usage)
			}
		})
	})
	rootCmd.AddCommand(diffCmd)
}
