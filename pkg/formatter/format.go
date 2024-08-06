package formatter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/yosssi/gohtml"
	"gopkg.in/yaml.v3"

	"github.com/zrcoder/ttoy/pkg/util"
)

func Html(data []byte) (*bytes.Buffer, error) {
	res := gohtml.FormatBytes(data)
	return bytes.NewBuffer(res), nil
}

func HtmlCli(origin []byte) {
	res := gohtml.FormatBytes(origin)
	util.ShowCode("html", res)
}

func Json(data []byte) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, data, "", "    ")
	return buf, err
}

func JsonCli(data []byte) {
	buf, err := Json(data)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("json", buf.Bytes())
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

func TomlCli(data []byte) {
	buf, err := Toml(data)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("toml", buf.Bytes())
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

func YamlCli(data []byte) {
	buf, err := Yaml(data)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("yaml", buf.Bytes())
}
