package encoder

import (
	"bytes"
	"html"
	"net/url"
	"os"

	decqr "github.com/tuotoo/qrcode"

	"github.com/zrcoder/ttoy/util"
)

func EncodeHtml(input []byte) {
	res := html.EscapeString(string(input))
	util.Show([]byte(res))
}

func DecodeHtml(input []byte) {
	res := html.UnescapeString(string(input))
	util.Show([]byte(res))
}

func EncodeUrl(input string) {
	res := url.QueryEscape(input)
	util.Show([]byte(res))
}

func DecodeUrl(input string) {
	if res, err := url.QueryUnescape(input); err != nil {
		util.ShowFatal(err)
	} else {
		util.Show([]byte(res))
	}
}

func DecodeQr(img []byte) {
	buf := bytes.NewBuffer(img)
	res, err := decqr.Decode(buf)
	if err != nil {
		util.ShowFatal(err)
	}
	if util.OutputFile == "" {
		util.Show([]byte(res.Content))
		return
	}
	err = os.WriteFile(util.OutputFile, []byte(res.Content), 0o640)
	if err != nil {
		util.ShowFatal(err)
	}
}
