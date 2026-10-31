package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
	"github.com/zrcoder/ttoy/service"
)

type DataView struct {
	input     *editor.State
	imageView *ImageViewer
	activeTab int
	svc       *service.View
}

func NewDataView() *DataView {
	p := &DataView{
		input: &editor.State{Buffer: editor.NewBuffer("")},
		svc:   &service.View{},
	}
	p.imageView = NewImageViewer()
	return p
}

func (p *DataView) DataView(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Tabs(c, &p.activeTab, "JSON", "YAML", "TOML").Margin(18)
		switch p.activeTab {
		case 0:
			p.viewTab(c, "json", p.svc.Json)
		case 1:
			p.viewTab(c, "yaml", p.svc.Yaml)
		default:
			p.viewTab(c, "toml", p.svc.Toml)
		}
	})
}

func (p *DataView) viewTab(c *ui.Context, lang string, fn func(string) ([]byte, error)) {
	p.input.ReadOnly = false
	ui.Row(c).Fill().Children(func() {
		ui.Column(c).Fill().Children(func() {
			editor.View(c, p.input).Fill()
		})
		ui.Column(c).Width(52).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
			btn := ui.ButtonBase(c).Width(32).Height(32).Radius(6).BorderWidth(2).BorderColor(c.Theme().Surface).
				OnClick(func() {
					res, err := fn(p.input.Buffer.String())
					if err != nil {
						Toast(c, err.Error(), func() {})
						return
					}
					p.imageView.SetSvg(res)
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
