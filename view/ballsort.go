package view

import (
	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/choice"
	"github.com/ELY-Open-Source/CrychicUI/sound"
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/ballsort"
)

// BallSort renders a playable Ball Sort puzzle using native mygo UI.
type BallSort struct {
	game             *ballsort.Game
	levels           []string
	level            int
	solvedDialogOpen bool

	// ballTex paints the palette's black and white, which would else
	// match the background.
	ballTex map[ui.Color]*ui.Bitmap
}

func NewBallSort() *BallSort {
	b := new(BallSort)
	b.game = ballsort.New()
	b.ballTex = map[ui.Color]*ui.Bitmap{
		ballsort.BlackBall: ui.NewBitmap(ballTexture(ballsort.BlackBall)),
		ballsort.WhiteBall: ui.NewBitmap(ballTexture(ballsort.WhiteBall)),
	}
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
			ui.Text(c, "Solved!").Focus().FontSize(18).FontWeight(700).TextColor(c.Theme().Success)
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
			choice.Segments(c, "Levels", p.levels, &p.level).OnChange(func() {
				p.game.SelectLevel(p.level)
			})
			ui.Spacer(c).Width(16)

			button.Outline(c, "Reset").OnClick(func() {
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
		borderWidth := []float32{0, 2, 2, 2}
		if b.IsDone() {
			borderColor = c.Theme().Success
			borderWidth = []float32{2, 2, 2, 2}
		}
		ui.Column(c).Width(70).Height(168).BorderWidth(borderWidth...).BorderColor(borderColor).Radius(8).AlignItems(ui.Center).Children(func() {
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
		if bmp := p.ballTex[ball.Color]; bmp != nil {
			ui.Image(c, bmp).Width(40).Height(40)
			return
		}
		color = ball.Color
	}
	ui.Box(c).Width(40).Height(40).Radius(20).Background(color)
}

func (p *BallSort) onBottleClick(c *ui.Context, i int) {
	if p.game.IsDone() {
		return
	}
	lift := p.game.ShiftBall
	p.game.SelectBottle(i)
	// One cue per kind of move, told apart by how ShiftBall moved.
	switch now := p.game.ShiftBall; {
	case lift == nil && now != nil:
		sound.Play(sound.Select)
	case lift != nil && now == nil:
		sound.Play(sound.Pop)
	case lift != nil:
		sound.Play(sound.Swipe)
	default:
		sound.Play(sound.Tick)
	}
	if !p.solvedDialogOpen && p.game.IsDone() {
		sound.Play(sound.Success)
	}
	p.solvedDialogOpen = p.game.IsDone()
	c.Invalidate()
}
