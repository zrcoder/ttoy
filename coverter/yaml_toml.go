package converter

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/huh"
	"github.com/zrcoder/ttoy/result"
	"gopkg.in/yaml.v3"
)

func yamlToml() error {
	input := ""
	text := huh.NewText().Title("yaml <> toml").Value(&input).Editor("nvim")
	if err := text.Run(); err != nil {
		return err
	}
	res, err1 := yaml2toml([]byte(input))
	if err1 != nil {
		res, err2 := toml2yaml(input)
		if err2 != nil {
			err := fmt.Errorf("invalid input(yaml/toml)\n%s\n%s", err1, err2)
			result.ShowError(err)
		} else {
			result.ShowCode("yaml", res)
		}
	} else {
		result.ShowCode("toml", res)
	}
	return nil
}

func yaml2toml(jsonData []byte) (string, error) {
	var obj any
	if err := yaml.Unmarshal(jsonData, &obj); err != nil {
		return "", err
	}
	writer := bytes.NewBuffer(nil)
	err := toml.NewEncoder(writer).Encode(obj)
	return writer.String(), err
}

// TODO: fix format bug
func toml2yaml(tomlData string) (string, error) {
	var obj any
	if _, err := toml.Decode(tomlData, &obj); err != nil {
		return "", err
	}
	yamlData, err := yaml.Marshal(obj)
	return string(yamlData), err
}
