package gwallet

import (
	"gioui.org/text"
	"gioui.org/widget"
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

func (g *GioWallet) sendBody() []func(gtx C) D {
	return []func(gtx C) D{
		func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "vflexb(start,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
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
				},
				func(gtx C) D {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X
					title := theme.H6(g.UI.Theme, "Send Header")
					title.Alignment = text.Start
					return title.Layout(gtx)
				})

		},
	}
}
