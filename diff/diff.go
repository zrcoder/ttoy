package diff

import (
	"fmt"
	"os"
	"strings"

	"github.com/zrcoder/ttoy/util"

	lg "github.com/charmbracelet/lipgloss"
	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

var (
	insertSty = lg.NewStyle().Background(util.Green)
	deleteSty = lg.NewStyle().Background(util.Red)
	blankSty  = lg.NewStyle().Foreground(util.Faint).Strikethrough(true)
	textSty   = lg.NewStyle().Width(60).Border(lg.NormalBorder())
)

func Show(srcFile, dstFile string, inline bool) {
	src, err := os.ReadFile(srcFile)
	if err != nil {
		util.ShowFatal(err)
	}
	dst, err := os.ReadFile(dstFile)
	if err != nil {
		util.ShowFatal(err)
	}

	differ := dmp.New()
	diffs := differ.DiffMain(string(src), string(dst), false)

	if inline {
		fmt.Println(differ.DiffPrettyText(diffs))
		return
	}

	srcView, dstView := diffView(diffs)
	res := lg.JoinVertical(lg.Center,
		srcFile+" -> "+dstFile,
		"",
		lg.JoinHorizontal(lg.Top, srcView, dstView),
	)
	fmt.Println(res)
}

func diffView(diffs []dmp.Diff) (string, string) {
	var src, dst strings.Builder
	for _, diff := range diffs {
		switch diff.Type {
		case dmp.DiffEqual:
			src.WriteString(diff.Text)
			dst.WriteString(diff.Text)
		case dmp.DiffDelete:
			src.WriteString(deleteSty.Render(diff.Text))
			dst.WriteString(blankSty.Render(diff.Text))
		case dmp.DiffInsert:
			src.WriteString(blankSty.Render(diff.Text))
			dst.WriteString(insertSty.Render(diff.Text))
		}
	}
	return textSty.Render(src.String()), textSty.Render(dst.String())
}
