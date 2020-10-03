package gwallet

import (
	"gioui.org/layout"
)

func (g *GioWallet) Resposnsivity(gtx C) {
	g.UI.res.Mode = "mobile"

	r := map[string]interface{}{
		"MainLayout":    "vflexb(start,r(_),f(1,_))",
		"ContentLayout": "vflexb(start,f(1,_),r(_))",
		"NavLayout":     "hflexs(start,r(_),f(1,_))",
		"NavSize":       128,
		"NavItemsAxis":  layout.Horizontal,

		"NavIconAndLabel": "hflexs(r(_),r(_))",
		"Logo":            g.UI.Theme.Icons["Logo"],

		"ContentBodyLayout": "vflexb(start,r(inset(0dp,_)),f(1,inset(0dp,_))))",
	}

	switch w := gtx.Constraints.Max.X; {
	//case w > 2560:
	//	g.UI.res.Mode = "desktopXL"
	//case w > 1920:
	//	g.UI.res.Mode = "desktopL"
	//case w > 1680:
	//	g.UI.res.Mode = "desktopM"
	//case w > 1440:
	//	g.UI.res.Mode = "desktopS"
	case w > 960:
		g.UI.res.Mode = "tablet"
		r["MainLayout"] = "hflexb(start,r(_),f(1,_))"
		r["NavLayout"] = "vflexs(start,r(_),f(1,_))"
		r["NavItemsAxis"] = layout.Vertical
		r["ContentBodyLayout"] = "hflexs(f(0.5,_),f(0.5,_))"
	case w > 600:
		r["ContentBodyLayout"] = "vflexb(start,r(inset(0dp,_)),f(1,inset(0dp,_))))"
	//case w > 1136:
	//	g.UI.res.Mode = "tabletS"
	//case w > 1024:
	//	g.UI.res.Mode = "tabletM"
	//case w > 960:
	//	g.UI.res.Mode = "mobileXXXL"
	//case w > 800:
	//	g.UI.res.Mode = "mobileXXL"
	//case w > 640:
	//	g.UI.res.Mode = "mobileXL"
	//case w > 480:
	//	g.UI.res.Mode = "mobileL"
	//case w > 320:
	//	g.UI.res.Mode = "mobileM"
	//case w > 240:
	//	g.UI.res.Mode = "mobileS"
	case w > 0:
		g.UI.res.Mode = "mobile"
		r["MainLayout"] = "vflexb(start,r(_),f(1,_))"
		r["ContentLayout"] = "vflexb(start,f(1,_),r(_))"
		r["NavLayout"] = "hflexs(start,r(_),f(1,_))"
		r["NavSize"] = 128
		r["NavIconSize"] = 128
		r["NavIconAndLabel"] = "hflexs(r(_),r(_))"
		r["Logo"] = g.UI.Theme.Icons["Logo"]
		r["NavItemsAxis"] = layout.Horizontal
		r["ContentBodyLayout"] = "vflexb(start,r(inset(0dp,_)),f(1,inset(0dp,_))))"
	}
	g.UI.res.mod = r
	//fmt.Println("MODE", g.UI.res.Mode)
	//fmt.Println("MAX", g.UI.Context.Constraints.Max.X)
	return
}
