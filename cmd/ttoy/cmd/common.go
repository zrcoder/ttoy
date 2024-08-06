package cmd

import (
	"errors"
	"os"

	"github.com/zrcoder/ttoy/pkg/encoder"
	"github.com/zrcoder/ttoy/pkg/util"
)

type Action func(input []byte)

const UnknowFormat = "unknow format"

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
			InputFormat = extWithoutDot(InputFile)
		}
		var err error
		Input, err = os.ReadFile(InputFile)
		if err != nil {
			util.ShowFatal(err)
		}
	} else {
		Input = util.ReadStdin()
	}
	if OutputFormat == "" && OutputFile != "" {
		OutputFormat = extWithoutDot(OutputFile)
	}
	util.OutputFile = OutputFile
}

func do(action Action) {
	if action == nil {
		util.ShowFatal(errors.New(UnknowFormat))
	}
	action(Input)
}

var url string

func encodeOrDecode(encode bool) {
	if InputFormat == "html" || OutputFormat == "html" {
		if encode {
			encoder.EncodeHtml(Input)
		} else {
			encoder.DecodeHtml(Input)
		}
	}
	if url == "" {
		url = string(Input)
	}
	if encode {
		encoder.EncodeUrl(url)
	} else {
		encoder.DecodeUrl(url)
	}
}

func extWithoutDot(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}
