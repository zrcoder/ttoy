package hanoi

var current *Game

// Bind registers the game that A/B/C should drive. The view calls this
// right before running user code.
func Bind(g *Game) {
	current = g
}

func Pick(i int) error {
	if current == nil {
		return nil
	}
	return current.SelectPileByIndex(i)
}
