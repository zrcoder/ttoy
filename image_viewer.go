package main

import (
	"sync"

	"github.com/egoist/mygo/ui"
)

type ImageViewer struct {
	mu     sync.Mutex
	bitmap *ui.Bitmap
	svg    *ui.SVG
}

func NewImageViewer() *ImageViewer {
	return &ImageViewer{}
}

// SetSvg shows SVG data; ones with text are rasterized in the background so
// their labels survive.
func (v *ImageViewer) SetSvg(c *ui.Context, data []byte) {
	var s *ui.SVG
	if parsed, err := ui.ParseSVG(data); err == nil {
		s = parsed
	}
	v.mu.Lock()
	v.bitmap, v.svg = nil, s
	v.mu.Unlock()
	if !needsRaster(data) {
		return
	}
	services := c.Services()
	go func() {
		img, err := rasterizeSVG(data)
		if err != nil {
			return
		}
		bm := ui.NewBitmap(img)
		v.mu.Lock()
		v.bitmap = bm
		v.mu.Unlock()
		services.Invalidate()
	}()
}

func (v *ImageViewer) SetPng(data []byte) {
	bm, err := ui.DecodeBitmap(data)
	if err != nil {
		return
	}
	v.mu.Lock()
	v.bitmap, v.svg = bm, nil
	v.mu.Unlock()
}

func (v *ImageViewer) View(c *ui.Context) {
	v.mu.Lock()
	bitmap, svg := v.bitmap, v.svg
	v.mu.Unlock()

	var img ui.ImageSource
	var w, h float32
	switch {
	case bitmap != nil:
		bw, bh := bitmap.Size()
		img, w, h = bitmap, float32(bw), float32(bh)
	case svg != nil:
		img = svg
		w, h = svg.Size()
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
