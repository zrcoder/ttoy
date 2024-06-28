package main

import (
	"errors"

	"github.com/zrcoder/ttoy/converter"
	"github.com/zrcoder/ttoy/encoder"
	"github.com/zrcoder/ttoy/formatter"
	"github.com/zrcoder/ttoy/generator"
	"github.com/zrcoder/ttoy/util"
)

type (
	OperAction func(input []byte)
	Action     func([]byte) error
)

func main() {
	const (
		converter = "converter"
		formatter = "formatter"
		generator = "generator"
		encoder   = "encoder"
		decoder   = "decoder"
	)

	oper := ""
	ops := []string{converter, formatter, generator, encoder, decoder}
	err := util.NewSelect("", ops, &oper).Run()
	if err != nil {
		util.ShowError(err)
		return
	}

	opers := map[string]OperAction{
		converter: convert,
		formatter: format,
		generator: generate,
		encoder:   encode,
		decoder:   decode,
	}
	opers[oper](util.Input)
}

func convert(input []byte) {
	const (
		jy = "json --> yaml"
		yj = "yaml --> json"
		yt = "yaml --> toml"
		ty = "toml --> yaml"
		jt = "json --> toml"
		tj = "toml --> json"
	)
	converters := map[string]Action{
		jy: converter.Json2Yaml,
		yj: converter.Yaml2Json,
		yt: converter.Yaml2Toml,
		ty: converter.Toml2Yaml,
		jt: converter.Json2Toml,
		tj: converter.Toml2Json,
	}
	srcformat := util.ParseInputExtension()
	key := ""
	if srcformat == "" {
		ops := []string{jy, jt, yj, yt, tj, ty}
		util.NewSelect("source format", ops, &key).Run()
	} else {
		set := map[string]bool{"json": true, "yaml": true, "toml": true}
		ops := make([]string, 0, 2)
		for k := range set {
			if srcformat == k {
				continue
			}
			ops = append(ops, k)
		}
		util.NewSelect("to", ops, &key).Run()
		key = srcformat + " --> " + key
	}

	if err := converters[key](input); err != nil {
		util.ShowError(err)
	}
}

func format(input []byte) {
	formatters := map[string]Action{
		"json": formatter.Json,
		"yaml": formatter.Yaml,
		"toml": formatter.Toml,
		"html": formatter.Html,
	}
	srcformat := util.ParseInputExtension()
	if srcformat == "" {
		util.NewSelect("source format", []string{"json", "yaml", "toml", "html"}, &srcformat).Run()
	}
	action, ok := formatters[srcformat]
	if !ok {
		util.ShowError(errors.New("not supported format"))
	}
	if err := action(input); err != nil {
		util.ShowError(err)
	}
}

func generate(input []byte) {
	const (
		jsvg    = "json --> svg"
		jstruct = "json --> struct"
		uuid    = "uuid"
		hash    = "hash"
	)
	generators := map[string]Action{
		uuid:    generator.UUID,
		hash:    generator.Hash,
		jsvg:    generator.Json2Svg,
		jstruct: generator.Json2Struct,
	}
	key := ""
	util.NewSelect("generator", []string{uuid, hash, jsvg, jstruct}, &key).Run()
	if err := generators[key](input); err != nil {
		util.ShowError(err)
	}
}

func encode(input []byte) {
	encodeOrDecode(input, true)
}

func decode(input []byte) {
	encodeOrDecode(input, false)
}

func encodeOrDecode(input []byte, encode bool) {
	srcformat := util.ParseInputExtension()
	var err error
	if srcformat == "html" || srcformat == "xml" {
		if encode {
			err = encoder.EncodeHtml(input)
		} else {
			err = encoder.DecodeHtml(input)
		}
	} else {
		url := string(input)
		if url == "" {
			util.NewInput("encode url", &url).Run()
		}
		if encode {
			err = encoder.EncodeUrl(url)
		} else {
			err = encoder.DecodeUrl(url)
		}
	}
	if err != nil {
		util.ShowError(err)
	}
}
