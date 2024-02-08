package util

import (
	"os"

	"github.com/charmbracelet/huh"
)

func NewText(title, extension string, bind *string) *huh.Text {
	return huh.NewText().Title(title).Value(bind).Editor(getEditor()).EditorExtension(extension)
}

func NewInput(title string, bind *string) *huh.Input {
	return huh.NewInput().Title(title).Value(bind)
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	return editor
}
