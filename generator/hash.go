package generator

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/zrcoder/ttoy/util"
)

func Hash() error {
	if stdinput := util.ReadStdin(); stdinput != "" {
		return hashAndShow(stdinput)
	}
	return hashRun()
}

func hashAndShow(input string) error {
	h := md5.New()
	if _, err := h.Write([]byte(input)); err != nil {
		return err
	}
	resMd5 := hex.EncodeToString(h.Sum(nil))

	h = sha1.New()
	if _, err := h.Write([]byte(input)); err != nil {
		return err
	}
	resSha1 := hex.EncodeToString(h.Sum(nil))

	h = sha256.New()
	if _, err := h.Write([]byte(input)); err != nil {
		return err
	}
	resSha256 := hex.EncodeToString(h.Sum(nil))

	h = sha512.New()
	if _, err := h.Write([]byte(input)); err != nil {
		return err
	}
	resSha512 := hex.EncodeToString(h.Sum(nil))

	util.ShowKVs(
		"md5:", resMd5,
		"sha1:", resSha1,
		"sha256:", resSha256,
		"sha512:", resSha512,
	)

	return nil
}

func hashRun() error {
	input := ""

	if err := util.NewText("Hash", "txt", &input).Run(); err != nil {
		return err
	}

	return hashAndShow(input)
}
