package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/yosssi/gohtml"
	"gopkg.in/yaml.v3"

	"github.com/zrcoder/ttoy/util"
)

func Html(origin []byte) error {
	res := gohtml.FormatBytes(origin)
	return util.ShowCode("html", res)
}

func Json(data []byte) error {
	var obj any
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		return util.ShowCode("json", data)
	}
}

func Toml(data []byte) error {
	var obj any
	if err := toml.Unmarshal(data, &obj); err != nil {
		return err
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		return err
	}
	return util.ShowCode("toml", writer.Bytes())
}

func Yaml(data []byte) error {
	var obj any
	if err := yaml.Unmarshal(data, &obj); err != nil {
		return err
	}
	if data, err := yaml.Marshal(obj); err != nil {
		return err
	} else {
		return util.ShowCode("yaml", data)
	}
}
