package main

import (
	"fmt"
	"os"

	"github.com/zrcoder/ttoy/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
