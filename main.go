package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/zrcoder/ttoy/service/convert"
	"github.com/zrcoder/ttoy/service/encode"
	"github.com/zrcoder/ttoy/service/format"
	"github.com/zrcoder/ttoy/service/generate"
	"github.com/zrcoder/ttoy/service/sort"
	"github.com/zrcoder/ttoy/service/plot"
	"github.com/zrcoder/ttoy/service/view"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "TToy",
		Description: "A smart APP contains dev tools",
		Services: []application.Service{
			application.NewService(convert.New()),
			application.NewService(format.New()),
			application.NewService(encode.New()),
			application.NewService(generate.New()),
			application.NewService(view.New()),
			application.NewService(plot.New()),
			application.NewService(sort.New()),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "TToy",
		Width:            1536,
		Height:           1024,
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
	})

	log.Fatal(app.Run())
}
