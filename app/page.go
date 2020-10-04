package gwallet

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/lyt"
)

func (g *GioWallet) getPages() pages {
	return pages{
		"Welcome": walletPage{
			Title:  "Welcome",
			Header: g.welcomeHeader(),
			Body:   g.welcomeBody(),
		},
		"Overview": walletPage{
			Title:  "Overview",
			Header: g.overviewHeader(),
			Body:   g.overviewBody(),
		},
		"Send": walletPage{
			Title:  "Send",
			Header: g.sendHeader(),
			Body:   g.sendBody(),
		},
		"Receive": walletPage{
			Title:  "Receive",
			Header: g.receiveHeader(),
			Body:   g.receiveBody(),
		},
		"Transactions": walletPage{
			Title:  "Transactions",
			Header: g.transactionsHeader(),
			Body:   g.transactionsBody(),
		},
		"Explore": walletPage{
			Title:  "Explore",
			Header: g.exploreHeader(),
			Body:   g.exploreBody(),
		},
		"Peers": walletPage{
			Title:  "Peers",
			Header: g.peersHeader(),
			Body:   g.peersBody(),
		},
		"Settings": walletPage{
			Title:  "Settings",
			Header: g.settingsHeader(),
			Body:   g.settingsBody(),
		},
	}
}

func (g *GioWallet) page(page walletPage) func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["White"], g.UI.Theme.Colors["White"], g.UI.Theme.Colors["White"], 0, 0, 0, func(gtx C) D {
		return lyt.Format(gtx, "vflex(start,r(inset(0dp0dp0dp0dp,_)),f(1,inset(0dp0dp0dp0dp,_)))", page.Header, page.Body)
	})
}

func ContainerLayout(outsideColor, borderColor, insideColor string, margin, border, padding int, itemContent func(gtx C) D) func(gtx C) D {
	//vmin := gtx.Constraints.Min.Y
	//if d.FullWidth {
	//	hmin = gtx.Constraints.Max.Y
	//}
	return stack(outsideColor, margin, stack(borderColor, border, stack(insideColor, padding, stack(insideColor, padding, itemContent))))
}

func (g *GioWallet) pageButton(b *widget.Clickable, f func(), icon, page string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(g.UI.Theme.T, b, g.UI.Theme.Icons[icon])
		btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
		btn.Size = unit.Dp(21)
		btn.Background = helper.HexARGB(g.UI.Theme.Colors["Secondary"])
		for b.Clicked() {
			f()
			currentPage = page
		}
		return btn.Layout(gtx)
	}
}
