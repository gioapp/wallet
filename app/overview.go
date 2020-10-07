package gwallet

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/btn"
	"github.com/gioapp/wallet/pkg/dap/box"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/nav"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	balances               Balances
	latestTransactions     []Tx
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

		latestTransactions = append(latestTransactions, t)
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
		gtx.Constraints.Min.X = 80
		gtx.Constraints.Min.X = 80
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

func (g *GioWallet) overviewBody() func(gtx C) D {
	return box.BoxBase(g.ui.Theme.Colors["PanelBg"], func(gtx C) D {
		return lyt.Format(gtx, g.ui.R.Mod["TwoEqual"].(string),
			balancesView(g.ui.Theme, g.Status),
			recentTxView(g.ui.Theme, g.ui.N))
	})
}

func balancesView(th *theme.Theme, s Status) func(gtx C) D {
	return panelView(th, "Balances", func(gtx C) D {
		return lyt.Format(gtx, "vflex(r(_),r(_),r(_),r(_),r(_))",
			overviewRow(th, "Available", row(th, s.bal.available)),
			overviewRow(th, "Pending", row(th, s.bal.pending)),
			overviewRow(th, "Immature", row(th, s.bal.immature)),
			helper.DuoUIline(false, 1, 0, 1, th.Colors["Border"]),
			overviewRow(th, "Total", row(th, s.bal.total)),
		)
	})
}

func recentTxView(th *theme.Theme, n *nav.Navigation) func(gtx C) D {
	return panelView(th, "Recent transactions", func(gtx C) D {
		return recentTransactionsList.Layout(gtx, len(latestTransactions), func(gtx C, i int) D {
			tx := latestTransactions[i]
			btn := btn.IconTextBtn(th, tx.Btn, "hflex(r(_),f(1,_)", false, func(gtx C) D {
				return lyt.Format(gtx, "vflex(r(_),r(_))",
					func(gtx C) D {
						return lyt.Format(gtx, "hflex(r(_),r(_))",
							func(gtx C) D {
								title := theme.H1(th, "title")
								title.Alignment = text.Start
								return title.Layout(gtx)
							},
							func(gtx C) D {
								title := theme.H1(th, "title")
								title.Alignment = text.Start
								return title.Layout(gtx)
							},
						)
					},
					func(gtx C) D {
						title := theme.H1(th, "title")
						title.Alignment = text.Start
						return title.Layout(gtx)
					},
				)
			})
			btn.TextSize = unit.Dp(12)
			btn.Icon = th.Icons["MapsLayers"]
			btn.CornerRadius = unit.Dp(0)
			btn.Background = th.Colors["NavBg"]
			for tx.Btn.Clicked() {
				//n.CurrentPage = tx.Id
			}
			return btn.Layout(gtx)
		})
	})
}

func panelView(th *theme.Theme, title string, content func(gtx C) D) func(gtx C) D {
	return box.BoxPanel(th, func(gtx C) D {
		return lyt.Format(gtx, "hmax(vflex(r(_),r(_)))",
			func(gtx C) D {
				title := theme.H3(th, title)
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			content,
		)
	})
}
