package service

import (
	"errors"

	"github.com/zrcoder/cdor"
	ndor "github.com/zrcoder/ndor/pkg"
)

type Plot struct{}

func (p *Plot) D2(data string) (string, error) {
	return wrapSvgData(cdorTransform(data, func(c *cdor.Cdor) {
		c.D2(data)
	}))
}

func (p *Plot) Ndor(input string) (string, error) {
	res, linErr := ndor.Run(0, 0, input)
	if linErr != nil {
		return "", errors.New(linErr.Msg)
	}
	return res, nil
}
