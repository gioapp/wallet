package dap

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/container"
	"github.com/gioapp/wallet/pkg/dap/mod"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

func Page(th *theme.Theme, page mod.Page) func(gtx C) D {
	return container.C().
		OutsideColor(th.Colors["PanelBg"]).
		BorderColor(th.Colors["Border"]).
		InsideColor(th.Colors["PanelBg"]).
		Margin(0).
		Border(0).
		Padding(0).
		Layout(func(gtx C) D {
			return lyt.Format(gtx, "vflex(start,r(inset(0dp0dp0dp0dp,_)),f(1,inset(0dp0dp0dp0dp,_)))", page.Header, page.Body)
		})
}

func pageButton(d *mod.Dap, b *widget.Clickable, f func(), icon, page string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(d.UI.Theme.T, b, d.UI.Theme.Icons[icon])
		btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
		btn.Size = unit.Dp(21)
		btn.Background = helper.HexARGB(d.UI.Theme.Colors["Secondary"])
		for b.Clicked() {
			f()
			d.UI.N.CurrentPage = page
		}
		return btn.Layout(gtx)
	}
}
