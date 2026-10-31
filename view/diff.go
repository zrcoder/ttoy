package view

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
	"github.com/zrcoder/ttoy/service"
)

// diffTab pairs two editable buffers with a diff viewer over them. The
// diff viewer's old side is read-only, so Edit shows both buffers as
// editors and Diff compares them with character-level highlights.
type diffTab struct {
	diff     *editor.DiffState
	oldEd    *editor.State
	newEd    *editor.State
	mode     int // 0: Edit, 1: Diff
	lastMode int
	svc      *service.Sort
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
		oldEd: &editor.State{Buffer: oldBuf},
		newEd: &editor.State{Buffer: newBuf},
	}
	if sortable {
		t.svc = &service.Sort{}
	}
	return t
}

func (t *diffTab) View(c *ui.Context) {
	// Entering Diff formats both texts so that key order does not diff.
	var sortErr string
	if t.svc != nil && t.mode == 1 && t.lastMode == 0 {
		if err := t.sort(); err != nil {
			t.mode = 0
			sortErr = err.Error()
		}
	}
	t.lastMode = t.mode

	added, deleted := t.diff.Stats()
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).AlignItems(ui.Center).Children(func() {
			ui.Segmented(c, &t.mode, "Edit", "Diff").Margin(18, 18, 18, 0)
			ui.Text(c, fmt.Sprintf("+%d −%d", added, deleted)).FontSize(12).TextColor(c.Theme().TextMuted)
		})
		// ui.Spacer(c).Height(8)
		if t.mode == 0 {
			t.editView(c)
			return
		}
		editor.DiffView(c, t.diff).Grow(1).FillWidth()
	})
	if sortErr != "" {
		Toast(c, sortErr, func() {})
	}
}

func (t *diffTab) editView(c *ui.Context) {
	ui.Row(c).Grow(1).FillWidth().Children(func() {
		ui.Column(c.Key("diff-original")).Fill().Children(func() {
			editor.View(c, t.oldEd).Fill()
		})
		ui.Column(c.Key("diff-modified")).Fill().Children(func() {
			editor.View(c, t.newEd).Fill()
		})
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
		ui.Tabs(c, &p.activeTab, "Text", "JSON").Margin(18)
		ui.Spacer(c).Height(8)
		if p.activeTab == 0 {
			p.text.View(c)
			return
		}
		p.json.View(c)
	})
}
