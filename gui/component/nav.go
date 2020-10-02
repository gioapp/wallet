package component

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/gui/rcd"
)

var (
	navButtonOverview    = new(widget.Clickable)
	navButtonSend        = new(widget.Clickable)
	navButtonReceive     = new(widget.Clickable)
	navButtonAddressBook = new(widget.Clickable)
	navButtonHistory     = new(widget.Clickable)
	mainNav              = &layout.List{
		Axis: layout.Vertical,
	}
)

func MainNavigation(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme, allPages *model.DuoUIpages, nav *model.DuoUInav) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//navButtons := navButtons(rc, gtx, th, allPages, nav)
		//gtx.Constraints.Width.Max = nav.Width
		//th.DuoUIcontainer(0, th.Colors["Dark"]).Layout(gtx, layout.NW, func(gtx layout.Context)layout.Dimensions{
		//mainNav.Layout(gtx, len(navButtons), func(i int) {
		//	layout.UniformInset(unit.Dp(0)).Layout(gtx, navButtons[i])
		//})
		//})
		return layout.Dimensions{}
	}
}

func navButtons(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, allPages *model.DuoUIpages, nav *model.DuoUInav) []func(gtx layout.Context) layout.Dimensions {
	//return []func(){
	//navMenuButton(rc, gtx, th, allPages.Theme["OVERVIEW"], nav, "OVERVIEW", "overviewIcon", navButtonOverview),
	//th.DuoUIline(gtx, 0, 0, 1, th.Colors["LightGrayIII"]),
	//navMenuButton(rc, gtx, th, allPages.Theme["SEND"], nav, "SEND", "sendIcon", navButtonSend),
	// navMenuLine(gtx, th),
	// navMenuButton(rc, gtx, th, allPages.Theme["RECEIVE"], "RECEIVE", "receiveIcon", navButtonReceive),
	//th.DuoUIline(gtx, 0, 0, 1, th.Colors["LightGrayIII"]),
	//navMenuButton(rc, gtx, th, allPages.Theme["ADDRESSBOOK"], nav, "ADDRESSBOOK", "addressBookIcon", navButtonAddressBook),
	//th.DuoUIline(gtx, 0, 0, 1, th.Colors["LightGrayIII"]),
	//navMenuButton(rc, gtx, th, allPages.Theme["HISTORY"], nav, "HISTORY", "historyIcon", navButtonHistory),
	//}
	return []func(gtx layout.Context) layout.Dimensions{}
}

func navMenuButton(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, page *page.DuoUIpage, nav *model.DuoUInav, title, icon string, navButton *widget.Clickable) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//var menuItem gelook.DuoUIbutton
			//menuItem = th.DuoUIbutton(th.Fonts["Secondary"],
			//	title, th.Colors["Dark"],
			//	th.Colors["LightGrayII"],
			//	th.Colors["LightGrayII"],
			//	th.Colors["Dark"], icon,
			//	CurrentCurrentPageColor(rc.ShowPage, title, navItemIconColor, th.Colors["Primary"]),
			//	nav.TextSize, nav.IconSize,
			//	nav.Width, nav.Height,
			//	nav.PaddingVertical, nav.PaddingHorizontal, nav.PaddingVertical, nav.PaddingHorizontal)
			for navButton.Clicked() {
				rc.ShowPage = title
				SetPage(rc, page)
			}
			//menuItem.MenuLayout(gtx, navButton)
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}

//func navMenuLine(gtx *layout.Context, th *theme.DuoUItheme) func(gtx layout.Context)layout.Dimensions{
//	return func(gtx layout.Context)layout.Dimensions{
//		gelook.DuoUIdrawRectangle(gtx, int(navItemWidth), 1, th.Colors["LightGrayIII"], [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
//	}
//}
