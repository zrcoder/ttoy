package view

import (
	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/choice"
	"github.com/ELY-Open-Source/CrychicUI/feedback"
	"github.com/ELY-Open-Source/CrychicUI/input"
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
	hashInput  string
	hashOut    []string
}

func NewGenerate() *Generate {
	art := ""
	p := &Generate{
		ascii:      &service.Generate{},
		asciiInput: &art,
		output:     &editor.State{Buffer: editor.NewBuffer("")},
		hashInput:  "",
		fontSel:    "standard",
	}
	return p
}

func (p *Generate) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		tabs(c, &p.activeTab, "ASCII Art", "Hash")
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
		input.TextField(c, "", p.asciiInput, input.TextFieldOptions{})
		ui.Row(c).AlignItems(ui.Center).Children(func() {
			fonts := []string{"standard", "slant", "block", "big", "starwars"}
			choice.Dropdown(c, "standard", fonts, &p.fontSel)
			button.Solid(c, "Generate").Margin(18).OnClick(func() {
				for _, ch := range *(p.asciiInput) {
					if ch > 127 {
						return
					}
				}
				res, err := p.ascii.AsciiArt(*p.asciiInput, p.fontSel)
				if err != nil {
					return
				}
				p.output.Buffer.Set(res)
			})
		})

		editor.View(c, p.output).Fill().Margin(0, 0, 18, 0)
	})
}

func (p *Generate) hashTab(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		input.AutoTextArea(c, &p.hashInput, input.AutoTextAreaOptions{}).Height(360).Margin(18, 0)
		errorAlert := false
		button.Solid(c, "Hash").Margin(18, 0).OnClick(func() {
			in := p.hashInput
			res, err := p.ascii.Hash(in)
			if err != nil {
				feedback.DismissibleAlert(c, feedback.ToneWarning, err.Error(), "", &errorAlert, feedback.DismissibleAlertOptions{})
				return
			}
			p.hashOut = []string{
				res["md5"], res["sha1"], res["sha256"], res["sha512"],
			}
		})
		if len(p.hashOut) > 0 {
			labels := []string{"MD5", "SHA1", "SHA256", "SHA512"}
			for i, label := range labels {
				input.TextField(c, label, &p.hashOut[i], input.TextFieldOptions{}).Margin(10, 0)
			}
		}
	})
}
