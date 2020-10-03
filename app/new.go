package gwallet

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/appdata"
	"github.com/gioapp/wallet/pkg/theme"
	rpcclient "github.com/p9c/pod/pkg/rpc/client"
)

func NewGioWallet(coinName string) *GioWallet {
	g := &GioWallet{
		dataDir: appdata.Dir(coinName, false),
		ctx:     context.Background(),
	}

	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:18444",
		User:         "bitnode",
		Pass:         "javazac",
		HTTPPostMode: true,  // Bitcoin core only supports HTTP POST mode
		TLS:          false, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {

	}

	g.rpc = client

	defer g.rpc.Shutdown()
	g.UI = walletUI{
		Theme: theme.NewTheme(),
		//mob:   make(chan bool),
	}
	currentPage = "Overview"
	//g.UI.Theme.Icons = icons.NewIPFSicons()
	g.UI.pages = g.getPages()
	g.UI.Theme.T.Color.Primary = helper.HexARGB(g.UI.Theme.Colors["Primary"])
	g.UI.Theme.T.Color.Text = helper.HexARGB(g.UI.Theme.Colors["Charcoal"])
	g.UI.Theme.T.Color.Hint = helper.HexARGB(g.UI.Theme.Colors["Silver"])
	g.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1024), unit.Dp(800)),
		app.Title("ParallelCoin - DUO - True Story"),
	)
	g.menuItems = g.getMenuItems()
	//for _, item :=range g.menuItems{
	//	for item.Btn.Clicked(){
	//		g.UI.currentPage = item.Title
	//		fmt.Println("ttt",item.Title)
	//
	//	}
	//}

	n := Navigation{
		Name:  "Navigacion",
		Bg:    g.UI.Theme.Colors["NavBg"],
		Items: g.menuItems,
	}
	g.UI.nav = n
	g.Status.bal = &Balances{}
	//g.Page = gipfsPage{
	//	Title:  "Status",
	//	Header: g.statusHeader(),
	//	Body:   g.statusBody(),
	//}
	//if g.sh.IsUp() {
	g.GetOverview()
	//g.GetFiles()
	//}

	//suffixes[0] = "B"
	//suffixes[1] = "KB"
	//suffixes[2] = "MB"
	//suffixes[3] = "GB"
	//suffixes[4] = "TB"

	return g
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func (g *GioWallet) getMenuItems() []Item {
	return []Item{
		Item{
			Title: "Overview",
			Icon:  g.UI.Theme.Icons["Overview"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Send",
			Icon:  g.UI.Theme.Icons["Send"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Receive",
			Icon:  g.UI.Theme.Icons["Receive"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Transactions",
			Icon:  g.UI.Theme.Icons["History"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Explore",
			Icon:  g.UI.Theme.Icons["Blocks"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Peers",
			Icon:  g.UI.Theme.Icons["Network"],
			Btn:   new(widget.Clickable),
		},
		Item{
			Title: "Settings",
			Icon:  g.UI.Theme.Icons["Settings"],
			Btn:   new(widget.Clickable),
		},
	}
}
