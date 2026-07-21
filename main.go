package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/zrcoder/ttoy/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "TToy",
		Description: "A smart APP contains dev tools",
		Services: []application.Service{
			application.NewService(service.New()),
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
	})

	log.Fatal(app.Run())
}
