package util

import (
	"os"

	"github.com/charmbracelet/huh"
)

func NewText(title, extension string, bind *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title(title).
				Value(bind).
				Editor(getEditor()).
				EditorExtension(extension),
		),
	)
}

func NewSelect(title string, options []string, bind *string) *huh.Form {
	ops := make([]huh.Option[string], len(options))
	for i, o := range options {
		ops[i] = huh.NewOption[string](o, o)
	}
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(title).
				Value(bind).
				Options(
					ops...,
				),
		),
	)
}

func NewInput(title string, bind *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Placeholder(title).
				Value(bind),
		),
	)
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	return editor
}
