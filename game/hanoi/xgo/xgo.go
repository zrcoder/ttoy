package xgo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/xgobuild"
	_ "github.com/zrcoder/ttoy/game/hanoi/exported/github.com/zrcoder/ttoy/game/hanoi/export"
)

// codePrelude is prepended to user code. The dot import lets user code call
// A/B/C unqualified; Go+ (xgo) supplies the implicit "package main".
const codePrelude = `import . "github.com/zrcoder/ttoy/game/hanoi/export"
var (
	a = A
	b = B
	c = C
)
`

// RunCode compiles and runs user code that drives the bound game via A, B, C.
func RunCode(code string) error {
	_, err := ixgo.RunFile("main.xgo", codePrelude+code, nil, 0)
	if err != nil {
		return parseBuildError(err)
	}
	return nil
}

// parseBuildError translates an ixgo compile error into a line number relative
// to the user code, since codePrelude adds lines before it.
func parseBuildError(err error) error {
	msg := err.Error()
	before, after, ok := strings.Cut(msg, " ")
	if !ok {
		return err
	}
	arr := strings.SplitN(before, ":", 3)
	if len(arr) < 2 {
		return err
	}
	line, convErr := strconv.Atoi(arr[1])
	if convErr != nil {
		return err
	}
	line -= strings.Count(codePrelude, "\n")
	if line < 1 {
		return err
	}
	return fmt.Errorf("line %d: %s", line, after)
}
