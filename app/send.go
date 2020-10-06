package gwallet

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	addressBookBtn    = new(widget.Clickable)
	pasteClipboardBtn = new(widget.Clickable)
	clearBtn          = new(widget.Clickable)
	allAvailableBtn   = new(widget.Clickable)
	chooseFeeBtn      = new(widget.Clickable)
	clearAllBtn       = new(widget.Clickable)
	sendBtn           = new(widget.Clickable)
	addressInput      = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	labelInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	amountInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func (g *GioWallet) GetSend() {

}

func (g *GioWallet) sendHeader() func(gtx C) D {
	return boxBase(g.ui.Theme, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.ui.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflex(middle,r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.ui.Theme, "Send Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}

func (g *GioWallet) sendBody() func(gtx C) D {
	return func(gtx C) D {
		//gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflex(start,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				return lyt.Format(gtx, "hflex(middle,f(1,inset(8dp8dp8dp8dp,_)),r(_),r(_),r(_))",
					boxBase(g.ui.Theme, func(gtx C) D {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						e := material.Editor(g.ui.Theme.T, footerSearchInput, "Hash")
						return e.Layout(gtx)
					}),
					func(gtx C) D {
						btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
						btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
						btn.Size = unit.Dp(21)
						btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
						for connectionsBtn.Clicked() {
							//ui.N.CurrentPage = "Welcome"
						}
						return btn.Layout(gtx)
					},
					func(gtx C) D {
						btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
						btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
						btn.Size = unit.Dp(21)
						btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
						for connectionsBtn.Clicked() {
							//ui.N.CurrentPage = "Welcome"
						}
						return btn.Layout(gtx)
					},
					func(gtx C) D {
						btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
						btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
						btn.Size = unit.Dp(21)
						btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
						for connectionsBtn.Clicked() {
							//ui.N.CurrentPage = "Welcome"
						}
						return btn.Layout(gtx)
					},
				)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.ui.Theme, "Send Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.ui.Theme, "Send Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	}
}
