package view

import (
	"strings"

	"github.com/ELY-Open-Source/CrychicUI/choice"
	"github.com/ELY-Open-Source/CrychicUI/code"
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
	"github.com/zrcoder/ttoy/service"
)

const (
	modeEdit = iota
	modeDiff
)

// diffTab pairs two editable buffers with a diff viewer over them. The
// diff viewer's old side is read-only, so Edit shows both buffers as
// editors and Diff compares them with character-level highlights.
type diffTab struct {
	diff       *editor.DiffState
	oldEd      *editor.State
	newEd      *editor.State
	mode       int // 0: Edit, 1: Diff
	lastMode   int
	svc        *service.Sort
	diffSplite bool
}

func newDiffTab(lang *editor.Language, sortable bool) *diffTab {
	oldBuf := editor.NewBuffer("")
	newBuf := editor.NewBuffer("")
	if lang != nil {
		oldBuf.SetLanguage(lang)
		newBuf.SetLanguage(lang)
	}
	t := &diffTab{
		diff: &editor.DiffState{
			Old:      oldBuf,
			New:      newBuf,
			Editable: true,
			Inline:   editor.DiffInlineChar,
		},
		oldEd:      &editor.State{Buffer: oldBuf},
		newEd:      &editor.State{Buffer: newBuf},
		mode:       modeEdit,
		lastMode:   modeEdit,
		diffSplite: true,
	}
	if sortable {
		t.svc = &service.Sort{}
	}
	return t
}

func (t *diffTab) View(c *ui.Context) {
	var sortErr string
	if t.svc != nil && t.mode == modeDiff && t.lastMode == modeEdit {
		if err := t.sort(); err != nil {
			t.mode = modeEdit
			sortErr = err.Error()
		}
	}
	if sortErr != "" {
		Toast(c, sortErr, func() {})
	}

	t.lastMode = t.mode

	ui.Column(c).Fill().Children(func() {
		choice.Segments(c, "Mode", []string{"Edit", "Diff"}, &t.mode).Margin(18, 0)
		if t.mode == modeEdit {
			t.editView(c)
			return
		}
		code.DiffView(c, t.diff, &t.diffSplite, code.DiffViewOptions{}).Fill().Margin(0, 0, 18, 0)
	})
}

func (t *diffTab) editView(c *ui.Context) {
	ui.Row(c).Fill().Margin(0, 0, 18, 0).Children(func() {
		code.CodeEditor(c, t.oldEd, code.CodeEditorOptions{}).Fill()
		ui.Spacer(c).Width(18)
		code.CodeEditor(c, t.newEd, code.CodeEditorOptions{}).Fill()
	})
}

// sort formats both texts as JSON so that key order does not diff.
func (t *diffTab) sort() error {
	bufs := []*editor.Buffer{t.diff.Old, t.diff.New}
	texts := make([]string, len(bufs))
	for i, b := range bufs {
		text := b.String()
		if strings.TrimSpace(text) == "" {
			continue
		}
		res, err := t.svc.JSON(text)
		if err != nil {
			return err
		}
		texts[i] = res
	}
	for i, b := range bufs {
		if texts[i] != "" && texts[i] != b.String() {
			b.Set(texts[i])
		}
	}
	return nil
}

type Diff struct {
	text      *diffTab
	json      *diffTab
	activeTab int
}

func NewDiff() *Diff {
	return &Diff{
		text: newDiffTab(nil, false),
		json: newDiffTab(editor.LanguageNamed("json"), true),
	}
}

func (p *Diff) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		tabs(c, &p.activeTab, "Text", "JSON")
		if p.activeTab == 0 {
			p.text.View(c)
			return
		}
		p.json.View(c)
	})
}
