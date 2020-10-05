package nav

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
	"image"
)

var (
	logoBtn = new(widget.Clickable)
)

type (
	D = layout.Dimensions
	C = layout.Context
	W = layout.Widget
)

type Logo struct {
	Title string
	Logo  string
}

func (n *Navigation) logoLayout(th *theme.Theme) func(gtx C) D {
	return func(gtx C) D {
		return material.Clickable(gtx, logoBtn, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			for logoBtn.Clicked() {
				switch n.wide {
				case false:
					n.wide = true
				default:
					n.wide = false
				}
			}
			inset := "inset(8dp18dp8dp8dp,_)"
			if n.wide {
				inset = "inset(8dp8dp8dp8dp,_)"
			}
			return lyt.Format(gtx, inset, func(gtx C) D {
				logo := th.Icons["Logo"]
				if n.wide {
					logo = th.Icons["Logo"]
				}
				var d D
				size := gtx.Px(unit.Dp(48))
				logo.Color = helper.HexARGB(th.Colors["PanelBg"])
				logo.Layout(gtx, unit.Px(float32(size)))
				d = D{
					Size: image.Point{X: size, Y: size},
				}
				return d
			})
		})
	}
}
