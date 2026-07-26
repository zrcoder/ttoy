package service

import "encoding/json"

type Sort struct{}

func New() *Sort {
	return &Sort{}
}

func (s *Sort) JSON(input string) (output string, err error) {
	jsonb := []byte(input)
	var obj any
	err = json.Unmarshal(jsonb, &obj)
	if err != nil {
		return "", err
	}
	out, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}
