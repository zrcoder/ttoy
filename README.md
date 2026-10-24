# ttoy

ttoy is a smart APP contains dev tools and mini-games.

> This version is a native desktop app based on [mygo](https://github.com/egoist/mygo) (pure Go, no JS/HTML/CSS).

## Commands

```bash
# Core workflow
task dev      # Run the application (go run .)
task build    # Build the application (mygo build -> build/)
task run      # Run the built binary
task test     # Run all tests
task vet      # Run go vet

# Direct go
go run .
go build .
go test ./...
go vet ./...
```

## Architecture

- **Entry point**: `main.go` — mygo app, creates window, mounts `Ttoy.View`
- **App model**: `app.go` — `Ttoy` struct holds all UI state; navigation + lazy-loaded panels
- **UI components** (root directory, package main): panels, editors, toasts, image viewer
  - `panels.go` — Format / Convert / Encode / Generate / Diff / View / Plot panels
  - `dual_editor.go` — dual editor for transform/encode
  - `image_viewer.go` — image display with save/copy
  - `ballsort_panel.go`, `hanoi_panel.go`, `ice_panel.go` — games
  - `config.go` — toast helpers and constants
- **Services**: `service/` — pure Go logic (Convert, Encode, Format, Generate, Plot, Save, Sort, View)
- **Games**: `game/` — pure Go logic (BallSort, Hanoi, IceMagic), UI polls state directly
- **Resources**: `resources/` — embedded icons, images, wiki docs (`//go:embed`)
- **Config**: `mygo.json` — mygo project config

## Build

Uses [go-task](https://taskfile.dev/). Production builds use `mygo build` (no Node, no Vite, no frontend).

```bash
task build    # output in build/
```

## Go Version

Go 1.27.3 (from `go.mod`). Key deps: `egoist/mygo v0.6.3`, `fogleman/gg`, `goplus/ixgo`.
