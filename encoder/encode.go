package encoder

import (
	"html"
	"net/url"

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
