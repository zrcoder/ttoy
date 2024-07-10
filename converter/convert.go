package converter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"

	"github.com/zrcoder/ttoy/util"
)

func Json2Toml(jsonData []byte) {
	var obj any
	if err := json.Unmarshal([]byte(jsonData), &obj); err != nil {
		util.ShowFatal(err)
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("toml", writer.Bytes())
}

func Toml2Json(tomlData []byte) {
	var obj any
	if _, err := toml.Decode(string(tomlData), &obj); err != nil {
		util.ShowFatal(err)
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		util.ShowFatal(err)
	} else {
		util.ShowCode("json", data)
	}
}

func Json2Yaml(jsonData []byte) {
	var obj any
	if err := json.Unmarshal(jsonData, &obj); err != nil {
		util.ShowFatal(err)
	}
	data, err := yaml.Marshal(obj)
	if err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("yaml", data)
}

func Yaml2Json(yamlData []byte) {
	var obj any
	if err := yaml.Unmarshal(yamlData, &obj); err != nil {
		util.ShowFatal(err)
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		util.ShowFatal(err)
	} else {
		util.ShowCode("json", data)
	}
}

func Yaml2Toml(jsonData []byte) {
	var obj any
	if err := yaml.Unmarshal(jsonData, &obj); err != nil {
		util.ShowFatal(err)
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		util.ShowFatal(err)
	}
	util.ShowCode("toml", writer.Bytes())
}

func Toml2Yaml(tomlData []byte) {
	var obj any
	if _, err := toml.Decode(string(tomlData), &obj); err != nil {
		util.ShowFatal(err)
	}
	if data, err := yaml.Marshal(obj); err != nil {
		util.ShowFatal(err)
	} else {
		util.ShowCode("yaml", data)
	}
}
