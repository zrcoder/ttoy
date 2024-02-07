package internal

import (
	"encoding/json"

	"github.com/zrcoder/ttoy/internal/pkg"
	"gopkg.in/yaml.v3"
)

type JsonYaml struct {
	*pkg.Base
}

func NewJsonYaml() *JsonYaml {
	b := pkg.New("json <> yaml", 2)
	jy := &JsonYaml{Base: b}
	b.ConvertAction = jy.convert
	return jy
}

func (jy JsonYaml) convert() error {
	jsonData := jy.Inputs[0].Value()
	yamlData := jy.Inputs[1].Value()
	m := map[string]any{}
	if jy.Inputs[0].Focused() {
		if err := json.Unmarshal([]byte(jsonData), &m); err != nil {
			return err
		}
		if data, err := yaml.Marshal(m); err != nil {
			return err
		} else {
			jy.Inputs[1].SetValue(string(data))
		}
	} else {
		if err := yaml.Unmarshal([]byte(yamlData), &m); err != nil {
			return err
		}
		if data, err := json.MarshalIndent(m, "", "  "); err != nil {
			return err
		} else {
			jy.Inputs[0].SetValue(string(data))
		}
	}
	return nil
}
