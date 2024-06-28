package converter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/zrcoder/ttoy/util"
	"gopkg.in/yaml.v2"
)

func Json2Toml(jsonData []byte) error {
	var obj any
	if err := json.Unmarshal([]byte(jsonData), &obj); err != nil {
		return err
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		return err
	}
	return util.ShowCode("toml", writer.Bytes())
}

func Toml2Json(tomlData []byte) error {
	var obj any
	if _, err := toml.Decode(string(tomlData), &obj); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		return util.ShowCode("json", data)
	}
}

func Json2Yaml(jsonData []byte) error {
	var obj any
	if err := json.Unmarshal(jsonData, &obj); err != nil {
		return err
	}
	data, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	return util.ShowCode("yaml", data)
}

func Yaml2Json(yamlData []byte) error {
	var obj any
	if err := yaml.Unmarshal(yamlData, &obj); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		return util.ShowCode("json", data)
	}
}

func Yaml2Toml(jsonData []byte) error {
	var obj any
	if err := yaml.Unmarshal(jsonData, &obj); err != nil {
		return err
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		return err
	}
	return util.ShowCode("toml", writer.Bytes())
}

func Toml2Yaml(tomlData []byte) error {
	var obj any
	if _, err := toml.Decode(string(tomlData), &obj); err != nil {
		return err
	}
	if data, err := yaml.Marshal(obj); err != nil {
		return err
	} else {
		return util.ShowCode("yaml", data)
	}
}
