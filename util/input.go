package util

import (
	"os"
	"path/filepath"
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
			InputFormat = filepath.Ext(InputFile)
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
		OutputFormat = filepath.Ext(OutputFile)
	}
}
