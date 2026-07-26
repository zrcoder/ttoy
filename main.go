package main

import (
	"embed"
	"log"

	"github.com/zrcoder/ttoy/service"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "TToy",
		Description: "A smart APP contains dev tools",
		Services: []application.Service{
			application.NewService(new(service.Convert)),
			application.NewService(new(service.Format)),
			application.NewService(new(service.Encode)),
			application.NewService(new(service.Generate)),
			application.NewService(new(service.View)),
			application.NewService(new(service.Plot)),
			application.NewService(new(service.Sort)),
			application.NewService(new(service.Save)),
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
