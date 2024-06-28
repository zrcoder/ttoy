package cmd

import (
	"github.com/zrcoder/ttoy/encoder"
	"github.com/zrcoder/ttoy/util"
)

type Action func(input []byte) error

var UnknowFormat = "unknow format"

func do(action Action) {
	if action == nil {
		util.ShowFatal(UnknowFormat)
	}
	err := action(util.Input)
	if err != nil {
		util.ShowFatal(err)
	}
}

var url string

func encodeOrDecode(encode bool) error {
	if util.InputFormat == "html" || util.OutputFormat == "html" {
		if encode {
			return encoder.EncodeHtml(util.Input)
		}
		return encoder.DecodeHtml(util.Input)
	}
	if url == "" {
		url = string(util.Input)
	}
	if encode {
		return encoder.EncodeUrl(url)
	}
	return encoder.DecodeUrl(url)
}
