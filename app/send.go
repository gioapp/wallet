package gwallet

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
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
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], 10, 10, 10, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.UI.Theme, "Send Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}

func (g *GioWallet) sendBody() func(gtx C) D {
	return func(gtx C) D {
		//gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflexb(start,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				return lyt.Format(gtx, "hflexb(middle,f(1,inset(8dp8dp8dp8dp,_)),r(_),r(_),r(_))",
					ContainerLayout(g.UI.Theme.Colors["Primary"], g.UI.Theme.Colors["Dark"], g.UI.Theme.Colors["White"], 1, 1, 1, func(gtx C) D {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						e := material.Editor(g.UI.Theme.T, footerSearchInput, "Hash")
						return e.Layout(gtx)
					}),
					func(gtx C) D {
						btn := material.IconButton(g.UI.Theme.T, connectionsBtn, g.UI.Theme.Icons["networkIcon"])
						btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
						btn.Size = unit.Dp(21)
						btn.Background = helper.HexARGB(g.UI.Theme.Colors["Secondary"])
						for connectionsBtn.Clicked() {
							currentPage = "Welcome"
						}
						return btn.Layout(gtx)
					},
					func(gtx C) D {
						btn := material.IconButton(g.UI.Theme.T, connectionsBtn, g.UI.Theme.Icons["networkIcon"])
						btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
						btn.Size = unit.Dp(21)
						btn.Background = helper.HexARGB(g.UI.Theme.Colors["Secondary"])
						for connectionsBtn.Clicked() {
							currentPage = "Welcome"
						}
						return btn.Layout(gtx)
					},
					func(gtx C) D {
						btn := material.IconButton(g.UI.Theme.T, connectionsBtn, g.UI.Theme.Icons["networkIcon"])
						btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
						btn.Size = unit.Dp(21)
						btn.Background = helper.HexARGB(g.UI.Theme.Colors["Secondary"])
						for connectionsBtn.Clicked() {
							currentPage = "Welcome"
						}
						return btn.Layout(gtx)
					},
				)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.UI.Theme, "Send Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.UI.Theme, "Send Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	}
}
