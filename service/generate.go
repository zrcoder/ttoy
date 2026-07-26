package service

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"

	"github.com/common-nighthawk/go-figure"
)

type Generate struct{}

func (g *Generate) AsciiArt(input, font string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}
	f := figure.NewFigure(input, font, true)
	return f.String(), nil
}

func (g *Generate) Hash(input string) (map[string]string, error) {
	data := []byte(input)
	h := md5.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resMd5 := hex.EncodeToString(h.Sum(nil))

	h = sha1.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resSha1 := hex.EncodeToString(h.Sum(nil))

	h = sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resSha256 := hex.EncodeToString(h.Sum(nil))

	h = sha512.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	resSha512 := hex.EncodeToString(h.Sum(nil))

	return map[string]string{
		"md5":    resMd5,
		"sha1":   resSha1,
		"sha256": resSha256,
		"sha512": resSha512,
	}, nil
}
