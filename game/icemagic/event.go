package icemagic

func (g *Game) updateUI() {
	if g.onMove != nil {
		g.onMove()
	}
}
