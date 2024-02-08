package converter

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/zrcoder/ttoy/util"
)

func Json2Toml() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return json2toml([]byte(stdinput))
	}
	return jsonTomlRun()
}

func Toml2Json() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return toml2json(stdinput)
	}
	return tomlJsonRun()
}

func jsonTomlRun() error {
	input := ""
	if err := util.NewText("json -> toml", "json", &input).Run(); err != nil {
		return err
	}
	return json2toml([]byte(input))
}

func tomlJsonRun() error {
	input := ""
	if err := util.NewText("toml -> json", "toml", &input).Run(); err != nil {
		return err
	}
	return toml2json(input)
}

func json2toml(jsonData []byte) error {
	var obj any
	if err := json.Unmarshal([]byte(jsonData), &obj); err != nil {
		return err
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		return err
	}
	util.ShowCode("toml", writer.String())
	return nil
}

func toml2json(tomlData string) error {
	var obj any
	if _, err := toml.Decode(tomlData, &obj); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		util.ShowCode("json", string(data))
	}
	return nil
}
