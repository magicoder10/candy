package main

import (
	"candy/ui"
	"candy/ui/ptr"
)

var signInButtonStyles = ui.StatefulStyle{Styles: map[ui.State]*ui.Style{
	ui.NormalState: {
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
	},
	ui.HoverState: {
		Background: &ui.Background{
			Color: &ui.Color{
				Red:   255,
				Green: 255,
				Blue:  255,
				Alpha: 255,
			},
		},
		FontStyle: &ui.FontStyle{
			Color: &ui.Color{
				Red:   22,
				Green: 107,
				Blue:  107,
				Alpha: 255,
			},
		},
	},
}}

var demoButtonStyles = ui.StatefulStyle{Styles: map[ui.State]*ui.Style{
	ui.NormalState: {
		Background: &ui.Background{
			ImagePath: ptr.String("test/image2.jpg"),
		},
	},
},
}

var boxStyles = ui.StatefulStyle{Styles: map[ui.State]*ui.Style{
	ui.NormalState: {
		Height: ptr.Int(Height),
		Alignment: &ui.Alignment{
			Horizontal: ui.AlignHorizontalCenter.Ptr(),
		},
		Background: &ui.Background{
			ImagePath: ptr.String("test/signin.png"),
		},
	},
},
}

var textStyles = ui.StatefulStyle{Styles: map[ui.State]*ui.Style{
	ui.NormalState: {
		FontStyle: &ui.FontStyle{
			Family:     ptr.String("Source Code Pro"),
			Weight:     ptr.String("ExtraLight"),
			Italic:     ptr.Bool(false),
			Size:       ptr.Int(20),
			LineHeight: ptr.Int(24),
			Color: &ui.Color{
				Red:   255,
				Green: 255,
				Blue:  255,
				Alpha: 255,
			},
		},
	},
}}
