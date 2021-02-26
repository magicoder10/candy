package graphics

import (
	"image"
	"time"

	"candy/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type WindowConfig struct {
	Width  int
	Height int
	Title  string
}

var _ ebiten.Game = (*EbitenWindow)(nil)

type EbitenWindow struct {
	sp            Sprite
	nanoPerUpdate int64
	updateTime    time.Duration
	windowConfig  WindowConfig
	prevTime      time.Time
	lag           int64
	ebiten        *Ebiten
}

func (e *EbitenWindow) Init() {
	ebiten.SetWindowSize(e.windowConfig.Width, e.windowConfig.Height)
	ebiten.SetWindowTitle(e.windowConfig.Title)
	e.ebiten.initBuffer(e.windowConfig.Width, e.windowConfig.Height)
	e.prevTime = time.Now()
}

func (e *EbitenWindow) Update() error {
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

func (e *EbitenWindow) Draw(screen *ebiten.Image) {
	e.sp.Draw()
	e.ebiten.Render(screen)
}

func (e *EbitenWindow) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return e.windowConfig.Width, e.windowConfig.Height
}

func (e *EbitenWindow) pollEvents() []input.Input {
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
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		inputs = append(inputs, input.Input{
			Action: input.SinglePress,
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
		x, y := ebiten.CursorPosition()
		inputs = append(inputs, input.Input{
			Action:         input.SinglePress,
			Device:         input.MouseLeftButton,
			CursorPosition: image.Point{X: x, Y: y},
		})
	}
	return inputs
}

func NewEbitenWindow(windowConfig WindowConfig, sp Sprite, framesPerSeconds int, eb *Ebiten) *EbitenWindow {
	nanoPerUpdate := time.Second.Nanoseconds() / int64(framesPerSeconds)
	return &EbitenWindow{
		sp:            sp,
		nanoPerUpdate: nanoPerUpdate,
		updateTime:    time.Duration(nanoPerUpdate),
		windowConfig:  windowConfig,
		ebiten:        eb,
	}
}
