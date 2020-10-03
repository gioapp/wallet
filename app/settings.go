package gwallet

import (
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

func (g *GioWallet) GetSettings() {
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

func (g *GioWallet) settingsHeader() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], 10, 10, 10, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflexb(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H5(g.UI.Theme, "LANGUAGE")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H4(g.UI.Theme, "BUTTONS")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.UI.Theme, "ANALYTICS")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.Body(g.UI.Theme, "Help improve this app by sending anonymous usage data")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.Body(g.UI.Theme, "No CIDs, filenames, or other personal information are collected. We want metrics to show us which features are useful to help us prioritise what to work on next, and system configuration information to guide our testing.")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.UI.Theme, "Configure what is collected")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}

func (g *GioWallet) settingsBody() []func(gtx C) D {
	return []func(gtx C) D{
		func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			title := theme.H5(g.UI.Theme, "IPFS CONFIG")
			title.Alignment = text.Start
			return title.Layout(gtx)
		},
		func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			title := theme.Body(g.UI.Theme, "The IPFS config file is a JSON document. It is read once when the IPFS daemon is started. Save your changes, then restart the IPFS daemon to apply them. Check the documentation for further information.")
			title.Alignment = text.Start
			return title.Layout(gtx)
		},
		func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			title := theme.H6(g.UI.Theme, "CONFIG")
			title.Alignment = text.Start
			return title.Layout(gtx)
		},
	}
}
