package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
	cvter "github.com/zrcoder/ttoy/coverter"
	"github.com/zrcoder/ttoy/decoder"
	"github.com/zrcoder/ttoy/formatter"
	"github.com/zrcoder/ttoy/generator"
	"github.com/zrcoder/ttoy/generator/svg"
	"github.com/zrcoder/ttoy/util"
)

var (
	json = &cli.Command{
		Name:   "json",
		Usage:  "format json or convert json to yaml, toml and graph(svg)",
		Action: wrapAction(formatter.Json),
		Commands: []*cli.Command{
			newYamlCmd(wrapAction(cvter.Json2Yaml)),
			newTomlCmd(wrapAction(cvter.Json2Toml)),
			newSvgCmd("generate json graph", wrapAction(svg.Json)),
		},
	}
	yaml = &cli.Command{
		Name:   "yaml",
		Usage:  "convert yaml to json, toml and graph(svg)",
		Action: wrapAction(formatter.Yaml),
		Commands: []*cli.Command{
			newJsonCmd(wrapAction(cvter.Yaml2Json)),
			newTomlCmd(wrapAction(cvter.Yaml2Toml)),
			newSvgCmd("generate yaml graph", wrapAction(svg.Yaml)),
		},
	}
	toml = &cli.Command{
		Name:   "toml",
		Usage:  "convert toml to json, yaml and graph(svg)",
		Action: wrapAction(formatter.Toml),
		Commands: []*cli.Command{
			newJsonCmd(wrapAction(cvter.Toml2Json)),
			newYamlCmd(wrapAction(cvter.Toml2Yaml)),
			newSvgCmd("generate toml graph", wrapAction(svg.Tomal)),
		},
	}
	xml = &cli.Command{
		Name:    "xml",
		Aliases: []string{"html"},
		Usage:   "format xml/html or convert to markdown",
		Action:  wrapAction(formatter.Html),
	}
	encode = &cli.Command{
		Name:  "encode",
		Usage: "encode url, html and so on",
		Commands: []*cli.Command{
			{
				Name:   "url",
				Action: wrapAction(decoder.UrlEncode),
			},
			{
				Name:    "html",
				Aliases: []string{"xml"},
				Action:  wrapAction(decoder.HtmlEncode),
			},
		},
	}
	decode = &cli.Command{
		Name:  "decode",
		Usage: "decode url, html and so on",
		Commands: []*cli.Command{
			{
				Name:   "url",
				Action: wrapAction(decoder.UrlDecode),
			},
			{
				Name:    "html",
				Aliases: []string{"xml"},
				Action:  wrapAction(decoder.HtmlDecode),
			},
		},
	}
	hash = &cli.Command{
		Name:   "hash",
		Usage:  "generate hashes",
		Action: wrapAction(generator.Hash),
	}
	uuid = &cli.Command{
		Name:   "uuid",
		Usage:  "generate uuid",
		Action: wrapAction(generator.UUID),
	}
)

func wrapAction(f func() error) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		return f()
	}
}

func newJsonCmd(action cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:   "json",
		Usage:  "convert to json",
		Action: action,
	}
}

func newYamlCmd(action cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:   "yaml",
		Usage:  "convert to yaml",
		Action: action,
	}
}

func newTomlCmd(action cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:   "toml",
		Usage:  "convert to toml",
		Action: action,
	}
}

func newSvgCmd(usage string, action cli.ActionFunc) *cli.Command {
	return &cli.Command{
		Name:   "svg",
		Usage:  usage,
		Action: action,
	}
}

func main() {
	root := &cli.Command{
		Name:  "ttoy",
		Usage: "terminal dev toys",
		Commands: []*cli.Command{
			json,
			yaml,
			toml,
			xml,
			encode,
			decode,
			hash,
			uuid,
		},
	}

	root.HideHelp = true

	if err := root.Run(context.Background(), os.Args); err != nil {
		util.ShowError(err)
		os.Exit(1)
	}
}
