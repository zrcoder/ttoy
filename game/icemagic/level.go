package icemagic

import (
	"bytes"
	"fmt"
	"log/slog"

	"github.com/zrcoder/ttoy/game/icemagic/levels"
	"github.com/zrcoder/ttoy/game/internal"
)

const (
	sceneName = "ice-magic"
)

type Level struct {
	HideMagicButton bool
}

func (g *Game) initLevels() {
	g.chapters = []internal.Chapter{
		{
			Children: []internal.Level{
				{Data: &Level{HideMagicButton: true}},
				{}, {}, {}, {}, {}, {}, {}, {},
			},
		},
		{
			Children: []internal.Level{
				{}, {}, {}, {}, {}, {}, {}, {}, {},
			},
		},
		{
			Children: []internal.Level{
				{}, {}, {}, {}, {}, {}, {}, {}, {},
			},
		},
		{
			Children: []internal.Level{
				{}, {}, {}, {}, {}, {}, {}, {}, {},
			},
		},
		{
			Children: []internal.Level{
				{}, {},
			},
		},
	}
	for i := range g.chapters {
		g.chapters[i].Label = fmt.Sprintf("Chapter %d", i+1)
		for j := range g.chapters[i].Children {
			g.chapters[i].Children[j].Label = fmt.Sprintf("%d-%d", i+1, j+1)
		}
	}
}

func (g *Game) parseGrid(chapter, level int) {
	data, err := levels.FS.ReadFile(fmt.Sprintf("%d/%d.txt", chapter+1, level+1))
	if err != nil {
		slog.Error("failed to read level file", "error", err)
		return
	}

	lines := bytes.Split(data, []byte{'\n'})
	g.grid = make([][]*Sprite, len(lines))
	for y, line := range lines {
		g.grid[y] = make([]*Sprite, len(line))
		for x := range line {
			typeFlag := line[x]
			sprite := &Sprite{kind: typeFlag, x: x, y: y, game: g}
			g.grid[y][x] = sprite
			switch typeFlag {
			case Blank:
			case Fire:
				g.fires++
			case Player:
				g.player = sprite
			case IceFixed, Ice, Wall:
				sprite.checkToFixLeft()
			}
			left := sprite.Left()
			if left != nil && left.kind == IceFixed {
				left.kind = Ice
			}
		}
	}
	for _, line := range g.grid {
		for _, sprite := range line {
			sprite.regularCell()
		}
	}
}

func (s *Sprite) checkToFixLeft() {
	left := s.Left()
	if left == nil {
		return
	}
	var fix = func(condition bool) {
		if condition {
			s.leftFixed = true
			left.rightFixed = true
		}
	}
	switch left.kind {
	case Wall:
		fix(s.kind == IceFixed || s.kind == Wall)
	case IceFixed:
		fix(s.kind == Ice || s.kind == IceFixed || s.kind == Wall)
	case Ice:
		fix(s.kind == IceFixed)
	}
}
