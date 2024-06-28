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

func ShowError(err error) {
	fmt.Println(err)
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
	if OutputFile != "" {
		return os.WriteFile(OutputFile, buf.Bytes(), 0o640)
	}

	fmt.Println(buf.String())
	return nil
}
