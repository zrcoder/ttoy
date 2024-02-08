package decoder

import (
	"net/url"

	"github.com/zrcoder/ttoy/util"
)

func UrlEncode() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return urlEncodeAndShow(stdinput)
	}
	return urlEncodeRun()
}

func UrlDecode() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return urlDecodeAndShow(stdinput)
	}
	return urlDecodeRun()
}

func urlEncodeAndShow(input string) error {
	res := url.QueryEscape(input)
	util.Show(res)
	return nil
}

func urlEncodeRun() error {
	input := ""
	if err := util.NewInput("Encode url", &input).Run(); err != nil {
		return err
	}
	return urlEncodeAndShow(input)
}

func urlDecodeAndShow(ori string) error {
	if res, err := url.QueryUnescape(ori); err != nil {
		return err
	} else {
		util.Show(res)
	}
	return nil
}

func urlDecodeRun() error {
	input := ""
	if err := util.NewInput("Docode url", &input).Run(); err != nil {
		return err
	}
	return urlDecodeAndShow(input)
}
