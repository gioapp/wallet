package gwallet

import (
	"fmt"
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	balances     Balances
	transactions []Tx
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
			time:    fmt.Sprint(tx.Time),
			address: fmt.Sprint(tx.Address),
			amount:  fmt.Sprint(tx.Amount),
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
		return lyt.Format(gtx, "hflexb(start,r(inset(0dp0dp0dp0dp,_)),f(1,_))",
			//gtx.Constraints.Min.X = gtx.Constraints.Max.X
			func(gtx C) D {
				//gtx.Constraints.Min.X = 100
				title := theme.Body(th, label)
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			content,
		)
	}
}

func (g *GioWallet) overviewBody() func(gtx C) D {
	return func(gtx C) D {
		helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["Success"]))

		return lyt.Format(gtx, g.UI.res.mod["ContentBodyLayout"].(string),
			g.balancesView(),
			g.recentTxView())
	}
}

func (g *GioWallet) balancesView() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["Border"], g.UI.Theme.Colors["PanelBg"], 4, 1, 8, func(gtx C) D {
		//gtx.Constraints.Min.X = gtx.Constraints.Max.X
		//helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflexb(r(_),r(_),r(_),r(_),r(_))",
			func(gtx C) D {
				//gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H1(g.UI.Theme, "Balances")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			overviewRow(g.UI.Theme, "Available", row(g.UI.Theme, g.Status.bal.available)),
			overviewRow(g.UI.Theme, "Pending", row(g.UI.Theme, g.Status.bal.pending)),
			overviewRow(g.UI.Theme, "Immature", row(g.UI.Theme, g.Status.bal.immature)),
			overviewRow(g.UI.Theme, "Total", row(g.UI.Theme, g.Status.bal.total)),
		)
	})
}

func (g *GioWallet) recentTxView() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["Border"], g.UI.Theme.Colors["PanelBg"], 4, 1, 8, func(gtx C) D {
		return lyt.Format(gtx, "vflexb(r(_),r(_),r(_),r(_),r(_))",
			func(gtx C) D {
				//gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H1(g.UI.Theme, "Recent transactions")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			overviewRow(g.UI.Theme, "Available", row(g.UI.Theme, g.Status.bal.available)),
			overviewRow(g.UI.Theme, "Pending", row(g.UI.Theme, g.Status.bal.pending)),
			overviewRow(g.UI.Theme, "Immature", row(g.UI.Theme, g.Status.bal.immature)),
			overviewRow(g.UI.Theme, "Total", row(g.UI.Theme, g.Status.bal.total)),
		)
	})
}
