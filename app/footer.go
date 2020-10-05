package gwallet

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/dap/mod"
	"github.com/gioapp/wallet/pkg/lyt"
)

var (
	syncBtn           = new(widget.Clickable)
	unitBtn           = new(widget.Clickable)
	connectionsBtn    = new(widget.Clickable)
	footerSearchInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func (g *GioWallet) footer(ui *mod.UserInterface) func(gtx C) D {
	//return ContainerLayout(d.UI.Theme.Colors["Info"], d.UI.Theme.Colors["Info"], d.UI.Theme.Colors["Info"], 0, 0, 0, func(gtx C) D {
	return boxBase(ui.Theme, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "hflex(middle,r(inset(0dp0dp0dp6dp,_)),r(inset(20dp30dp20dp3dp,_)))",
			g.footerSearch(ui),
			g.footerMenu(ui),
		)
	})
}

func (g *GioWallet) footerMenu(ui *mod.UserInterface) func(gtx C) D {
	return func(gtx C) D {
		for syncBtn.Clicked() {
			//ui.N.CurrentPage = "Welcome"
		}
		return lyt.Format(gtx, "hflex(middle,r(_),r(_),r(_))",
			//g.pageButton(d, syncBtn, func() {}, "StrokeCase", ""),
			noReturn,
			helper.DuoUIline(true, 0, 2, 2, ui.Theme.Colors["DarkGrayI"]),
			//g.pageButton(tourBtn, func() {}, "GlyphPencil", "Welcome"),
			func(gtx C) D {
				btn := material.IconButton(ui.Theme.T, connectionsBtn, ui.Theme.Icons["networkIcon"])
				btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
				btn.Size = unit.Dp(21)
				btn.Background = helper.HexARGB(ui.Theme.Colors["Secondary"])
				for connectionsBtn.Clicked() {
					//f()
					//ui.N.CurrentPage = "Welcome"
				}
				return btn.Layout(gtx)
			},
		)
	}

}

func (g *GioWallet) footerSearch(ui *mod.UserInterface) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflex(middle,r(inset(8dp8dp8dp8dp,_)),r(_))",
			//ContainerLayout(d.UI.Theme.Colors["White"], d.UI.Theme.Colors["Dark"], d.UI.Theme.Colors["White"], 1, 1, 1, func(gtx C) D {
			boxBase(ui.Theme, func(gtx C) D {
				gtx.Constraints.Min.X = 430
				e := material.Editor(ui.Theme.T, footerSearchInput, "Hash")
				return e.Layout(gtx)
			}),
			func(gtx C) D {
				e := material.Button(ui.Theme.T, unitBtn, "Browse")
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
