package converter

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/zrcoder/ttoy/util"
	"gopkg.in/yaml.v3"
)

func Yaml2Toml() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return yaml2toml([]byte(stdinput))
	}
	return yamlTomlRun()
}

func Toml2Yaml() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return toml2yaml(stdinput)
	}
	return tomlYamlRun()
}

func yamlTomlRun() error {
	input := ""
	if err := util.NewText("yaml -> toml", "yaml", &input).Run(); err != nil {
		return err
	}
	return yaml2toml([]byte(input))
}

func tomlYamlRun() error {
	input := ""
	if err := util.NewText("toml -> yaml", "toml", &input).Run(); err != nil {
		return err
	}
	return toml2yaml(input)
}

func yaml2toml(jsonData []byte) error {
	var obj any
	if err := yaml.Unmarshal(jsonData, &obj); err != nil {
		return err
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		return err
	}
	util.ShowCode("toml", writer.String())
	return nil
}

func toml2yaml(tomlData string) error {
	var obj any
	if _, err := toml.Decode(tomlData, &obj); err != nil {
		return err
	}
	if data, err := yaml.Marshal(obj); err != nil {
		return err
	} else {
		util.ShowCode("yaml", string(data))
	}
	return nil
}
