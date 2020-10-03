package gwallet

import (
	"gioui.org/layout"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *GioWallet) AppMain(gtx C) {
	g.Resposnsivity(gtx)

	g.UI.nav.mode = g.UI.res.Mode
	g.UI.nav.navLayout = g.UI.res.mod["NavLayout"].(string)
	g.UI.nav.itemLayout = g.UI.res.mod["NavIconAndLabel"].(string)
	g.UI.nav.axis = g.UI.res.mod["NavItemsAxis"].(layout.Axis)
	g.UI.nav.size = g.UI.res.mod["NavSize"].(int)
	g.UI.nav.noText = g.UI.nav.wide
	g.UI.nav.logo = g.logo()

	lyt.Format(gtx, g.UI.res.mod["MainLayout"].(string),
		func(gtx C) D {
			g.UI.nav.size = 128
			g.UI.nav.noText = true
			if g.UI.nav.wide {
				g.UI.nav.size = 64
				g.UI.nav.noText = false
			}
			return g.UI.nav.Nav(g.UI.Theme, gtx)
		},
		func(gtx C) D {
			return lyt.Format(gtx, g.UI.res.mod["ContentLayout"].(string),
				//g.header(),
				g.page(g.UI.pages[currentPage]),
				g.footer(),
			)
		})
}

func (g *GioWallet) BeforeMain() {

}

func (g *GioWallet) AfterMain() {

}

func (g *GioWallet) Tik() func() {
	return func() {
		g.GetBalances()
		g.GetLatestTransactions()
	}
}
