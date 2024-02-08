package formatter

import (
	"encoding/json"

	"github.com/charmbracelet/huh"
	"github.com/zrcoder/ttoy/result"
)

func formatJson() error {
	data := ""
	if err := huh.NewText().Title("json formatter").Value(&data).Editor("nvim").Run(); err != nil {
		return err
	}
	var obj any
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return err
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		result.ShowCode("json", string(data))
	}
	return nil
}
