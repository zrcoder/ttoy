package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
	"github.com/zrcoder/ttoy/service"
)

type Generate struct {
	asciiInput *string
	fontSel    string
	output     *editor.State
	ascii      *service.Generate
	activeTab  int
	hashIn     *editor.State
	hashOut    []string
}

func NewGenerate() *Generate {
	art := ""
	p := &Generate{
		ascii:      &service.Generate{},
		asciiInput: &art,
		output:     &editor.State{Buffer: editor.NewBuffer("")},
		hashIn:     &editor.State{Buffer: editor.NewBuffer("")},
		fontSel:    "standard",
	}
	return p
}

func (p *Generate) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Tabs(c, &p.activeTab, "ASCII Art", "Hash").Margin(18)
		switch p.activeTab {
		case 0:
			p.asciiTab(c)
		default:
			p.hashTab(c)
		}
	})
}

func (p *Generate) asciiTab(c *ui.Context) {
	p.output.ReadOnly = true
	ui.Column(c).Fill().Children(func() {
		ui.TextInput(c, p.asciiInput).Margin(18)
		ui.Row(c).AlignItems(ui.Center).Children(func() {
			fonts := []string{"standard", "slant", "block", "big", "starwars"}
			ui.Combobox(c, &p.fontSel, fonts).Width(160).Margin(18)
			ui.ButtonBase(c).Width(80).Height(30).Radius(4).Background(c.Theme().Accent).OnClick(func() {
				res, err := p.ascii.AsciiArt(*p.asciiInput, p.fontSel)
				if err != nil {
					Toast(c, err.Error(), func() {})
					return
				}
				p.output.Buffer.Set(res)
			}).Children(func() {
				ui.Text(c, "Generate").TextColor(c.Theme().AccentText)
			})
		})

		editor.View(c, p.output).Fill()
	})
}

func (p *Generate) hashTab(c *ui.Context) {
	p.hashIn.ReadOnly = false
	ui.Column(c).Fill().Children(func() {
		editor.View(c, p.hashIn).Height(360)
		ui.ButtonBase(c).Width(80).Height(30).Radius(4).Margin(18).Background(c.Theme().Accent).OnClick(func() {
			in := p.hashIn.Buffer.String()
			res, err := p.ascii.Hash(in)
			if err != nil {
				Toast(c, err.Error(), func() {})
				return
			}
			p.hashOut = []string{
				res["md5"], res["sha1"], res["sha256"], res["sha512"],
			}
		}).Children(func() {
			ui.Text(c, "Hash").TextColor(c.Theme().AccentText)
		})
		if len(p.hashOut) > 0 {
			labels := []string{"MD5", "SHA1", "SHA256", "SHA512"}
			for i, label := range labels {
				ui.Row(c).Fill().Height(24).Children(func() {
					ui.Text(c, label).Width(60).FontSize(13)
					ui.Text(c, p.hashOut[i]).FontSize(13)
				})
			}
		}
	})
}
