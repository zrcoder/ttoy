package converter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"

	"github.com/zrcoder/ttoy/pkg/util"
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

func Json2TomlCli(jsonData []byte) {
	buf, err := Json2Toml(jsonData)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("toml", buf.Bytes())
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

func Toml2JsonCli(tomlData []byte) {
	buf, err := Toml2Json(tomlData)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("json", buf.Bytes())
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

func Json2YamlCli(jsonData []byte) {
	buf, err := Json2Yaml(jsonData)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("yaml", buf.Bytes())
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

func Yaml2JsonCli(yamlData []byte) {
	buf, err := Yaml2Json(yamlData)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("json", buf.Bytes())
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

func Yaml2TomlCli(yamlData []byte) {
	buf, err := Yaml2Toml(yamlData)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("toml", buf.Bytes())
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

func Toml2YamlCli(tomlData []byte) {
	buf, err := Toml2Yaml(tomlData)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("yaml", buf.Bytes())
}
