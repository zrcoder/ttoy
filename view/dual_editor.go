package view

import (
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
	return &DualEditor{
		leftBuf:       &editor.State{Buffer: editor.NewBuffer("")},
		rightBuf:      &editor.State{Buffer: editor.NewBuffer("")},
		leftLang:      leftLang,
		rightLang:     rightLang,
		leftReadOnly:  false,
		rightReadOnly: true,
	}
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
	ui.Row(c).Fill().Children(func() {
		// Left editor
		ui.Column(c.Key("dual-left")).Fill().Children(func() {
			d.leftBuf.ReadOnly = d.leftReadOnly
			editor.View(c, d.leftBuf).Fill()
		})

		// Transform buttons + copy
		ui.Column(c).Width(72).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
			if d.leftFn != nil {
				btn := ui.ButtonBase(c).Width(32).Height(32).Radius(6).BorderWidth(2).BorderColor(c.Theme().Surface).
					OnClick(func() {
						res, err := d.leftFn(d.Left())
						if err != nil {
							Toast(c, err.Error(), func() {})
							return
						}
						d.SetRight(res)
					}).Children(func() {
					ui.Text(c, "→").FontSize(14)
				})
				ui.Spacer(c).Height(8)

				if btn.Hovered() {
					btn.Background(c.Theme().SurfaceHover)
				}
			}
			if d.rightFn != nil {
				btn := ui.ButtonBase(c).Width(32).Height(32).Radius(6).BorderWidth(2).BorderColor(c.Theme().Surface).
					OnClick(func() {
						res, err := d.rightFn(d.Right())
						if err != nil {
							Toast(c, err.Error(), func() {})
							return
						}
						d.SetLeft(res)
					}).Children(func() {
					ui.Text(c, "←").FontSize(14)
				})

				if btn.Hovered() {
					btn.Background(c.Theme().SurfaceHover)
				}
			}
			ui.Spacer(c).Height(16)
		})

		// Right editor
		ui.Column(c.Key("dual-right")).Fill().Children(func() {
			d.rightBuf.ReadOnly = d.rightReadOnly
			editor.View(c, d.rightBuf).Fill()
		})
	})
}
