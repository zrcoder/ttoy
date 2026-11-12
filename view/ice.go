package view

import (
	"embed"
	"fmt"
	"log"

	"github.com/ELY-Open-Source/CrychicUI/button"
	"github.com/ELY-Open-Source/CrychicUI/sound"
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/common"
	"github.com/zrcoder/ttoy/game/icemagic"
)

// IceMagic renders the IceMagic puzzle grid.
type IceMagic struct {
	game       *icemagic.Game
	chapterSel int
	levelSel   int

	bitmaps map[string]*ui.Bitmap
	gifs    map[string]*ui.VideoPlayer

	chapterPopoverOpen bool
	levelPopoverOpen   bool

	// judged latches that the game over was shown, until the next game.
	judged   bool
	diagOpen bool

	note int // next move sound, climbing a pentatonic scale
}

func NewIceMagic(assertsFS embed.FS) *IceMagic {
	res := &IceMagic{game: icemagic.New()}
	res.loadBitmaps(assertsFS)
	res.loadGifs(assertsFS)
	return res
}

func (p *IceMagic) loadBitmaps(assertsFS embed.FS) {
	p.bitmaps = make(map[string]*ui.Bitmap)
	for _, k := range []string{
		"images/icemagic/wall1.png",
		"images/icemagic/wall2.png",
		"images/icemagic/wall3.png",
		"images/icemagic/wall4.png",
		"images/icemagic/wall5.png",
		"images/icemagic/ice.png",
	} {
		data, err := assertsFS.ReadFile("resources/" + k)
		if err != nil {
			log.Fatal(err)
		}
		b, err := ui.DecodeBitmap(data)
		if err != nil {
			log.Fatal(err)
		}
		p.bitmaps[k] = b
	}
}

// loadGifs loops the GIF sprites; DecodeBitmap shows only a first frame.
func (p *IceMagic) loadGifs(assertsFS embed.FS) {
	p.gifs = make(map[string]*ui.VideoPlayer)
	for _, k := range []string{"images/icemagic/fire.gif", "images/icemagic/player.gif"} {
		data, err := assertsFS.ReadFile("resources/" + k)
		if err != nil {
			log.Fatal(err)
		}
		dec, err := ui.NewGIFDecoder(data)
		if err != nil {
			log.Fatal(err)
		}
		player, err := ui.NewVideoPlayer(dec)
		if err != nil {
			log.Fatal(err)
		}
		player.SetLoop(true)
		p.gifs[k] = player
	}
}

// moveNotes is the C major pentatonic the walk keys climb.
var moveNotes = [5]float64{523.25, 587.33, 659.25, 783.99, 880}

// moveSfx plays the scale's next note.
func (p *IceMagic) moveSfx() {
	sound.PlayNote(moveNotes[p.note])
	p.note = (p.note + 1) % len(moveNotes)
}

// magicSfx plays the next note a fifth above itself.
func (p *IceMagic) magicSfx() {
	n := moveNotes[p.note]
	sound.PlayNote(n)
	sound.PlayNote(n * 1.5)
}

func (p *IceMagic) View(c *ui.Context) {
	c.OnShortcut(0, ui.KeyJ, func() { p.moveSfx(); p.game.MoveLeft() })
	c.OnShortcut(0, ui.KeyL, func() { p.moveSfx(); p.game.MoveRight() })
	c.OnShortcut(0, ui.KeyA, func() { p.magicSfx(); p.game.MagicLeft() })
	c.OnShortcut(0, ui.KeyD, func() { p.magicSfx(); p.game.MagicRight() })
	ui.Row(c).FillWidth().Height(44).Children(func() {
		ui.Text(c, "IceMagic").FontSize(20).FontWeight(700)
	})
	ui.Column(c).Fill().AlignItems(ui.Center).Children(func() {
		p.menu(c)
		ui.Column(c).Margin(30, 0).Fill().Focus().AlignItems(ui.Center).Children(func() {
			p.board(c)
			p.controls(c)
		})
	})

	p.gameOver(c)
}

// gameOver shows a one-title modal on a level's end, closed by Escape, click outside or Enter.
func (p *IceMagic) gameOver(c *ui.Context) {
	if !p.diagOpen {
		return
	}
	title, tint := "Level Cleared!", c.Theme().Success
	if p.game.State() != common.StateSucceed {
		title, tint = "Burned!", c.Theme().Danger
	}
	if ui.Modal(c, &p.diagOpen, func() {
		panel := ui.Box(c).Focusable().Children(func() { // the modal takes the focus as it opens
			ui.Text(c, title).FontSize(20).Bold().TextColor(tint)
		})
		if panel.Shortcut(0, ui.KeyEnter) {
			p.diagOpen = false
			c.Invalidate()
			p.over()
		}
	}).Dismissed() {
		p.over()
	}
}
func (p *IceMagic) over() {
	if p.game.State() == common.StateSucceed {
		p.nextLevel()
	} else {
		p.game.Reset()
	}
}

func (p *IceMagic) pick(ch, lv int) {
	p.chapterSel, p.levelSel = ch, lv
	p.game.SelectLevel(ch+1, lv+1)
}

// nextLevel advances the chapter's levels, then goes to the next chapter's level 1.
func (p *IceMagic) nextLevel() {
	chs := p.game.Chapters()
	if p.levelSel++; p.levelSel < chs[p.chapterSel] {
		p.pick(p.chapterSel, p.levelSel)
		return
	}
	if p.chapterSel++; p.chapterSel == len(chs) {
		p.chapterSel, p.levelSel = len(chs)-1, chs[len(chs)-1]-1
	}
	p.pick(p.chapterSel, 0)
}

func (p *IceMagic) menu(c *ui.Context) {
	chapters := p.game.Chapters()

	ui.Row(c).Margin(30, 0).Children(func() {
		ui.MenuButton(c, fmt.Sprintf("Chapter %d", p.chapterSel+1), func(m *ui.Menu) {
			for i := range chapters {
				if m.Item(fmt.Sprintf("Chapter %d", i+1)).Chosen() {
					if i == p.chapterSel {
						return
					}
					sound.Play(sound.Select)
					p.pick(i, 0)
				}
			}
		})
		ui.Spacer(c).Width(18)
		ui.MenuButton(c, fmt.Sprintf("Level %d", p.levelSel+1), func(m *ui.Menu) {
			for i := range chapters[p.chapterSel] {
				if m.Item(fmt.Sprintf("Label %d", i+1)).Chosen() {
					if i == p.levelSel {
						return
					}
					sound.Play(sound.Select)
					p.pick(p.chapterSel, i)
				}
			}
		})
		ui.Spacer(c).Width(18)
		button.Solid(c, "Reset").OnClick(p.game.Reset)
	})
}

// board draws the frame the replay shows right now.
func (p *IceMagic) board(c *ui.Context) {
	cells, nextAt, playing := p.game.Present(c.Now())
	if len(cells) == 0 {
		ui.Text(c, "Loading grid...").TextColor(c.Theme().TextMuted)
		return
	}
	ui.Column(c).Margin(28, 0).Children(func() {
		for _, row := range cells {
			ui.Row(c).Children(func() {
				for _, cell := range row {
					p.cell(c, cell)
				}
			})
		}
	})
	if playing {
		if d := nextAt.Sub(c.Now()); d > 0 {
			c.After(d)
		} else {
			c.AnimationFrame()
		}
	}
	// latch a shown game over until the next game begins.
	if st := p.game.State(); st == common.StatePlaying {
		p.judged = false
	} else if !p.judged && !playing {
		p.judged, p.diagOpen = true, true
		if st == common.StateSucceed {
			sound.Play(sound.Success)
		} else {
			sound.Play(sound.Error)
		}
	}
}

func (p *IceMagic) cell(c *ui.Context, cell *icemagic.Cell) {
	ui.Box(c).Width(34).Height(34).BorderColor(ui.RGB(240, 240, 240)).BorderWidth(
		b2f(cell.BorderTop),
		b2f(cell.BorderRight),
		b2f(cell.BorderBottom),
		b2f(cell.BorderLeft),
	).Children(func() {
		if gif := p.gifs[cell.Image]; gif != nil {
			ui.Video(c, gif).Fill().Fit(ui.Cover)
			return
		}
		ui.Image(c, p.bitmaps[cell.Image]).Fill().Fit(ui.Cover)
	})
}

// b2f reports a bool border side as a border width.
func b2f(b bool) float32 {
	if b {
		return 0.5
	}
	return 0
}

func (p *IceMagic) controls(c *ui.Context) {
	ui.Row(c).Children(func() {
		p.controlBtn(c, "J", "arrow-left", func() { p.moveSfx(); p.game.MoveLeft() })
		p.controlBtn(c, "A", "arrow-down-left", func() { p.magicSfx(); p.game.MagicLeft() })
		p.controlBtn(c, "D", "arrow-down-right", func() { p.magicSfx(); p.game.MagicRight() })
		p.controlBtn(c, "L", "arrow-right", func() { p.moveSfx(); p.game.MoveRight() })
	})
}

func (p *IceMagic) controlBtn(c *ui.Context, key, icon string, action func()) {
	ui.Column(c).AlignItems(ui.Center).Margin(10).Children(func() {
		button.IconButton(c, key, icon).OnClick(action)
		ui.Text(c, key).FontSize(12).TextColor(c.Theme().TextMuted)
	})
}
