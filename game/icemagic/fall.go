package icemagic

import "time"

type Bar struct {
	Left, Right *Sprite
}

func (b *Bar) canFall() bool {
	g := b.Left.game
	if b.Left.y >= len(g.grid)-1 {
		return false
	}
	if b.iceFixed() {
		return false
	}
	down := g.grid[b.Left.y+1]
	for x := b.Left.x; x <= b.Right.x; x++ {
		switch down[x].kind {
		case Fire, Blank:
		default:
			return false
		}
	}
	return true
}

func (b *Bar) getUpFallBars() []*Bar {
	if b.Left.y == 0 {
		return nil
	}
	upRow := b.Left.game.grid[b.Left.y-1]
	var res []*Bar
	preX := -1
	for x := b.Left.x; x <= b.Right.x; x++ {
		up := upRow[x]
		switch up.kind {
		case Wall, Blank:
		default:
			bar := up.bar()
			if bar.Left.x > preX {
				res = append(res, bar)
				preX = bar.Right.x
			}
		}
	}
	return res
}

func (b *Bar) fallBar1StepQuietly() []*Bar {
	g := b.Left.game
	y := b.Left.y
	if y >= len(g.grid)-1 {
		return nil
	}
	row := g.grid[y]
	downRow := g.grid[y+1]
	var res []*Bar
	preX := b.Left.x
	for x := b.Left.x; x <= b.Right.x; x++ {
		cur, down := row[x], downRow[x]
		if down.kind == Fire {
			switch cur.kind {
			case Player:
				cur.PlayerDie()
				g.updateUI()
				return nil
			case Ice, IceFixed:
				cur.IceDie()
				down.FireDie()
				if preX < x {
					res = append(res, &Bar{row[preX], row[x-1]})
				}
				preX = x + 1
			}
		}
	}
	if preX <= b.Right.x {
		res = append(res, &Bar{row[preX], row[b.Right.x]})
	}
	for x := b.Left.x; x <= b.Right.x; x++ {
		g.swapQuietly(row[x], downRow[x])
	}
	return res
}

func (b *Bar) iceFixed() bool {
	return b.Left.leftFixed || b.Right.rightFixed
}

func (g *Game) checkFall(s *Sprite) {
	if s == nil || s.kind == Blank || s.kind == Wall {
		return
	}
	g.fallBars(s.bar())
}

func (g *Game) fallBars(bars ...*Bar) bool {
	if len(bars) == 0 {
		return false
	}
	var upBars, newBars []*Bar
	res := false
	for _, b := range bars {
		if !b.canFall() {
			continue
		}
		res = true
		upBars = append(upBars, b.getUpFallBars()...)
		newBars = append(newBars, b.fallBar1StepQuietly()...)
	}
	if !res {
		return res
	}
	time.Sleep(stepTime)
	g.updateUI()
	g.fallBars(newBars...)
	g.fallBars(upBars...)
	return res
}

func (s *Sprite) bar() *Bar {
	if s == nil {
		return nil
	}
	if s.IsIce() {
		return s.getIceBar()
	}
	return &Bar{Left: s, Right: s}
}

func (s *Sprite) getIceBar() *Bar {
	if s == nil || !s.IsIce() {
		return nil
	}
	x1, x2 := s.x, s.x
	row := s.game.grid[s.y]
	for x1 >= 0 && row[x1].IsIce() && row[x1].leftFixed {
		x1--
	}
	if !row[x1].IsIce() {
		x1++
	}
	for x2 < len(row) && row[x2].IsIce() && row[x2].rightFixed {
		x2++
	}
	if !row[x2].IsIce() {
		x2--
	}
	return &Bar{Left: row[x1], Right: row[x2]}
}

func (s *Sprite) fall() bool {
	return s.game.fallBars(&Bar{s, s})
}
