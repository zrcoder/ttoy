package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/yosssi/gohtml"
	"gopkg.in/yaml.v3"
)

func Html(data []byte) (*bytes.Buffer, error) {
	res := gohtml.FormatBytes(data)
	return bytes.NewBuffer(res), nil
}

func Json(data []byte) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, data, "", "    ")
	return buf, err
}

func Toml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := toml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	err = toml.NewEncoder(buf).Encode(obj)
	return buf, err
}

func Yaml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	err = encoder.Encode(&obj)
	return buf, err
}
