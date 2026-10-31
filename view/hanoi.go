package view

import (
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/hanoi"
)

// Hanoi renders a playable Towers of Hanoi puzzle.
type Hanoi struct {
	game *hanoi.Game

	levelLabels []string
	diskColors  []ui.Color
}

func NewHanoi() *Hanoi {
	p := new(Hanoi)
	p.game = hanoi.New()
	p.diskColors = []ui.Color{
		{R: 255, G: 0, B: 0, A: 255},
		{R: 0, G: 128, B: 0, A: 255},
		{R: 0, G: 0, B: 255, A: 255},
		{R: 255, G: 255, B: 0, A: 255},
		{R: 165, G: 42, B: 42, A: 255},
		{R: 255, G: 192, B: 203, A: 255},
	}
	levels := p.game.Levels()
	p.levelLabels = make([]string, len(levels))
	for i := range p.levelLabels {
		p.levelLabels[i] = levels[i].Label
	}
	return p
}

func (p *Hanoi) View(c *ui.Context) {
	switch {
	case c.Shortcut(0, ui.KeyA):
		p.onPileClick(c, 0)
	case c.Shortcut(0, ui.KeyB):
		p.onPileClick(c, 1)
	case c.Shortcut(0, ui.KeyC):
		p.onPileClick(c, 2)
	}
	ui.Column(c).Fill().Children(func() {
		ui.Text(c, "Hanoi").FontSize(20).FontWeight(700)

		ui.Row(c).FillWidth().AlignItems(ui.Stretch).Margin(18, 0, 18, 0).Children(func() {
			level := -1
			ui.Segmented(c, &level, p.levelLabels...).OnChange(func() {
				p.game.SelectLevel(level)
			})

			ui.Spacer(c)

			ui.Button(c, "Reset").OnClick(func() {
				p.game.Reset()
				c.Invalidate()
			})
		})

		ui.Column(c).FillWidth().AlignItems(ui.Center).Children(func() {
			if p.game.IsDone() {
				ui.Text(c, "Solved!").FontWeight(800).TextColor(c.Theme().Success)
			} else {
				ui.Text(c, p.game.State()).TextColor(c.Theme().TextMuted)
			}

			ui.Row(c).Children(func() {
				piles := p.game.Piles()
				for _, pile := range piles {
					p.PileView(c, &pile)
				}
			})
		})
	})
}

func (p *Hanoi) PileView(c *ui.Context, pile *hanoi.PileInfo) {
	ui.Column(c).AlignItems(ui.Center).Margin(38).Children(func() {
		ui.Column(c).Width(140).Height(260).AlignItems(ui.Center).Children(func() {
			if p.game.ShiftPileIndex() == pile.Index {
				p.diskView(c, p.game.ShiftDisk.ID)
			}

			ui.Spacer(c)

			// Disks (bottom-up), stored as IDs
			maxDisks := p.game.CurrentLevel().Disks
			for j := range maxDisks {
				idx := maxDisks - 1 - j
				if idx < len(pile.Disks) {
					ui.Row(c).Height(24).Children(func() {
						p.diskView(c, pile.Disks[idx])
					})
				}
			}
		})

		ui.Row(c).Height(24).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
			ui.Text(c, string(rune('A'+pile.Index))).FontSize(14).FontWeight(700).TextColor(c.Theme().Text)
		})
	}).OnClick(func() {
		p.onPileClick(c, pile.Index)
	})
}

func (p *Hanoi) diskView(c *ui.Context, diskID int) {
	w := float32(diskID*20) + 40
	color := p.diskColors[diskID%len(p.diskColors)]
	ui.Column(c).Width(w).Height(20).Radius(4).Background(color)
}

// onPileClick plays at pile i, whether the pointer or the keyboard's
// Enter or Space put it there. The game picks up on one call and drops on
// the next, so no selection state is kept here.
func (p *Hanoi) onPileClick(c *ui.Context, i int) {
	if p.game.IsDone() {
		return
	}
	_ = p.game.SelectPileByIndex(i)
	c.Invalidate()
}
