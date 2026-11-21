package view

import (
	"slices"
	"time"

	"github.com/zrcoder/ttoy/game/icemagic"
)

// mygo draws after an input handler returns, so cascading mutations are
// replayed from snapshots recorded via SetOnMove.
const iceFrameStep = 100 * time.Millisecond

type iceTimeline struct {
	start  time.Time
	frames [][][]*icemagic.Cell
}

func (t *iceTimeline) reset(base [][]*icemagic.Cell) {
	t.start = time.Now()
	t.frames = [][][]*icemagic.Cell{base}
}

// ensureBase stores the live grid only when it drifted from the last frame.
func (t *iceTimeline) ensureBase(live [][]*icemagic.Sprite) {
	n := len(t.frames)
	if n == 0 {
		t.reset(snapshotCells(live))
		return
	}
	if !iceSameCellsGrid(t.frames[n-1], live) {
		t.push(snapshotCells(live))
	}
}

// push skips frames identical to the last.
func (t *iceTimeline) push(live [][]*icemagic.Cell) {
	n := len(t.frames)
	if n > 0 && iceSameCells(t.frames[n-1], live) {
		return
	}
	now := time.Now()
	if n > 0 && now.Sub(t.start) >= time.Duration(n)*iceFrameStep {
		t.start = now.Add(-time.Duration(n) * iceFrameStep)
	}
	t.frames = append(t.frames, live)
}

func iceSameCells(a, b [][]*icemagic.Cell) bool {
	rowEq := func(x, y []*icemagic.Cell) bool {
		return slices.EqualFunc(x, y, func(i, j *icemagic.Cell) bool { return *i == *j })
	}
	return slices.EqualFunc(a, b, rowEq)
}

// iceSameCellsGrid compares the live grid by reference, no copy.
func iceSameCellsGrid(a [][]*icemagic.Cell, b [][]*icemagic.Sprite) bool {
	if len(a) != len(b) {
		return false
	}
	for y := range a {
		if len(a[y]) != len(b[y]) {
			return false
		}
		for x := range a[y] {
			if *a[y][x] != b[y][x].Cell {
				return false
			}
		}
	}
	return true
}

// present returns nil when there is nothing to replay; caller renders live.
func (t *iceTimeline) present(now time.Time, live [][]*icemagic.Sprite) ([][]*icemagic.Cell, time.Time, bool) {
	n := len(t.frames) - 1
	if n >= 0 && !iceSameCellsGrid(t.frames[n], live) {
		t.frames = nil
		n = -1
	}
	if n < 0 {
		return nil, time.Time{}, false
	}
	k := max(int(now.Sub(t.start)/iceFrameStep), 0)
	if k >= n {
		return t.frames[n], time.Time{}, false
	}
	return t.frames[k], t.start.Add(time.Duration(k+1) * iceFrameStep), true
}

// snapshotCells copies only when storing a frame; live rendering reads by reference.
func snapshotCells(grid [][]*icemagic.Sprite) [][]*icemagic.Cell {
	cells := make([][]*icemagic.Cell, len(grid))
	for y := range grid {
		cells[y] = make([]*icemagic.Cell, len(grid[y]))
		for x := range grid[y] {
			c := grid[y][x].Cell
			cells[y][x] = &c
		}
	}
	return cells
}
