package main

import (
	_ "embed"

	"github.com/zrcoder/ttoy/view"

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
	KeyDataView = "dataview"
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
	convert  *view.Convert
	encode   *view.Encode
	format   *view.Format
	generate *view.Generate
	diff     *view.Diff
	dataView *view.DataView
	plot     *view.Plot
	ballSort *view.BallSort
	hanoi    *view.Hanoi
	iceMagic *view.IceMagic
}

func New() *Ttoy {
	return &Ttoy{
		navSelected: KeyHome,
		convert:     view.NewConvert(),
		encode:      view.NewEncode(),
		format:      view.NewFormat(),
		generate:    view.NewGenerate(),
		diff:        view.NewDiff(),
		dataView:    view.NewDataView(),
		plot:        view.NewPlot(),
		ballSort:    view.NewBallSort(),
		hanoi:       view.NewHanoi(),
		iceMagic:    view.NewIceMagic(),
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
				ui.SidebarItem(c, string(KeyDataView), nil, "View")
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
			t.format.View(c)
		case KeyConvert:
			t.convert.View(c)
		case KeyEncode:
			t.encode.View(c)
		case KeyGenerate:
			t.generate.View(c)
		case KeyDiff:
			t.diff.View(c)
		case KeyDataView:
			t.dataView.DataView(c)
		case KeyPlot:
			t.plot.View(c)
		case KeyIce:
			t.iceMagic.View(c)
		case KeyBall:
			t.ballSort.View(c)
		case KeyHanoi:
			t.hanoi.View(c)
		}
	})
}

func (t *Ttoy) home(c *ui.Context) {
	ui.Column(c).Fill().AlignItems(ui.Center).Justify(ui.Center).Children(func() {
		ui.Text(c, "TToy").FontSize(32).FontWeight(800).TextColor(c.Theme().TextMuted)
		ui.Text(c, "My dev tools.").FontSize(16).TextColor(c.Theme().TextMuted).Margin(8, 0, 0, 0)
	})
}
