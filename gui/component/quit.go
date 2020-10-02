package component

import (
	"gioui.org/widget"
)

var (
	buttonQuit = new(widget.Clickable)
)

//func QuitButton(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context)layout.Dimensions{
//	return func(gtx layout.Context)layout.Dimensions{
//		layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Constraints)layout.Dimensions {
//			var closeMeniItem gelook.DuoUIbutton
//			closeMeniItem = th.DuoUIbutton("", "", "", th.Colors["Dark"], "", "", "closeIcon", CurrentCurrentPageColor(rc.ShowPage, "CLOSE", th.Colors["Light"], th.Colors["Primary"]), footerMenuItemTextSize, footerMenuItemIconSize, footerMenuItemWidth, footerMenuItemHeight, 0, 0, 0, 0)
//			for buttonQuit.Clicked() {
//				rc.Dialog.Show = true
//				rc.Dialog = &model.DuoUIdialog{
//					Show: true,
//					Green: func(gtx layout.Context)layout.Dimensions{
//						//interrupt.Request()
//						// TODO make this close the window or at least switch to a shutdown screen
//						rc.Dialog.Show = false
//					},
//					GreenLabel: "QUIT",
//					Orange: func(gtx layout.Context)layout.Dimensions{
//						//interrupt.RequestRestart()
//						// TODO make this close the window or at least switch to a shutdown screen
//						rc.Dialog.Show = false
//					},
//					OrangeLabel: "RESTART",
//					Red:         func(gtx layout.Context)layout.Dimensions{ rc.Dialog.Show = false },
//					RedLabel:    "CANCEL",
//					CustomField: func(gtx layout.Context)layout.Dimensions{},
//					Title:       "Are you sure?",
//					Text:        "Confirm ParallelCoin close",
//				}
//			}
//			closeMeniItem.IconLayout(gtx, buttonQuit)
//		})
//	}
//}
