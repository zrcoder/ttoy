package service

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Format struct{}

func (f *Format) Html(data string) (string, error) {
	// Simple native HTML formatting: indent block tags.
	var buf strings.Builder
	indent := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '<' && i+1 < len(data) && data[i+1] != '/' {
			j := strings.IndexByte(data[i:], '>')
			if j == -1 {
				buf.WriteByte(data[i])
				continue
			}
			tag := data[i+1 : i+j]
			if strings.HasPrefix(tag, "/") {
				indent--
				if indent < 0 {
					indent = 0
				}
			}
			for k := 0; k < indent; k++ {
				buf.WriteString("  ")
			}
			buf.WriteString(data[i : i+j+1])
			buf.WriteByte('\n')
			if !strings.HasPrefix(tag, "/") && !isVoidTag(tag) {
				indent++
			}
			i += j
		} else {
			buf.WriteByte(data[i])
		}
	}
	return buf.String(), nil
}

func isVoidTag(tag string) bool {
	switch tag {
	case "area", "base", "br", "col", "embed", "hr", "img", "input",
		"link", "meta", "param", "source", "track", "wbr":
		return true
	}
	return false
}

func (f *Format) Json(data string) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, []byte(data), "", "    ")
	return buf.String(), err
}

func (f *Format) Toml(data string) (string, error) {
	var obj any
	err := toml.Unmarshal([]byte(data), &obj)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	err = toml.NewEncoder(buf).Encode(obj)
	return buf.String(), err
}

func (f *Format) Yaml(data string) (string, error) {
	var obj any
	err := yaml.Unmarshal([]byte(data), &obj)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	err = encoder.Encode(&obj)
	return buf.String(), err
}
