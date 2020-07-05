package component

import (
	"gioui.org/layout"

	"github.com/gioapp/wallet/gui/rcd"
	"github.com/gioapp/wallet/pkg/gelook"
)

func iconButton(rc *rcd.RcVar, gtx *layout.Context, th *gelook.DuoUItheme, page *gelook.DuoUIpage) func() {
	return func() {
		var logMenuItem gelook.DuoUIbutton
		logMenuItem = th.DuoUIbutton("", "", "", th.Colors["Dark"], "", "", "traceIcon", CurrentCurrentPageColor(rc.ShowPage, "LOG", th.Colors["Light"], th.Colors["Primary"]), footerMenuItemTextSize, footerMenuItemIconSize, footerMenuItemWidth, footerMenuItemHeight, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal)
		for buttonLog.Clicked(gtx) {
			SetPage(rc, page)
			rc.ShowPage = "LOG"
		}
		logMenuItem.IconLayout(gtx, buttonLog)
	}
}
