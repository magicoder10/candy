package main

import (
	"log"
	"time"

	"candy/assets"
	"candy/env"
	"candy/graphics"
	"candy/input"
	"candy/observability"
	"candy/screen"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var _ ebiten.Game = (*ebitenGame)(nil)


type ebitenGame struct {
	sp graphics.Sprite
	nanoPerUpdate int64
	updateTime time.Duration
	windowConfig WindowConfig
	prevTime time.Time
	lag int64
	ebiten *graphics.Ebiten
}

func (e *ebitenGame) Init()  {
	ebiten.SetWindowSize(e.windowConfig.Width, e.windowConfig.Height)
	ebiten.SetWindowTitle(e.windowConfig.Title)
	e.prevTime = time.Now()
}

func (e *ebitenGame) Update() error {
	now := time.Now()
	elapsed := now.Sub(e.prevTime)
	e.lag += elapsed.Nanoseconds()
	e.prevTime = now

	inputs := e.pollEvents()
	for _, in := range inputs {
		e.sp.HandleInput(in)
	}

	for e.lag >= e.nanoPerUpdate {
		e.sp.Update(e.updateTime)
		e.lag -= e.nanoPerUpdate
	}
	return nil
}

func (e ebitenGame) pollEvents() []input.Input {
	inputs := make([]input.Input, 0)
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.LeftArrowKey,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.RightArrowKey,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.UpArrowKey,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.DownArrowKey,
		})
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		inputs = append(inputs, input.Input{
			Action: input.Press,
			Device: input.RKey,
		})
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		inputs = append(inputs, input.Input{
			Action: input.SinglePress,
			Device: input.SpaceKey,
		})
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.LeftArrowKey,
		})
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.RightArrowKey,
		})
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.UpArrowKey,
		})
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.DownArrowKey,
		})
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyR) {
		inputs = append(inputs, input.Input{
			Action: input.Release,
			Device: input.RKey,
		})
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		inputs = append(inputs, input.Input{
			Action: input.SinglePress,
			Device: input.MouseLeftButton,
		})
	}
	return inputs
}

func (e ebitenGame) Draw(screen *ebiten.Image) {
	e.sp.Draw()
	e.ebiten.Render(screen)
}

func (e ebitenGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return e.windowConfig.Width, e.windowConfig.Height
}

type WindowConfig struct {
	Width int
	Height int
	Title string
}

func newEbitenGame(windowConfig WindowConfig, sp graphics.Sprite, framesPerSeconds int, ebiten *graphics.Ebiten) ebitenGame {
	nanoPerUpdate := time.Second.Nanoseconds() / int64(framesPerSeconds)
	return ebitenGame{
		sp:              sp,
		nanoPerUpdate:   nanoPerUpdate,
		updateTime: time.Duration(nanoPerUpdate),
		windowConfig: windowConfig,
		ebiten: ebiten,
	}
}

func main() {
	env.AutoLoad()

	eb := graphics.NewEbiten()

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

	g := newEbitenGame(WindowConfig{
		Width:  screen.Width,
		Height: screen.Height,
		Title:  "Candy",
	}, app, 24, &eb)
	g.Init()

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
