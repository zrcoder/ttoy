package icemagic

import (
	"math/rand"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Game struct {
	app      *application.App
	grid     [][]*Sprite
	player   *Sprite
	failed   bool // failed if the play is burned
	fires    int
	rd       *rand.Rand
	chapter  int
	level    int
	chpaters []int
}

func New(app *application.App) *Game {
	g := &Game{
		app:      app,
		chpaters: []int{9, 9, 9, 9, 9, 2},
		rd:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	g.Reset()
	return g
}

func (g *Game) Grid() [][]*Sprite {
	return g.grid
}

func (g *Game) Chapters() []int {
	return g.chpaters
}

func (g *Game) MoveLeft() bool {
	player := g.player
	left := player.left()
	if left == nil {
		return false
	}
	switch left.kind {
	case Blank:
		up := player.up()
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
		player.playerDie()
		g.updateUI()
	case Wall:
		return g.player.climbLeft()
	default:
	}
	return false
}

func (g *Game) MoveRight() bool {
	player := g.player
	right := player.right()
	switch right.kind {
	case Blank:
		up := player.up()
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
		player.playerDie()
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

func (g *Game) Reset() {
	g.fires = 0
	g.failed = false
	g.parseGrid(g.chapter, g.level)
}

func (g *Game) SelectLevel(chapter, level int) {
	g.chapter = chapter - 1
	g.level = level - 1
	g.Reset()
}

func (g *Game) swap(src, dst *Sprite, duration time.Duration) bool {
	if !g.swapQuietly(src, dst) {
		return false
	}
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
