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

func Hash(input []byte) {
	h := md5.New()
	if _, err := h.Write(input); err != nil {
		util.ShowFatal(err)
	}
	resMd5 := hex.EncodeToString(h.Sum(nil))

	h = sha1.New()
	if _, err := h.Write(input); err != nil {
		util.ShowFatal(err)
	}
	resSha1 := hex.EncodeToString(h.Sum(nil))

	h = sha256.New()
	if _, err := h.Write(input); err != nil {
		util.ShowFatal(err)
	}
	resSha256 := hex.EncodeToString(h.Sum(nil))

	h = sha512.New()
	if _, err := h.Write(input); err != nil {
		util.ShowFatal(err)
	}
	resSha512 := hex.EncodeToString(h.Sum(nil))

	util.ShowKVs(
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

func Json2Struct(input []byte) {
	buf := bytes.NewBuffer(input)
	out, err := gojson.Generate(buf, gojson.ParseJson,
		StructOption.Name,
		StructOption.Pkg,
		StructOption.Tags,
		StructOption.SubStruct,
		StructOption.ConvertFloats)
	if err != nil {
		util.ShowFatal(err)
	}
	util.Show(out)
}

func Json2Svg(inputJson []byte) {
	c := cdor.Ctx()
	c.Json(string(inputJson))
	data, err := c.Gen()
	if err != nil {
		util.ShowFatal(err)
	}
	util.Show(data)
}

func Uuid(input []byte) {
	res := uuid.NewString()
	util.Show([]byte(res))
}

func D2(input []byte) {
	c := cdor.Ctx().D2(string(input))
	data, err := c.Gen()
	if err != nil {
		util.ShowFatal(err)
	}
	util.Show(data)
}

func Ndor(input []byte) {
	data, lineErr := ndor.Gen(0, 0, string(input))
	if lineErr != nil {
		util.ShowFatal(fmt.Errorf("line %d: %s", lineErr.Number, lineErr.Msg))
	}
	if util.OutputFile == "" {
		util.OutputFile = "ttoy.ndor.png"
		fmt.Println("output:", util.OutputFile)
	}
	err := os.WriteFile(util.OutputFile, data, 0o640)
	if err != nil {
		util.ShowFatal(err)
	}
}

func Qrcode(input []byte) {
	qr, err := genqr.New(string(input), genqr.Medium)
	if err != nil {
		util.ShowFatal(err)
	}
	if util.OutputFile == "" {
		util.Show([]byte(qr.ToSmallString(false)))
		return
	}
	img, err := qr.PNG(-1)
	if err != nil {
		util.ShowFatal(err)
	}
	err = os.WriteFile(util.OutputFile, img, 0o640)
	if err != nil {
		util.ShowFatal(err)
	}
}
