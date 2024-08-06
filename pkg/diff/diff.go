package diff

import (
	"fmt"
	"os"
	"strings"

	"github.com/zrcoder/ttoy/pkg/util"

	lg "github.com/charmbracelet/lipgloss"
	dmp "github.com/sergi/go-diff/diffmatchpatch"
)

const textWidth = 66

var (
	insertSty = lg.NewStyle().Background(util.Green)
	deleteSty = lg.NewStyle().Background(util.Red)
	blankSty  = lg.NewStyle().Foreground(util.Faint)
	wideSty   = lg.NewStyle().Width(textWidth)
	textSty   = wideSty.Border(lg.NormalBorder())
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
		text := diff.Text
		if len(text) > textWidth {
			text = wideSty.Render(text)
		}
		switch diff.Type {
		case dmp.DiffEqual:
			src.WriteString(text)
			dst.WriteString(text)
		case dmp.DiffDelete:
			src.WriteString(deleteSty.Render(text))
			dst.WriteString(blankSty.Render(text))
		case dmp.DiffInsert:
			src.WriteString(blankSty.Render(text))
			dst.WriteString(insertSty.Render(text))
		}
	}
	return textSty.Render(src.String()), textSty.Render(dst.String())
}
