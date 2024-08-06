package main

import (
	"bytes"
	"context"

	"github.com/zrcoder/ttoy/pkg/converter"
	"github.com/zrcoder/ttoy/pkg/formatter"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

type Transformer = func([]byte) (*bytes.Buffer, error)

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ConvertJsonToYaml(input string) (string, error) {
	return a.transform([]byte(input), converter.Json2Yaml)
}

func (a *App) ConvertYamlToJson(input string) (string, error) {
	return a.transform([]byte(input), converter.Yaml2Json)
}

func (a *App) ConvertJsonToToml(input string) (string, error) {
	return a.transform([]byte(input), converter.Json2Toml)
}

func (a *App) ConvertTomlToJson(input string) (string, error) {
	return a.transform([]byte(input), converter.Toml2Json)
}

func (a *App) ConvertYamlToToml(input string) (string, error) {
	return a.transform([]byte(input), converter.Yaml2Toml)
}

func (a *App) ConvertTomlToYaml(input string) (string, error) {
	return a.transform([]byte(input), converter.Toml2Yaml)
}

func (a *App) FormatJson(input string) (string, error) {
	return a.transform([]byte(input), formatter.Json)
}

func (a *App) FormatYaml(input string) (string, error) {
	return a.transform([]byte(input), formatter.Yaml)
}

func (a *App) FormatToml(input string) (string, error) {
	return a.transform([]byte(input), formatter.Toml)
}

func (a *App) FormatHtml(input string) (string, error) {
	return a.transform([]byte(input), formatter.Html)
}

func (a *App) transform(input []byte, transformer Transformer) (string, error) {
	buf, err := transformer(input)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
