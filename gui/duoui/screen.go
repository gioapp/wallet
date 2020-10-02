package duoui

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/container"
)

// Main wallet screen
func (ui *DuoUI) DuoUIsplashScreen() {
	//ui.ly.Theme.DuoUIcontainer(0, ui.ly.Theme.Colors["Dark"]).Layout(gtx, layout.Center, func(gtx layout.Context)layout.Dimensions{
	//	logo, _ := gelook.NewDuoUIicon(svg.ParallelCoin)
	//	layout.Flex{Axis: layout.Vertical}.Layout(gtx,
	//		layout.Rigid(func(gtx layout.Context)layout.Dimensions {
	//			layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
	//				layout.Rigid(func(gtx layout.Context)layout.Dimensions {
	//					layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context)layout.Dimensions{
	//						size := gtx.Px(unit.Dp(256)) - 2*gtx.Px(unit.Dp(8))
	//						if logo != nil {
	//							logo.Color = gelook.HexARGB(ui.ly.Theme.Colors["Light"])
	//							logo.Layout(gtx, unit.Px(float32(size)))
	//						}
	//						gtx.Dimensions = layout.Dimensions{
	//							Size: image.Point{X: size, Y: size},
	//						}
	//					})
	//				}),
	//				layout.Flexed(1, func(gtx layout.Context)layout.Dimensions{
	//					layout.UniformInset(unit.Dp(60)).Layout(gtx, func(gtx layout.Context)layout.Dimensions{
	//						txt := ui.ly.Theme.H1("PLAN NINE FROM FAR, FAR AWAY SPACE")
	//						txt.Font.Typeface = ui.ly.Theme.Fonts["Secondary"]
	//						txt.Color = ui.ly.Theme.Colors["Light"]
	//						txt.Layout(gtx)
	//					})
	//				}),
	//			)
	//		}),
	//		layout.Flexed(1, component.DuoUIlogger(ui.rc, gtx, ui.ly.Theme)),
	//	)
	//})
}

// Main wallet screen
func (ui *DuoUI) DuoUImainScreen() {
	container.DuoUIcontainer(ui.ly.Theme, 0, ui.ly.Theme.Colors["Dark"]).Layout(ui.ly.Context, layout.Center, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			//layout.Rigid(ui.DuoUIheader()),
			layout.Flexed(1, ui.DuoUIbody()),
			//layout.Rigid(ui.DuoUIfooter()),
		)
	})
}
