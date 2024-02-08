package formatter

import "github.com/charmbracelet/huh"

var actions = map[string]func() error{
	"js": formatJson,
	"xm": formatXML,
}

func Run() error {
	action := ""

	err := huh.NewSelect[string]().
		Title("Fomatters").
		Options(
			huh.NewOption("json", "js"),
			huh.NewOption("xml/html", "xm"),
		).
		Value(&action).
		Run()

	if err != nil {
		return err
	}

	return actions[action]()
}
