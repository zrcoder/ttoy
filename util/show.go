package util

import (
	"bytes"
	"fmt"
	"os"

	lg "github.com/charmbracelet/lipgloss"
)

var OutputFile string

func Show(data []byte) {
	if OutputFile != "" {
		err := os.WriteFile(OutputFile, data, 0o640)
		if err != nil {
			ShowFatal(err)
		}
	}
	fmt.Println(string(data))
}

func ShowCode(code string, data []byte) {
	Show(data)
}

func ShowFatal(err error) {
	fmt.Println(lg.NewStyle().Foreground(Red).Render(err.Error()))
	os.Exit(1)
}

func ShowKVs(res ...string) {
	buf := bytes.NewBuffer(nil)
	for i, v := range res {
		buf.WriteString(v)
		if i%2 == 0 {
			buf.WriteString("\t")
		} else {
			buf.WriteByte('\n')
		}
	}
	Show(buf.Bytes())
}
