package formatter

import (
	"github.com/charmbracelet/huh"
	"github.com/yosssi/gohtml"
	"github.com/zrcoder/ttoy/result"
)

func formatXML() error {
	data := ""

	if err := huh.NewText().Title("xml/html formatter").Value(&data).Editor("nvim").Run(); err != nil {
		return err
	}

	res := gohtml.Format(data)

	result.ShowCode("html", res)
	return nil
}
