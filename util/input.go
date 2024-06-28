package util

import (
	"os"
	"strings"
)

var (
	InputFile    string
	OutputFile   string
	InputFormat  string
	OutputFormat string

	Input []byte
)

func Init() {
	if InputFile != "" {
		if InputFormat == "" {
			InputFormat = parseFormat(InputFile)
		}
		var err error
		Input, err = os.ReadFile(InputFile)
		if err != nil {
			ShowFatal(err)
		}
	} else {
		Input = ReadStdin()
	}
	if OutputFormat == "" && OutputFile != "" {
		OutputFormat = parseFormat(OutputFile)
	}
}

func parseFormat(path string) string {
	i := strings.LastIndex(path, ".")
	if i < 0 {
		return ""
	}
	return path[i+1:]
}
