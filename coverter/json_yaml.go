package converter

import (
	"encoding/json"

	"github.com/zrcoder/ttoy/util"
	"gopkg.in/yaml.v3"
)

func Json2Yaml() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return json2yaml([]byte(stdinput))
	}
	return jsonYamlRun()
}

func Yaml2Json() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return yaml2json([]byte(stdinput))
	}
	return yamlJsonRun()
}

func jsonYamlRun() error {
	input := ""
	if err := util.NewText("json -> yaml", "json", &input).Run(); err != nil {
		return err
	}
	return json2yaml([]byte(input))
}

func yamlJsonRun() error {
	input := ""
	if err := util.NewText("yaml -> json", "yaml", &input).Run(); err != nil {
		return err
	}
	return yaml2json([]byte(input))
}

func json2yaml(jsonData []byte) error {
	var obj any
	if err := json.Unmarshal(jsonData, &obj); err != nil {
		return err
	}
	data, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	util.ShowCode("yaml", string(data))
	return nil
}

func yaml2json(yamlData []byte) error {
	var obj any
	if err := yaml.Unmarshal(yamlData, &obj); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		util.ShowCode("json", string(data))
	}
	return nil
}
