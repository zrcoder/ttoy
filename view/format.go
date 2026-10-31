package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/service"
)

type Format struct {
	json      *DualEditor
	yaml      *DualEditor
	toml      *DualEditor
	html      *DualEditor
	svc       *service.Format
	activeTab int
}

func NewFormat() *Format {
	p := &Format{svc: &service.Format{}}
	p.json = NewDualEditor("json", "json").
		LeftFn(p.svc.Json).RightFn(p.svc.Json).RightReadOnly(false)
	p.yaml = NewDualEditor("yaml", "yaml").
		LeftFn(p.svc.Yaml).RightFn(p.svc.Yaml).RightReadOnly(false)
	p.toml = NewDualEditor("toml", "toml").
		LeftFn(p.svc.Toml).RightFn(p.svc.Toml).RightReadOnly(false)
	p.html = NewDualEditor("html", "html").
		LeftFn(p.svc.Html).RightFn(p.svc.Html).RightReadOnly(false)
	return p
}

func (p *Format) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Tabs(c, &p.activeTab, "JSON", "YAML", "TOML", "HTML").Margin(18)
		ui.Spacer(c).Height(8)
		switch p.activeTab {
		case 0:
			p.json.View(c)
		case 1:
			p.yaml.View(c)
		case 2:
			p.toml.View(c)
		default:
			p.html.View(c)
		}
	})
}
