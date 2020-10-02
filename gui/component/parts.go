package component

import (
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"

	"github.com/gioapp/wallet/gui/rcd"
)

func SetPage(rc *rcd.RcVar, p *page.DuoUIpage) {
	//page.Command()
	rc.CurrentPage = p
}

func CurrentCurrentPageColor(showPage, p, color, currentPageColor string) (c string) {
	if showPage == p {
		c = currentPageColor
	} else {
		c = color
	}
	return
}

func fill(gtx *layout.Context, col color.RGBA) {
	//cs := gtx.Constraints
	//d := image.Point{X: cs.Width.Min, Y: cs.Height.Min}
	dr := f32.Rectangle{
		//Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	//gtx.Dimensions = layout.Dimensions{Size: d}
}
func Editor(gtx *layout.Context, th *theme.DuoUItheme, editorController *widget.Editor, label string, handler func(widget.EditorEvent)) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		container.DuoUIcontainer(th, 8, "ffffffff").Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
			//width := gtx.Constraints.Width.Max
			//e := th.DuoUIeditor(label, "ff000000", "ffffffff", width)
			//e.Font.Typeface = th.Fonts["Mono"]
			//e.TextSize = unit.Dp(12)
			//layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//	e.Layout(gtx, editorController)
			//})
			//for _, e := range editorController.Events(gtx) {
			//	switch e.(type) {
			//	case gel.ChangeEvent:
			//		handler(e)
			//	}
			//}
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}

func StringsArrayEditor(gtx *layout.Context, th *theme.DuoUItheme, editorController *widget.Editor, label string, handler func(widget.EditorEvent)) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		container.DuoUIcontainer(th, 8, "ffffffff").Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
			//e := th.DuoUIeditor(label, "ff000000", "ffffffff", 16)
			//e.Font.Typeface = th.Fonts["Mono"]
			//// e.Font.Style = text.Italic
			//layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//	e.Layout(gtx, editorController)
			//})
			//for _, e := range editorController.Events(gtx) {
			//	switch e.(type) {
			//	case gel.ChangeEvent:
			//		handler(e)
			//	}
			//}
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}

//
//func ConsoleInput(gtx *layout.Context, th *theme.DuoUItheme, editorController *widget.Editor, label string, handler func(gel.SubmitEvent)) func(gtx layout.Context) layout.Dimensions {
//	return func(gtx layout.Context) layout.Dimensions {
//		layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
//			//e := th.DuoUIeditor(label, "Dark", "Light", 120)
//			//e.Font.Typeface = th.Fonts["Primary"]
//			//e.Color = gelook.HexARGB(th.Colors["Light"])
//			//e.Font.Style = text.Italic
//			//e.Layout(gtx, editorController)
//			//for _, e := range editorController.Events(gtx) {
//			//	if e, ok := e.(gel.SubmitEvent); ok {
//			//		handler(e)
//			//		editorController.SetText("")
//			//	}
//			//}
//			return layout.Dimensions{}
//		})
//		return layout.Dimensions{}
//	}
//}

func Button(gtx *layout.Context, th *theme.DuoUItheme, buttonController *widget.Clickable, font text.Typeface, textSize, pt, pr, pb, pl int, color, bgColor, label string, handler func(gtx layout.Context) layout.Dimensions) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		container.DuoUIcontainer(th, 0, th.Colors[""])
		//button := th.DuoUIbutton(font, label, color, bgColor, "", "", "", "", textSize, 0, 128, 48, pt, pr, pb, pl)
		for buttonController.Clicked() {
			//handler()
		}
		//button.Layout(gtx, buttonController)
		return layout.Dimensions{}
	}

}

func MonoButton(gtx *layout.Context, th *theme.DuoUItheme, buttonController *widget.Clickable, textSize int, color, bgColor, font, label string, handler func(gtx layout.Context) layout.Dimensions) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//var button gelook.Button
			button := material.Button(th.T, buttonController, label)
			if font != "" {
				button.Font.Typeface = th.Fonts[font]
			}
			if color != "" {
				button.Color = helper.HexARGB(th.Colors[color])
			}
			if textSize != 0 {
				button.TextSize = unit.Dp(float32(textSize))
			}
			if bgColor != "" {
				button.Background = helper.HexARGB(th.Colors[bgColor])
			}
			for buttonController.Clicked() {
				//handler()
			}
			return button.Layout(gtx)
		})
		return layout.Dimensions{}
	}
}

func Label(gtx layout.Context, th *theme.DuoUItheme, font text.Typeface, size float32, color, label string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//l := th.DuoUIlabel(unit.Dp(size), label)
		//l.Font.Typeface = font
		//l.Color = color
		//l.Layout(gtx)
		return layout.Dimensions{}
	}
}
