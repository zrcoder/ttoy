package view

import (
	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/code"
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
		tabs(c, &p.activeTab, "JSON", "YAML", "TOML")
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
	p.input.Buffer.SetLanguage(editor.LanguageNamed(lang))
	ui.Row(c).Fill().Children(func() {
		ui.Column(c).Fill().Children(func() {
			code.CodeEditor(c, p.input, code.CodeEditorOptions{}).Fill().Margin(18, 0)
		})
		ui.Column(c).Width(52).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
			if button.IconButton(c, "gen", "arrow-right").Radius(999).Clicked() {
				res, err := fn(p.input.Buffer.String())
				if err != nil {
					Toast(c, err.Error(), func() {})
					return
				}
				p.imageView.SetSvg(res)
			}
		})
		ui.Column(c).Fill().Children(func() {
			p.imageView.View(c)
		})
	})
}
