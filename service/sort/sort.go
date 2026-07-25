package sort

import "encoding/json"

type Sorter struct{}

func New() *Sorter {
	return &Sorter{}
}

func (s *Sorter) JSON(input string) (output string, err error) {
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
