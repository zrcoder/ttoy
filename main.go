package main

import (
	"os"

	"github.com/charmbracelet/huh"
	converter "github.com/zrcoder/ttoy/coverter"
	"github.com/zrcoder/ttoy/decoder"
	"github.com/zrcoder/ttoy/formatter"
	"github.com/zrcoder/ttoy/generator"
	"github.com/zrcoder/ttoy/result"
)

var (
	actions = map[string]func() error{
		"cv": converter.Run,
		"ed": decoder.Run,
		"ft": formatter.Run,
		"gr": generator.Run,
	}
)

func main() {
	err := run()
	if err != nil {
		result.ShowError(err)
		os.Exit(1)
	}
}

func run() error {
	var action string

	err := huh.NewSelect[string]().
		Title("ttoy").
		Options(
			huh.NewOption("Converters", "cv"),
			huh.NewOption("Encoders/Decoders", "ed"),
			huh.NewOption("Fomatters", "ft"),
			huh.NewOption("Generators", "gr"),
		).
		Value(&action).
		Run()

	if err != nil {
		return err
	}

	return actions[action]()
}
