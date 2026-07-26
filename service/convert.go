package service

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Convert struct{}

func (c *Convert) Json2Toml(data string) (string, error) {
	var obj any
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(obj); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *Convert) Toml2Json(data string) (string, error) {
	var obj any
	if _, err := toml.Decode(data, &obj); err != nil {
		return "", err
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(obj); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *Convert) Json2Yaml(data string) (string, error) {
	var obj any
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := yaml.NewEncoder(&buf).Encode(obj); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *Convert) Yaml2Json(data string) (string, error) {
	var obj any
	if err := yaml.Unmarshal([]byte(data), &obj); err != nil {
		return "", err
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(obj); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *Convert) Yaml2Toml(data string) (string, error) {
	var obj any
	if err := yaml.Unmarshal([]byte(data), &obj); err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(obj); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *Convert) Toml2Yaml(data string) (string, error) {
	var obj any
	if _, err := toml.Decode(data, &obj); err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := yaml.NewEncoder(&buf).Encode(obj); err != nil {
		return "", err
	}
	return buf.String(), nil
}
