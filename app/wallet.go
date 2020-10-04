package gwallet

import (
	"gioui.org/layout"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *GioWallet) AppMain() func(gtx C) D {
	return func(gtx C) D {
		g.Resposnsivity(gtx)
		g.UI.nav.mode = g.UI.Res.Mode
		g.UI.nav.navLayout = g.UI.Res.Mod["NavLayout"].(string)
		g.UI.nav.itemLayout = g.UI.Res.Mod["NavIconAndLabel"].(string)
		g.UI.nav.axis = g.UI.Res.Mod["NavItemsAxis"].(layout.Axis)
		g.UI.nav.size = g.UI.Res.Mod["NavSize"].(int)
		g.UI.nav.noContent = g.UI.nav.wide
		g.UI.nav.logo = g.logo()
		return lyt.Format(gtx, g.UI.Res.Mod["MainLayout"].(string),
			func(gtx C) D {
				g.UI.nav.size = 128
				g.UI.nav.noContent = true
				if g.UI.nav.wide {
					g.UI.nav.size = 64
					g.UI.nav.noContent = false
				}
				return g.UI.nav.Nav(g.UI.Theme, gtx)
			},
			func(gtx C) D {
				return lyt.Format(gtx, g.UI.Res.Mod["ContentLayout"].(string),
					//g.header(),
					//g.page(g.UI.pages[currentPage]),
					//g.footer(),
					noReturn, noReturn,
				)
			})
	}
}

func (g *GioWallet) BeforeMain() {
	//g.Resposnsivity(gtx)
	//g.UI.nav.mode = g.UI.Res.Mode
	//g.UI.nav.navLayout = g.UI.Res.Mod["NavLayout"].(string)
	//g.UI.nav.itemLayout = g.UI.Res.Mod["NavIconAndLabel"].(string)
	//g.UI.nav.axis = g.UI.Res.Mod["NavItemsAxis"].(layout.Axis)
	//g.UI.nav.size = g.UI.Res.Mod["NavSize"].(int)
	//g.UI.nav.noContent = g.UI.nav.wide
	//g.UI.nav.logo = g.logo()
}

func (g *GioWallet) AfterMain() {

}

func (g *GioWallet) Tik() func() {
	return func() {
		g.GetBalances()
		g.GetLatestTransactions()
	}
}
