package generator

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/charmbracelet/huh"
	"github.com/zrcoder/ttoy/result"
)

func hash() error {
	ori := ""

	if err := huh.NewText().Title("Hash").Value(&ori).Editor("nvim").Run(); err != nil {
		return err
	}

	h := md5.New()
	if _, err := h.Write([]byte(ori)); err != nil {
		return err
	}
	resMd5 := hex.EncodeToString(h.Sum(nil))

	h = sha1.New()
	if _, err := h.Write([]byte(ori)); err != nil {
		return err
	}
	resSha1 := hex.EncodeToString(h.Sum(nil))

	h = sha256.New()
	if _, err := h.Write([]byte(ori)); err != nil {
		return err
	}
	resSha256 := hex.EncodeToString(h.Sum(nil))

	h = sha512.New()
	if _, err := h.Write([]byte(ori)); err != nil {
		return err
	}
	resSha512 := hex.EncodeToString(h.Sum(nil))

	result.ShowKVs(
		"md5:", resMd5,
		"sha1:", resSha1,
		"sha256:", resSha256,
		"sha512:", resSha512,
	)

	return nil
}
