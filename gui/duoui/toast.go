package duoui

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

var (
	buttonToastOK = new(widget.Clickable)
	listToasts    = &layout.List{
		Axis:        layout.Vertical,
		ScrollToEnd: false,
		Alignment:   0,
		Position:    layout.Position{},
	}
)

// func (ui *DuoUI) DuoUItoastSys() {
//	layout.Align(layout.NE).Layout(gtx, func(gtx layout.Context)layout.Dimensions{
//		listToasts.Layout(gtx, len(ui.rc.Toasts), func(i int) {
//			layout.UniformInset(unit.Dp(16)).Layout(gtx, ui.rc.Toasts[i])
//		})
//	})
// }

func (ui *DuoUI) toastButton(text, txtColor, bgColor, txtHoverColor, bgHoverColor, icon, iconColor string, button *widget.Clickable) func(gtx layout.Context) layout.Dimensions {
	//var b material.Button
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Inset{Top: unit.Dp(8), Bottom: unit.Dp(8), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx, func(gtx layout.Context)layout.Dimensions{
		//b = ui.ly.Theme.DuoUIbutton(ui.ly.Theme.Fonts["Primary"], text, txtColor, bgColor, txtHoverColor, bgHoverColor, icon, iconColor, 16, 24, 120, 60, 0, 0, 0, 0)
		//for button.Clicked() {
		// ui.rc.ShowToast = false
		//}
		//b.Layout(gtx, button)
		//})
		return layout.Dimensions{}
	}
}

// func (ui *DuoUI) toastAdd() {
//	ui.rc.Toasts = append(ui.rc.Toasts, func(gtx layout.Context)layout.Dimensions{
//		//iconOK, _ := view.NewDuoUIicon(icons.NavigationCheck)
//		theme.DuoUIdrawRectangle(gtx, 418, 160, "aa000000", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
//		theme.DuoUIdrawRectangle(gtx, 408, 150, ui.ly.Theme.Color.Primary, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
//		layout.Flex{
//			Axis:      layout.Vertical,
//			Alignment: layout.Middle,
//		}.Layout(gtx,
//			layout.Rigid(func(gtx layout.Context)layout.Dimensions {
//				layout.Flex{
//					Axis:      layout.Horizontal,
//					Alignment: layout.Middle,
//				}.Layout(gtx,
//					layout.Rigid(func(gtx layout.Context)layout.Dimensions {
//						layout.Align(layout.Center).Layout(gtx, func(gtx layout.Context)layout.Dimensions{
//							layout.Inset{Top: unit.Dp(24), Bottom: unit.Dp(8), Left: unit.Dp(0), Right: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context)layout.Dimensions{
//								cur := ui.ly.Theme.H4("TOAST MESSAGE!!!")
//								cur.Color = ui.ly.Theme.Color.Light
//								cur.Alignment = text.Start
//								cur.Layout(gtx)
//							})
//						})
//					}),
//				)
//			}),
//			layout.Rigid(func(gtx layout.Context)layout.Dimensions {
//				layout.Flex{
//					Axis:      layout.Horizontal,
//					Alignment: layout.Middle,
//				}.Layout(gtx)
//				//layout.Rigid(toastButton("OK", "ffcfcfcf", "ff308030", "ffcfcfcf", duo, rc, buttonToastOK, iconOK)),
//
//			}),
//		)
//	})
//	go func(ui *DuoUI) {
//		time.Sleep(3 * time.Second)
//		ui.rc.Toasts[len(ui.rc.Toasts)-1] = nil // or the zero value of T
//		ui.rc.Toasts = ui.rc.Toasts[:len(ui.rc.Toasts)-1]
//		op.InvalidateOp{}.Add(gtx.Ops)
//	}(ui)
// }
