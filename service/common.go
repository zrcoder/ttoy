package service

import (
	"github.com/zrcoder/cdor"
)

func cdorTransform(data string, fn func(*cdor.Cdor)) ([]byte, error) {
	c := cdor.New()
	fn(c)
	c.Fill("transparent")
	c.Theme(200)
	return c.Gen()
}
