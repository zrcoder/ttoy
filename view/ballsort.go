package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/ballsort"
)

// BallSort renders a playable Ball Sort puzzle using native mygo UI.
type BallSort struct {
	game             *ballsort.Game
	levels           []string
	level            int
	solvedDialogOpen bool
}

func NewBallSort() *BallSort {
	b := new(BallSort)
	b.game = ballsort.New()
	b.level = b.game.LevelIndex()
	b.levels = make([]string, len(b.game.Levels()))
	for i, level := range b.game.Levels() {
		b.levels[i] = level.Label
	}
	return b
}

func (p *BallSort) View(c *ui.Context) {
	bottles := p.game.GetBottles()
	// Covering Reset and level switches: a new game has no dialog.
	if p.game.IsDone() {
		ui.Modal(c, &p.solvedDialogOpen, func() {
			ui.Text(c, "Solved!").FontSize(18).FontWeight(700).TextColor(c.Theme().Success)
		})
	} else {
		p.solvedDialogOpen = false
	}
	for i := range bottles {
		if c.Shortcut(0, ui.KeyA+ui.Key(i)) {
			p.onBottleClick(c, i)
			break
		}
	}
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).Height(44).Children(func() {
			ui.Text(c, "Ball Sort").FontSize(20).FontWeight(700)
		})
		// Level selector
		ui.Row(c).AlignItems(ui.Center).Margin(18, 0, 18, 0).Children(func() {
			ui.Segmented(c, &p.level, p.levels...).OnChange(func() {
				p.game.SelectLevel(p.level)
			})
			ui.Spacer(c).Width(16)

			ui.Button(c, "Reset").OnClick(func() {
				p.game.Reset()
				c.Invalidate()
			})

		})
		ui.Column(c).AlignItems(ui.Center).Children(func() {
			ui.Text(c, p.game.State())
			// Bottles
			ui.Row(c).Children(func() {
				for _, bottle := range bottles {
					p.bottleView(c, bottle)
				}
			})
		})
	})
}

func (p *BallSort) bottleView(c *ui.Context, b *ballsort.Bottle) {
	ui.Column(c).AlignItems(ui.Center).Margin(12).Children(func() {
		if b.HasShiftBall() {
			p.ballView(c, b.ShiftBall)
		} else {
			p.ballView(c, nil)
		}
		borderColor := ui.RGB(190, 190, 190)
		if b.IsDone() {
			borderColor = c.Theme().Success
		}
		ui.Column(c).Width(70).Height(168).BorderWidth(0, 2, 2, 2).BorderColor(borderColor).Radius(8).AlignItems(ui.Center).Children(func() {
			for j := ballsort.BottleBallCount - 1; j >= 0; j-- {
				ui.Spacer(c).Height(2)
				if j < len(b.Balls) {
					p.ballView(c, b.Balls[j])
				} else {
					p.ballView(c, nil)
				}
			}
		})
		ui.Row(c).Height(24).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
			ui.Text(c, string(rune('A'+b.Index))).FontSize(14).FontWeight(700).TextColor(c.Theme().Text)
		})
	}).OnClick(func() {
		p.onBottleClick(c, b.Index)
	})
}

func (p *BallSort) ballView(c *ui.Context, ball *ballsort.Ball) {
	color := ui.Transparent
	if ball != nil {
		color = ball.Color
	}
	ui.Box(c).Width(40).Height(40).Radius(20).Background(color)
}

func (p *BallSort) onBottleClick(c *ui.Context, i int) {
	if p.game.IsDone() {
		return
	}
	p.game.SelectBottle(i)
	p.solvedDialogOpen = p.game.IsDone()
	c.Invalidate()
}
