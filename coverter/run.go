package converter

import "github.com/charmbracelet/huh"

var actions = map[string]func() error{
	"jy": jsonYaml,
	"jt": jsonToml,
	"yt": yamlToml,
}

func Run() error {
	action := ""

	err := huh.NewSelect[string]().
		Title("Converters").
		Options(
			huh.NewOption("json <> yaml", "jy"),
			huh.NewOption("json <> toml", "jt"),
			huh.NewOption("yaml <> toml", "yt"),
		).
		Value(&action).
		Run()

	if err != nil {
		return err
	}

	return actions[action]()
}
