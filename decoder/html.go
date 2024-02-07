package decoder

import (
	"html"

	"github.com/zrcoder/ttoy/util"
)

func HtmlEncode() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return htmlEncodeAndShow(stdinput)
	}
	return htmlEncodeRun()
}

func HtmlDecode() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return htmlDecodeAndShow(stdinput)
	}
	return htmlDecodeRun()
}

func htmlEncodeAndShow(input string) error {
	res := html.EscapeString(input)
	util.Show(res)
	return nil
}

func htmlEncodeRun() error {
	input := ""
	if err := util.NewText("Encode html", "html", &input).Run(); err != nil {
		return err
	}
	return htmlEncodeAndShow(input)
}

func htmlDecodeAndShow(ori string) error {
	res := html.UnescapeString(ori)
	util.Show(res)
	return nil
}

func htmlDecodeRun() error {
	ori := ""
	if err := util.NewText("Decode html", "html", &ori).Run(); err != nil {
		return err
	}
	return htmlDecodeAndShow(ori)
}
