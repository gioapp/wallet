package component

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/rcd"
)

func iconButton(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, page *page.DuoUIpage) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//var logMenuItem gelook.DuoUIbutton
		//logMenuItem = th.DuoUIbutton("", "", "", th.Colors["Dark"], "", "", "traceIcon", CurrentCurrentPageColor(rc.ShowPage, "LOG", th.Colors["Light"], th.Colors["Primary"]), footerMenuItemTextSize, footerMenuItemIconSize, footerMenuItemWidth, footerMenuItemHeight, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal)
		//for buttonLog.Clicked() {
		//	SetPage(rc, page)
		//	rc.ShowPage = "LOG"
		//}
		//logMenuItem.IconLayout(gtx, buttonLog)
		return layout.Dimensions{}
	}
}
