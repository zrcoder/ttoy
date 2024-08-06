package converter

import (
	"bytes"
	"strings"
	"testing"
)

type Case struct {
	name    string
	input   []byte
	want    string
	wantErr bool
	fn      func([]byte) (*bytes.Buffer, error)
}

func Test_Base(t *testing.T) {
	const (
		jsonStr = `
{
    "name": "Tom"
}
`
		yamlStr = `
name: Tom
`
		tomlStr = `
name = "Tom"
`
	)
	cases := []Case{
		{
			name:  "json2toml",
			fn:    Json2Toml,
			input: []byte(jsonStr),
			want:  tomlStr,
		},
		{
			name:  "toml2json",
			fn:    Toml2Json,
			input: []byte(tomlStr),
			want:  jsonStr,
		},
		{
			name:  "json2yaml",
			fn:    Json2Yaml,
			input: []byte(jsonStr),
			want:  yamlStr,
		},
		{
			name:  "yaml2json",
			fn:    Yaml2Json,
			input: []byte(yamlStr),
			want:  jsonStr,
		},
		{
			name:  "yaml2toml",
			fn:    Yaml2Toml,
			input: []byte(yamlStr),
			want:  tomlStr,
		},
		{
			name:  "toml2yaml",
			fn:    Toml2Yaml,
			input: []byte(tomlStr),
			want:  yamlStr,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf, err := c.fn(c.input)
			if (err != nil) != c.wantErr {
				t.Errorf("got error = %v, wantErr %v", err, c.wantErr)
				return
			}
			got := strings.TrimSpace(buf.String())
			want := strings.TrimSpace(c.want)
			if got != want {
				t.Errorf("got\n---\n%v\n---\nwant\n---\n%v\n---", got, want)
			}
		})
	}
}
