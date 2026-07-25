package format

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/yosssi/gohtml"
	"gopkg.in/yaml.v3"
)

type Formatter struct{}

func New() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Html(data string) (string, error) {
	res := gohtml.FormatBytes([]byte(data))
	return string(res), nil
}

func (f *Formatter) Json(data string) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, []byte(data), "", "    ")
	return buf.String(), err
}

func (f *Formatter) Toml(data string) (string, error) {
	var obj any
	err := toml.Unmarshal([]byte(data), &obj)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	err = toml.NewEncoder(buf).Encode(obj)
	return buf.String(), err
}

func (f *Formatter) Yaml(data string) (string, error) {
	var obj any
	err := yaml.Unmarshal([]byte(data), &obj)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	err = encoder.Encode(&obj)
	return buf.String(), err
}
