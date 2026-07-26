package service

import (
	"encoding/base64"
	"html"
	"net/url"
)

type Encode struct{}

func (e *Encode) Base64Encode(input string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}

func (e *Encode) Base64Decode(input string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(input)
	return string(res), err
}

func (e *Encode) URLEncode(input string) (string, error) {
	return url.QueryEscape(input), nil
}

func (e *Encode) URLDecode(input string) (string, error) {
	return url.QueryUnescape(input)
}

func (e *Encode) HTMLEncode(input string) (string, error) {
	return html.EscapeString(input), nil
}

func (e *Encode) HTMLDecode(input string) (string, error) {
	return html.UnescapeString(input), nil
}
