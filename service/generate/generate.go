package generate

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"github.com/common-nighthawk/go-figure"
	"github.com/zrcoder/cdor"
	ndor "github.com/zrcoder/ndor/pkg"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) AsciiArt(input, font string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}
	f := figure.NewFigure(input, font, true)
	return f.String(), nil
}

func (g *Generator) Hash(input string) (map[string]string, error) {
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

func (g *Generator) Json2Svg(data string) (string, error) {
	return svgData(g.svgTransform(data, func(c *cdor.Cdor) {
		c.Json(data)
	}))
}

func (g *Generator) Yaml2Svg(data string) (string, error) {
	return svgData(g.svgTransform(data, func(c *cdor.Cdor) {
		c.Yaml(data)
	}))
}

func (g *Generator) Toml2Svg(data string) (string, error) {
	return svgData(g.svgTransform(data, func(c *cdor.Cdor) {
		c.Toml(data)
	}))
}

func (g *Generator) D2Svg(data string) (string, error) {
	return svgData(g.svgTransform(data, func(c *cdor.Cdor) {
		c.D2(data)
	}))
}

func (g *Generator) NdorPng(input string) (string, error) {
	res, linErr := ndor.Run(0, 0, input)
	if linErr != nil {
		return "", errors.New(linErr.Msg)
	}
	return res, nil
}

func (g *Generator) svgTransform(data string, fn func(*cdor.Cdor)) ([]byte, error) {
	c := cdor.Ctx()
	fn(c)
	c.Fill("transparent")
	c.Theme(200)
	return c.Gen()
}

func svgData(input []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(input), nil
}
