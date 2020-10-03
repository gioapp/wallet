package gwallet

import (
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

func (g *GioWallet) GetReceive() {

}

func (g *GioWallet) receiveHeader() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], 10, 10, 10, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.UI.Theme, "Receive Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}

func (g *GioWallet) receiveBody() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		title := theme.H5(g.UI.Theme, "Receive Body")
		title.Alignment = text.Start
		return title.Layout(gtx)
	}
}
