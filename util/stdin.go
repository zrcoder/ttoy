package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadStdin read standard input string
func ReadStdin() string {
	if IsStdinEmpty() {
		return ""
	}

	reader := bufio.NewReader(os.Stdin)
	var buf strings.Builder

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if _, err = buf.WriteRune(r); err != nil {
			panic(err)
		}
	}

	return buf.String()
}

// IsStdinEmpty returns whether stdin is empty.
func IsStdinEmpty() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return true
	}

	if stat.Mode()&os.ModeNamedPipe == 0 && stat.Size() == 0 {
		return true
	}

	return false
}
