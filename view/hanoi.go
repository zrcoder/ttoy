package view

import (
	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/choice"
	"github.com/ELY-Open-Source/CrychicUI/sound"
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/hanoi"
)

// Hanoi renders a playable Towers of Hanoi puzzle.
type Hanoi struct {
	game             *hanoi.Game
	solvedDialogOpen bool

	level       int
	levelLabels []string
	diskColors  []ui.Color
}

func NewHanoi() *Hanoi {
	p := new(Hanoi)
	p.game = hanoi.New()
	p.level = p.game.LevelIndex()
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
	// Covering Reset and level switches: a new game has no dialog.
	if p.game.IsDone() {
		ui.Modal(c, &p.solvedDialogOpen, func() {
			ui.Text(c, "Solved!").FontSize(18).FontWeight(700).TextColor(c.Theme().Success)
		})
	} else {
		p.solvedDialogOpen = false
	}
	switch {
	case c.Shortcut(0, ui.KeyA):
		p.onPileClick(c, 0)
	case c.Shortcut(0, ui.KeyB):
		p.onPileClick(c, 1)
	case c.Shortcut(0, ui.KeyC):
		p.onPileClick(c, 2)
	}
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).Height(44).Children(func() {
			ui.Text(c, "Hanoi").FontSize(20).FontWeight(700)
		})

		ui.Column(c).Fill().AlignItems(ui.Center).Children(func() {
			ui.Row(c).Margin(18, 0, 18, 0).Children(func() {
				choice.Segments(c, "Levels", p.levelLabels, &p.level).OnChange(func() {
					p.game.SelectLevel(p.level)
				})
				ui.Spacer(c).Width(18)
				button.Solid(c, "Reset").OnClick(func() {
					p.game.Reset()
					c.Invalidate()
				})
			})
			ui.Column(c).FillWidth().AlignItems(ui.Center).Children(func() {
				ui.Text(c, p.game.State()).TextColor(c.Theme().TextMuted).Margin(18, 0, 0, 0)
				ui.Row(c).Children(func() {
					piles := p.game.Piles()
					for _, pile := range piles {
						p.PileView(c, &pile)
					}
				})
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
	lift := p.game.ShiftPileIndex()
	_ = p.game.SelectPileByIndex(i)
	// One cue per kind of move, told apart by ShiftPileIndex before and
	// after: -1 while no disk is held.
	switch hold := p.game.ShiftPileIndex(); {
	case lift == -1 && hold != -1:
		sound.Play(sound.Select)
	case lift != -1 && hold == -1:
		sound.Play(sound.Pop)
	default:
		sound.Play(sound.Tick) // an invalid move, or nothing to lift
	}
	if !p.solvedDialogOpen && p.game.IsDone() {
		sound.Play(sound.Success)
	}
	p.solvedDialogOpen = p.game.IsDone()
	c.Invalidate()
}
