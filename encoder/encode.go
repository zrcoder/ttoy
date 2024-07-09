package encoder

import (
	"bytes"
	"html"
	"net/url"
	"os"

	decqr "github.com/tuotoo/qrcode"

	"github.com/zrcoder/ttoy/util"
)

func EncodeHtml(input []byte) error {
	res := html.EscapeString(string(input))
	return util.Show([]byte(res))
}

func DecodeHtml(input []byte) error {
	res := html.UnescapeString(string(input))
	return util.Show([]byte(res))
}

func EncodeUrl(input string) error {
	res := url.QueryEscape(input)
	return util.Show([]byte(res))
}

func DecodeUrl(input string) error {
	if res, err := url.QueryUnescape(input); err != nil {
		return err
	} else {
		return util.Show([]byte(res))
	}
}

func DecodeQr(img []byte) error {
	buf := bytes.NewBuffer(img)
	res, err := decqr.Decode(buf)
	if err != nil {
		return err
	}
	if util.OutputFile == "" {
		return util.Show([]byte(res.Content))
	}

	return os.WriteFile(util.OutputFile, []byte(res.Content), 0o640)
}
