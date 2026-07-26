package service

import (
	"encoding/base64"

	"github.com/zrcoder/cdor"
)

func wrapSvgData(input []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(input), nil
}

func cdorTransform(data string, fn func(*cdor.Cdor)) ([]byte, error) {
	c := cdor.Ctx()
	fn(c)
	c.Fill("transparent")
	c.Theme(200)
	return c.Gen()
}
