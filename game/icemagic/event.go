package icemagic

import (
	"time"

	"github.com/zrcoder/ttoy/game/common"
)

const stepTime = 100 * time.Millisecond

type Event struct {
	Grid  [][]*Sprite
	State common.State
}

func (g *Game) updateUI() {
	time.Sleep(stepTime)
	g.app.Event.Emit("icemagic:update", Event{
		Grid:  g.Grid(),
		State: g.state(),
	})
}

func (g *Game) state() common.State {
	switch {
	case g.failed:
		return common.StateFailed
	case g.fires == 0:
		return common.StateSucceed
	default:
		return common.StatePlaying
	}
}
