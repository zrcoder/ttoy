package hanoi

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

const (
	PileCount = 3

	MaxDisks = 5
)

type Game struct {
	PileA      *Pile
	PileB      *Pile
	PileC      *Pile
	ShiftDisk  *Disk
	steps      int
	rand       *rand.Rand
	levelIndex int
	levels     []Level

	onMove func()
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
	Label        string
	Disks        int
	SampleCode   string
	InitCodeLine int
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
		levels: []Level{
			{
				Label: "Easy",
				Disks: 2,
				SampleCode: `// Move disks with A, B or C.
// The first call picks a disk up, the next drops it.
A
B

A
C
// TODO: finish the last move here

`,
				InitCodeLine: 7,
			},
			{
				Label: "Medium",
				Disks: 3,
				SampleCode: `// We can define a recursive function to solve this problem.
func solve(n int, A, B, C func()) {
	if n == 0 {
		return
	}
	solve n-1, A, C, B
	
	// TODO: move the only one disk from area A to area C
	
	
	solve n-1, B, A, C
}

// Call the solve function to solve the problem.
solve 3, A, B, C
`,
				InitCodeLine: 8,
			},
			{
				Label: "Hard",
				Disks: 5,
				SampleCode: `// Use the recursive function to solve this problem.
func solve(n int, A, B, C func()) {
	
}
`,
				InitCodeLine: 2,
			},
		},
		rand: rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}
	g.PileA = NewPile(g, 0)
	g.PileB = NewPile(g, 1)
	g.PileC = NewPile(g, 2)

	g.Reset()
	return g
}

func NewPile(g *Game, index int) *Pile {
	return &Pile{
		game:  g,
		Index: index,
		Disks: make([]*Disk, 0, MaxDisks),
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
		g.notifyMove()
		return nil
	}
	if g.ShiftDisk.pile == pile || pile.Empty() || pile.Top().ID > g.ShiftDisk.ID {
		pile.Push(g.ShiftDisk)
		g.ShiftDisk = nil
		g.steps++
		g.notifyMove()
		return nil
	}
	g.notifyMove()
	return errors.New("invalid move")
}

func (g *Game) MinSteps() int {
	return (1 << g.CurrentLevel().Disks) - 1
}

func (g *Game) Steps() int {
	return g.steps
}

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
}

func (g *Game) CurrentLevel() Level {
	return g.levels[g.levelIndex]
}

func (g *Game) Piles() []PileInfo {
	return []PileInfo{
		{Index: 0, Disks: g.pileIDs(g.PileA)},
		{Index: 1, Disks: g.pileIDs(g.PileB)},
		{Index: 2, Disks: g.pileIDs(g.PileC)},
	}
}

func (g *Game) SetOnMove(f func()) {
	g.onMove = f
}

func (g *Game) notifyMove() {
	if g.onMove != nil {
		g.onMove()
	}
}

func (g *Game) pileIDs(p *Pile) []int {
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
