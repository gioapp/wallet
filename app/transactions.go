package gwallet

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/btn"
	"github.com/gioapp/wallet/pkg/dap/box"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	transactions     []Tx
	transactionsList = &layout.List{}
)

func (g *GioWallet) GetTransactions() {

}

func (g *GioWallet) transactionsHeader() func(gtx C) D {
	return box.BoxBase(g.ui.Theme.Colors["PanelBg"], func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.ui.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflex(middle,r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.ui.Theme, "Transactions Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}

func (g *GioWallet) transactionsBody() func(gtx C) D {
	return box.BoxPanel(g.ui.Theme, func(gtx C) D {
		return transactionsList.Layout(gtx, len(transactions), func(gtx C, i int) D {
			tx := transactions[i]
			btn := btn.IconTextBtn(g.ui.Theme, tx.Btn, "hflex(r(_),f(1,_)", false, func(gtx C) D {
				return lyt.Format(gtx, "vflex(r(_),r(_))",
					func(gtx C) D {
						return lyt.Format(gtx, "hflex(r(_),r(_))",
							func(gtx C) D {
								title := theme.H1(g.ui.Theme, "title")
								title.Alignment = text.Start
								return title.Layout(gtx)
							},
							func(gtx C) D {
								title := theme.H1(g.ui.Theme, "title")
								title.Alignment = text.Start
								return title.Layout(gtx)
							},
						)
					},
					func(gtx C) D {
						title := theme.H1(g.ui.Theme, "title")
						title.Alignment = text.Start
						return title.Layout(gtx)
					},
				)
			})
			btn.TextSize = unit.Dp(12)
			btn.Icon = g.ui.Theme.Icons["MapsLayers"]
			btn.CornerRadius = unit.Dp(0)
			btn.Background = g.ui.Theme.Colors["NavBg"]
			for tx.Btn.Clicked() {
				//n.CurrentPage = tx.Id
			}
			return btn.Layout(gtx)
		})
	})
}
