package decoder

import (
	"html"

	"github.com/charmbracelet/huh"
	"github.com/zrcoder/ttoy/result"
)

func encodeHtml() error {
	ori := ""
	if err := huh.NewText().Title("Encode html").Value(&ori).Editor("nvim").Run(); err != nil {
		return err
	}
	res := html.EscapeString(ori)
	result.Show(res)
	return nil
}

func decodeHtml() error {
	ori := ""
	if err := huh.NewText().Title("Decode html").Value(&ori).Editor("nvim").Run(); err != nil {
		return err
	}
	res := html.UnescapeString(ori)
	result.Show(res)
	return nil

}
