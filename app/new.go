package gwallet

import (
	"fmt"
	"gioui.org/widget"
	"github.com/gioapp/wallet/pkg/nav"
	"github.com/gioapp/wallet/pkg/theme"
	rpcclient "github.com/p9c/pod/pkg/rpc/client"
)

func NewGioWallet(coinName string) *GioWallet {
	g := &GioWallet{}

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

	//th.Icons = icons.NewIPFSicons()
	//g.UI.pages = g.getPages()
	//g.menuItems = g.getMenuItems()

	g.Status.bal = &Balances{}
	g.GetOverview()
	return g
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}

func (g *GioWallet) getMenuItems(th *theme.Theme) []nav.Item {
	return []nav.Item{
		nav.Item{
			Title: "Overview",
			Icon:  th.Icons["Overview"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Send",
			Icon:  th.Icons["Send"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Receive",
			Icon:  th.Icons["Receive"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Transactions",
			Icon:  th.Icons["History"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Explore",
			Icon:  th.Icons["Blocks"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Peers",
			Icon:  th.Icons["Network"],
			Btn:   new(widget.Clickable),
		},
		nav.Item{
			Title: "Settings",
			Icon:  th.Icons["Settings"],
			Btn:   new(widget.Clickable),
		},
	}
}
