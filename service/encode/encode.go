package encode

import (
	"encoding/base64"
	"html"
	"net/url"
)

type Encoder struct{}

func New() *Encoder {
	return &Encoder{}
}

func (e *Encoder) Base64Encode(input string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}

func (e *Encoder) Base64Decode(input string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(input)
	return string(res), err
}

func (e *Encoder) URLEncode(input string) (string, error) {
	return url.QueryEscape(input), nil
}

func (e *Encoder) URLDecode(input string) (string, error) {
	return url.QueryUnescape(input)
}

func (e *Encoder) HTMLEncode(input string) (string, error) {
	return html.EscapeString(input), nil
}

func (e *Encoder) HTMLDecode(input string) (string, error) {
	return html.UnescapeString(input), nil
}
