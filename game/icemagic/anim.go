package icemagic

import (
	"slices"
	"time"
)

// frameStep is how long one replayed frame stays on screen.
const frameStep = 100 * time.Millisecond

// moveAnim replays pushed board steps over time for the view: the game
// logic runs to the end in one go, while mygo only draws frames after
// an input handler returns, so intermediate states live here.
type moveAnim struct {
	start  time.Time
	frames [][][]*Cell
}

// beginAnim records the board before a move changes it, unless the last
// frame already shows it.
func (g *Game) beginAnim() {
	n := len(g.anim.frames)
	if n == 0 {
		g.anim = moveAnim{start: time.Now(), frames: [][][]*Cell{g.gridCells()}}
		return
	}
	if !sameCells(g.anim.frames[n-1], g.gridCells()) {
		g.anim.frames = append(g.anim.frames, g.gridCells())
	}
}

// frame records the board as the replay's next frame.
func (g *Game) frame() {
	n := len(g.anim.frames)
	now := time.Now()
	if n > 0 && now.Sub(g.anim.start) >= time.Duration(n)*frameStep {
		// The replay has ended: let the new frames start from now.
		g.anim.start = now.Add(-time.Duration(n) * frameStep)
	}
	g.anim.frames = append(g.anim.frames, g.gridCells())
}

// gridCells copies each cell of the board by value; sprites mutate in
// place, so a snapshot can only keep Cell copies, behind a pointer.
func (g *Game) gridCells() [][]*Cell {
	cells := make([][]*Cell, len(g.grid))
	for y := range g.grid {
		cells[y] = make([]*Cell, len(g.grid[y]))
		for x := range g.grid[y] {
			c := g.grid[y][x].Cell
			cells[y][x] = &c
		}
	}
	return cells
}

// sameCells reports whether two board frames match cell by cell.
func sameCells(a, b [][]*Cell) bool {
	rowEq := func(x, y []*Cell) bool {
		return slices.EqualFunc(x, y, func(i, j *Cell) bool { return *i == *j })
	}
	return slices.EqualFunc(a, b, rowEq)
}

// Present returns what the board shows at time t, when the next frame
// is due, and whether the replay is still playing.
func (g *Game) Present(t time.Time) ([][]*Cell, time.Time, bool) {
	n := len(g.anim.frames) - 1
	if n >= 0 && !sameCells(g.anim.frames[n], g.gridCells()) {
		// The board changed outside the replay: drop it and go live.
		g.anim = moveAnim{}
		n = -1
	}
	if n < 0 {
		return g.gridCells(), time.Time{}, false
	}
	k := max(int(t.Sub(g.anim.start)/frameStep), 0)
	if k >= n {
		return g.anim.frames[n], time.Time{}, false
	}
	return g.anim.frames[k], g.anim.start.Add(time.Duration(k+1) * frameStep), true
}
