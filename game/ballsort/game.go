package ballsort

import (
	"math/rand/v2"

	"github.com/zrcoder/ttoy/game/common"
)

const (
	BottleBallCount = 4
	EmptyBottles    = 2
)

type Game struct {
	Bottles          []*Bottle
	ShiftBall        *Ball
	DoneBottlesCount int
	colors           []string
	levelIndex       int
	levels           []Level
	rand             *rand.Rand
	cursor           int
}

type Level struct {
	Label       string
	FullBottles int
}

type Bottle struct {
	*Game `json:"-"`
	Index int
	Balls []*Ball
}

type Ball struct {
	BottleIndex int
	Color       string
}

type EventData struct {
	State          string    `json:"State"`
	Bottles        []*Bottle `json:"Bottles"`
	DoneCount      int       `json:"DoneCount"`
	TotalCount     int       `json:"TotalCount"`
	Cursor         int       `json:"Cursor"`
	ShiftBallIndex int       `json:"ShiftBallIndex"`
	ShiftBallColor string    `json:"ShiftBallColor"`
}

func New() *Game {
	g := &Game{
		colors: []string{
			"red", "#7FFF00", "blue", "yellow", "violet", "pink", "green", "orange",
		},
		levels: []Level{
			{Label: "EASY", FullBottles: 5},
			{Label: "MEDIUM", FullBottles: 6},
			{Label: "HARD", FullBottles: 7},
			{Label: "EXPERT", FullBottles: 8},
		},
		rand: rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}
	g.Reset()
	return g
}

func (g *Game) Reset() {
	g.ShiftBall = nil
	g.DoneBottlesCount = 0
	g.shuffleColors()
	n := g.currentBallTubes()
	balls := g.generateBalls(n)
	g.Bottles = g.createBottles(n, balls)
	g.ShiftBall = nil
	for _, bottle := range g.Bottles {
		if bottle.IsDone() {
			g.DoneBottlesCount++
		}
	}
}

func (g *Game) shuffleColors() {
	g.rand.Shuffle(len(g.colors), func(i, j int) {
		g.colors[i], g.colors[j] = g.colors[j], g.colors[i]
	})
}

func (g *Game) generateBalls(n int) []*Ball {
	balls := make([]*Ball, 0, n*BottleBallCount)
	for i := range n {
		for range BottleBallCount {
			balls = append(balls, &Ball{Color: g.colors[i]})
		}
	}
	g.rand.Shuffle(len(balls), func(i, j int) {
		balls[i], balls[j] = balls[j], balls[i]
	})
	return balls
}

func (g *Game) createBottles(n int, balls []*Ball) []*Bottle {
	bottles := make([]*Bottle, 0, n+EmptyBottles)
	i := 0
	for range n {
		bottle := &Bottle{Game: g}
		bottle.Balls = balls[i : i+BottleBallCount]
		for _, ball := range bottle.Balls {
			ball.BottleIndex = bottle.Index
		}
		i += BottleBallCount
		bottles = append(bottles, bottle)
	}
	for range EmptyBottles {
		bottle := &Bottle{
			Game:  g,
			Balls: make([]*Ball, 0, BottleBallCount),
		}
		bottles = append(bottles, bottle)
	}
	g.rand.Shuffle(len(bottles), func(i, j int) {
		bottles[i], bottles[j] = bottles[j], bottles[i]
	})
	for i, bottle := range bottles {
		bottle.Index = i
	}
	return bottles
}

func (g *Game) SelectBottle(i int) {
	if i < 0 || i >= len(g.Bottles) {
		return
	}
	bottle := g.Bottles[i]
	if bottle.IsDone() {
		return
	}
	if bottle.IsEmpty() {
		if g.ShiftBall == nil {
			return
		}
		bottle.Push(g.ShiftBall)
		g.ShiftBall = nil
		g.updateUI()
		return
	}
	if g.ShiftBall == nil {
		g.ShiftBall = bottle.Pop()
		g.updateUI()
		return
	}
	if g.ShiftBall.BottleIndex != i &&
		(g.ShiftBall.Color != bottle.Top().Color || bottle.IsFull()) {
		g.Bottles[g.ShiftBall.BottleIndex].Push(g.ShiftBall)
		g.ShiftBall = bottle.Pop()
		g.updateUI()
		return
	}
	bottle.Push(g.ShiftBall)
	g.ShiftBall = nil
	g.DoneBottlesCount += bottle.checkDone()
	g.updateUI()
}

func (g *Game) IsDone() bool {
	return g.DoneBottlesCount == g.currentBallTubes()
}

func (g *Game) SelectCurrent() {
	g.SelectBottle(g.cursor)
}

func (g *Game) Cursor() int {
	return g.cursor
}

func (g *Game) currentBallTubes() int {
	return g.levels[g.levelIndex].FullBottles
}

func (g *Game) Levels() []Level {
	return g.levels
}

func (g *Game) GetBottles() []*Bottle {
	return g.Bottles
}

func (g *Game) LevelIndex() int {
	return g.levelIndex
}

func (g *Game) SelectLevel(index int) {
	g.levelIndex = index
	g.cursor = 0
	g.Reset()
	g.updateUI()
}

func (g *Game) State() string {
	return "Done: " + string(rune('0'+g.DoneBottlesCount)) + "/" + string(rune('0'+g.currentBallTubes()))
}

func (g *Game) updateUI() {
	// No-op: UI polls state directly.
	_ = g.ShiftBall
}

func (g *Game) state() common.State {
	if g.IsDone() {
		return common.StateSucceed
	}
	return common.StatePlaying
}

func (b *Bottle) Pop() *Ball {
	n := len(b.Balls)
	if n == 0 {
		return nil
	}
	res := b.Balls[n-1]
	b.Balls = b.Balls[:n-1]
	return res
}

func (b *Bottle) Push(ball *Ball) {
	b.Balls = append(b.Balls, ball)
	ball.BottleIndex = b.Index
}

func (b *Bottle) Top() *Ball {
	if len(b.Balls) == 0 {
		return nil
	}
	return b.Balls[len(b.Balls)-1]
}

func (b *Bottle) IsEmpty() bool {
	return len(b.Balls) == 0
}

func (b *Bottle) IsFull() bool {
	return len(b.Balls) == BottleBallCount
}

func (b *Bottle) checkDone() int {
	if !b.IsFull() {
		return 0
	}
	for i := 1; i < len(b.Balls); i++ {
		if b.Balls[i].Color != b.Balls[0].Color {
			return 0
		}
	}
	return 1
}

func (b *Bottle) IsDone() bool {
	return b.checkDone() == 1
}

func (b *Bottle) IsShiftBall() bool {
	return b.Game.ShiftBall != nil && b.Game.ShiftBall.BottleIndex == b.Index
}
