package formatter

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/zrcoder/ttoy/util"
)

func Toml() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return tomlFormatAndShow(stdinput)
	}
	return tomlRun()
}

func tomlFormatAndShow(data string) error {
	var obj any
	if err := toml.Unmarshal([]byte(data), &obj); err != nil {
		return err
	}
	writer := bytes.NewBuffer(nil)
	if err := toml.NewEncoder(writer).Encode(obj); err != nil {
		return err
	}
	util.ShowCode("toml", writer.String())
	return nil
}

func tomlRun() error {
	data := ""

	if err := util.NewText("toml formatter", "toml", &data).Run(); err != nil {
		return err
	}

	return tomlFormatAndShow(data)
}
