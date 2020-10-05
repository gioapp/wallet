package gwallet

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	browseBtn = new(widget.Clickable)
)

func (g *GioWallet) GetPeers() {
	//f, err := g.sh.ID()
	//checkError(err)
	//fv, ft, err := g.sh.Version()
	//checkError(err)
	//g.Status = Status{
	//	Title: "string",
	//	//HostingSize: "uint",
	//	PeerId:  f.ID,
	//	Version: fv + " " + ft,
	//	//Gateway:   ,
	//	//Api:       "string",
	//	Addresses: f.Addresses,
	//	Pub:       f.PublicKey,
	//}
}

func (g *GioWallet) peersHeader(th *theme.Theme) func(gtx C) D {
	return boxBase(th, func(gtx C) D {
		e := material.Button(th.T, browseBtn, "Submit")
		e.Inset = layout.Inset{
			Top:    unit.Dp(4),
			Right:  unit.Dp(4),
			Bottom: unit.Dp(4),
			Left:   unit.Dp(4),
		}
		e.CornerRadius = unit.Dp(4)
		return e.Layout(gtx)
	})
}

func (g *GioWallet) peersBody() func(gtx C) D {
	return func(gtx C) D {
		//return lyt.Format(gtx, "vflex(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp30dp0dp,_)),r(inset(5dp0dp5dp0dp,_),r(inset(5dp0dp5dp0dp,_)))",
		//	statusRow(g.UI.Theme, "RateIn: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.RateIn))),
		//	statusRow(g.UI.Theme, "RateOut: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.RateOut))),
		//	statusRow(g.UI.Theme, "TotalIn: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.TotalIn))),
		//	statusRow(g.UI.Theme, "TotalOut: ", row(g.UI.Theme, fmt.Sprint(g.Status.Live.TotalOut))),
		//)
		return D{}
	}
}
