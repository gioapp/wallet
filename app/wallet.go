package gwallet

import (
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (g *GioWallet) AppMain() {
	lyt.Format(g.UI.Context, "hflexb(start,r(_),f(1,_))",
		func(gtx C) D {
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			width := 128
			if g.UI.nav.wide {
				width = 64
			}
			return g.UI.nav.Nav(g.UI.Theme, gtx, width, g.UI.nav.wide, g.logo())
		},
		func(gtx C) D {
			return lyt.Format(gtx, g.UI.res.mod["MainLayout"].(string),
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
	}
}
