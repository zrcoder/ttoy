package view

import (
	"fmt"

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
	msg          string
}

func NewIceMagic() *IceMagic {
	g := icemagic.New()
	return &IceMagic{
		game:       g,
		chapterSel: 0,
		levelSel:   0,
	}
}

func (p *IceMagic) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).Fill().Height(44).Children(func() {
			ui.Text(c, "IceMagic").FontSize(20).FontWeight(700)
			ui.Spacer(c)
			ui.Text(c, p.msg).FontSize(13).TextColor(c.Theme().TextMuted)
		})
		ui.Spacer(c).Height(12)

		ui.Row(c).AlignItems(ui.Center).Children(func() {
			ui.Text(c, "Chapter").FontSize(13)
			chapters := p.game.Chapters()
			chLabels := make([]string, len(chapters))
			for i := range chapters {
				chLabels[i] = fmt.Sprintf("Ch %d", i+1)
			}
			ui.Combobox(c, &p.chapterLabel, chLabels).Width(100)
			ui.Spacer(c).Width(12)
			levelCount := chapters[p.chapterSel]
			ui.Text(c, "Level").FontSize(13)
			lvLabels := make([]string, levelCount)
			for i := range levelCount {
				lvLabels[i] = fmt.Sprintf("%d", i+1)
			}
			ui.Combobox(c, &p.levelLabel, lvLabels).Width(80)
			ui.Spacer(c).Width(12)
			ui.ButtonBase(c).Width(60).Height(28).Radius(4).Background(c.Theme().Accent).OnClick(func() {
				p.game.SelectLevel(p.chapterSel+1, p.levelSel+1)
				p.msg = ""
				c.Invalidate()
			}).Children(func() {
				ui.Text(c, "Go").TextColor(c.Theme().AccentText)
			})
		})

		ui.Spacer(c).Height(12)

		// Grid
		grid := p.game.Grid()
		if len(grid) == 0 {
			ui.Text(c, "Loading grid...").TextColor(c.Theme().TextMuted)
			return
		}

		ui.Row(c).AlignItems(ui.Center).Children(func() {
			for r, row := range grid {
				ui.Column(c).Children(func() {
					for cIdx, sprite := range row {
						r2, c2 := r, cIdx
						p.cell(c, sprite, r2, c2)
					}
				})
				if r < len(grid)-1 {
					ui.Spacer(c).Height(2)
				}
			}
		})

		ui.Spacer(c).Height(16)

		// Controls
		ui.Row(c).AlignItems(ui.Center).Children(func() {
			p.controlBtn(c, "A", "↙", func() { p.game.MagicLeft(); p.msg = "" })
			ui.Spacer(c).Width(16)
			p.controlBtn(c, "J", "←", func() { p.game.MoveLeft(); p.msg = "" })
			p.controlBtn(c, "L", "→", func() { p.game.MoveRight(); p.msg = "" })
			ui.Spacer(c).Width(16)
			p.controlBtn(c, "D", "↘", func() { p.game.MagicRight(); p.msg = "" })
		})
		ui.Spacer(c).Width(16)
		ui.ButtonBase(c).Width(70).Height(28).Radius(4).Background(c.Theme().Accent).OnClick(func() {
			p.game.SelectLevel(p.chapterSel+1, p.levelSel+1)
			p.msg = ""
			c.Invalidate()
		}).Children(func() {
			ui.Text(c, "Reset").TextColor(c.Theme().AccentText)
		})
	})
}

var _ = fmt.Sprintf

func (p *IceMagic) cell(c *ui.Context, sprite *icemagic.Sprite, r, col int) {
	cell := sprite.Cell
	img := ""
	if cell.Image != "" {
		img = cell.Image
	}

	ui.Column(c.Key(fmt.Sprintf("cell-%d-%d", r, col))).Width(34).Height(34).OnClick(func() {
		c.Invalidate()
	}).Children(func() {
		ui.Row(c).Fill().Children(func() {
			if img != "" {
				ui.Text(c, img).FontSize(20)
			} else {
				ui.Text(c, " ").FontSize(20)
			}
		})
	})
}

func (p *IceMagic) controlBtn(c *ui.Context, key, icon string, action func()) {
	ui.Column(c).AlignItems(ui.Center).Children(func() {
		ui.ButtonBase(c).Width(44).Height(44).Radius(8).Background(c.Theme().Surface).OnClick(action).Children(func() {
			ui.Text(c, icon).FontSize(20)
		})
		ui.Text(c, key).FontSize(11).TextColor(c.Theme().TextMuted)
	})
}
