package service

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"html"
	"net/url"

	"github.com/common-nighthawk/go-figure"
	"github.com/zrcoder/ttoy/service/converter"
	"github.com/zrcoder/ttoy/service/formatter"
	"github.com/zrcoder/ttoy/service/generator"
	"github.com/zrcoder/ttoy/service/sort"

	ndor "github.com/zrcoder/ndor/pkg"
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

func (s *Service) GenD2Svg(input string) (string, error) {
	return s.regularSvgData(s.transform([]byte(input), generator.D2Svg))
}

func (s *Service) GenNdorPng(input string) (string, error) {
	res, linErr := ndor.Run(0, 0, input)
	if linErr != nil {
		return "", errors.New(linErr.Msg)
	}
	return res, nil
}

func (s *Service) JSONSort(input string) (string, error) {
	return sort.JSON(input)
}

func (s *Service) EncodeBase64(input string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}

func (s *Service) DecodeBase64(input string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(input)
	return string(res), err
}

func (s *Service) EncodeUrl(input string) (string, error) {
	return url.QueryEscape(input), nil
}

func (s *Service) DecodeUrl(input string) (string, error) {
	return url.QueryUnescape(input)
}

func (s *Service) EncodeHtml(input string) (string, error) {
	return html.EscapeString(input), nil
}

func (s *Service) DecodeHtml(input string) (string, error) {
	return html.UnescapeString(input), nil
}

func (s *Service) Hash(input string) (map[string]string, error) {
	data := []byte(input)
	h := md5.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resMd5 := hex.EncodeToString(h.Sum(nil))

	h = sha1.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resSha1 := hex.EncodeToString(h.Sum(nil))

	h = sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resSha256 := hex.EncodeToString(h.Sum(nil))

	h = sha512.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resSha512 := hex.EncodeToString(h.Sum(nil))

	return map[string]string{
		"md5":    resMd5,
		"sha1":   resSha1,
		"sha256": resSha256,
		"sha512": resSha512,
	}, nil
}

func (s *Service) AsciiArt(input, font string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}
	f := figure.NewFigure(input, font, true)
	return f.String(), nil
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
