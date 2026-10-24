package main

import (
	_ "embed"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

//go:embed resources/icons/moon.svg
var moonIconData []byte

//go:embed resources/icons/sun.svg
var sunIconData []byte

const (
	KeyHome     = "home"
	KeyFormat   = "fmt"
	KeyConvert  = "cvt"
	KeyEncode   = "codec"
	KeyGenerate = "gen"
	KeyDiff     = "diff"
	KeyView     = "view"
	KeyPlot     = "graph"
	KeyIce      = "icemagic"
	KeyBall     = "ballsort"
	KeyHanoi    = "hanoi"
)

// Ttoy is the application model. All UI state lives here; components are
// stateless and receive config/callbacks through their config structs.
type Ttoy struct {
	navSelected string

	// Feature panels (lazy-initialized, own their state).
	convert  *ConvertPanel
	encode   *EncodePanel
	format   *FormatPanel
	generate *GeneratePanel
	diff     *DiffPanel
	view     *ViewPanel
	plot     *PlotPanel
	ballSort *BallSortPanel
	hanoi    *HanoiPanel
	iceMagic *IceMagicPanel
}

func New() *Ttoy {
	return &Ttoy{
		navSelected: KeyHome,
	}
}

func (t *Ttoy) View(c *ui.Context) {
	ui.Row(c.Key("root")).Fill().Children(func() {
		t.nav(c)
		t.content(c)
	})
}

func (t *Ttoy) toggleTheme(c *ui.Context) {
	if mygo.Theme.IsDark() {
		mygo.Theme.SetSource(mygo.ThemeLight)
	} else {
		mygo.Theme.SetSource(mygo.ThemeDark)
	}
	c.Invalidate()
}

func (t *Ttoy) nav(c *ui.Context) {
	ui.Column(c.Key("nav")).Width(240).FillHeight().Children(func() {
		ui.Sidebar(c, &t.navSelected, func() {
			ui.SidebarItem(c, string(KeyHome), nil, "Home")

			ui.SidebarSection(c, "Tools", nil, func() {
				ui.SidebarItem(c, string(KeyFormat), nil, "Format")
				ui.SidebarItem(c, string(KeyConvert), nil, "Convert")
				ui.SidebarItem(c, string(KeyEncode), nil, "Encode")
				ui.SidebarItem(c, string(KeyGenerate), nil, "Generate")
				ui.SidebarItem(c, string(KeyDiff), nil, "Diff")
				ui.SidebarItem(c, string(KeyView), nil, "View")
				ui.SidebarItem(c, string(KeyPlot), nil, "Plot")
			})
			ui.SidebarSection(c, "Games", nil, func() {
				ui.SidebarItem(c, string(KeyIce), nil, "IceMagic")
				ui.SidebarItem(c, string(KeyBall), nil, "BallSort")
				ui.SidebarItem(c, string(KeyHanoi), nil, "Hanoi")
			})
		}).Fill()

		ui.Row(c).Absolute().Bottom(0).Left(0).Height(40).Children(func() {
			themeIcon := moonIconData
			if mygo.Theme.IsDark() {
				themeIcon = sunIconData
			}
			btn := ui.ButtonBase(c).Margin(8).Padding(8).Radius(4).OnClick(func() {
				t.toggleTheme(c)
			})
			if btn.Hovered() {
				btn.Background(c.Theme().SurfaceHover)
			}
			btn.Children(func() {
				ui.Icon(c, ui.MustParseSVG(themeIcon)).Size(16, 16).TextColor(c.Theme().Text)
			})
		})
	})
}

func (t *Ttoy) content(c *ui.Context) {
	ui.Column(c.Key("content")).Margin(18).Fill().Children(func() {
		switch t.navSelected {
		case KeyHome:
			t.home(c)
		case KeyFormat:
			t.ensureFormat().View(c)
		case KeyConvert:
			t.ensureConvert().View(c)
		case KeyEncode:
			t.ensureEncode().View(c)
		case KeyGenerate:
			t.ensureGenerate().View(c)
		case KeyDiff:
			t.ensureDiff().View(c)
		case KeyView:
			t.ensureView().View(c)
		case KeyPlot:
			t.ensurePlot().View(c)
		case KeyIce:
			t.ensureIceMagic().View(c)
		case KeyBall:
			t.ensureBallSort().View(c)
		case KeyHanoi:
			t.ensureHanoi().View(c)
		}
	})
}

func (t *Ttoy) home(c *ui.Context) {
	ui.Column(c).Fill().AlignItems(ui.Center).Justify(ui.Center).Children(func() {
		ui.Text(c, "TToy").FontSize(32).FontWeight(800).TextColor(c.Theme().TextMuted)
		ui.Text(c, "My dev tools.").FontSize(16).TextColor(c.Theme().TextMuted).Margin(8, 0, 0, 0)
	})
}

func (t *Ttoy) ensureFormat() *FormatPanel {
	if t.format == nil {
		t.format = NewFormatPanel()
	}
	return t.format
}

func (t *Ttoy) ensureConvert() *ConvertPanel {
	if t.convert == nil {
		t.convert = NewConvertPanel()
	}
	return t.convert
}

func (t *Ttoy) ensureEncode() *EncodePanel {
	if t.encode == nil {
		t.encode = NewEncodePanel()
	}
	return t.encode
}

func (t *Ttoy) ensureGenerate() *GeneratePanel {
	if t.generate == nil {
		t.generate = NewGeneratePanel()
	}
	return t.generate
}

func (t *Ttoy) ensureDiff() *DiffPanel {
	if t.diff == nil {
		t.diff = NewDiffPanel()
	}
	return t.diff
}

func (t *Ttoy) ensureView() *ViewPanel {
	if t.view == nil {
		t.view = NewViewPanel()
	}
	return t.view
}

func (t *Ttoy) ensurePlot() *PlotPanel {
	if t.plot == nil {
		t.plot = NewPlotPanel()
	}
	return t.plot
}

func (t *Ttoy) ensureBallSort() *BallSortPanel {
	if t.ballSort == nil {
		t.ballSort = NewBallSortPanel()
	}
	return t.ballSort
}

func (t *Ttoy) ensureHanoi() *HanoiPanel {
	if t.hanoi == nil {
		t.hanoi = NewHanoiPanel()
	}
	return t.hanoi
}

func (t *Ttoy) ensureIceMagic() *IceMagicPanel {
	if t.iceMagic == nil {
		t.iceMagic = NewIceMagicPanel()
	}
	return t.iceMagic
}
