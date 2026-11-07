package view

import (
	"embed"
	"fmt"
	"log"
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/game/icemagic"
)

// IceMagic renders the IceMagic puzzle grid.
type IceMagic struct {
	game       *icemagic.Game
	chapterSel int
	levelSel   int

	bitmaps map[string]*ui.Bitmap
	gifs    map[string]*ui.VideoPlayer

	// hidden until the matching button is clicked
	chapterPopoverOpen bool
	levelPopoverOpen   bool
}

func NewIceMagic(assertsFS embed.FS) *IceMagic {
	res := &IceMagic{game: icemagic.New()}
	res.loadBitmaps(assertsFS)
	res.loadGifs(assertsFS)
	return res
}

// loadBitmaps decodes the still cell images once.
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

// loadGifs gives the animated cell sprites looping players: DecodeBitmap
// shows only the first frame of a GIF.
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

func (p *IceMagic) View(c *ui.Context) {
	c.OnShortcut(0, ui.KeyJ, func() { p.game.MoveLeft() })
	c.OnShortcut(0, ui.KeyL, func() { p.game.MoveRight() })
	c.OnShortcut(0, ui.KeyA, func() { p.game.MagicLeft() })
	c.OnShortcut(0, ui.KeyD, func() { p.game.MagicRight() })

	ui.Column(c).Fill().Focus().Children(func() {
		ui.Row(c).Fill().Height(44).Children(func() {
			ui.Text(c, "IceMagic").FontSize(20).FontWeight(700)
		})
		ui.Column(c).FillWidth().AlignItems(ui.Center).Children(func() {
			p.menu(c)
			p.board(c)
			p.controls(c)
		})
	})
}

// menu shows the chapter and level pickers and the reset button.
func (p *IceMagic) menu(c *ui.Context) {
	chapters := p.game.Chapters()
	ui.Row(c).Children(func() {
		p.pickMenu(c, fmt.Sprintf("Chapter %d", p.chapterSel+1), &p.chapterPopoverOpen, len(chapters), func(i int) {
			p.chapterSel = i
			p.game.SelectLevel(p.chapterSel+1, p.levelSel+1)
		})
		p.pickMenu(c, fmt.Sprintf("Level %d", p.levelSel+1), &p.levelPopoverOpen, chapters[p.chapterSel], func(i int) {
			p.levelSel = i
			p.game.SelectLevel(p.chapterSel+1, p.levelSel+1)
		}).Margin(8)
		ui.PrimaryButton(c, "Reset").Margin(18).OnClick(p.game.Reset)
	})
}

// pickMenu is a labeled button that opens a popover listing 1..n, where
// picking i calls pick with the zero-based index and closes the popover.
// The anchor button is returned for the caller to decorate.
func (p *IceMagic) pickMenu(c *ui.Context, label string, open *bool, n int, pick func(i int)) ui.Element {
	btn := ui.Button(c, label+" ▾")
	if btn.Clicked() {
		*open = !*open
		c.Invalidate()
	}
	ui.Popover(c, btn, open, func() {
		for i := range n {
			name := strconv.Itoa(i + 1)
			entry := ui.Row(c.Key(name)).Padding(6, 12).Radius(5).Children(func() {
				ui.Text(c, name)
			})
			if entry.Hovered() {
				entry.Background(c.Theme().Accent).TextColor(c.Theme().AccentText)
			}
			if entry.Clicked() {
				pick(i)
				*open = false
				c.Invalidate()
			}
		}
	})
	return btn
}

// board draws the frame the replay shows right now, and while it is
// still playing asks for the next one.
func (p *IceMagic) board(c *ui.Context) {
	images, nextAt, playing := p.game.Present(c.Now())
	if len(images) == 0 {
		ui.Text(c, "Loading grid...").TextColor(c.Theme().TextMuted)
		return
	}
	ui.Column(c).Margin(28, 0).Children(func() {
		for _, row := range images {
			ui.Row(c).Children(func() {
				for _, img := range row {
					p.cell(c, img)
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
}

// cell draws one 34×34 board cell by its image path: a looping video for
// the GIF sprites, a bitmap for the rest.
func (p *IceMagic) cell(c *ui.Context, img string) {
	ui.Box(c).Width(34).Height(34).Children(func() {
		if gif := p.gifs[img]; gif != nil {
			ui.Video(c, gif).Fill().Fit(ui.Cover)
			return
		}
		ui.Image(c, p.bitmaps[img]).Fill().Fit(ui.Cover)
	})
}

// controls shows the touch buttons mirrored by the shortcuts above.
func (p *IceMagic) controls(c *ui.Context) {
	ui.Row(c).Children(func() {
		p.controlBtn(c, "J", "←", func() { p.game.MoveLeft() })
		p.controlBtn(c, "L", "→", func() { p.game.MoveRight() })
	})
	ui.Row(c).Children(func() {
		p.controlBtn(c, "A", "↙", func() { p.game.MagicLeft() })
		p.controlBtn(c, "D", "↘", func() { p.game.MagicRight() })
	})
}

func (p *IceMagic) controlBtn(c *ui.Context, key, icon string, action func()) {
	ui.Column(c).AlignItems(ui.Center).Margin(10).Children(func() {
		ui.Button(c, icon).OnClick(action)
		ui.Text(c, key).FontSize(11).TextColor(c.Theme().TextMuted)
	})
}
