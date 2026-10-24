# ttoy

A native desktop app of developer tools and mini-games, built with [mygo](https://github.com/egoist/mygo) — pure Go, no JS/HTML/CSS.

## Commands

```bash
task dev      # go run .
task build    # mygo build -> build/
task run      # run the built binary
task test     # go test ./...
task vet      # go vet ./...
```

## Structure

```
.
├── main.go            # entry: mygo app + window
├── app.go             # Ttoy model: sidebar nav + lazy panels
├── panels.go          # Format / Convert / Encode / Generate / Diff / View / Plot
├── dual_editor.go     # shared two-pane editor
├── image_viewer.go    # SVG/PNG viewer (SVG rasterized for text)
├── svg_raster.go      # macOS Quick Look SVG -> PNG
├── config.go          # toast helpers
├── service/           # pure Go: convert, encode, format, generate, plot, save, sort, view
└── game/              # pure Go: IceMagic, BallSort, Hanoi
```

## Build

Requires [go-task](https://taskfile.dev/). `task build` runs `mygo build` — no Node, no Vite, no frontend.
