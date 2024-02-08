package generator

import "github.com/charmbracelet/huh"

var actions = map[string]func() error{
	"hs": hash,
	"ui": genUuid,
}

func Run() error {
	action := ""

	err := huh.NewSelect[string]().
		Title("Generators").
		Options(
			huh.NewOption("hash", "hs"),
			huh.NewOption("uuid", "ui"),
		).
		Value(&action).
		Run()

	if err != nil {
		return err
	}

	return actions[action]()
}
