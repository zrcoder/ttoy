package generator

import (
	"bytes"

	"github.com/zrcoder/cdor"
)

var StructOption = struct {
	Name          string
	Pkg           string
	Tags          []string
	ConvertFloats bool
	SubStruct     bool
}{}

func Json2Svg(data []byte) (*bytes.Buffer, error) {
	return toSvg(func(c *cdor.Cdor) {
		c.Json(string(data))
	})
}

func Yaml2Svg(data []byte) (*bytes.Buffer, error) {
	return toSvg(func(c *cdor.Cdor) {
		c.Yaml(string(data))
	})
}

func Tomal2Svg(data []byte) (*bytes.Buffer, error) {
	return toSvg(func(c *cdor.Cdor) {
		c.Toml(string(data))
	})
}

func D2Svg(data []byte) (*bytes.Buffer, error) {
	return toSvg(func(c *cdor.Cdor) {
		c.D2(string(data))
	})
}

func toSvg(fn func(*cdor.Cdor)) (*bytes.Buffer, error) {
	c := cdor.Ctx()
	fn(c)
	c.Fill("transparent")
	c.Theme(200)
	data, err := c.Gen()
	return bytes.NewBuffer(data), err
}
