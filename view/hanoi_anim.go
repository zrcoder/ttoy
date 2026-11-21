package view

import (
	"time"

	"github.com/zrcoder/ttoy/game/hanoi"
)

// mygo draws after an input handler returns, so cascading moves are replayed
// from snapshots recorded via SetOnMove.
const hanoiFrameStep = 300 * time.Millisecond

type board struct {
	piles     [3][]int // disk IDs, bottom to top
	shiftDisk int      // held disk ID, -1 if none
	shiftFrom int      // source pile of the held disk, -1 if none
	steps     int
}

type hanoiTimeline struct {
	start  time.Time
	frames []board
}

func (t *hanoiTimeline) reset(b board) {
	t.start = time.Now()
	t.frames = []board{b}
}

// ensureBase stores the live board only when it drifted from the last frame.
func (t *hanoiTimeline) ensureBase(g *hanoi.Game) {
	n := len(t.frames)
	if n == 0 {
		t.reset(snapshotBoard(g))
		return
	}
	if !frameMatchesLive(t.frames[n-1], g) {
		t.push(snapshotBoard(g))
	}
}

func (t *hanoiTimeline) push(b board) {
	n := len(t.frames)
	now := time.Now()
	if n > 0 && now.Sub(t.start) >= time.Duration(n)*hanoiFrameStep {
		t.start = now.Add(-time.Duration(n) * hanoiFrameStep)
	}
	t.frames = append(t.frames, b)
}

// present returns idx -1 when there is nothing to replay.
func (t *hanoiTimeline) present(now time.Time, g *hanoi.Game) (board, time.Time, bool, int) {
	n := len(t.frames) - 1
	if n >= 0 && !frameMatchesLive(t.frames[n], g) {
		t.frames = nil
		n = -1
	}
	if n < 0 {
		return board{}, time.Time{}, false, -1
	}
	k := max(int(now.Sub(t.start)/hanoiFrameStep), 0)
	if k >= n {
		return t.frames[n], time.Time{}, false, n
	}
	return t.frames[k], t.start.Add(time.Duration(k+1) * hanoiFrameStep), true, k
}

// frameMatchesLive compares the game by reference, no copy.
func frameMatchesLive(b board, g *hanoi.Game) bool {
	live := [3][]*hanoi.Disk{g.PileA.Disks, g.PileB.Disks, g.PileC.Disks}
	for i := range 3 {
		if len(b.piles[i]) != len(live[i]) {
			return false
		}
		for j := range b.piles[i] {
			if b.piles[i][j] != live[i][j].ID {
				return false
			}
		}
	}
	switch {
	case b.shiftDisk >= 0:
		if g.ShiftDisk == nil || b.shiftDisk != g.ShiftDisk.ID || b.shiftFrom != g.ShiftPileIndex() {
			return false
		}
	case g.ShiftDisk != nil:
		return false
	}
	return b.steps == g.Steps()
}

// snapshotBoard copies only when storing a frame; live rendering reads by reference.
func snapshotBoard(g *hanoi.Game) board {
	piles := g.Piles()
	b := board{shiftDisk: -1, shiftFrom: -1, steps: g.Steps()}
	for i := range 3 {
		b.piles[i] = piles[i].Disks
	}
	if d := g.ShiftDisk; d != nil {
		b.shiftDisk = d.ID
		b.shiftFrom = g.ShiftPileIndex()
	}
	return b
}
