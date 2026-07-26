package service

import (
	"encoding/base64"
	"os"
	"strings"
)

type Save struct{}

func (s *Save) Image(path, dataURL string) error {
	const sep = ";base64,"
	_, after, ok := strings.Cut(dataURL, sep)
	if !ok {
		return nil
	}

	data, err := base64.StdEncoding.DecodeString(after)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
