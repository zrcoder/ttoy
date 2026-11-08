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
	svc := &service.Format{}
	return &Format{
		svc:  svc,
		json: NewDualEditor("json", "json").LeftFn(svc.Json),
		yaml: NewDualEditor("yaml", "yaml").LeftFn(svc.Yaml),
		toml: NewDualEditor("toml", "toml").LeftFn(svc.Toml),
		html: NewDualEditor("html", "html").LeftFn(svc.Html),
	}
}

func (p *Format) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		tabs(c, &p.activeTab, "JSON", "YAML", "TOML", "HTML")
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
