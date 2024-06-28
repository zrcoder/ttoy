package generator

import (
	"github.com/zrcoder/cdor"
	"github.com/zrcoder/ttoy/util"
)

func Json2Svg(inputJson []byte) error {
	c := cdor.Ctx()
	c.Json(string(inputJson))
	data, err := c.Gen()
	if err != nil {
		return err
	}
	return util.Show(data)
}
