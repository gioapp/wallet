package gwallet

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
)

var (
	syncBtn           = new(widget.Clickable)
	unitBtn           = new(widget.Clickable)
	connectionsBtn    = new(widget.Clickable)
	footerSearchInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}

	currentPage string
)

func (g *GioWallet) footer() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["Info"], g.UI.Theme.Colors["Info"], g.UI.Theme.Colors["Info"], 0, 0, 0, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "hflexb(middle,r(inset(0dp0dp0dp6dp,_)),r(inset(20dp30dp20dp3dp,_)))",
			g.footerSearch(),
			g.footerMenu(),
		)
	})
}
func (g *GioWallet) footerMenu() func(gtx C) D {
	return func(gtx C) D {
		for syncBtn.Clicked() {
			currentPage = "Welcome"
		}
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
			g.pageButton(syncBtn, func() {}, "StrokeCase", ""),
			helper.DuoUIline(true, 0, 2, 2, g.UI.Theme.Colors["DarkGrayI"]),
			//g.pageButton(tourBtn, func() {}, "GlyphPencil", "Welcome"),
			func(gtx C) D {
				btn := material.IconButton(g.UI.Theme.T, connectionsBtn, g.UI.Theme.Icons["networkIcon"])
				btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
				btn.Size = unit.Dp(21)
				btn.Background = helper.HexARGB(g.UI.Theme.Colors["Secondary"])
				for connectionsBtn.Clicked() {
					//f()
					currentPage = "Welcome"
				}
				return btn.Layout(gtx)
			},
		)
	}

}

func (g *GioWallet) footerSearch() func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(inset(8dp8dp8dp8dp,_)),r(_))",
			ContainerLayout(g.UI.Theme.Colors["White"], g.UI.Theme.Colors["Dark"], g.UI.Theme.Colors["White"], 1, 1, 1, func(gtx C) D {
				gtx.Constraints.Min.X = 430
				e := material.Editor(g.UI.Theme.T, footerSearchInput, "Hash")
				return e.Layout(gtx)
			}),
			func(gtx C) D {
				e := material.Button(g.UI.Theme.T, unitBtn, "Browse")
				e.Inset = layout.Inset{
					Top:    unit.Dp(4),
					Right:  unit.Dp(4),
					Bottom: unit.Dp(4),
					Left:   unit.Dp(4),
				}
				e.CornerRadius = unit.Dp(4)
				return e.Layout(gtx)
			},
		)
	}
}
