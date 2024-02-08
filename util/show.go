package util

import (
	"fmt"
)

func Show(s string) {
	fmt.Println(s)
}

func ShowCode(code, content string) {
	fmt.Println(content)
}

func ShowError(err error) {
	fmt.Println(err)
}

func ShowKVs(res ...string) {
	for i, v := range res {
		fmt.Print(v)
		if i%2 == 0 {
			fmt.Print("\t")
		} else {
			fmt.Println()
		}
	}
}
