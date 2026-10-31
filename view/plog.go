package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
	"github.com/zrcoder/ttoy/service"
)

type Plot struct {
	input     *editor.State
	imageView *ImageViewer
	activeTab int
	svc       *service.Plot
}

func NewPlot() *Plot {
	p := &Plot{
		input: &editor.State{Buffer: editor.NewBuffer("")},
		svc:   &service.Plot{},
	}
	p.imageView = NewImageViewer()
	return p
}

func (p *Plot) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Tabs(c, &p.activeTab, "D2", "Ndor").Margin(18)
		switch p.activeTab {
		case 0:
			p.plotTab(c, p.svc.D2, "d2")
		default:
			p.plotTab(c, p.svc.Ndor, "ndor")
		}
	})
}

func (p *Plot) plotTab(c *ui.Context, fn func(string) ([]byte, error), codeType string) {
	p.input.ReadOnly = false
	lang := "c"
	if codeType == "d2" {
		lang = "python"
	}
	ui.Row(c).Fill().Children(func() {
		ui.Column(c).Fill().Children(func() {
			editor.View(c, p.input).Lang(lang).Fill()
		})
		ui.Column(c).Width(52).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
			btn := ui.ButtonBase(c).Width(32).Height(32).Radius(6).BorderWidth(2).BorderColor(c.Theme().Surface).
				OnClick(func() {
					res, err := fn(p.input.Buffer.String())
					if err != nil {
						Toast(c, err.Error(), func() {})
						return
					}
					if codeType == "d2" {
						p.imageView.SetSvg(res)
					} else {
						p.imageView.SetPng(res)
					}
				}).Children(func() {
				ui.Text(c, "→").FontSize(14)
			})

			if btn.Hovered() {
				btn.Background(c.Theme().SurfaceHover)
			}
		})
		ui.Column(c).Fill().Children(func() {
			p.imageView.View(c)
		})
	})
}
