package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/hanoi"
)

// HanoiPanel renders a playable Towers of Hanoi puzzle.
type HanoiPanel struct {
	game     *hanoi.Game
	selected int // selected pile index, -1 = none
	msg      string
}

func NewHanoiPanel() *HanoiPanel {
	return &HanoiPanel{
		game:     hanoi.New(),
		selected: -1,
	}
}

func (p *HanoiPanel) View(c *ui.Context) {
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).Fill().Height(44).Children(func() {
			ui.Text(c, "Hanoi").FontSize(20).FontWeight(700)
			ui.Spacer(c)
			ui.Text(c, p.msg).FontSize(13).TextColor(c.Theme().TextMuted)
			ui.Spacer(c).Width(16)
			ui.ButtonBase(c).Width(70).Height(28).Radius(4).Background(c.Theme().Accent).OnClick(func() {
				p.game = hanoi.New()
				p.selected = -1
				p.msg = ""
				c.Invalidate()
			}).Children(func() {
				ui.Text(c, "Reset").TextColor(c.Theme().AccentText)
			})
		})
		ui.Spacer(c).Height(12)

		if p.game.IsDone() {
			ui.Column(c).AlignItems(ui.Center).Children(func() {
				ui.Text(c, "Solved!").FontSize(24).FontWeight(800).TextColor(c.Theme().Success)
				ui.Text(c, p.game.State()).FontSize(14).TextColor(c.Theme().TextMuted)
			})
			return
		}

		ui.Row(c).AlignItems(ui.Center).Margin(0, 0, 12, 0).Children(func() {
			for i, lv := range p.game.Levels() {
				idx := i
				bg := ui.Color{}
				if idx == p.game.LevelIndex() {
					bg = c.Theme().Accent
				}
				ui.ButtonBase(c).Width(70).Height(26).Radius(4).Background(bg).OnClick(func() {
					p.game.SelectLevel(idx)
					p.selected = -1
					p.msg = ""
					c.Invalidate()
				}).Children(func() {
					ui.Text(c, lv.Label).TextColor(c.Theme().Text).FontSize(12)
				})
				if i < len(p.game.Levels())-1 {
					ui.Spacer(c).Width(6)
				}
			}
			ui.Spacer(c).Width(16)
			ui.Text(c, p.game.State()).FontSize(13).TextColor(c.Theme().TextMuted)
		})

		ui.Row(c).AlignItems(ui.Center).Children(func() {
			piles := p.game.Piles()
			diskColors := []ui.Color{
				{R: 255, G: 0, B: 0, A: 255},
				{R: 0, G: 128, B: 0, A: 255},
				{R: 0, G: 0, B: 255, A: 255},
				{R: 255, G: 255, B: 0, A: 255},
				{R: 165, G: 42, B: 42, A: 255},
				{R: 255, G: 192, B: 203, A: 255},
			}
			for i, pile := range piles {
				idx := i
				ui.Column(c.Key(fmt.Sprintf("pile-%d", i))).AlignItems(ui.Center).OnClick(func() {
					p.onPileClick(idx)
					c.Invalidate()
				}).Children(func() {
					ui.Column(c).Width(140).Height(260).BorderWidth(1).BorderColor(c.Theme().Border).Radius(8).
						Background(c.Theme().Surface).Children(func() {
						// Disks (bottom-up), stored as IDs
						disks := pile.Disks
						n := len(disks)
						maxDisks := p.game.CurrentLevel().Disks
						for j := range maxDisks {
							idx := maxDisks - 1 - j
							ui.Row(c).Height(24).Children(func() {
								if idx < n {
									dID := disks[idx]
									w := float32(dID*20 + 40)
									ui.Column(c).Width(w).Height(20).Radius(4).Background(diskColors[dID%len(diskColors)])
								} else {
									ui.Column(c).Width(40).Height(20).Radius(4).Background(ui.Color{R: 240, G: 240, B: 240, A: 60})
								}
							})
						}
					})
					ui.Row(c).Height(24).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
						ui.Text(c, string(rune('A'+i))).FontSize(14).FontWeight(700).TextColor(c.Theme().Text)
					})
				})
				if i < len(piles)-1 {
					ui.Spacer(c).Width(20)
				}
			}
		})
	})
}

func (p *HanoiPanel) onPileClick(i int) {
	if p.game.IsDone() {
		return
	}
	if p.selected == i {
		p.selected = -1
		return
	}
	if p.selected == -1 {
		p.selected = i
		return
	}
	err := p.game.SelectPileByIndex(i)
	if err != nil {
		p.msg = err.Error()
	} else {
		p.msg = ""
	}
	p.selected = -1
	if p.game.IsDone() {
		p.msg = "Solved!"
	}
}

var _ = fmt.Sprintf
