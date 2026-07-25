package view

import "github.com/zrcoder/ttoy/service/generate"

type Viewer struct {
	gen *generate.Generator
}

func New() *Viewer {
	return &Viewer{gen: generate.New()}
}

func (v *Viewer) Json(data string) (string, error) {
	return v.gen.Json2Svg(data)
}

func (v *Viewer) Yaml(data string) (string, error) {
	return v.gen.Yaml2Svg(data)
}

func (v *Viewer) Toml(data string) (string, error) {
	return v.gen.Toml2Svg(data)
}
