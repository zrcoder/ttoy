package decoder

import (
	"net/url"

	"github.com/charmbracelet/huh"
	"github.com/zrcoder/ttoy/result"
)

func encodeUrl() error {
	ori := ""
	input := huh.NewInput().Title("Encode url").Value(&ori)
	if err := input.Run(); err != nil {
		return err
	}
	res := url.QueryEscape(ori)
	result.Show(res)
	return nil
}
func decodeUrl() error {
	ori := ""
	input := huh.NewInput().Title("Docode url").Value(&ori)
	if err := input.Run(); err != nil {
		return err
	}
	res, err := url.QueryUnescape(ori)
	if err != nil {
		return err
	}
	result.Show(res)
	return nil
}
