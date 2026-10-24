package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/ballsort"
)

// BallSortPanel renders a playable Ball Sort puzzle using native mygo UI.
type BallSortPanel struct {
	game     *ballsort.Game
	selected int // index of selected bottle, -1 = none
	msg      string
	levelSel int
	levels   []ballsort.Level
}

func NewBallSortPanel() *BallSortPanel {
	g := ballsort.New()
	return &BallSortPanel{
		game:     g,
		selected: -1,
		levels:   g.Levels(),
	}
}

func (p *BallSortPanel) View(c *ui.Context) {
	// Refresh levels in case they changed
	p.levels = p.game.Levels()

	ui.Column(c).Fill().Children(func() {
		ui.Row(c).Fill().Height(44).Children(func() {
			ui.Text(c, "Ball Sort").FontSize(20).FontWeight(700)
			ui.Spacer(c)
			ui.Text(c, p.msg).FontSize(13).TextColor(c.Theme().TextMuted)
			ui.Spacer(c).Width(16)
			ui.ButtonBase(c).Width(70).Height(28).Radius(4).Background(c.Theme().Accent).OnClick(func() {
				p.game = ballsort.New()
				p.selected = -1
				p.msg = ""
				p.levelSel = 0
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

		ui.Scroll(c).Fill().Children(func() {
			ui.Column(c).AlignItems(ui.Center).Children(func() {
				// Level selector
				ui.Row(c).AlignItems(ui.Center).Margin(0, 0, 12, 0).Children(func() {
					ui.Text(c, "Level").FontSize(13)
					for i, lv := range p.levels {
						idx := i
						active := idx == p.levelSel
						bg := ui.Color{}
						if active {
							bg = c.Theme().Accent
						}
						ui.ButtonBase(c).Width(70).Height(26).Radius(4).Background(bg).OnClick(func() {
							p.levelSel = idx
							p.game.SelectLevel(idx)
							p.selected = -1
							p.msg = ""
							c.Invalidate()
						}).Children(func() {
							ui.Text(c, lv.Label).TextColor(c.Theme().Text).FontSize(12)
						})
						if i < len(p.levels)-1 {
							ui.Spacer(c).Width(6)
						}
					}
					ui.Spacer(c).Width(16)
					ui.Text(c, p.game.State()).FontSize(13).TextColor(c.Theme().TextMuted)
				})

				// Bottles
				ui.Row(c).AlignItems(ui.Center).Children(func() {
					bottles := p.game.GetBottles()
					for i, bottle := range bottles {
						p.bottle(c, i, bottle)
						if i < len(bottles)-1 {
							ui.Spacer(c).Width(10)
						}
					}
				})
			})
		})
	})
}

func (p *BallSortPanel) bottle(c *ui.Context, i int, b *ballsort.Bottle) {
	active := p.selected == i
	done := b.IsDone()
	border := c.Theme().Border
	if active {
		border = c.Theme().Accent
	} else if done {
		border = c.Theme().Success
	}
	ui.Column(c.Key(fmt.Sprintf("bottle-%d", i))).AlignItems(ui.Center).OnClick(func() {
		p.onBottleClick(i)
		c.Invalidate()
	}).Children(func() {
		ui.Column(c).Width(70).Height(190).BorderWidth(2).BorderColor(border).Radius(8).
			Background(c.Theme().Surface).Children(func() {
			balls := b.Balls
			n := len(balls)
			for j := range ballsort.BottleBallCount {
				idx := n - 1 - j
				ui.Row(c).Height(34).Children(func() {
					if idx >= 0 && idx < n && balls[idx] != nil {
						col := parseColor(balls[idx].Color)
						ui.Column(c).Width(40).Height(40).Radius(20).Background(col)
					} else {
						ui.Column(c).Width(40).Height(40).Radius(20).Background(ui.Color{R: 240, G: 240, B: 240, A: 60})
					}
				})
			}
			ui.Row(c).Height(24).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
				ui.Text(c, string(rune('A'+i))).FontSize(14).FontWeight(700).TextColor(c.Theme().Text)
			})
		})
	})
}

func (p *BallSortPanel) onBottleClick(i int) {
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
	src := p.selected
	p.game.SelectBottle(i)
	p.selected = -1
	if p.game.IsDone() {
		p.msg = "Solved!"
	}
	_ = src
}

func parseColor(s string) ui.Color {
	s = strings.ToLower(strings.TrimSpace(s))
	switch s {
	case "red":
		return ui.Color{R: 255, G: 0, B: 0, A: 255}
	case "green":
		return ui.Color{R: 0, G: 128, B: 0, A: 255}
	case "blue":
		return ui.Color{R: 0, G: 0, B: 255, A: 255}
	case "yellow":
		return ui.Color{R: 255, G: 255, B: 0, A: 255}
	case "violet":
		return ui.Color{R: 238, G: 130, B: 238, A: 255}
	case "pink":
		return ui.Color{R: 255, G: 192, B: 203, A: 255}
	case "orange":
		return ui.Color{R: 255, G: 165, B: 0, A: 255}
	case "#7fff00":
		return ui.Color{R: 127, G: 255, B: 0, A: 255}
	default:
		if strings.HasPrefix(s, "#") {
			hex := s[1:]
			if len(hex) == 6 {
				r, _ := strconv.ParseUint(hex[0:2], 16, 8)
				g, _ := strconv.ParseUint(hex[2:4], 16, 8)
				b, _ := strconv.ParseUint(hex[4:6], 16, 8)
				return ui.Color{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
			}
		}
		return ui.Color{R: 128, G: 128, B: 128, A: 255}
	}
}
