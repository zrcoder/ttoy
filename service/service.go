package service

import (
	"bytes"
	"encoding/base64"

	"github.com/zrcoder/ttoy/pkg/converter"
	"github.com/zrcoder/ttoy/pkg/formatter"
	"github.com/zrcoder/ttoy/pkg/generator"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

type Transformer = func([]byte) (*bytes.Buffer, error)

func (s *Service) ConvertJsonToYaml(input string) (string, error) {
	return s.transform([]byte(input), converter.Json2Yaml)
}

func (s *Service) ConvertYamlToJson(input string) (string, error) {
	return s.transform([]byte(input), converter.Yaml2Json)
}

func (s *Service) ConvertJsonToToml(input string) (string, error) {
	return s.transform([]byte(input), converter.Json2Toml)
}

func (s *Service) ConvertTomlToJson(input string) (string, error) {
	return s.transform([]byte(input), converter.Toml2Json)
}

func (s *Service) ConvertYamlToToml(input string) (string, error) {
	return s.transform([]byte(input), converter.Yaml2Toml)
}

func (s *Service) ConvertTomlToYaml(input string) (string, error) {
	return s.transform([]byte(input), converter.Toml2Yaml)
}

func (s *Service) FormatJson(input string) (string, error) {
	return s.transform([]byte(input), formatter.Json)
}

func (s *Service) FormatYaml(input string) (string, error) {
	return s.transform([]byte(input), formatter.Yaml)
}

func (s *Service) FormatToml(input string) (string, error) {
	return s.transform([]byte(input), formatter.Toml)
}

func (s *Service) FormatHtml(input string) (string, error) {
	return s.transform([]byte(input), formatter.Html)
}

func (s *Service) GenJsonSvg(input string) (string, error) {
	return s.regularSvgData(s.transform([]byte(input), generator.Json2Svg))
}

func (s *Service) GenYamlSvg(input string) (string, error) {
	return s.regularSvgData(s.transform([]byte(input), generator.Yaml2Svg))
}

func (s *Service) GenTomlSvg(input string) (string, error) {
	return s.regularSvgData(s.transform([]byte(input), generator.Tomal2Svg))
}

func (s *Service) regularSvgData(input string, err error) (string, error) {
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(input)), err
}

func (s *Service) transform(input []byte, transformer Transformer) (string, error) {
	buf, err := transformer(input)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
