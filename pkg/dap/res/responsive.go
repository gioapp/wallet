package res

import (
	"gioui.org/layout"
)

type Responsivity struct {
	Mode string
	Mod  map[string]interface{}
	S    Screen
}
type Screen struct {
	X, Y int
}

func (r *Responsivity) Resposnsivity() {
	switch w := r.S.X; {
	//case w > 2560:
	//	g.UI.Res.Mode = "desktopXL"
	//case w > 1920:
	//	g.UI.Res.Mode = "desktopL"
	//case w > 1680:
	//	g.UI.Res.Mode = "desktopM"
	//case w > 1440:
	//	g.UI.Res.Mode = "desktopS"
	case w > 960:
		r.Mode = "tablet"
		r.Mod["Screen"] = "max(inset(0dp90dp50dp10dp,_))"
		r.Mod["Main"] = "max(hflex(start,r(_),f(1,_)))"
		r.Mod["Nav"] = "vflex(start,r(_),f(1,_))"
		r.Mod["NavItemsAxis"] = layout.Vertical
		r.Mod["Content"] = "max(vflex(start,r(_),f(1,_),r(_)))"
		r.Mod["TwoEqual"] = "max(hflex(start,f(0.5,_),f(0.5,_)))"
	case w > 600:
		r.Mod["TwoEqual"] = "vflex(start,f(0.5,_),f(0.5,_))"
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

	}
	return
}

func Resposnsivity(r *Responsivity) {

	return
}
