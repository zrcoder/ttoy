package util

import (
	"bytes"
	"fmt"
	"os"
)

func Show(data []byte) error {
	if OutputFile != "" {
		return os.WriteFile(OutputFile, data, 0o640)
	}
	fmt.Println(string(data))
	return nil
}

func ShowCode(code string, data []byte) error {
	return Show(data)
}

func ShowFatal(err any) {
	fmt.Println(err)
	os.Exit(1)
}

func ShowKVs(res ...string) error {
	buf := bytes.NewBuffer(nil)
	for i, v := range res {
		buf.WriteString(v)
		if i%2 == 0 {
			buf.WriteString("\t")
		} else {
			buf.WriteByte('\n')
		}
	}
	return Show(buf.Bytes())
}
