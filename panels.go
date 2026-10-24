package main

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
	"github.com/zrcoder/ttoy/service"
)

// ---------------------------------------------------------------------------
// Format panel
// ---------------------------------------------------------------------------

type FormatPanel struct {
	json      *DualEditor
	yaml      *DualEditor
	toml      *DualEditor
	html      *DualEditor
	svc       *service.Format
	activeTab int
}

func NewFormatPanel() *FormatPanel {
	p := &FormatPanel{svc: &service.Format{}}
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

func (p *FormatPanel) View(c *ui.Context) {
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

// ---------------------------------------------------------------------------
// Convert panel
// ---------------------------------------------------------------------------

type ConvertPanel struct {
	json2yaml *DualEditor
	json2toml *DualEditor
	yaml2toml *DualEditor
	svc       *service.Convert
	activeTab int
}

func NewConvertPanel() *ConvertPanel {
	p := &ConvertPanel{svc: &service.Convert{}}
	p.json2yaml = NewDualEditor("json", "yaml").
		LeftFn(p.svc.Json2Yaml).RightFn(p.svc.Yaml2Json).RightReadOnly(false)
	p.json2toml = NewDualEditor("json", "toml").
		LeftFn(p.svc.Json2Toml).RightFn(p.svc.Toml2Json).RightReadOnly(false)
	p.yaml2toml = NewDualEditor("yaml", "toml").
		LeftFn(p.svc.Yaml2Toml).RightFn(p.svc.Toml2Yaml).RightReadOnly(false)
	return p
}

func (p *ConvertPanel) View(c *ui.Context) {
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

// ---------------------------------------------------------------------------
// Encode panel
// ---------------------------------------------------------------------------

type EncodePanel struct {
	base64    *DualEditor
	html      *DualEditor
	url       *DualEditor
	svc       *service.Encode
	activeTab int
}

func NewEncodePanel() *EncodePanel {
	p := &EncodePanel{svc: &service.Encode{}}
	p.base64 = NewDualEditor("text", "text").
		LeftFn(p.svc.Base64Encode).RightFn(p.svc.Base64Decode).RightReadOnly(false)
	p.html = NewDualEditor("html", "html").
		LeftFn(p.svc.HTMLEncode).RightFn(p.svc.HTMLDecode).RightReadOnly(false)
	p.url = NewDualEditor("text", "text").
		LeftFn(p.svc.URLEncode).RightFn(p.svc.URLDecode).RightReadOnly(false)
	return p
}

func (p *EncodePanel) View(c *ui.Context) {
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

// ---------------------------------------------------------------------------
// Generate panel
// ---------------------------------------------------------------------------

type GeneratePanel struct {
	asciiInput *string
	fontSel    string
	output     *editor.State
	ascii      *service.Generate
	activeTab  int
	hashIn     *editor.State
	hashOut    []string
}

func NewGeneratePanel() *GeneratePanel {
	art := ""
	p := &GeneratePanel{
		ascii:      &service.Generate{},
		asciiInput: &art,
		output:     &editor.State{Buffer: editor.NewBuffer("")},
		hashIn:     &editor.State{Buffer: editor.NewBuffer("")},
		fontSel:    "standard",
	}
	return p
}

func (p *GeneratePanel) View(c *ui.Context) {
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

func (p *GeneratePanel) asciiTab(c *ui.Context) {
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

func (p *GeneratePanel) hashTab(c *ui.Context) {
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

// ---------------------------------------------------------------------------
// Diff panel
// ---------------------------------------------------------------------------

// diffTab pairs two editable buffers with a diff viewer over them. The
// diff viewer's old side is read-only, so Edit shows both buffers as
// editors and Diff compares them with character-level highlights.
type diffTab struct {
	diff     *editor.DiffState
	oldEd    *editor.State
	newEd    *editor.State
	mode     int // 0: Edit, 1: Diff
	lastMode int
	svc      *service.Sort
}

func newDiffTab(lang *editor.Language, sortable bool) *diffTab {
	oldBuf := editor.NewBuffer("")
	newBuf := editor.NewBuffer("")
	if lang != nil {
		oldBuf.SetLanguage(lang)
		newBuf.SetLanguage(lang)
	}
	t := &diffTab{
		diff: &editor.DiffState{
			Old:      oldBuf,
			New:      newBuf,
			Editable: true,
			Inline:   editor.DiffInlineChar,
		},
		oldEd: &editor.State{Buffer: oldBuf},
		newEd: &editor.State{Buffer: newBuf},
	}
	if sortable {
		t.svc = &service.Sort{}
	}
	return t
}

func (t *diffTab) View(c *ui.Context) {
	// Entering Diff formats both texts so that key order does not diff.
	var sortErr string
	if t.svc != nil && t.mode == 1 && t.lastMode == 0 {
		if err := t.sort(); err != nil {
			t.mode = 0
			sortErr = err.Error()
		}
	}
	t.lastMode = t.mode

	added, deleted := t.diff.Stats()
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).AlignItems(ui.Center).Children(func() {
			ui.Segmented(c, &t.mode, "Edit", "Diff").Margin(18, 18, 18, 0)
			ui.Text(c, fmt.Sprintf("+%d −%d", added, deleted)).FontSize(12).TextColor(c.Theme().TextMuted)
		})
		// ui.Spacer(c).Height(8)
		if t.mode == 0 {
			t.editView(c)
			return
		}
		editor.DiffView(c, t.diff).Grow(1).FillWidth()
	})
	if sortErr != "" {
		Toast(c, sortErr, func() {})
	}
}

func (t *diffTab) editView(c *ui.Context) {
	ui.Row(c).Grow(1).FillWidth().Children(func() {
		ui.Column(c.Key("diff-original")).Fill().Children(func() {
			editor.View(c, t.oldEd).Fill()
		})
		ui.Column(c.Key("diff-modified")).Fill().Children(func() {
			editor.View(c, t.newEd).Fill()
		})
	})
}

// sort formats both texts as JSON so that key order does not diff.
func (t *diffTab) sort() error {
	bufs := []*editor.Buffer{t.diff.Old, t.diff.New}
	texts := make([]string, len(bufs))
	for i, b := range bufs {
		text := b.String()
		if strings.TrimSpace(text) == "" {
			continue
		}
		res, err := t.svc.JSON(text)
		if err != nil {
			return err
		}
		texts[i] = res
	}
	for i, b := range bufs {
		if texts[i] != "" && texts[i] != b.String() {
			b.Set(texts[i])
		}
	}
	return nil
}

type DiffPanel struct {
	text      *diffTab
	json      *diffTab
	activeTab int
}

func NewDiffPanel() *DiffPanel {
	return &DiffPanel{
		text: newDiffTab(nil, false),
		json: newDiffTab(editor.LanguageNamed("json"), true),
	}
}

func (p *DiffPanel) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Tabs(c, &p.activeTab, "Text", "JSON").Margin(18)
		ui.Spacer(c).Height(8)
		if p.activeTab == 0 {
			p.text.View(c)
			return
		}
		p.json.View(c)
	})
}

// ---------------------------------------------------------------------------
// View panel
// ---------------------------------------------------------------------------

type ViewPanel struct {
	input     *editor.State
	imageView *ImageViewer
	activeTab int
	svc       *service.View
}

func NewViewPanel() *ViewPanel {
	p := &ViewPanel{
		input: &editor.State{Buffer: editor.NewBuffer("")},
		svc:   &service.View{},
	}
	p.imageView = NewImageViewer()
	return p
}

func (p *ViewPanel) View(c *ui.Context) {
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

func (p *ViewPanel) viewTab(c *ui.Context, lang string, fn func(string) ([]byte, error)) {
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
					p.imageView.SetSvg(c, res)
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

// ---------------------------------------------------------------------------
// Plot panel
// ---------------------------------------------------------------------------

type PlotPanel struct {
	input     *editor.State
	imageView *ImageViewer
	activeTab int
	svc       *service.Plot
}

func NewPlotPanel() *PlotPanel {
	p := &PlotPanel{
		input: &editor.State{Buffer: editor.NewBuffer("")},
		svc:   &service.Plot{},
	}
	p.imageView = NewImageViewer()
	return p
}

func (p *PlotPanel) View(c *ui.Context) {
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

func (p *PlotPanel) plotTab(c *ui.Context, fn func(string) ([]byte, error), codeType string) {
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
						p.imageView.SetSvg(c, res)
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
