package main

import (
	"github.com/egoist/mygo/ui"
)

// Toast dimensions for inline toasts.
const (
	ToastWidth  = 340.0
	ToastHeight = 80.0
	ToastRadius = 8.0
)

// Inline toast background/text colors (dark fallback).
var (
	ErrorBgColor   = ui.Color{R: 50, G: 25, B: 25, A: 255}
	ErrorTextColor = ui.Color{R: 255, G: 255, B: 255, A: 255}
)

// Toast shows a brief centered inline message.
func Toast(c *ui.Context, message string, onDismiss func()) {
	ui.Row(c).Top(0).Bottom(0).Fill().Absolute().OnClick(func() {
		if onDismiss != nil {
			onDismiss()
		}
		c.Invalidate()
	}).Children(func() {
		ui.Spacer(c)
		ui.Column(c).Width(ToastWidth).Height(ToastHeight).Radius(ToastRadius).Background(ErrorBgColor).Children(func() {
			ui.Row(c).Fill().Height(ToastHeight).AlignItems(ui.Center).Justify(ui.Center).Children(func() {
				ui.Text(c, message).TextColor(ErrorTextColor).FontSize(14)
			})
		})
		ui.Spacer(c)
	})
}
