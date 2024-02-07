package pkg

import (
	lg "github.com/charmbracelet/lipgloss"
)

const (
	errorHeight = 2
	helpHeight  = 5
	sepWidth    = 2
	sep         = "  "
)

var (
	cursorStyle = lg.NewStyle().
			Foreground(lg.Color("212"))

	cursorLineStyle = lg.NewStyle().
			Background(lg.Color("57")).
			Foreground(lg.Color("230"))

	placeholderStyle = lg.NewStyle().
				Foreground(lg.Color("238"))

	endOfBufferStyle = lg.NewStyle().
				Foreground(lg.Color("235"))

	focusedPlaceholderStyle = lg.NewStyle().
				Foreground(lg.Color("99"))

	focusedBorderStyle = lg.NewStyle().
				Border(lg.RoundedBorder()).
				BorderForeground(lg.Color("238"))

	blurredBorderStyle = lg.NewStyle().
				Border(lg.HiddenBorder())
)
