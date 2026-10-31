package view

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/zrcoder/ttoy/svg"
)

// ImageViewer shows one picture. MyGo runs views and input handlers on the
// main thread, and nothing else in the app touches it, so it carries no
// lock.
type ImageViewer struct {
	bitmap *ui.Bitmap
	svg    *ui.SVG
	// data is the SVG as it came, to draw again when the theme flips: a
	// diagram carries a color scheme for each.
	data []byte
	dark bool
}

func NewImageViewer() *ImageViewer {
	return &ImageViewer{}
}

// SetSvg shows SVG data, with its labels drawn as paths from the fonts it
// carries, as ui.SVG draws shapes but no text.
func (v *ImageViewer) SetSvg(data []byte) {
	dark := mygo.Theme.IsDark()
	v.bitmap, v.svg = nil, readySVG(data, dark)
	v.data, v.dark = data, dark
}

// readySVG reads an SVG as ui.SVG can draw it, readying it first, and
// falls back to the document as it is when that fails.
func readySVG(data []byte, dark bool) *ui.SVG {
	src := data
	if prepared, err := svg.Prepare(data, dark); err == nil {
		src = prepared
	}
	s, _ := ui.ParseSVG(src)
	return s
}

func (v *ImageViewer) SetPng(data []byte) {
	bm, err := ui.DecodeBitmap(data)
	if err != nil {
		return
	}
	v.bitmap, v.svg = bm, nil
	v.data = nil
}

func (v *ImageViewer) View(c *ui.Context) {
	dark := mygo.Theme.IsDark()
	if v.data != nil && v.dark != dark {
		v.dark, v.svg = dark, readySVG(v.data, dark)
	}
	bitmap, svgSrc := v.bitmap, v.svg

	var img ui.ImageSource
	var w, h float32
	switch {
	case bitmap != nil:
		bw, bh := bitmap.Size()
		img, w, h = bitmap, float32(bw), float32(bh)
	case svgSrc != nil:
		img = svgSrc
		w, h = svgSrc.Size()
	default:
		ui.Text(c, "No image").TextColor(c.Theme().TextMuted).Fill()
		return
	}

	maxW := float32(700)
	maxH := float32(500)
	scale := float32(1)
	if w > maxW {
		scale = maxW / w
	}
	if h*scale > maxH {
		scale = maxH / h
	}

	ui.Column(c).Fill().AlignItems(ui.Center).AlignContent(ui.Center).Children(func() {
		ui.Image(c, img).Width(w * scale).Height(h * scale)
	})
}
