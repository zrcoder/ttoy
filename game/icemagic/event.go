package icemagic

import (
	"github.com/zrcoder/ttoy/game/common"
)

type Event struct {
	Grid  [][]*Sprite
	State common.State
}

func (g *Game) updateUI() {
	// No-op: UI polls state directly.
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
