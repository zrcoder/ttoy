package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/service"
)

type Encode struct {
	base64    *DualEditor
	html      *DualEditor
	url       *DualEditor
	svc       *service.Encode
	activeTab int
}

func NewEncode() *Encode {
	p := &Encode{svc: &service.Encode{}}
	p.base64 = NewDualEditor("text", "text").
		LeftFn(p.svc.Base64Encode).RightFn(p.svc.Base64Decode).RightReadOnly(false)
	p.html = NewDualEditor("html", "html").
		LeftFn(p.svc.HTMLEncode).RightFn(p.svc.HTMLDecode).RightReadOnly(false)
	p.url = NewDualEditor("text", "text").
		LeftFn(p.svc.URLEncode).RightFn(p.svc.URLDecode).RightReadOnly(false)
	return p
}

func (p *Encode) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Tabs(c, &p.activeTab, "Base64", "HTML", "URL").Margin(18)
		ui.Spacer(c).Height(8)
		switch p.activeTab {
		case 0:
			p.base64.View(c)
		case 1:
			p.html.View(c)
		default:
			p.url.View(c)
		}
	})
}
