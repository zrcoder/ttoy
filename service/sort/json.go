package sort

import "encoding/json"

// JSON takes a given JSON string as input, sorts the keys in lexicographic order and outputs a new JSON string.
// The JSON input must be a valid JSON object, such as a dictionary or map, in order for the function to work.
// The function returns an empty string with the error if there is a parsing or sorting error.
func JSON(input string) (output string, err error) {
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
