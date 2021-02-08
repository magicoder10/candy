package main

import (
    "log"

    "candy/assets"
    "candy/graphics"
    "candy/observability"
    "candy/screen"
    "github.com/hajimehoshi/ebiten/v2"
)

func main() {
    eb := graphics.NewEbiten(true)

    ass, err := assets.LoadAssets("public")
    if err != nil {
        panic(err)
    }

    logger := observability.NewLogger(observability.Info)

    app, err := screen.NewApp(&logger, ass, &eb)
    if err != nil {
        panic(err)
    }
    err = app.Launch()
    if err != nil {
        panic(err)
    }

    g := graphics.NewEbitenWindow(graphics.WindowConfig{
        Width:  screen.Width,
        Height: screen.Height,
        Title:  "Candy",
    }, app, 24, &eb)
    g.Init()

    if err := ebiten.RunGame(&g); err != nil {
        log.Fatal(err)
    }
}
