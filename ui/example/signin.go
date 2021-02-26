package main

import (
	"fmt"

	"candy/assets"
	"candy/audio"
	"candy/ui"
	"candy/ui/ptr"
)

var _ ui.Component = (*SignIn)(nil)

type SignIn struct {
	ui.SharedComponent
	backgroundMusic audio.Audio
	router          *ui.Router
}

func (s SignIn) Init() {
	s.SharedComponent.Init()
	s.backgroundMusic.Play()
}

func (s SignIn) Destroy() {
	s.SharedComponent.Destroy()
	s.backgroundMusic.Stop()
}

func NewSignIn(
	router *ui.Router,
	assets assets.Assets,
) *SignIn {
	return &SignIn{
		backgroundMusic: assets.GetAudio("screen/signin_bg.mp3"),
		SharedComponent: ui.SharedComponent{
			Name:   "SignIn",
			Layout: ui.NewLayout(ui.BoxLayoutType),
			Children: []ui.Component{
				ui.NewBox(
					nil,
					[]ui.Component{
						ui.NewButton(&ui.ButtonProps{
							Text: ptr.String("Sign In with Github"),
							OnClick: func() {
								router.Navigate("/demo", nil)
								fmt.Println("Sign In button clicked")
							},
						}, &ui.Style{
							Alignment: &ui.Alignment{
								Horizontal: ui.AlignHorizontalCenter.Ptr(),
							},
							FontStyle: &ui.FontStyle{
								Weight:     ptr.String("medium"),
								LineHeight: ptr.Int(20),
								Size:       ptr.Int(20),
							},
							Margin: &ui.EdgeSpacing{
								Top: ptr.Int(469),
							},
							Padding: &ui.EdgeSpacing{
								All:    nil,
								Top:    ptr.Int(12),
								Bottom: ptr.Int(12),
							},
							Width: ptr.Int(356),
							Background: &ui.Background{
								Color: &ui.Color{
									Red:   22,
									Green: 107,
									Blue:  107,
									Alpha: 255,
								},
							},
						}),
					},
					&ui.Style{
						Height: ptr.Int(Height),
						Alignment: &ui.Alignment{
							Horizontal: ui.AlignHorizontalCenter.Ptr(),
						},
						Background: &ui.Background{
							ImagePath: ptr.String("test/signin.png"),
						},
					},
				),
			},
		},
	}
}
