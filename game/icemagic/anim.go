package icemagic

import "time"

// frameStep is how long one replayed frame stays on screen.
const frameStep = 100 * time.Millisecond

// moveAnim replays pushed board steps over time for the view: the game
// logic runs to the end in one go, while mygo only draws frames after
// an input handler returns, so intermediate states live here.
type moveAnim struct {
	start  time.Time
	frames [][][]string
}

// beginAnim records the board before a move changes it, unless the last
// frame already shows it.
func (g *Game) beginAnim() {
	n := len(g.anim.frames)
	if n == 0 {
		g.anim = moveAnim{start: time.Now(), frames: [][][]string{g.gridImages()}}
		return
	}
	if !sameImages(g.anim.frames[n-1], g.gridImages()) {
		g.anim.frames = append(g.anim.frames, g.gridImages())
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
	g.anim.frames = append(g.anim.frames, g.gridImages())
}

// gridImages copies the image path of each cell of the board.
func (g *Game) gridImages() [][]string {
	images := make([][]string, len(g.grid))
	for y := range g.grid {
		images[y] = make([]string, len(g.grid[y]))
		for x := range g.grid[y] {
			images[y][x] = g.grid[y][x].Cell.Image
		}
	}
	return images
}

// sameImages reports whether two board frames match cell by cell.
func sameImages(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for y := range a {
		if len(a[y]) != len(b[y]) {
			return false
		}
		for x := range a[y] {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}
	return true
}

// Present returns what the board shows at time t, when the next frame
// is due, and whether the replay is still playing.
func (g *Game) Present(t time.Time) ([][]string, time.Time, bool) {
	n := len(g.anim.frames) - 1
	if n >= 0 && !sameImages(g.anim.frames[n], g.gridImages()) {
		// The board changed outside the replay: drop it and go live.
		g.anim = moveAnim{}
		n = -1
	}
	if n < 0 {
		return g.gridImages(), time.Time{}, false
	}
	k := max(int(t.Sub(g.anim.start)/frameStep), 0)
	if k >= n {
		return g.anim.frames[n], time.Time{}, false
	}
	return g.anim.frames[k], g.anim.start.Add(time.Duration(k+1) * frameStep), true
}
