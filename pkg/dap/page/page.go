package page

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/container"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

type (
	D = layout.Dimensions
	C = layout.Context
	W = layout.Widget
)

type Page struct {
	Title  string
	Header layout.Widget
	Body   layout.Widget
	Footer layout.Widget
}

func (page Page) P(th *theme.Theme) func(gtx C) D {
	return container.C().
		OutsideColor(th.Colors["PanelBg"]).
		BorderColor(th.Colors["Border"]).
		InsideColor(th.Colors["PanelBg"]).
		Margin(0).
		Border(0).
		Padding(0).
		Layout(func(gtx C) D {
			fmt.Println("TestRRR11sdsadasdasd11")
			//return D{}
			return lyt.Format(gtx, "max(hflex(start,r(_),f(1,_)))", page.Header, page.Body)
		})
}

func pageButton(th *theme.Theme, b *widget.Clickable, f func(), icon, page string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(th.T, b, th.Icons[icon])
		btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
		btn.Size = unit.Dp(21)
		btn.Background = helper.HexARGB(th.Colors["Secondary"])
		for b.Clicked() {
			f()
			//d.UI.N.CurrentPage = page
		}
		return btn.Layout(gtx)
	}
}