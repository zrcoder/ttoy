package converter

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/huh"
	"github.com/zrcoder/ttoy/result"
)

func jsonToml() error {
	input := ""
	text := huh.NewText().Title("json <> toml").Value(&input).Editor("nvim")
	if err := text.Run(); err != nil {
		return err
	}
	res, err1 := json2toml([]byte(input))
	if err1 != nil {
		res, err2 := toml2Json(input)
		if err2 != nil {
			err := fmt.Errorf("invalid input(json/yaml)\n%s\n%s", err1, err2)
			result.ShowError(err)
		} else {
			result.ShowCode("json", res)
		}
	} else {
		result.ShowCode("toml", res)
	}
	return nil
}

func json2toml(jsonData []byte) (string, error) {
	var obj any
	if err := json.Unmarshal([]byte(jsonData), &obj); err != nil {
		return "", err
	}
	writer := bytes.NewBuffer(nil)
	err := toml.NewEncoder(writer).Encode(obj)
	return writer.String(), err
}

func toml2Json(tomlData string) (string, error) {
	var obj any
	if _, err := toml.Decode(tomlData, &obj); err != nil {
		return "", err
	}
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	return string(jsonData), err
}
