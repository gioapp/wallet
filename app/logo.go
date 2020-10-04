package gwallet

import (
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/lyt"
	"image"
)

var (
	logoBtn = new(widget.Clickable)
)

func (g *GioWallet) logo() func(gtx C) D {
	return func(gtx C) D {
		return material.Clickable(gtx, logoBtn, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			for logoBtn.Clicked() {
				switch g.UI.nav.wide {
				case false:
					g.UI.nav.wide = true
				default:
					g.UI.nav.wide = false
				}
			}
			inset := "inset(8dp18dp8dp8dp,_)"
			if g.UI.nav.wide {
				inset = "inset(8dp8dp8dp8dp,_)"
			}
			return lyt.Format(gtx, inset, func(gtx C) D {
				logo := g.UI.Theme.Icons["Logo"]
				if g.UI.nav.wide {
					logo = g.UI.Theme.Icons["Logo"]
				}
				var d D
				size := gtx.Px(unit.Dp(48))
				logo.Color = helper.HexARGB(g.UI.Theme.Colors["PanelBg"])
				logo.Layout(gtx, unit.Px(float32(size)))
				d = D{
					Size: image.Point{X: size, Y: size},
				}
				return d
			})
		})
	}
}
