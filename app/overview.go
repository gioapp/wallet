package gwallet

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/btn"
	"github.com/gioapp/wallet/pkg/dap/mod"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	balances               Balances
	transactions           []Tx
	recentTransactionsList = &layout.List{
		Axis: layout.Vertical,
	}
)

func (g *GioWallet) GetOverview() {
	//checkError(err)
	g.Status = Status{
		bal: &balances,
	}
	return
}

func (g *GioWallet) GetBalances() {
	b, err := g.rpc.GetBalance("")
	checkError(err)
	fmt.Println("Ssss", b)
	balances = Balances{
		available: fmt.Sprint(b),
		pending:   fmt.Sprint(b),
		immature:  fmt.Sprint(b),
		total:     fmt.Sprint(b),
	}
}

func (g *GioWallet) GetLatestTransactions() {
	txs, err := g.rpc.ListTransactionsCount("", 5)
	checkError(err)
	fmt.Println("txstxstxs", txs)

	for _, tx := range txs {
		t := Tx{
			Id:      fmt.Sprint(tx.TxID),
			Time:    fmt.Sprint(tx.Time),
			Address: fmt.Sprint(tx.Address),
			Amount:  fmt.Sprint(tx.Amount),
			Btn:     new(widget.Clickable),
		}
		fmt.Println("Sstrtrtrtrtrtss", t)

		transactions = append(transactions, t)
	}

}

func (g *GioWallet) overviewHeader() func(gtx C) D {
	return func(gtx C) D {
		return D{}
	}
}

func row(th *theme.Theme, label string) func(gtx C) D {
	return func(gtx C) D {
		t := theme.Body(th, label)
		t.Alignment = text.Start
		return t.Layout(gtx)
	}
}

func overviewRow(th *theme.Theme, label string, content func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflex(start,r(inset(0dp0dp0dp0dp,_)),f(1,_))",
			func(gtx C) D {
				title := theme.Body(th, label)
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			content,
		)
	}
}

func (g *GioWallet) overviewBody(d *mod.Dap) func(gtx C) D {
	return boxBase(d.UI.Theme, func(gtx C) D {
		return lyt.Format(gtx, d.UI.R.Mod["TwoEqual"].(string),
			g.balancesView(d),
			g.recentTxView(d))
	})
}

func (g *GioWallet) balancesView(d *mod.Dap) func(gtx C) D {
	return g.panelView(d.UI.Theme, "Balances", func(gtx C) D {
		return lyt.Format(gtx, "vflex(r(_),r(_),r(_),r(_),r(_))",
			overviewRow(d.UI.Theme, "Available", row(d.UI.Theme, g.Status.bal.available)),
			overviewRow(d.UI.Theme, "Pending", row(d.UI.Theme, g.Status.bal.pending)),
			overviewRow(d.UI.Theme, "Immature", row(d.UI.Theme, g.Status.bal.immature)),
			helper.DuoUIline(false, 1, 0, 1, d.UI.Theme.Colors["Border"]),
			overviewRow(d.UI.Theme, "Total", row(d.UI.Theme, g.Status.bal.total)),
		)
	})
}

func (g *GioWallet) recentTxView(d *mod.Dap) func(gtx C) D {
	return g.panelView(d.UI.Theme, "Recent transactions", func(gtx C) D {
		return recentTransactionsList.Layout(gtx, len(transactions), func(gtx C, i int) D {
			tx := transactions[i]
			btn := btn.IconTextBtn(d.UI.Theme, tx.Btn, "hflex(r(_),f(1,_)", false, func(gtx C) D {
				return lyt.Format(gtx, "vflex(r(_),r(_))",
					func(gtx C) D {
						return lyt.Format(gtx, "hflex(r(_),r(_))",
							func(gtx C) D {
								title := theme.H1(d.UI.Theme, "title")
								title.Alignment = text.Start
								return title.Layout(gtx)
							},
							func(gtx C) D {
								title := theme.H1(d.UI.Theme, "title")
								title.Alignment = text.Start
								return title.Layout(gtx)
							},
						)
					},
					func(gtx C) D {
						title := theme.H1(d.UI.Theme, "title")
						title.Alignment = text.Start
						return title.Layout(gtx)
					},
				)
			})
			btn.TextSize = unit.Dp(12)
			btn.Icon = d.UI.Theme.Icons["MapsLayers"]
			btn.CornerRadius = unit.Dp(0)
			btn.Background = d.UI.Theme.Colors["NavBg"]
			for tx.Btn.Clicked() {
				d.UI.N.CurrentPage = tx.Id
			}
			return btn.Layout(gtx)
		})
	})
}

func (g *GioWallet) panelView(th *theme.Theme, title string, content func(gtx C) D) func(gtx C) D {
	return boxPanel(th, func(gtx C) D {
		return lyt.Format(gtx, "vflex(r(_),r(_))",
			func(gtx C) D {
				title := theme.H1(th, title)
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			content,
		)
	})
}
