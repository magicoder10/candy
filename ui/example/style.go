package main

import (
	"candy/ui"
	"candy/ui/ptr"
)

// SL: Stylesheet Language
//
// signInButtonStyle {
//   Alignment: {
//	   Horizontal: center;
//   }
//   FontStyle: {
//     Weight: medium;
//	   LineHeight: 20px;
//	   Size: 20px;
//   }
//   Margin: {
//     Top: 469px;
//   }
//   Padding: {
//     Top: 12px;
//     Bottom: 12px;
//   }
//	 Width: 356px;
//	 Background: {
//     Color: rgba(22, 107, 107, 255);
//	 }
// }
var signInButtonStyle = ui.Style{
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
}

// SL: Stylesheet Language
//
// demoButtonStyle {
//  Background: {
//	  ImagePath: "test/image2.jpg";
//  }
// }
var demoButtonStyle = ui.Style{
	Background: &ui.Background{
		ImagePath: ptr.String("test/image2.jpg"),
	},
}

// SL: Stylesheet Language
//
// boxStyle {
//   Height: $height;
//	 Alignment: {
//     Horizontal: center;
//	 }
//   Background: {
//	   ImagePath: "test/signin.png";
//   }
// }
var boxStyle = ui.Style{
	Height: ptr.Int(Height),
	Alignment: &ui.Alignment{
		Horizontal: ui.AlignHorizontalCenter.Ptr(),
	},
	Background: &ui.Background{
		ImagePath: ptr.String("test/signin.png"),
	},
}

// SL: Stylesheet Language
//
// boxStyle {
//   FontStyle: {
//     Family: "Source Code Pro";
//     Weight: extraLight;
//     Size: 20px;
//     LineHeight: 24px;
//     Color: rgba(255, 255, 255, 255);
//   }
// }
var textStyle = ui.Style{FontStyle: &ui.FontStyle{
	Family:     ptr.String("Source Code Pro"),
	Weight:     ptr.String("ExtraLight"),
	Size:       ptr.Int(20),
	LineHeight: ptr.Int(24),
	Color: &ui.Color{
		Red:   255,
		Green: 255,
		Blue:  255,
		Alpha: 255,
	}}}
