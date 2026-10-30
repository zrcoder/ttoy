package main

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/hanoi"
)

// HanoiPanel renders a playable Towers of Hanoi puzzle.
type HanoiPanel struct {
	game *hanoi.Game
	// cursor is the keyboard cursor over the piles: arrows move it,
	// Enter or Space plays there, as clicking does.
	cursor int
	msg    string
}

func NewHanoiPanel() *HanoiPanel {
	return &HanoiPanel{
		game: hanoi.New(),
	}
}

func (p *HanoiPanel) View(c *ui.Context) {
	// The board takes the keyboard as the panel appears and when clicked;
	// its shortcuts reach it while a button of the panel holds the focus.
	board := ui.Column(c.Key("hanoi-board")).Fill().Focusable().AutoFocus()
	p.handleKeys(board)
	board.Children(func() {
		ui.Row(c).Fill().Height(44).Children(func() {
			ui.Text(c, "Hanoi").FontSize(20).FontWeight(700)
			ui.Text(c, p.msg).FontSize(13).TextColor(c.Theme().TextMuted)
			ui.Spacer(c)
			ui.ButtonBase(c).Width(70).Height(28).Radius(4).Background(c.Theme().Accent).OnClick(func() {
				p.game = hanoi.New()
				p.cursor = 0
				p.msg = ""
				c.Invalidate()
			}).Children(func() {
				ui.Text(c, "Reset").TextColor(c.Theme().AccentText)
			})
		})

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
					p.cursor = 0
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
				// The cursor marker shows where the keyboard plays; the
				// accent border marks the pile the held disk came from.
				border := c.Theme().Border
				if p.game.ShiftPileIndex() == idx {
					border = c.Theme().Accent
				}
				ui.Column(c.Key(fmt.Sprintf("pile-%d", i))).AlignItems(ui.Center).Children(func() {
					ui.Row(c).Height(18).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
						if p.cursor == idx {
							ui.Text(c, "▾").FontSize(14).TextColor(c.Theme().Accent)
						} else {
							ui.Text(c, " ").FontSize(14)
						}
					})
					ui.Column(c).Width(140).Height(260).AlignItems(ui.Center).BorderWidth(1).BorderColor(border).Radius(8).
						Background(c.Theme().Surface).Children(func() {
						ui.Spacer(c)
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
				}).OnClick(func() {
					p.onPileClick(idx)
					c.Invalidate()
				})
				if i < len(piles)-1 {
					ui.Spacer(c).Width(20)
				}
			}
		})
	})
}

// onPileClick plays at pile i, whether the pointer or the keyboard's
// Enter or Space put it there. The game picks up on one call and drops on
// the next, so no selection state is kept here.
func (p *HanoiPanel) onPileClick(i int) {
	if p.game.IsDone() {
		return
	}
	p.cursor = i
	p.msg = ""
	if err := p.game.SelectPileByIndex(i); err != nil {
		p.msg = err.Error()
	}
	if p.game.IsDone() {
		p.msg = "Solved!"
	}
}

// handleKeys takes the panel's keys: Left and Right move the cursor over
// the piles, Enter or Space plays there, and 1 to 3 jump to a pile and
// play.
func (p *HanoiPanel) handleKeys(board ui.Element) {
	if board.Shortcut(0, ui.KeyLeft) {
		p.cursor = (p.cursor + hanoi.PileCount - 1) % hanoi.PileCount
	}
	if board.Shortcut(0, ui.KeyRight) {
		p.cursor = (p.cursor + 1) % hanoi.PileCount
	}
	if board.Shortcut(0, ui.KeyEnter) || board.Shortcut(0, ui.KeySpace) {
		p.onPileClick(p.cursor)
	}
	for d := range hanoi.PileCount {
		if board.Shortcut(0, ui.Key1+ui.Key(d)) {
			p.onPileClick(d)
		}
	}
}
