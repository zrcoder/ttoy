package internal

import (
	"math/rand/v2"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Base struct {
	app            *application.App
	chapters       []Chapter
	chapterOptions []any
	chapterIndex   int
	levels         []Level
	levelOptions   []any
	levelIndex     int
	reset          func()
	sceneName      string
	rd             *rand.Rand
	successMsgs    []string
}

func New(app *application.App, opts ...Option) *Base {
	seed1, seed2 := uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())
	b := &Base{
		app:         app,
		rd:          rand.New(rand.NewPCG(seed1, seed2)),
		successMsgs: []string{"Wanderful!", "Brilliant!", "Excellent!", "Fantastic!", "Awesome!"},
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

type ChapterLevelOptionInput struct {
	LevelSelect LevelMeta `json:"levelSelect"`
}

func (b *Base) EmitEvent(eventName string, data any) {
	b.app.Event.Emit(eventName, data)
}
