package decoder

import "github.com/charmbracelet/huh"

var actions = map[string]func() error{
	"eu": encodeUrl,
	"du": decodeUrl,
	"eh": encodeHtml,
	"dh": decodeHtml,
}

func Run() error {
	action := ""

	err := huh.NewSelect[string]().
		Title("Encoders/Decoders").
		Options(
			huh.NewOption("encode url", "eu"),
			huh.NewOption("decode url", "du"),
			huh.NewOption("encode html", "eh"),
			huh.NewOption("decode html", "dh"),
		).
		Value(&action).
		Run()

	if err != nil {
		return err
	}

	return actions[action]()
}
