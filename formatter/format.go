package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/yosssi/gohtml"
	"gopkg.in/yaml.v3"

	"github.com/zrcoder/ttoy/util"
)

func Html(origin []byte) {
	res := gohtml.FormatBytes(origin)
	util.ShowCode("html", res)
}

func Json(data []byte) {
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, data, "", "  ")
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("json", buf.Bytes())
}

func Toml(data []byte) {
	var obj any
	if err := toml.Unmarshal(data, &obj); err != nil {
		util.ShowFatal(err)
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("toml", writer.Bytes())
}

func Yaml(data []byte) {
	var obj any
	if err := yaml.Unmarshal(data, &obj); err != nil {
		util.ShowFatal(err)
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(&obj); err != nil {
		util.ShowFatal(err)
	} else {
		util.ShowCode("yaml", buf.Bytes())
	}
}
