package view

import (
	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/code"
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
)

// DualEditor renders two editors side by side with transform buttons between
// them. leftFn/rightFn are the service callbacks; when non-nil, clicking the
// corresponding arrow button transforms the opposite editor's content.
type DualEditor struct {
	leftBuf       *editor.State
	rightBuf      *editor.State
	leftLang      string
	rightLang     string
	leftReadOnly  bool
	rightReadOnly bool
	leftFn        func(string) (string, error)
	rightFn       func(string) (string, error)
}

func NewDualEditor(leftLang, rightLang string) *DualEditor {
	d := &DualEditor{
		leftBuf:       &editor.State{Buffer: editor.NewBuffer("")},
		rightBuf:      &editor.State{Buffer: editor.NewBuffer("")},
		leftLang:      leftLang,
		rightLang:     rightLang,
		leftReadOnly:  false,
		rightReadOnly: true,
	}
	d.leftBuf.Buffer.SetLanguage(editor.LanguageNamed(d.leftLang))
	d.rightBuf.Buffer.SetLanguage(editor.LanguageNamed(d.rightLang))
	return d
}

func (d *DualEditor) LeftReadOnly(v bool) *DualEditor                     { d.leftReadOnly = v; return d }
func (d *DualEditor) RightReadOnly(v bool) *DualEditor                    { d.rightReadOnly = v; return d }
func (d *DualEditor) LeftFn(fn func(string) (string, error)) *DualEditor  { d.leftFn = fn; return d }
func (d *DualEditor) RightFn(fn func(string) (string, error)) *DualEditor { d.rightFn = fn; return d }
func (d *DualEditor) SetLeft(s string)                                    { d.leftBuf.Buffer.Set(s) }
func (d *DualEditor) SetRight(s string)                                   { d.rightBuf.Buffer.Set(s) }
func (d *DualEditor) Left() string                                        { return d.leftBuf.Buffer.String() }
func (d *DualEditor) Right() string                                       { return d.rightBuf.Buffer.String() }

func (d *DualEditor) View(c *ui.Context) {
	ui.Row(c).Fill().Margin(18, 0).Children(func() {

		d.leftBuf.ReadOnly = d.leftReadOnly
		d.rightBuf.ReadOnly = d.rightReadOnly

		code.CodeEditor(c, d.leftBuf, code.CodeEditorOptions{}).Fill()

		ui.Spacer(c).Width(18)

		ui.Column(c).AlignItems(ui.Center).Children(func() {
			if d.leftFn != nil {
				if button.IconButton(c, "left-right", "arrow-right").Radius(999).Clicked() {
					res, err := d.leftFn(d.Left())
					if err != nil {
						Toast(c, err.Error(), func() {})
						return
					}
					d.SetRight(res)
				}
			}
			ui.Spacer(c).Height(18)
			if d.rightFn != nil {
				if button.IconButton(c, "right-left", "arrow-left").Radius(999).Clicked() {
					res, err := d.rightFn(d.Right())
					if err != nil {
						Toast(c, err.Error(), func() {})
						return
					}
					d.SetLeft(res)
				}
			}
		})

		ui.Spacer(c).Width(18)

		code.CodeEditor(c, d.rightBuf, code.CodeEditorOptions{}).Fill()
	})
}
