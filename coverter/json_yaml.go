package converter

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/zrcoder/ttoy/result"
	"gopkg.in/yaml.v3"
)

func jsonYaml() error {
	input := ""
	text := huh.NewText().Title("json <> yaml").Value(&input).Editor("nvim")
	if err := text.Run(); err != nil {
		return err
	}
	inputData := []byte(input)
	res, err1 := json2Yaml(inputData)
	if err1 != nil {
		res, err2 := yaml2Json(inputData)
		if err2 != nil {
			err := fmt.Errorf("invalid input(json/yaml)\n%s\n%s", err1, err2)
			result.ShowError(err)
		} else {
			result.ShowCode("json", res)
		}
	} else {
		result.ShowCode("yaml", res)
	}
	return nil
}

func json2Yaml(jsonData []byte) (string, error) {
	var obj any
	if err := json.Unmarshal(jsonData, &obj); err != nil {
		return "", err
	}
	yamlData, err := yaml.Marshal(obj)
	return string(yamlData), err
}

func yaml2Json(yamlData []byte) (string, error) {
	var obj any
	if err := yaml.Unmarshal(yamlData, &obj); err != nil {
		return "", err
	}
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	return string(jsonData), err
}
