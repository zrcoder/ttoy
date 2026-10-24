package service

import (
	"github.com/zrcoder/cdor"
)

type View struct{}

func (v *View) Json(data string) ([]byte, error) {
	return cdorTransform(data, func(c *cdor.Cdor) {
		c.Json(data)
	})
}

func (v *View) Yaml(data string) ([]byte, error) {
	return cdorTransform(data, func(c *cdor.Cdor) {
		c.Yaml(data)
	})
}

func (v *View) Toml(data string) ([]byte, error) {
	return cdorTransform(data, func(c *cdor.Cdor) {
		c.Toml(data)
	})
}
