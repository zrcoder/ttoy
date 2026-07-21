package converter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

func Json2Toml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	err = toml.NewEncoder(writer).Encode(obj)
	return writer, err
}

func Toml2Json(data []byte) (*bytes.Buffer, error) {
	var obj any
	_, err := toml.Decode(string(data), &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(obj)
	return writer, err
}

func Json2Yaml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	return writer, err
}

func Yaml2Json(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(obj)
	return writer, err
}

func Yaml2Toml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	err = toml.NewEncoder(writer).Encode(obj)
	return writer, err
}

func Toml2Yaml(data []byte) (*bytes.Buffer, error) {
	var obj any
	_, err := toml.Decode(string(data), &obj)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	err = encoder.Encode(obj)
	return buf, err
}
