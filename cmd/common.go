package cmd

import (
	"os"
	"path/filepath"

	"github.com/zrcoder/ttoy/encoder"
	"github.com/zrcoder/ttoy/util"
)

type Action func(input []byte) error

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
			InputFormat = filepath.Ext(InputFile)
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
		OutputFormat = filepath.Ext(OutputFile)
	}
	util.OutputFile = OutputFile
}

func do(action Action) {
	if action == nil {
		util.ShowFatal(UnknowFormat)
	}
	err := action(Input)
	if err != nil {
		util.ShowFatal(err)
	}
}

var url string

func encodeOrDecode(encode bool) error {
	if InputFormat == "html" || OutputFormat == "html" {
		if encode {
			return encoder.EncodeHtml(Input)
		}
		return encoder.DecodeHtml(Input)
	}
	if url == "" {
		url = string(Input)
	}
	if encode {
		return encoder.EncodeUrl(url)
	}
	return encoder.DecodeUrl(url)
}
