package icemagic

import (
	"math/rand"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/zrcoder/ttoy/game/internal"
)

type Game struct {
	base     *internal.Base
	chapters []internal.Chapter
	grid     [][]*Sprite
	player   *Sprite
	failed   bool // failed if the play is burned
	fires    int
	rd       *rand.Rand
}

func New(app *application.App) *Game {
	g := &Game{}
	g.initLevels()
	base := internal.New(
		app,
		internal.WithChapters(g.chapters, g.reset),
	)
	g.base = base
	g.rd = rand.New(rand.NewSource(time.Now().UnixNano()))
	g.reset()
	return g
}

func (g *Game) Grid() [][]*Sprite {
	return g.grid
}

func (g *Game) MoveLeft() bool {
	player := g.player
	left := player.Left()
	if left == nil {
		return false
	}
	switch left.kind {
	case Blank:
		up := player.Up()
		if ok := g.swap(left, player, stepTime); !ok {
			return false
		}
		g.checkFall(up)
		return player.fall()
	case Ice:
		if !left.iceSlideLeft() {
			return g.player.climbLeft()
		}
		return false
	case Fire:
		player.PlayerDie()
		g.updateUI()
	case Wall:
		return g.player.climbLeft()
	default:
	}
	return false
}

func (g *Game) MoveRight() bool {
	player := g.player
	right := player.Right()
	switch right.kind {
	case Blank:
		up := player.Up()
		if ok := g.swap(player, right, stepTime); !ok {
			return false
		}
		g.checkFall(up)
		return player.fall()
	case Ice:
		if !right.iceSlideRight() {
			return g.player.climbRight()
		}
		return false
	case Fire:
		player.PlayerDie()
		g.updateUI()
	case Wall:
		return g.player.climbRight()
	default:
	}
	return false
}

func (g *Game) MagicLeft() {
	g.player.magicLeft()
}

func (g *Game) MagicRight() {
	g.player.magicRight()
}

func (g *Game) reset() {
	chapter, level := g.base.ChapterIndex(), g.base.LevelIndex()
	g.fires = 0
	g.failed = false
	g.parseGrid(chapter, level)
}

func (g *Game) swap(src, dst *Sprite, duration time.Duration) bool {
	if !g.swapQuietly(src, dst) {
		return false
	}
	time.Sleep(duration)
	g.updateUI()
	return true
}

func (g *Game) swapQuietly(src, dst *Sprite) bool {
	if src == nil || dst == nil {
		return false
	}
	sRow := g.grid[src.y]
	dRow := g.grid[dst.y]
	sRow[src.x], dRow[dst.x] = dst, src
	src.x, dst.x = dst.x, src.x
	src.y, dst.y = dst.y, src.y
	src.regularCell()
	dst.regularCell()
	return true
}
