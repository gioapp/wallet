package gwallet

import (
	"fmt"
	"gioui.org/widget"
	"github.com/gioapp/wallet/pkg/dap/mod"
	"github.com/gioapp/wallet/pkg/dap/page"
	"github.com/gioapp/wallet/pkg/nav"
	rpcclient "github.com/p9c/pod/pkg/rpc/client"
)

func NewGioWallet(d *mod.Dap) mod.Sap {
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
	d.UI.N.MenuItems = append(g.getMenuItems(&d.UI))

	return mod.Sap{
		Title: "ParallelCoin",
		App:   g,
	}
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
				Footer: noReturn,
			},
		},
		nav.Item{
			Title: "Send",
			Icon:  ui.Theme.Icons["Send"],
			Btn:   new(widget.Clickable),
			Page: page.Page{
				Title:  "Send",
				Header: g.sendHeader(ui.Theme),
				Body:   g.sendBody(ui),
				Footer: noReturn,
			},
		},
		//nav.Item{
		//	Title: "Receive",
		//	Icon:  ui.Theme.Icons["Receive"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Receive",
		//		Header: g.receiveHeader(th),
		//		Body:   g.receiveBody(th),
		//	},
		//},
		//nav.Item{
		//	Title: "Transactions",
		//	Icon:  ui.Theme.Icons["History"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Transactions",
		//		Header: g.transactionsHeader(th),
		//		Body:   g.transactionsBody(th),
		//	},
		//},
		//nav.Item{
		//	Title: "Explore",
		//	Icon:  ui.Theme.Icons["Blocks"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Explore",
		//		Header: g.exploreHeader(),
		//		Body:   g.exploreBody(),
		//	},
		//},
		//nav.Item{
		//	Title: "Peers",
		//	Icon:  ui.Theme.Icons["Network"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title:  "Peers",
		//		Header: g.peersHeader(th),
		//		Body:   g.peersBody(),
		//	},
		//},
		//nav.Item{
		//	Title: "Settings",
		//	Icon:  ui.Theme.Icons["Settings"],
		//	Btn:   new(widget.Clickable),
		//	Page: page.Page{
		//		Title: "Settings",
		//		//			Header: g.settingsHeader(th),
		//		//			Body:   g.settingsBody(th),
		//	},
		//},
	}
}
