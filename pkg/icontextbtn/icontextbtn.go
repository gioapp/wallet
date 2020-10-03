package icontextbtn

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
	"github.com/gioapp/wallet/pkg/theme"
	"image"
	"image/color"
)

type (
	D = layout.Dimensions
	C = layout.Context
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
	lay          string
	noText       bool
}

func IconTextBtn(t *theme.Theme, b *widget.Clickable, i *widget.Icon, is unit.Value, c, w, l string, noText bool) IconTextButton {
	return IconTextButton{
		Theme:     t,
		Button:    b,
		Icon:      i,
		IconColor: c,
		IconSize:  is,
		Text:      w,
		TextSize:  unit.Dp(16),
		noText:    noText,
		lay:       l,
	}
}

func (b IconTextButton) Layout(gtx C) D {
	return material.Clickable(gtx, b.Button, func(gtx C) D {
		var leftMargin float32 = 12
		if b.noText {
			leftMargin = 8
		}
		return layout.Inset{
			Top:    unit.Dp(8),
			Right:  unit.Dp(8),
			Bottom: unit.Dp(8),
			Left:   unit.Dp(leftMargin),
		}.Layout(gtx, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X

			layIcon := func(gtx C) D {
				return layout.Inset{}.Layout(gtx, func(gtx C) D {
					var d D
					if b.Icon != nil {
						//size := gtx.Px(b.IconSize) - 2*gtx.Px(unit.Dp(16))
						size := gtx.Px(b.IconSize)
						b.Icon.Color = helper.HexARGB(b.Theme.Colors["NavItem"])
						b.Icon.Layout(gtx, unit.Px(float32(size)))
						d = D{
							Size: image.Point{X: size, Y: size},
						}
					}
					return d
				})
			}
			layLabel := func(gtx C) D {
				return layout.Inset{
					Top:    unit.Dp(5),
					Right:  unit.Dp(8),
					Bottom: unit.Dp(0),
					Left:   unit.Dp(8),
				}.Layout(gtx, func(gtx C) D {
					l := theme.Body(b.Theme, b.Text)
					l.TextSize = b.TextSize
					l.Alignment = text.Middle
					l.Color = helper.HexARGB(b.Theme.Colors["NavItem"])
					return l.Layout(gtx)
				})
			}

			if !b.noText {
				layLabel = func(gtx C) D { return D{} }
			}

			return lyt.Format(gtx, b.lay,
				layIcon,
				layLabel)
		})
	})
}
