package component

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/gui/rcd"
)

var (
	buttonHeader = new(widget.Clickable)
)

// func ContentHeader(gtx *layout.Context, th *theme.DuoUItheme, b func()) func(gtx layout.Context)layout.Dimensions{
//	return func(gtx layout.Context)layout.Dimensions{
//		hmin := gtx.Constraints.Width.Min
//		vmin := gtx.Constraints.Height.Min
//		layout.Stack{Alignment: layout.Center}.Layout(gtx,
//			layout.Expanded(func(gtx layout.Context)layout.Dimensions{
//				clip.Rect{
//					Rect: f32.Rectangle{Max: f32.Point{
//						X: float32(gtx.Constraints.Width.Min),
//						Y: float32(gtx.Constraints.Height.Min),
//					}},
//				}.Op(gtx.Ops).Add(gtx.Ops)
//				fill(gtx, gelook.HexARGB(th.Colors["Primary"]))
//			}),
//			layout.Stacked(func(gtx layout.Context)layout.Dimensions{
//				gtx.Constraints.Width.Min = hmin
//				gtx.Constraints.Height.Min = vmin
//				layout.UniformInset(unit.Dp(0)).Layout(gtx, b)
//			}),
//		)
//	}
// }

func HeaderMenu(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme, allPages *model.DuoUIpages) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//	headerNav := []func(){
		//		//headerMenuButton(rc, gtx, th, "", "CommunicationImportExport", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "NotificationNetworkCheck", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "NotificationSync", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "NotificationSyncDisabled", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "NotificationSyncProblem", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "NotificationVPNLock", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "MapsLayers", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "MapsLayersClear", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "ImageTimer", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "ImageRemoveRedEye", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "DeviceSignalCellular0Bar", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "ActionTimeline", buttonHeader),
		//		//headerMenuButton(rc, gtx, th, "", "HardwareWatch", buttonHeader),
		//	}
		//	footerNav.Layout(gtx, len(headerNav), func(i int) {
		//		layout.UniformInset(unit.Dp(0)).Layout(gtx, headerNav[i])
		//	})
		return layout.Dimensions{}
		//})
	}
}

func headerMenuButton(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, text, icon string, headerButton *widget.Clickable) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//var footerMenuItem gelook.DuoUIbutton
			//footerMenuItem = th.DuoUIbutton("", "", "", "", "", th.Colors["Dark"], icon, CurrentCurrentPageColor(rc.ShowPage, text, navItemIconColor, th.Colors["Primary"]), footerMenuItemTextSize, footerMenuItemIconSize, footerMenuItemWidth, footerMenuItemHeight, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal, footerMenuItemPaddingVertical, footerMenuItemPaddingHorizontal)
			for headerButton.Clicked() {
				rc.ShowPage = text
			}
			//footerMenuItem.IconLayout(gtx, headerButton)
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}
