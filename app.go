package main

import (
	"embed"

	"github.com/zrcoder/ttoy/view"

	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/choice"
	"github.com/ELY-Open-Source/CrychicUI/icon"
	"github.com/ELY-Open-Source/CrychicUI/sound"
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

//go:embed resources/icons/moon.svg
var moonIconData []byte

//go:embed resources/icons/sun.svg
var sunIconData []byte

//go:embed resources/*
var assertsFS embed.FS

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

	musicEnabled bool
}

func New() *Ttoy {
	return &Ttoy{
		navSelected:  KeyHome,
		convert:      view.NewConvert(),
		encode:       view.NewEncode(),
		format:       view.NewFormat(),
		generate:     view.NewGenerate(),
		diff:         view.NewDiff(),
		dataView:     view.NewDataView(),
		plot:         view.NewPlot(),
		ballSort:     view.NewBallSort(),
		hanoi:        view.NewHanoi(),
		iceMagic:     view.NewIceMagic(assertsFS),
		musicEnabled: true,
	}
}

func (t *Ttoy) View(c *ui.Context) {
	ui.Row(c.Key("root")).Fill().Children(func() {
		t.nav(c)
		t.content(c)
	})
}

func (t *Ttoy) nav(c *ui.Context) {
	sidebarItem := func(c *ui.Context, key string, icon *ui.SVG, label string) {
		if ui.SidebarItem(c, key, icon, label).Clicked() {
			sound.Play(sound.Select)
		}
	}
	ui.Column(c.Key("nav")).Width(240).FillHeight().Children(func() {
		ui.Sidebar(c, &t.navSelected, func() {
			sidebarItem(c, string(KeyHome), icon.Must("home"), "Home")
			ui.SidebarSection(c, "Tools", nil, func() {
				sidebarItem(c, KeyFormat, icon.Must("cog"), "Format")
				sidebarItem(c, KeyConvert, icon.Must("code"), "Convert")
				sidebarItem(c, KeyEncode, icon.Must("refresh-ccw"), "Encode")
				sidebarItem(c, KeyGenerate, icon.Must("sparkles"), "Generate")
				sidebarItem(c, KeyDiff, icon.Must("database-zap"), "Diff")
				sidebarItem(c, KeyDataView, icon.Must("grid"), "Data View")
				sidebarItem(c, KeyPlot, icon.Must("chart-bar"), "Plot")
			})
			ui.SidebarSection(c, "Games", nil, func() {
				sidebarItem(c, KeyIce, icon.Must("snowflake"), "IceMagic")
				sidebarItem(c, KeyBall, icon.Must("dice-6"), "BallSort")
				sidebarItem(c, KeyHanoi, icon.Must("radio-tower"), "Hanoi")
			})
		}).Fill()

		ui.Row(c).Absolute().Bottom(10).Left(10).Width(200).Height(40).Children(func() {
			dark := mygo.Theme.IsDark()
			if button.ThemeToggle(c, &dark, button.ThemeToggleOptions{}).Changed() {
				if mygo.Theme.IsDark() {
					mygo.Theme.SetSource(mygo.ThemeLight)
				} else {
					mygo.Theme.SetSource(mygo.ThemeDark)
				}
			}

			ui.Spacer(c)

			if choice.MuteToggle(c, "music", &t.musicEnabled).Width(36).Height(36).Changed() {
				sound.SetEnabled(t.musicEnabled)
			}
		})
	})
}

func (t *Ttoy) content(c *ui.Context) {
	ui.Box(c.Key("content")).Margin(18).Fill().Children(func() {
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
