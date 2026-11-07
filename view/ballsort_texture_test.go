package view

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/ballsort"
)

func TestBallTextures(t *testing.T) {
	p := NewBallSort()
	if len(p.ballTex) != 2 || p.ballTex[ballsort.BlackBall] == nil || p.ballTex[ballsort.WhiteBall] == nil {
		t.Fatalf("the panel's textures are %v, want exactly black and white", p.ballTex)
	}
	for _, color := range liveColors(t) {
		if color == ballsort.BlackBall || color == ballsort.WhiteBall {
			continue
		}
		if p.ballTex[color] != nil {
			t.Errorf("color %v unexpectedly textured", color)
		}
	}
	for _, c := range []ui.Color{ballsort.BlackBall, ballsort.WhiteBall} {
		img := ballTexture(c)
		b := img.Bounds()
		if b.Dx() != int(ballSize) || b.Dy() != int(ballSize) {
			t.Errorf("color %v: texture is %dx%d, want %dx%d", c, b.Dx(), b.Dy(), int(ballSize), int(ballSize))
		}
		i := img.PixOffset(b.Dx()/2, b.Dy()/2)
		if img.Pix[i+3] != 255 {
			t.Errorf("color %v: the ball's center is not opaque", c)
		}
		if img.Pix[3] != 0 {
			t.Errorf("color %v: the disc's corner is not transparent", c)
		}
	}
}

// liveColors collects the palette from a live game.
func liveColors(t *testing.T) []ui.Color {
	t.Helper()
	seen := map[ui.Color]bool{}
	var out []ui.Color
	for _, bottle := range ballsort.New().GetBottles() {
		for _, ball := range bottle.Balls {
			if !seen[ball.Color] {
				seen[ball.Color] = true
				out = append(out, ball.Color)
			}
		}
	}
	return out
}
