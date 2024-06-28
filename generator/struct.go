package generator

import (
	"bytes"

	"github.com/ChimeraCoder/gojson"
	"github.com/zrcoder/ttoy/util"
)

func Json2Struct(input []byte) error {
	structName := "Ttoy"
	pkgName := "ttoy"
	buf := bytes.NewBuffer(input)
	tags := []string{"json", "yaml"}
	out, err := gojson.Generate(buf, gojson.ParseJson, structName, pkgName, tags, false, true)
	if err != nil {
		return err
	}
	return util.Show(out)
}
