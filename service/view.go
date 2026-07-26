package service

import (
	"github.com/zrcoder/cdor"
)

type View struct{}

func (v *View) Json(data string) (string, error) {
	return wrapSvgData(cdorTransform(data, func(c *cdor.Cdor) {
		c.Json(data)
	}))
}

func (v *View) Yaml(data string) (string, error) {
	return wrapSvgData(cdorTransform(data, func(c *cdor.Cdor) {
		c.Yaml(data)
	}))
}

func (v *View) Toml(data string) (string, error) {
	return wrapSvgData(cdorTransform(data, func(c *cdor.Cdor) {
		c.Toml(data)
	}))
}
