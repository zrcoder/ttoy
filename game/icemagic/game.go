package icemagic

import (
	"math/rand/v2"

	"github.com/zrcoder/ttoy/game/common"
)

type Game struct {
	grid     [][]*Sprite
	player   *Sprite
	failed   bool // failed if the play is burned
	fires    int
	rand     *rand.Rand
	chapter  int
	level    int
	chpaters []int
}

func New() *Game {
	g := &Game{
		chpaters: []int{9, 9, 9, 9, 9, 2},
		rand:     rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}
	g.Reset()
	return g
}

func (g *Game) Grid() [][]*Sprite {
	return g.grid
}

// PlayerX returns the player's column.
func (g *Game) PlayerX() int {
	if g.player == nil {
		return -1
	}
	return g.player.x
}

// PlayerY returns the player's row.
func (g *Game) PlayerY() int {
	if g.player == nil {
		return -1
	}
	return g.player.y
}

// Kind returns the cell type at (x, y), or Blank if out of bounds.
func (g *Game) KindAt(x, y int) byte {
	if y < 0 || y >= len(g.grid) {
		return Blank
	}
	row := g.grid[y]
	if x < 0 || x >= len(row) {
		return Blank
	}
	if row[x] == nil {
		return Blank
	}
	return row[x].kind
}

// State returns the current game state.
func (g *Game) State() common.State {
	switch {
	case g.failed:
		return common.StateFailed
	case g.fires == 0:
		return common.StateSucceed
	default:
		return common.StatePlaying
	}
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
		if ok := g.swap(left, player); !ok {
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
		if ok := g.swap(player, right); !ok {
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

func (g *Game) swap(src, dst *Sprite) bool {
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
