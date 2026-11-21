package view

import (
	"fmt"

	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/choice"
	"github.com/ELY-Open-Source/CrychicUI/sound"
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/editor"
	"github.com/zrcoder/ttoy/game/hanoi"
	"github.com/zrcoder/ttoy/game/hanoi/xgo"
)

// 3-disk solution; also solves the default 2-disk level (moves after IsDone are ignored).
const sampleHanoiCode = `// Move disks with A, B or C.
// The first call picks a disk up, the next drops it.
A
B
A
C
B
C
`

// Hanoi renders a playable Towers of Hanoi puzzle.
type Hanoi struct {
	game             *hanoi.Game
	solvedDialogOpen bool

	level       int
	levelLabels []string
	diskColors  []ui.Color

	codeBuf *editor.State
	runErr  string

	timeline hanoiTimeline

	lastFrameIdx int // highest frame whose sound fired; -1 = none shown
	lastBoard    board
	solvedShown  bool
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
	p.codeBuf = &editor.State{Buffer: editor.NewBuffer(sampleHanoiCode)}
	p.codeBuf.Buffer.SetLanguage(editor.LanguageNamed("go"))
	hanoi.Bind(p.game)
	p.game.SetOnMove(p.onMove)
	p.resetView()
	return p
}

func (p *Hanoi) View(c *ui.Context) {
	board, nextAt, playing, idx := p.timeline.present(c.Now(), p.game)
	if idx < 0 {
		board = snapshotBoard(p.game)
	}
	if idx > p.lastFrameIdx {
		if p.lastFrameIdx >= 0 {
			p.onFrameTransition(p.lastBoard, board)
		}
		p.lastBoard = board
		p.lastFrameIdx = idx
	}
	if playing {
		if d := nextAt.Sub(c.Now()); d > 0 {
			c.After(d)
		} else {
			c.AnimationFrame()
		}
	}

	if p.solvedDialogOpen {
		ui.Modal(c, &p.solvedDialogOpen, func() {
			ui.Text(c, "Solved!").FontSize(18).FontWeight(700).TextColor(c.Theme().Success)
		})
	}

	ui.Column(c).Fill().Children(func() {
		ui.Row(c).Height(44).Children(func() {
			ui.Text(c, "Hanoi").FontSize(20).FontWeight(700)
		})

		ui.Column(c).Fill().AlignItems(ui.Center).Children(func() {
			ui.Row(c).Margin(18, 0, 18, 0).Children(func() {
				choice.Segments(c, "Levels", p.levelLabels, &p.level).OnChange(func() {
					p.game.SelectLevel(p.level)
					p.resetView()
				})
				button.Solid(c, "Reset").OnClick(func() {
					p.game.Reset()
					p.resetView()
					c.Invalidate()
				})
			})
			state := fmt.Sprintf("Steps: %d, Minimum: %d", board.steps, p.game.MinSteps())
			ui.Text(c, state).TextColor(c.Theme().TextMuted).Margin(30, 0)
			ui.Column(c).FillWidth().Margin(10, 0).AlignItems(ui.Center).Children(func() {
				ui.Row(c).Children(func() {
					for i := range 3 {
						p.PileView(c, board, i)
					}
				})
			})

			p.codePanel(c)
		})

	})

	if p.runErr != "" {
		Toast(c, p.runErr, func() { p.runErr = "" })
	}
}

func (p *Hanoi) PileView(c *ui.Context, board board, index int) {
	ui.Column(c).AlignItems(ui.Center).Children(func() {
		if board.shiftFrom == index && board.shiftDisk >= 0 {
			p.diskView(c, board.shiftDisk)
		} else {
			ui.Row(c).Height(20)
		}

		disks := board.piles[index]
		col := ui.Column(c).Width(180).AlignItems(ui.Center).Margin(18).Children(func() {
			ui.Column(c).Width(110).Height(130).AlignItems(ui.Center).Children(func() {
				ui.Spacer(c)
				maxDisks := p.game.CurrentLevel().Disks
				for j := range maxDisks {
					idx := maxDisks - 1 - j
					if idx < len(disks) {
						ui.Row(c).Height(24).Children(func() {
							p.diskView(c, disks[idx])
						})
					}
				}
			})
			ui.Row(c).Height(24).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
				ui.Text(c, string(rune('A'+index))).FontSize(14).FontWeight(700).TextColor(c.Theme().Text)
			})
		})

		if col.Clicked() {
			p.onPileClick(c, index)
		}
		col.Border(2, c.Theme().Border).Radius(10)
		if index < 2 {
			col.BorderStyle(ui.BorderDashed)
		}
	})
}

func (p *Hanoi) diskView(c *ui.Context, diskID int) {
	w := float32(diskID*20) + 40
	color := p.diskColors[diskID%len(p.diskColors)]
	ui.Column(c).Width(w).Height(20).Radius(4).Background(color)
}

func (p *Hanoi) onPileClick(c *ui.Context, i int) {
	p.timeline.ensureBase(p.game)
	_ = p.game.SelectPileByIndex(i)
	c.Invalidate()
}

func (p *Hanoi) onMove() {
	p.timeline.push(snapshotBoard(p.game))
}

func (p *Hanoi) onFrameTransition(prev, cur board) {
	switch {
	case prev.shiftDisk < 0 && cur.shiftDisk >= 0:
		sound.Play(sound.Select)
	case prev.shiftDisk >= 0 && cur.shiftDisk < 0:
		sound.Play(sound.Pop)
		if !p.solvedShown && p.isBoardSolved(cur) {
			sound.Play(sound.Success)
			p.solvedDialogOpen = true
			p.solvedShown = true
		}
	default:
		sound.Play(sound.Tick)
	}
}

func (p *Hanoi) isBoardSolved(b board) bool {
	return len(b.piles[2]) == p.game.CurrentLevel().Disks
}

func (p *Hanoi) resetView() {
	p.timeline.reset(snapshotBoard(p.game))
	p.lastFrameIdx = -1
	p.solvedShown = false
	p.solvedDialogOpen = false
}

func (p *Hanoi) codePanel(c *ui.Context) {
	ui.Column(c).Fill().AlignItems(ui.Center).Children(func() {
		if button.IconButton(c, "Run", "arrow-up").Margin(10).Clicked() {
			p.runCode(c)
		}
		// code.CodeEditor(c, p.codeBuf, code.CodeEditorOptions{}).Fill()
		editor.View(c, p.codeBuf).Fill()
	})
}

func (p *Hanoi) runCode(c *ui.Context) {
	p.game.Reset()
	p.resetView()
	p.runErr = ""
	if err := xgo.RunCode(p.codeBuf.Buffer.String()); err != nil {
		p.runErr = err.Error()
	}
	c.Invalidate()
}
