package icemagic

import (
	"strconv"
	"time"

	"github.com/zrcoder/ttoy/game/internal"
)

const (
	Blank    = ' '
	Wall     = '='
	Fire     = 'F'
	Player   = 'M'
	Ice      = 'i'
	IceFixed = 'I'
)

var (
	stepTime = 100 * time.Millisecond
)

type Sprite struct {
	game       *Game
	kind       byte
	x          int
	y          int
	leftFixed  bool
	rightFixed bool
	Cell       internal.Cell
}

func (s *Sprite) iceSlideLeft() bool {
	left := s.Left()
	if left == nil {
		return false
	}
	switch left.kind {
	case Fire:
		left.FireDie()
		s.IceDie()
		time.Sleep(stepTime)
		s.game.updateUI()
		s.game.checkFall(left.Up())
		s.game.checkFall(s.Up())
		return true
	case Blank:
		up := s.Up()
		if ok := s.game.swap(left, s, stepTime); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			s.iceSlideLeft()
		}
		s.game.checkFall(up)
		return true
	}
	return false
}
func (s *Sprite) iceSlideRight() bool {
	right := s.Right()
	if right == nil {
		return false
	}
	switch right.kind {
	case Fire:
		right.FireDie()
		s.IceDie()
		time.Sleep(stepTime)
		s.game.updateUI()
		s.game.checkFall(right.Up())
		s.game.checkFall(s.Up())
		return true
	case Blank:
		up := s.Up()
		if ok := s.game.swap(s, right, stepTime); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			s.iceSlideRight()
		}
		s.game.checkFall(up)
		return true
	}
	return false
}

func (s *Sprite) climbLeft() bool {
	return s.climb(s.LeftUp())
}

func (s *Sprite) climbRight() bool {
	return s.climb(s.RightUp())
}

func (s *Sprite) climb(dst *Sprite) bool {
	up := s.Up()
	if up != nil && up.kind != Blank {
		return false
	}
	if dst == nil || dst.kind != Blank {
		return false
	}
	return s.game.swap(s, dst, 2*stepTime)
}

func (s *Sprite) Left() *Sprite {
	if s.x == 0 {
		return nil
	}
	return s.game.grid[s.y][s.x-1]
}
func (s *Sprite) Right() *Sprite {
	row := s.game.grid[s.y]
	n := len(row)
	if s.x == n-1 {
		return nil
	}
	return row[s.x+1]
}
func (s *Sprite) Up() *Sprite {
	if s.y == 0 {
		return nil
	}
	return s.game.grid[s.y-1][s.x]
}
func (s *Sprite) Down() *Sprite {
	if s.y == len(s.game.grid)-1 {
		return nil
	}
	return s.game.grid[s.y+1][s.x]
}
func (s *Sprite) LeftUp() *Sprite {
	left := s.Left()
	if left == nil {
		return nil
	}
	return left.Up()
}
func (s *Sprite) RightUp() *Sprite {
	right := s.Right()
	if right == nil {
		return nil
	}
	return right.Up()
}
func (s *Sprite) LeftDown() *Sprite {
	left := s.Left()
	if left == nil {
		return nil
	}
	return left.Down()
}
func (s *Sprite) RightDown() *Sprite {
	right := s.Right()
	if right == nil {
		return nil
	}
	return right.Down()
}

func (s *Sprite) IsIce() bool {
	return s.kind == Ice || s.kind == IceFixed
}

func (s *Sprite) IceDie() {
	s.kind = Blank
	s.UnFix()
	s.regularCell()
}
func (s *Sprite) FireDie() {
	s.game.fires--
	s.kind = Blank
	s.regularCell()
}
func (s *Sprite) PlayerDie() {
	s.game.failed = true
	s.kind = Blank
	s.regularCell()
}
func (s *Sprite) UnFix() {
	s.leftFixed = false
	s.rightFixed = false
	left, right := s.Left(), s.Right()
	if left != nil {
		left.rightFixed = false
	}
	if right != nil {
		right.leftFixed = false
	}
}

func (s *Sprite) magicLeft() {
	if s.kind != Player {
		return
	}
	s.magic(s.LeftDown())

}
func (s *Sprite) magicRight() {
	if s.kind != Player {
		return
	}
	s.magic((s.RightDown()))
}

func (s *Sprite) magic(dst *Sprite) {
	if dst == nil {
		return
	}
	switch dst.kind {
	case Blank:
		dst.kind = Ice
		left := dst.Left()
		right := dst.Right()
		if left != nil && (left.kind == Ice || left.kind == Wall) {
			left.rightFixed = true
			dst.leftFixed = true
		}
		if right != nil && (right.kind == Ice || right.kind == Wall) {
			right.leftFixed = true
			dst.rightFixed = true
		}
		s.game.updateUI()
	case Ice:
		up, left, right := dst.Up(), dst.Left(), dst.Right()
		dst.IceDie()
		s.game.updateUI()
		s.game.checkFall(up)
		s.game.checkFall(left)
		s.game.checkFall(right)
	}
}

func (s *Sprite) regularCell() {
	imgPath := ""
	switch s.kind {
	case Wall:
		i := s.game.rd.Intn(5) + 1
		imgPath = "images/icemagic/wall" + strconv.Itoa(i) + ".png"
	case Fire:
		imgPath = "images/icemagic/fire.gif"
	case Player:
		imgPath = "images/icemagic/player.gif"
	case Ice, IceFixed:
		imgPath = "images/icemagic/ice.png"
	}
	s.Cell.Images = []string{imgPath}
	if s.kind == Wall || s.kind == Ice || s.kind == IceFixed {
		s.Cell.BorderTop = true
		s.Cell.BorderBottom = true
	}
	if s.kind == Ice || s.kind == IceFixed {
		s.Cell.BorderLeft = !s.leftFixed
		s.Cell.BorderRight = !s.rightFixed
	}
}
