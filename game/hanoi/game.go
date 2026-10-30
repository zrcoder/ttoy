package hanoi

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/zrcoder/ttoy/game/common"
)

const (
	PileCount = 3
)

type Game struct {
	PileA      *Pile
	PileB      *Pile
	PileC      *Pile
	ShiftDisk  *Disk
	colors     []string
	steps      int
	rand       *rand.Rand
	levelIndex int
	levels     []Level
}

type Pile struct {
	game  *Game
	Index int
	Disks []*Disk
}

type Disk struct {
	pile *Pile
	ID   int
}

type Level struct {
	Label string
	Disks int
}

type PileInfo struct {
	Index int   `json:"Index"`
	Disks []int `json:"Disks"`
}

type EventData struct {
	Piles []PileInfo `json:"Piles"`
	State string     `json:"State"`
	Steps int        `json:"Steps"`
}

func New() *Game {
	g := &Game{
		colors: []string{
			"red", "green", "blue", "yellow", "brown", "pink",
		},
		levels: []Level{
			{Label: "Easy", Disks: 2},
			{Label: "Medium", Disks: 3},
			{Label: "Hard", Disks: 5},
			{Label: "Expert", Disks: 6},
		},
		rand: rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}
	g.PileA = NewPile(g, 0)
	g.PileB = NewPile(g, 1)
	g.PileC = NewPile(g, 2)
	g.Shuffle(len(g.colors), func(i, j int) {
		g.colors[i], g.colors[j] = g.colors[j], g.colors[i]
	})
	g.Reset()
	return g
}

func NewPile(g *Game, index int) *Pile {
	return &Pile{
		game:  g,
		Index: index,
		Disks: make([]*Disk, 0, len(g.colors)),
	}
}

func NewDisk(p *Pile, id int) *Disk {
	return &Disk{
		pile: p,
		ID:   id,
	}
}

func (g *Game) Reset() {
	g.PileA.renewDisks()
	g.PileB.ClearDisks()
	g.PileC.ClearDisks()
	g.ShiftDisk = nil
	g.steps = 0
}

func (g *Game) IsDone() bool {
	return len(g.PileC.Disks) == g.CurrentLevel().Disks
}

func (g *Game) SelectPile(pile *Pile) error {
	if g.IsDone() {
		return errors.New("game is done")
	}
	if g.ShiftDisk == nil {
		g.ShiftDisk = pile.Pop()
		g.updateUI()
		return nil
	}
	if g.ShiftDisk.pile == pile || pile.Empty() || pile.Top().ID > g.ShiftDisk.ID {
		pile.Push(g.ShiftDisk)
		g.ShiftDisk = nil
		g.steps++
		g.updateUI()
		return nil
	}
	return errors.New("invalid move")
}

func (g *Game) MinSteps() int {
	return (1 << g.CurrentLevel().Disks) - 1
}

// ShiftPileIndex returns the index of the pile the held disk was taken
// from, or -1 while no disk is held.
func (g *Game) ShiftPileIndex() int {
	if g.ShiftDisk == nil || g.ShiftDisk.pile == nil {
		return -1
	}
	return g.ShiftDisk.pile.Index
}

func (g *Game) State() string {
	return fmt.Sprintf("Steps: %d, Minimum: %d", g.steps, g.MinSteps())
}

func (g *Game) Levels() []Level {
	return g.levels
}

func (g *Game) LevelIndex() int {
	return g.levelIndex
}

func (g *Game) SelectLevel(index int) {
	g.levelIndex = index
	g.Reset()
	g.updateUI()
}

func (g *Game) CurrentLevel() Level {
	return g.levels[g.levelIndex]
}

func (g *Game) Piles() []PileInfo {
	return []PileInfo{
		{Index: 0, Disks: g.pileToIDs(g.PileA)},
		{Index: 1, Disks: g.pileToIDs(g.PileB)},
		{Index: 2, Disks: g.pileToIDs(g.PileC)},
	}
}

func (g *Game) pileToIDs(p *Pile) []int {
	ids := make([]int, len(p.Disks))
	for i, d := range p.Disks {
		ids[i] = d.ID
	}
	return ids
}

func (g *Game) SelectPileByIndex(index int) error {
	var pile *Pile
	switch index {
	case 0:
		pile = g.PileA
	case 1:
		pile = g.PileB
	case 2:
		pile = g.PileC
	default:
		return errors.New("invalid pile index")
	}
	return g.SelectPile(pile)
}

func (g *Game) updateUI() {
	// No-op: UI polls state directly.
	_ = g.ShiftDisk
}

func (g *Game) state() common.State {
	if g.IsDone() {
		return common.StateSucceed
	}
	return common.StatePlaying
}

func (p *Pile) Push(d *Disk) {
	p.Disks = append(p.Disks, d)
	d.pile = p
}

func (p *Pile) Empty() bool {
	return len(p.Disks) == 0
}

func (p *Pile) Top() *Disk {
	return p.Disks[len(p.Disks)-1]
}

func (p *Pile) Pop() *Disk {
	if p.Empty() {
		return nil
	}
	disk := p.Disks[len(p.Disks)-1]
	p.Disks = p.Disks[:len(p.Disks)-1]
	return disk
}

func (p *Pile) renewDisks() {
	p.Disks = p.Disks[:0]
	n := p.game.CurrentLevel().Disks
	for i := range n {
		p.Disks = append(p.Disks, NewDisk(p, n-1-i))
	}
}

func (p *Pile) ClearDisks() {
	p.Disks = p.Disks[:0]
}

func (g *Game) Shuffle(n int, swap func(i, j int)) {
	g.rand.Shuffle(n, swap)
}
