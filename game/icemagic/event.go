package icemagic

import "github.com/zrcoder/ttoy/game/internal"

type Event struct {
	Grid  internal.GridUI
	State internal.State
}

func (g *Game) updateUI() bool {
	return g.base.EmitEvent("icemagic:update", Event{
		Grid:  g.UI(),
		State: g.state(),
	})
}

func (g *Game) state() internal.State {
	switch {
	case g.failed:
		return internal.StateFailed
	case g.fires == 0:
		return internal.StateSucceed
	default:
		return internal.StatePlaying
	}
}
