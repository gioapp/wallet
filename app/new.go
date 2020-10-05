package gwallet

import (
	"fmt"
	"gioui.org/widget"
	"github.com/gioapp/wallet/pkg/dap/mod"
	"github.com/gioapp/wallet/pkg/dap/page"
	"github.com/gioapp/wallet/pkg/nav"
	rpcclient "github.com/p9c/pod/pkg/rpc/client"
)

func NewGioWallet() *mod.Sap {
	s := &mod.Sap{
		Title: "ParallelCoin Wallet",
	}
	g := &GioWallet{}

	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:18444",
		User:         "bitnode",
		Pass:         "44444",
		HTTPPostMode: true,  // Bitcoin core only supports HTTP POST mode
		TLS:          false, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {

	}

	g.rpc = client

	defer g.rpc.Shutdown()

	//g.UI.pages = g.getPages()
	//g.menuItems = g.getMenuItems()

	g.Status.bal = &Balances{}
	g.GetOverview()
	s.App = g
	//s.UI.N.MenuItems = append(g.getMenuItems(s.UI))
	return s
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func (g *GioWallet) getMenuItems(ui *mod.UserInterface) []nav.Item {
	return []nav.Item{
		nav.Item{
			Title: "Overview",
			Icon:  ui.Theme.Icons["Overview"],
			Btn:   new(widget.Clickable),
			Page: page.Page{
				Title:  "Overview",
				Header: g.overviewHeader(),
				Body:   g.overviewBody(ui),
			},
		},
		//nav.Item{
		//	Title: "Send",
		//	Icon:  th.Icons["Send"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Send",
		//		Header: g.sendHeader(th),
		//		//Body:   g.sendBody(d.UI),
		//	},
		//},
		//nav.Item{
		//	Title: "Receive",
		//	Icon:  th.Icons["Receive"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Receive",
		//		Header: g.receiveHeader(th),
		//		Body:   g.receiveBody(th),
		//	},
		//},
		//nav.Item{
		//	Title: "Transactions",
		//	Icon:  th.Icons["History"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Transactions",
		//		Header: g.transactionsHeader(th),
		//		Body:   g.transactionsBody(th),
		//	},
		//},
		//nav.Item{
		//	Title: "Explore",
		//	Icon:  th.Icons["Blocks"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Explore",
		//		Header: g.exploreHeader(),
		//		Body:   g.exploreBody(),
		//	},
		//},
		//nav.Item{
		//	Title: "Peers",
		//	Icon:  th.Icons["Network"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Peers",
		//		Header: g.peersHeader(th),
		//		Body:   g.peersBody(),
		//	},
		//},
		//nav.Item{
		//	Title: "Settings",
		//	Icon:  th.Icons["Settings"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title: "Settings",
		//		//			Header: g.settingsHeader(th),
		//		//			Body:   g.settingsBody(th),
		//	},
		//},
	}
}
