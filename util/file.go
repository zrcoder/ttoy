package util

import (
	"flag"
	"os"
	"strings"
)

var (
	InputFile  string
	OutputFile string

	Input []byte
)

func init() {
	flag.StringVar(&InputFile, "i", "", "input file")
	flag.StringVar(&OutputFile, "o", "", "output file")
	flag.Parse()
	var err error
	if InputFile != "" {
		Input, err = os.ReadFile(InputFile)
	} else {
		Input = ReadStdin()
	}
	if err != nil {
		ShowError(err)
		os.Exit(1)
	}
}

func ParseInputExtension() string {
	i := strings.LastIndex(InputFile, ".")
	if i < 0 {
		return ""
	}
	return InputFile[i+1:]
}
