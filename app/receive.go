package gwallet

import (
	"gioui.org/text"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

func (g *GioWallet) GetReceive() {

}

func (g *GioWallet) receiveHeader() func(gtx C) D {
	return boxBase(g.ui.Theme.Colors["PanelBg"], func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflex(middle,r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.ui.Theme, "Receive Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}

func (g *GioWallet) receiveBody() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		title := theme.H5(g.ui.Theme, "Receive Body")
		title.Alignment = text.Start
		return title.Layout(gtx)
	}
}
