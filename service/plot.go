package service

import (
	"errors"

	"github.com/zrcoder/cdor"
	ndor "github.com/zrcoder/ndor/pkg"
)

type Plot struct{}

func (p *Plot) D2(data string) ([]byte, error) {
	return cdorTransform(data, func(c *cdor.Cdor) {
		c.D2(data)
	})
}

func (p *Plot) Ndor(input string) ([]byte, error) {
	res, linErr := ndor.Gen(0, 0, input)
	if linErr != nil {
		return nil, errors.New(linErr.Msg)
	}
	return res, nil
}
