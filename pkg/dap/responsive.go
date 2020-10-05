package dap

import (
	"gioui.org/layout"
	"github.com/gioapp/wallet/pkg/dap/mod"
)

func Resposnsivity(d *mod.Dap) {
	d.UI.R.Mode = "mobile"

	r := map[string]interface{}{
		"MainLayout":    "vflex(start,r(_),f(1,_))",
		"ContentLayout": "vflex(start,f(1,_),r(_))",
		"NavLayout":     "hflex(start,r(_),f(1,_))",
		"NavSize":       128,
		"NavItemsAxis":  layout.Horizontal,

		"NavIconAndLabel": "hflex(r(_),r(_))",
		"Logo":            d.UI.Theme.Icons["Logo"],

		"ContentBodyLayout": "vflex(start,r(inset(0dp,_)),f(1,inset(0dp,_))))",
	}

	switch w := d.UI.G.Constraints.Max.X; {
	//case w > 2560:
	//	g.UI.Res.Mode = "desktopXL"
	//case w > 1920:
	//	g.UI.Res.Mode = "desktopL"
	//case w > 1680:
	//	g.UI.Res.Mode = "desktopM"
	//case w > 1440:
	//	g.UI.Res.Mode = "desktopS"
	case w > 960:
		d.UI.R.Mode = "tablet"
		r["Screen"] = "max(inset(0dp90dp50dp10dp,_))"
		r["Main"] = "max(hflex(start,r(_),f(1,_)))"
		r["Nav"] = "vflex(start,r(_),f(1,_))"
		r["NavItemsAxis"] = layout.Vertical
		r["Content"] = "max(vflex(start,r(_),f(1,_),r(_)))"
		r["TwoEqual"] = "max(hflex(start,f(0.5,_),f(0.5,_)))"
	case w > 600:
		r["TwoEqual"] = "vflex(start,f(0.5,_),f(0.5,_))"
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
		d.UI.R.Mode = "mobile"
		r["Screen"] = "max(inset(80dp0dp80dp0dp,_))"
		r["Main"] = "max(inset(80dp0dp80dp0dp,_))"
		r["Content"] = "max(vflex(f(1,_),r(_))"
		r["Nav"] = "hflex(start,r(_),f(1,_))"
		r["NavSize"] = 128
		r["NavIconSize"] = 128
		r["NavIconAndLabel"] = "hflex(r(_),r(_))"
		r["Logo"] = d.UI.Theme.Icons["Logo"]
		r["NavItemsAxis"] = layout.Horizontal

		r["TwoEqual"] = "vflex(start,r(inset(0dp,_)),f(1,inset(0dp,_))))"
	}
	d.UI.R.Mod = r
	//fmt.Println("MODE", g.UI.Res.Mode)
	//fmt.Println("MAX", g.UI.Context.Constraints.Max.X)
	return
}
