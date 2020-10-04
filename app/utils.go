package gwallet

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/lyt"
)

func (g *GioWallet) twoPanels(l, r int, left, right func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflex(start,f(0.60,inset(30dp10dp30dp10dp,_)),f(0.30,inset(30dp10dp30dp10dp,_)))",
			ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], l, l, l, left),
			ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], r, r, r, right))
	}
}

func stack(color string, size int, content func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return layout.Stack{Alignment: layout.W}.Layout(gtx,
			layout.Expanded(func(gtx C) D {
				return helper.Fill(gtx, helper.HexARGB(color))
			}),
			layout.Stacked(func(gtx C) D {
				return layout.Inset{
					Top:    unit.Dp(float32(size)),
					Right:  unit.Dp(float32(size)),
					Bottom: unit.Dp(float32(size)),
					Left:   unit.Dp(float32(size)),
				}.Layout(gtx, content)
			}),
		)
	}
}
