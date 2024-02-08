package formatter

import (
	"github.com/yosssi/gohtml"
	"github.com/zrcoder/ttoy/util"
)

func Html() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return htmlFormatAndShow(stdinput)
	}
	return htmlRun()
}

func htmlFormatAndShow(origin string) error {
	res := gohtml.Format(origin)
	util.ShowCode("html", res)
	return nil
}

func htmlRun() error {
	data := ""

	if err := util.NewText("xml/html formatter", "html", &data).Run(); err != nil {
		return err
	}

	return htmlFormatAndShow(data)
}
