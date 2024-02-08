package formatter

import (
	"encoding/json"

	"github.com/zrcoder/ttoy/util"
)

func Json() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return jsonFormatAndShow(stdinput)
	}
	return jsonRun()
}

func jsonFormatAndShow(data string) error {
	var obj any
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		util.ShowCode("json", string(data))
	}
	return nil
}

func jsonRun() error {
	data := ""

	if err := util.NewText("json formatter", "json", &data).Run(); err != nil {
		return err
	}

	return jsonFormatAndShow(data)
}
