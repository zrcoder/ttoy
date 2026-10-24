package main

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// needsRaster reports whether an SVG has to be rasterized before showing.
// mygo's SVG renderer draws shapes but no text, which is most of what cdor
// diagrams are made of.
func needsRaster(data []byte) bool {
	return bytes.Contains(data, []byte("<text"))
}

// rasterizeSVG draws an SVG into an image with the system renderer, which
// keeps the text that ui.SVG leaves out. Quick Look paints on a white canvas,
// so white turns transparent and the margins are trimmed.
func rasterizeSVG(data []byte) (image.Image, error) {
	if runtime.GOOS != "darwin" {
		return nil, errors.New("svg rasterization is only supported on macOS")
	}

	dir, err := os.MkdirTemp("", "ttoy-svg-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "diagram.svg")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return nil, err
	}
	if out, err := exec.Command("/usr/bin/qlmanage", "-t", "-s", "1600", "-o", dir, path).CombinedOutput(); err != nil {
		return nil, errors.New("qlmanage: " + strings.TrimSpace(string(out)))
	}

	f, err := os.Open(path + ".png")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return trimWhite(img), nil
}

// trimWhite makes the white canvas transparent and crops the empty margins.
func trimWhite(img image.Image) image.Image {
	b := img.Bounds()
	m := image.NewNRGBA(b)
	minX, minY := b.Max.X, b.Max.Y
	maxX, maxY := b.Min.X-1, b.Min.Y-1
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			if c.R >= 250 && c.G >= 250 && c.B >= 250 {
				c.A = 0
			}
			m.SetNRGBA(x, y, c)
			if c.A > 0 {
				minX, minY = min(minX, x), min(minY, y)
				maxX, maxY = max(maxX, x), max(maxY, y)
			}
		}
	}
	if maxX < minX || maxY < minY {
		return m
	}
	const pad = 12
	return m.SubImage(image.Rect(
		max(b.Min.X, minX-pad),
		max(b.Min.Y, minY-pad),
		min(b.Max.X, maxX+1+pad),
		min(b.Max.Y, maxY+1+pad),
	))
}
