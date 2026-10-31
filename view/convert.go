package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/service"
)

type Convert struct {
	json2yaml *DualEditor
	json2toml *DualEditor
	yaml2toml *DualEditor
	svc       *service.Convert
	activeTab int
}

func NewConvert() *Convert {
	p := &Convert{svc: &service.Convert{}}
	p.json2yaml = NewDualEditor("json", "yaml").
		LeftFn(p.svc.Json2Yaml).RightFn(p.svc.Yaml2Json).RightReadOnly(false)
	p.json2toml = NewDualEditor("json", "toml").
		LeftFn(p.svc.Json2Toml).RightFn(p.svc.Toml2Json).RightReadOnly(false)
	p.yaml2toml = NewDualEditor("yaml", "toml").
		LeftFn(p.svc.Yaml2Toml).RightFn(p.svc.Toml2Yaml).RightReadOnly(false)
	return p
}

func (p *Convert) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Tabs(c, &p.activeTab, "JSON-YAML", "JSON-TOML", "YAML-TOML").Margin(18)
		ui.Spacer(c).Height(8)
		switch p.activeTab {
		case 0:
			p.json2yaml.View(c)
		case 1:
			p.json2toml.View(c)
		default:
			p.yaml2toml.View(c)
		}
	})
}
