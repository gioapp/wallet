package icontextbtn

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/theme"
	"image"
	"image/color"
)

type IconTextButton struct {
	Theme        *theme.Theme
	Button       *widget.Clickable
	Background   color.RGBA
	Icon         *widget.Icon
	IconColor    string
	IconSize     unit.Value
	Text         string
	TextSize     unit.Value
	CornerRadius unit.Value
	axis         layout.Axis
	noText       bool
}

func IconTextBtn(t *theme.Theme, b *widget.Clickable, i *widget.Icon, is unit.Value, c, w string, noText bool) IconTextButton {
	return IconTextButton{
		Theme:     t,
		Button:    b,
		Icon:      i,
		IconColor: c,
		IconSize:  is,
		Text:      w,
		TextSize:  unit.Dp(16),
		noText:    noText,
	}
}

func (b IconTextButton) Layout(gtx layout.Context) layout.Dimensions {
	return material.Clickable(gtx, b.Button, func(gtx layout.Context) layout.Dimensions {
		var leftMargin float32 = 12
		if b.noText {
			leftMargin = 8
		}
		return layout.Inset{
			Top:    unit.Dp(16),
			Right:  unit.Dp(8),
			Bottom: unit.Dp(16),
			Left:   unit.Dp(leftMargin),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			iconAndLabel := layout.Flex{Alignment: layout.Middle, Axis: b.axis}
			layIcon := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					var d layout.Dimensions
					if b.Icon != nil {
						//size := gtx.Px(b.IconSize) - 2*gtx.Px(unit.Dp(16))
						size := gtx.Px(b.IconSize) - 2*gtx.Px(unit.Dp(16))
						b.Icon.Color = helper.HexARGB(b.IconColor)
						b.Icon.Layout(gtx, unit.Px(float32(size)))
						d = layout.Dimensions{
							Size: image.Point{X: size, Y: size},
						}
					}
					return d
				})
			})
			layLabel := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{
					Top:    unit.Dp(4),
					Right:  unit.Dp(8),
					Bottom: unit.Dp(4),
					Left:   unit.Dp(8),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					l := theme.Body(b.Theme, b.Text)
					l.TextSize = b.TextSize
					l.Alignment = text.Middle
					l.Color = helper.HexARGB(b.Theme.Colors["Black"])
					return l.Layout(gtx)
				})
			})
			var lays []layout.FlexChild
			lays = append(lays, layIcon)
			if !b.noText {
				lays = append(lays, layLabel)

			}
			//for  b.Button.Clicked(){
			//	g.UI.currentPage = b.Title
			//	fmt.Println("ttt",b.Title)
			//}
			return iconAndLabel.Layout(gtx, lays...)
		})
	})
}
