package icemagic

import (
	"strconv"

	"github.com/zrcoder/ttoy/game/common"
)

const (
	Blank    = ' '
	Wall     = '='
	Fire     = 'F'
	Player   = 'M'
	Ice      = 'i'
	IceFixed = 'I'
)

type Sprite struct {
	game       *Game
	kind       byte
	x          int
	y          int
	leftFixed  bool
	rightFixed bool
	Cell       common.Cell
}

func (s *Sprite) iceSlideLeft() bool {
	left := s.left()
	if left == nil {
		return false
	}
	switch left.kind {
	case Fire:
		s.iceDie()
		s.game.updateUI()
		left.fireDie()
		s.game.updateUI()
		s.game.checkFall(left.up())
		s.game.checkFall(s.up())
		return true
	case Blank:
		up := s.up()
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
	right := s.right()
	if right == nil {
		return false
	}
	switch right.kind {
	case Fire:
		s.iceDie()
		s.game.updateUI()
		right.fireDie()
		s.game.updateUI()
		s.game.checkFall(right.up())
		s.game.checkFall(s.up())
		return true
	case Blank:
		up := s.up()
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
	return s.climb(s.leftUp())
}

func (s *Sprite) climbRight() bool {
	return s.climb(s.rightUp())
}

func (s *Sprite) climb(dst *Sprite) bool {
	up := s.up()
	if up != nil && up.kind != Blank {
		return false
	}
	if dst == nil || dst.kind != Blank {
		return false
	}
	return s.game.swap(s, dst, 2*stepTime)
}

func (s *Sprite) left() *Sprite {
	if s.x == 0 {
		return nil
	}
	return s.game.grid[s.y][s.x-1]
}
func (s *Sprite) right() *Sprite {
	row := s.game.grid[s.y]
	n := len(row)
	if s.x == n-1 {
		return nil
	}
	return row[s.x+1]
}
func (s *Sprite) up() *Sprite {
	if s.y == 0 {
		return nil
	}
	return s.game.grid[s.y-1][s.x]
}
func (s *Sprite) down() *Sprite {
	if s.y == len(s.game.grid)-1 {
		return nil
	}
	return s.game.grid[s.y+1][s.x]
}
func (s *Sprite) leftUp() *Sprite {
	left := s.left()
	if left == nil {
		return nil
	}
	return left.up()
}
func (s *Sprite) rightUp() *Sprite {
	right := s.right()
	if right == nil {
		return nil
	}
	return right.up()
}
func (s *Sprite) leftDown() *Sprite {
	left := s.left()
	if left == nil {
		return nil
	}
	return left.down()
}
func (s *Sprite) rightDown() *Sprite {
	right := s.right()
	if right == nil {
		return nil
	}
	return right.down()
}

func (s *Sprite) isIce() bool {
	return s.kind == Ice || s.kind == IceFixed
}

func (s *Sprite) iceDie() {
	s.kind = Blank
	s.unFix()
	s.regularCell()
}
func (s *Sprite) fireDie() {
	s.game.fires--
	s.kind = Blank
	s.regularCell()
}
func (s *Sprite) playerDie() {
	s.game.failed = true
	s.kind = Blank
	s.regularCell()
}
func (s *Sprite) unFix() {
	s.leftFixed = false
	s.rightFixed = false
	left, right := s.left(), s.right()
	if left != nil {
		left.rightFixed = false
		left.regularCell()
	}
	if right != nil {
		right.leftFixed = false
		right.regularCell()
	}
	s.regularCell()
}

func (s *Sprite) fix() {
	left := s.left()
	right := s.right()
	if left != nil && (left.kind == Ice || left.kind == Wall) {
		left.rightFixed = true
		s.leftFixed = true
		left.regularCell()
	}
	if right != nil && (right.kind == Ice || right.kind == Wall) {
		right.leftFixed = true
		s.rightFixed = true
		right.regularCell()
	}
	s.regularCell()
}

func (s *Sprite) magicLeft() {
	if s.kind != Player {
		return
	}
	s.magic(s.leftDown())

}
func (s *Sprite) magicRight() {
	if s.kind != Player {
		return
	}
	s.magic((s.rightDown()))
}

func (s *Sprite) magic(dst *Sprite) {
	if dst == nil {
		return
	}
	switch dst.kind {
	case Blank:
		dst.kind = Ice
		dst.fix()
		s.game.updateUI()
	case Ice:
		up, left, right := dst.up(), dst.left(), dst.right()
		dst.iceDie()
		s.game.updateUI()
		s.game.checkFall(up)
		s.game.checkFall(left)
		s.game.checkFall(right)
	}
	dst.regularCell()
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
	s.Cell.BorderBottom = false
	s.Cell.BorderTop = false
	s.Cell.BorderLeft = false
	s.Cell.BorderRight = false
	if s.kind == Wall || s.kind == Ice || s.kind == IceFixed {
		s.Cell.BorderTop = true
		s.Cell.BorderBottom = true
		s.Cell.BorderLeft = !s.leftFixed
		s.Cell.BorderRight = !s.rightFixed
	}
}
