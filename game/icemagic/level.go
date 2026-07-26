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
			sprite := &Sprite{Kind: typeFlag, X: x, Y: y, Game: g}
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
			if left != nil && left.Kind == IceFixed {
				left.Kind = Ice
			}
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
			s.LeftFixed = true
			left.RightFixed = true
		}
	}
	switch left.Kind {
	case Wall:
		fix(s.Kind == IceFixed || s.Kind == Wall)
	case IceFixed:
		fix(s.Kind == Ice || s.Kind == IceFixed || s.Kind == Wall)
	case Ice:
		fix(s.Kind == IceFixed)
	}
}
