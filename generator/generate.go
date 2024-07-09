package generator

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ChimeraCoder/gojson"
	"github.com/google/uuid"
	genqr "github.com/skip2/go-qrcode"
	"github.com/zrcoder/cdor"
	ndor "github.com/zrcoder/ndor/pkg"

	"github.com/zrcoder/ttoy/util"
)

func Hash(input []byte) error {
	h := md5.New()
	if _, err := h.Write(input); err != nil {
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

	return util.ShowKVs(
		"md5:", resMd5,
		"sha1:", resSha1,
		"sha256:", resSha256,
		"sha512:", resSha512,
	)
}

var StructOption = struct {
	Name          string
	Pkg           string
	Tags          []string
	ConvertFloats bool
	SubStruct     bool
}{}

func Json2Struct(input []byte) error {
	buf := bytes.NewBuffer(input)
	out, err := gojson.Generate(buf, gojson.ParseJson,
		StructOption.Name,
		StructOption.Pkg,
		StructOption.Tags,
		StructOption.SubStruct,
		StructOption.ConvertFloats)
	if err != nil {
		return err
	}
	return util.Show(out)
}

func Json2Svg(inputJson []byte) error {
	c := cdor.Ctx()
	c.Json(string(inputJson))
	data, err := c.Gen()
	if err != nil {
		return err
	}
	return util.Show(data)
}

func Uuid(input []byte) error {
	res := uuid.NewString()
	return util.Show([]byte(res))
}

func D2(input []byte) error {
	c := cdor.Ctx().D2(string(input))
	data, err := c.Gen()
	if err != nil {
		return err
	}
	return util.Show(data)
}

func Ndor(input []byte) error {
	data, err := ndor.Gen(0, 0, string(input))
	if err != nil {
		return fmt.Errorf("line %d: %s", err.Number, err.Msg)
	}
	if util.OutputFile == "" {
		util.OutputFile = "ttoy.ndor.png"
		fmt.Println("output:", util.OutputFile)
	}
	return os.WriteFile(util.OutputFile, data, 0o640)
}

func Qrcode(input []byte) error {
	qr, err := genqr.New(string(input), genqr.Medium)
	if err != nil {
		return err
	}
	if util.OutputFile == "" {
		return util.Show([]byte(qr.ToSmallString(false)))
	}
	img, err := qr.PNG(-1)
	if err != nil {
		return err
	}
	return os.WriteFile(util.OutputFile, img, 0o640)
}
