package formatter

import (
	"github.com/zrcoder/ttoy/util"
	"gopkg.in/yaml.v3"
)

func Yaml() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return yamlFormatAndShow(stdinput)
	}
	return yamlRun()
}

func yamlFormatAndShow(data string) error {
	var obj any
	if err := yaml.Unmarshal([]byte(data), &obj); err != nil {
		return err
	}
	if data, err := yaml.Marshal(obj); err != nil {
		return err
	} else {
		util.ShowCode("yaml", string(data))
	}
	return nil
}

func yamlRun() error {
	data := ""

	if err := util.NewText("yaml formatter", "yaml", &data).Run(); err != nil {
		return err
	}

	return yamlFormatAndShow(data)
}
