package view

import (
	"image"
	"math"

	"github.com/egoist/mygo/ui"
)

const (
	ballSize   = 40.0
	ballSAA    = 3
	ballPeriod = 9.0
)

// ballTexture patterns the palette's exact black or white ball, which
// its bare color may otherwise match the background.
func ballTexture(base ui.Color) *image.NRGBA {
	if ballShade(base) < 0.2 {
		return ballPattern(base, ui.RGB(236, 239, 244), func(x, y float64) bool {
			return ballCell((x+y)/math.Sqrt2, ballPeriod) < 3.5
		})
	}
	return ballPattern(base, ui.RGB(90, 99, 112), func(x, y float64) bool {
		dx, dy := ballCell(x, ballPeriod)-4.5, ballCell(y, ballPeriod)-4.5
		return dx*dx+dy*dy < 8.4 // a dot of radius 2.9
	})
}

// ballCell folds x into a period-wide cell, 0 to period.
func ballCell(x, period float64) float64 {
	return x - math.Floor(x/period)*period
}

// ballShade reports a color's perceived luminance, 0 to 1.
func ballShade(c ui.Color) float64 {
	return (0.2126*float64(c.R) + 0.7152*float64(c.G) + 0.0722*float64(c.B)) / 255
}

// ballPattern renders base/accent by pattern, antialiased with ballSAA²
// jittered subpixel samples per pixel. Per-channel averages stay straight
// colors, so NRGBA rather than a premultiplied RGBA.
func ballPattern(base, accent ui.Color, pattern func(x, y float64) bool) *image.NRGBA {
	n := int(ballSize)
	hf := ballSize / 2
	img := image.NewNRGBA(image.Rect(0, 0, n, n))
	for py := range n {
		for px := range n {
			var r, g, b, a float64
			for sy := range ballSAA {
				for sx := range ballSAA {
					x := float64(px) + (float64(sx)+0.5)/ballSAA - hf
					y := float64(py) + (float64(sy)+0.5)/ballSAA - hf
					if x*x+y*y > hf*hf {
						continue
					}
					c := base
					if pattern(x, y) {
						c = accent
					}
					r += float64(c.R)
					g += float64(c.G)
					b += float64(c.B)
					a += float64(c.A)
				}
			}
			k := 1.0 / (ballSAA * ballSAA)
			i := (py*n + px) * 4
			img.Pix[i] = byte(r*k + 0.5)
			img.Pix[i+1] = byte(g*k + 0.5)
			img.Pix[i+2] = byte(b*k + 0.5)
			img.Pix[i+3] = byte(a*k + 0.5)
		}
	}
	return img
}
