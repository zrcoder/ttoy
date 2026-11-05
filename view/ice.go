package view

import (
	"embed"
	"fmt"
	"log"
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/icemagic"
)

// IceMagic renders the IceMagic puzzle grid.
type IceMagic struct {
	game         *icemagic.Game
	chapterSel   int
	levelSel     int
	chapterLabel string
	levelLabel   string

	bitmaps map[string]*ui.Bitmap
	gifs    map[string]*ui.VideoPlayer

	// Popover state (keeps popover hidden until button clicked)
	chapterPopoverOpen bool
	levelPopoverOpen   bool
}

func NewIceMagic(assertsFS embed.FS) *IceMagic {
	g := icemagic.New()
	res := &IceMagic{
		game:               g,
		chapterSel:         0,
		levelSel:           0,
		chapterPopoverOpen: false,
		levelPopoverOpen:   false,
	}

	res.bitmaps = map[string]*ui.Bitmap{
		"images/icemagic/wall1.png": nil,
		"images/icemagic/wall2.png": nil,
		"images/icemagic/wall3.png": nil,
		"images/icemagic/wall4.png": nil,
		"images/icemagic/wall5.png": nil,
		"images/icemagic/ice.png":   nil,
	}
	for k := range res.bitmaps {
		data, err := assertsFS.ReadFile("resources/" + k)
		if err != nil {
			log.Fatal(err)
		}
		res.bitmaps[k], err = ui.DecodeBitmap(data)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Animated GIFs: DecodeBitmap shows only the first frame, so play
	// them via VideoPlayer (created once, like bitmaps).
	res.gifs = map[string]*ui.VideoPlayer{
		"images/icemagic/fire.gif":   nil,
		"images/icemagic/player.gif": nil,
	}
	for k := range res.gifs {
		data, err := assertsFS.ReadFile("resources/" + k)
		if err != nil {
			log.Fatal(err)
		}
		dec, err := ui.NewGIFDecoder(data)
		if err != nil {
			log.Fatal(err)
		}
		p, err := ui.NewVideoPlayer(dec)
		if err != nil {
			log.Fatal(err)
		}
		p.SetLoop(true)
		res.gifs[k] = p
	}

	return res
}

func (p *IceMagic) View(c *ui.Context) {
	c.OnShortcut(0, ui.KeyJ, func() { p.game.MoveLeft() })
	c.OnShortcut(0, ui.KeyL, func() { p.game.MoveRight() })
	c.OnShortcut(0, ui.KeyA, func() { p.game.MagicLeft() })
	c.OnShortcut(0, ui.KeyD, func() { p.game.MagicRight() })

	ui.Column(c).Fill().Focus().Children(func() {
		ui.Row(c).Fill().Height(44).Children(func() {
			ui.Text(c, "IceMagic").FontSize(20).FontWeight(700)
		})

		// Note: focus is managed at app level (sidebar).
		// Keys work after clicking because focus moves to game controls.
		// The c.Key() approach may not work due to sidebar focus management.

		ui.Column(c).FillWidth().AlignItems(ui.Center).Children(func() {
			ui.Row(c).Children(func() {
				chapters := p.game.Chapters()
				chBtn := ui.Button(c, fmt.Sprintf("Chapter %d ▾", p.chapterSel+1))
				if chBtn.Clicked() {
					p.chapterPopoverOpen = !p.chapterPopoverOpen
					c.Invalidate()
				}
				ui.Popover(c, chBtn, &p.chapterPopoverOpen, func() {
					for i := range chapters {
						label := strconv.Itoa(i + 1)
						entry := ui.Row(c.Key(label)).Padding(6, 12).Radius(5).Children(func() { ui.Text(c, label) })
						if entry.Hovered() {
							entry.Background(c.Theme().Accent).TextColor(c.Theme().AccentText)
						}
						if entry.Clicked() {
							p.chapterSel = i
							p.game.SelectLevel(p.chapterSel+1, p.levelSel+1)
							p.chapterPopoverOpen = false
							c.Invalidate()
						}
					}
				})
				levelCount := chapters[p.chapterSel]
				lvBtn := ui.Button(c, fmt.Sprintf("Level %d ▾", p.levelSel+1)).Margin(8)
				if lvBtn.Clicked() {
					p.levelPopoverOpen = !p.levelPopoverOpen
					c.Invalidate()
				}
				ui.Popover(c, lvBtn, &p.levelPopoverOpen, func() {
					for i := range levelCount {
						label := strconv.Itoa(i + 1)
						entry := ui.Row(c.Key(label)).Padding(6, 12).Radius(5).Children(func() { ui.Text(c, label) })
						if entry.Hovered() {
							entry.Background(c.Theme().Accent).TextColor(c.Theme().AccentText)
						}
						if entry.Clicked() {
							p.levelSel = i
							p.game.SelectLevel(p.chapterSel+1, p.levelSel+1)
							p.levelPopoverOpen = false
							c.Invalidate()
						}
					}
				})
				ui.PrimaryButton(c, "Reset").Margin(18).OnClick(func() {
					p.game.Reset()
					c.Invalidate()
				})
			})

			grid := p.game.Grid()
			if len(grid) == 0 {
				ui.Text(c, "Loading grid...").TextColor(c.Theme().TextMuted)
				return
			}

			ui.Column(c).Margin(28, 0).Children(func() {
				for _, row := range grid {
					ui.Row(c).Children(func() {
						for _, sprite := range row {
							p.cell(c, sprite)
						}
					})
				}
			})

			ui.Row(c).Children(func() {
				p.controlBtn(c, "J", "←", func() { p.game.MoveLeft() })
				p.controlBtn(c, "L", "→", func() { p.game.MoveRight() })
			})
			ui.Row(c).Children(func() {
				p.controlBtn(c, "A", "↙", func() { p.game.MagicLeft() })
				p.controlBtn(c, "D", "↘", func() { p.game.MagicRight() })
			})
		})
	})
}

var _ = fmt.Sprintf

func (p *IceMagic) cell(c *ui.Context, sprite *icemagic.Sprite) {
	cell := sprite.Cell

	ui.Box(c).Width(34).Height(34).BorderColor(c.Theme().Border).BorderWidth(
		b2f(cell.BorderTop),
		b2f(cell.BorderRight),
		b2f(cell.BorderBottom),
		b2f(cell.BorderLeft),
	).Children(func() {
		if gif := p.gifs[cell.Image]; gif != nil {
			ui.Video(c, gif).Fill().Fit(ui.Cover)
			return
		}
		ui.Image(c, p.bitmaps[cell.Image]).Fill().Fit(ui.Cover)
	})
}

// b2f reports a bool border side as a border width.
func b2f(b bool) float32 {
	if b {
		return 1
	}
	return 0
}

func (p *IceMagic) controlBtn(c *ui.Context, key, icon string, action func()) {
	ui.Column(c).AlignItems(ui.Center).Margin(10).Children(func() {
		ui.Button(c, icon).OnClick(action)
		ui.Text(c, key).FontSize(11).TextColor(c.Theme().TextMuted)
	})
}
