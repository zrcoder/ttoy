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
	*Game
	Kind       byte
	X          int
	Y          int
	LeftFixed  bool
	RightFixed bool
}

func (s *Sprite) moveLeft() bool {
	switch s.Kind {
	case Player:
		return s.MoveLeft()
	case Ice:
		return s.iceSlideLeft()
	}
	return false
}

func (s *Sprite) moveRight() bool {
	switch s.Kind {
	case Player:
		return s.MoveRight()
	case Ice:
		return s.iceSlideRight()
	}
	return false
}

func (s *Sprite) iceSlideLeft() bool {
	left := s.Left()
	if left == nil {
		return false
	}
	switch left.Kind {
	case Fire:
		left.FireDie()
		s.IceDie()
		time.Sleep(stepTime)
		s.Game.updateUI()
		s.Game.checkFall(left.Up())
		s.Game.checkFall(s.Up())
		return true
	case Blank:
		up := s.Up()
		if ok := s.Game.swap(left, s, stepTime); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			s.iceSlideLeft()
		}
		s.Game.checkFall(up)
		return true
	}
	return false
}
func (s *Sprite) iceSlideRight() bool {
	right := s.Right()
	if right == nil {
		return false
	}
	switch right.Kind {
	case Fire:
		right.FireDie()
		s.IceDie()
		time.Sleep(stepTime)
		s.Game.updateUI()
		s.Game.checkFall(right.Up())
		s.Game.checkFall(s.Up())
		return true
	case Blank:
		up := s.Up()
		if ok := s.Game.swap(s, right, stepTime); !ok {
			return false
		}
		if ok := s.fall(); !ok {
			s.iceSlideRight()
		}
		s.Game.checkFall(up)
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
	if up != nil && up.Kind != Blank {
		return false
	}
	if dst == nil || dst.Kind != Blank {
		return false
	}
	return s.Game.swap(s, dst, 2*stepTime)
}

func (s *Sprite) Left() *Sprite {
	if s.X == 0 {
		return nil
	}
	return s.Game.grid[s.Y][s.X-1]
}
func (s *Sprite) Right() *Sprite {
	row := s.Game.grid[s.Y]
	n := len(row)
	if s.X == n-1 {
		return nil
	}
	return row[s.X+1]
}
func (s *Sprite) Up() *Sprite {
	if s.Y == 0 {
		return nil
	}
	return s.Game.grid[s.Y-1][s.X]
}
func (s *Sprite) Down() *Sprite {
	if s.Y == len(s.Game.grid)-1 {
		return nil
	}
	return s.Game.grid[s.Y+1][s.X]
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
	return s.Kind == Ice || s.Kind == IceFixed
}

func (s *Sprite) IceDie() {
	s.Kind = Blank
	s.UnFix()
}
func (s *Sprite) FireDie() {
	s.Game.fires--
	s.Kind = Blank
}
func (s *Sprite) PlayerDie() {
	s.Game.failed = true
	s.Kind = Blank
}
func (s *Sprite) UnFix() {
	s.LeftFixed = false
	s.RightFixed = false
	left, right := s.Left(), s.Right()
	if left != nil {
		left.RightFixed = false
	}
	if right != nil {
		right.LeftFixed = false
	}
}

func (s *Sprite) magicLeft() {
	if s.Kind != Player {
		return
	}
	s.magic(s.LeftDown())

}
func (s *Sprite) magicRight() {
	if s.Kind != Player {
		return
	}
	s.magic((s.RightDown()))
}

func (s *Sprite) magic(dst *Sprite) {
	if dst == nil {
		return
	}
	switch dst.Kind {
	case Blank:
		dst.Kind = Ice
		left := dst.Left()
		right := dst.Right()
		if left != nil && (left.Kind == Ice || left.Kind == Wall) {
			left.RightFixed = true
			dst.LeftFixed = true
		}
		if right != nil && (right.Kind == Ice || right.Kind == Wall) {
			right.LeftFixed = true
			dst.RightFixed = true
		}
		s.Game.updateUI()
	case Ice:
		up, left, right := dst.Up(), dst.Left(), dst.Right()
		dst.IceDie()
		s.Game.updateUI()
		s.Game.checkFall(up)
		s.Game.checkFall(left)
		s.Game.checkFall(right)
	}
}

func (s *Sprite) uiCell() internal.Cell {
	imgPath := ""
	switch s.Kind {
	case Wall:
		i := s.Game.rd.Intn(5) + 1
		imgPath = "images/icemagic/wall" + strconv.Itoa(i) + ".png"
	case Fire:
		imgPath = "images/icemagic/fire.gif"
	case Player:
		imgPath = "images/icemagic/player.gif"
	case Ice, IceFixed:
		imgPath = "images/icemagic/ice.png"
	}
	cell := internal.Cell{
		Images: []string{imgPath},
	}
	if s.Kind == Wall || s.Kind == Ice || s.Kind == IceFixed {
		cell.BorderTop = true
		cell.BorderBottom = true
	}
	if s.Kind == Ice || s.Kind == IceFixed {
		cell.BorderLeft = !s.LeftFixed
		cell.BorderRight = !s.RightFixed
	}
	return cell
}
