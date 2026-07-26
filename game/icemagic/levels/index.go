package levels

import (
	"embed"
)

const (
	Rows = 14
	Cols = 16
)

//go:embed */*
var FS embed.FS
