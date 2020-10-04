package gwallet

import (
	"gioui.org/layout"
)

func (g *GioWallet) Resposnsivity(gtx C) {
	g.UI.Res.Mode = "mobile"

	r := map[string]interface{}{
		"MainLayout":    "vflex(start,r(_),f(1,_))",
		"ContentLayout": "vflex(start,f(1,_),r(_))",
		"NavLayout":     "hflex(start,r(_),f(1,_))",
		"NavSize":       128,
		"NavItemsAxis":  layout.Horizontal,

		"NavIconAndLabel": "hflex(r(_),r(_))",
		"Logo":            g.UI.Theme.Icons["Logo"],

		"ContentBodyLayout": "vflex(start,r(inset(0dp,_)),f(1,inset(0dp,_))))",
	}

	switch w := gtx.Constraints.Max.X; {
	//case w > 2560:
	//	g.UI.Res.Mode = "desktopXL"
	//case w > 1920:
	//	g.UI.Res.Mode = "desktopL"
	//case w > 1680:
	//	g.UI.Res.Mode = "desktopM"
	//case w > 1440:
	//	g.UI.Res.Mode = "desktopS"
	case w > 960:
		g.UI.Res.Mode = "tablet"
		r["ScreenLayout"] = "max(inset(0dp0dp0dp0dp,_))"
		r["MainLayout"] = "hflex(start,r(_),f(1,_))"
		r["NavLayout"] = "vflex(start,r(_),f(1,_))"
		r["NavItemsAxis"] = layout.Vertical
		r["ContentBodyLayout"] = "hflex(start,f(0.5,_),f(0.5,_))"
	case w > 600:
		r["ContentBodyLayout"] = "vflex(start,f(0.5,_),f(0.5,_))"
	//case w > 1136:
	//	g.UI.Res.Mode = "tabletS"
	//case w > 1024:
	//	g.UI.Res.Mode = "tabletM"
	//case w > 960:
	//	g.UI.Res.Mode = "mobileXXXL"
	//case w > 800:
	//	g.UI.Res.Mode = "mobileXXL"
	//case w > 640:
	//	g.UI.Res.Mode = "mobileXL"
	//case w > 480:
	//	g.UI.Res.Mode = "mobileL"
	//case w > 320:
	//	g.UI.Res.Mode = "mobileM"
	//case w > 240:
	//	g.UI.Res.Mode = "mobileS"
	case w > 0:
		g.UI.Res.Mode = "mobile"
		r["ScreenLayout"] = "max(inset(80dp0dp80dp0dp,_))"
		r["MainLayout"] = "vflex(start,r(_),f(1,_))"
		r["ContentLayout"] = "vflex(start,f(1,_),r(_))"
		r["NavLayout"] = "hflex(start,r(_),f(1,_))"
		r["NavSize"] = 128
		r["NavIconSize"] = 128
		r["NavIconAndLabel"] = "hflex(r(_),r(_))"
		r["Logo"] = g.UI.Theme.Icons["Logo"]
		r["NavItemsAxis"] = layout.Horizontal
		r["ContentBodyLayout"] = "vflex(start,r(inset(0dp,_)),f(1,inset(0dp,_))))"
	}
	g.UI.Res.Mod = r
	//fmt.Println("MODE", g.UI.Res.Mode)
	//fmt.Println("MAX", g.UI.Context.Constraints.Max.X)
	return
}
