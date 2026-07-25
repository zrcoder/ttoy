package convert

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Converter struct{}

func New() *Converter {
	return &Converter{}
}

func (c *Converter) Json2Toml(data string) (string, error) {
	return transform(data, func(b []byte) (*bytes.Buffer, error) {
		var obj any
		err := json.Unmarshal(b, &obj)
		if err != nil {
			return nil, err
		}
		writer := bytes.NewBuffer(nil)
		err = toml.NewEncoder(writer).Encode(obj)
		return writer, err
	})
}

func (c *Converter) Toml2Json(data string) (string, error) {
	return transform(data, func(b []byte) (*bytes.Buffer, error) {
		var obj any
		_, err := toml.Decode(string(b), &obj)
		if err != nil {
			return nil, err
		}
		writer := bytes.NewBuffer(nil)
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(obj)
		return writer, err
	})
}

func (c *Converter) Json2Yaml(data string) (string, error) {
	return transform(data, func(b []byte) (*bytes.Buffer, error) {
		var obj any
		err := json.Unmarshal(b, &obj)
		if err != nil {
			return nil, err
		}
		writer := bytes.NewBuffer(nil)
		encoder := yaml.NewEncoder(writer)
		err = encoder.Encode(obj)
		return writer, err
	})
}

func (c *Converter) Yaml2Json(data string) (string, error) {
	return transform(data, func(b []byte) (*bytes.Buffer, error) {
		var obj any
		err := yaml.Unmarshal(b, &obj)
		if err != nil {
			return nil, err
		}
		writer := bytes.NewBuffer(nil)
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "    ")
		err = encoder.Encode(obj)
		return writer, err
	})
}

func (c *Converter) Yaml2Toml(data string) (string, error) {
	return transform(data, func(b []byte) (*bytes.Buffer, error) {
		var obj any
		err := yaml.Unmarshal(b, &obj)
		if err != nil {
			return nil, err
		}
		writer := bytes.NewBuffer(nil)
		err = toml.NewEncoder(writer).Encode(obj)
		return writer, err
	})
}

func (c *Converter) Toml2Yaml(data string) (string, error) {
	return transform(data, func(b []byte) (*bytes.Buffer, error) {
		var obj any
		_, err := toml.Decode(string(b), &obj)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(nil)
		encoder := yaml.NewEncoder(buf)
		err = encoder.Encode(obj)
		return buf, err
	})
}

type Transformer = func([]byte) (*bytes.Buffer, error)

func transform(input string, fn Transformer) (string, error) {
	buf, err := fn([]byte(input))
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
