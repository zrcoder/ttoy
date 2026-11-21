package export

import "github.com/zrcoder/ttoy/game/hanoi"

//go:generate go tool qexp -outdir ../exported .

func A() { _ = hanoi.Pick(0) }
func B() { _ = hanoi.Pick(1) }
func C() { _ = hanoi.Pick(2) }
