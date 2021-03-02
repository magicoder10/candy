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
	//s.backgroundMusic.Play()
}

func (s SignIn) Destroy() {
	s.SharedComponent.Destroy()
	//s.backgroundMusic.Stop()
}

func NewSignIn(
	router *ui.Router,
	assets assets.Assets,
) *SignIn {
	return &SignIn{
		backgroundMusic: assets.GetAudio("screen/signin_bg.mp3"),
		SharedComponent: ui.SharedComponent{
			Name:   "SignIn",
			StatefulStyle: ui.NewStatefulStyleWithLayout(ui.BoxLayoutType),
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
						}, &signInButtonStyles),
					},
					&boxStyles,
				),
			},
		},
	}
}
