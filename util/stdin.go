package util

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// ReadStdin read standard input
func ReadStdin() []byte {
	if IsStdinEmpty() {
		return nil
	}

	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

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

	return buf.Bytes()
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
