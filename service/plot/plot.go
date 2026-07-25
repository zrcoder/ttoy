package plot

import "github.com/zrcoder/ttoy/service/generate"

type Plotter struct {
	gen *generate.Generator
}

func New() *Plotter {
	return &Plotter{gen: generate.New()}
}

func (p *Plotter) D2(data string) (string, error) {
	return p.gen.D2Svg(data)
}

func (p *Plotter) Ndor(input string) (string, error) {
	return p.gen.NdorPng(input)
}
