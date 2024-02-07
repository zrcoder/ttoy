package svg

import (
	"os"

	"github.com/zrcoder/cdor"
	"github.com/zrcoder/ttoy/util"
)

func Json() error {
	return svg("json")
}

func Yaml() error {
	return svg("yaml")
}

func Tomal() error {
	return svg("toml")
}

func svg(kind string) error {
	input := util.ReadStdin()
	if input == "" {
		if err := util.NewText(kind+" graph", kind, &input).Run(); err != nil {
			return err
		}
	}
	c := cdor.Ctx()
	switch kind {
	case "json":
		c.Json(input)
	case "yaml":
		c.Yaml(input)
	case "toml":
		c.Toml(input)
	}
	data, err := c.Gen()
	if err != nil {
		return err
	}
	file := ""
	if err := util.NewInput("file name", &file).Run(); err != nil {
		return err
	}
	return os.WriteFile(file, data, 0o600)
}
