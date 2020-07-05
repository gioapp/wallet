package duoui

import (
	svg "github.com/gioapp/wallet/pkg/ico/svg"
	"image"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/gioapp/wallet/gui/component"
	"github.com/gioapp/wallet/pkg/gelook"
)

// Main wallet screen
func (ui *DuoUI) DuoUIsplashScreen() {
	ui.ly.Theme.DuoUIcontainer(0, ui.ly.Theme.Colors["Dark"]).Layout(ui.ly.Context, layout.Center, func() {
		logo, _ := gelook.NewDuoUIicon(svg.ParallelCoin)
		layout.Flex{Axis: layout.Vertical}.Layout(ui.ly.Context,
			layout.Rigid(func() {
				layout.Flex{Axis: layout.Horizontal}.Layout(ui.ly.Context,
					layout.Rigid(func() {
						layout.UniformInset(unit.Dp(8)).Layout(ui.ly.Context, func() {
							size := ui.ly.Context.Px(unit.Dp(256)) - 2*ui.ly.Context.Px(unit.Dp(8))
							if logo != nil {
								logo.Color = gelook.HexARGB(ui.ly.Theme.Colors["Light"])
								logo.Layout(ui.ly.Context, unit.Px(float32(size)))
							}
							ui.ly.Context.Dimensions = layout.Dimensions{
								Size: image.Point{X: size, Y: size},
							}
						})
					}),
					layout.Flexed(1, func() {
						layout.UniformInset(unit.Dp(60)).Layout(ui.ly.Context, func() {
							txt := ui.ly.Theme.H1("PLAN NINE FROM FAR, FAR AWAY SPACE")
							txt.Font.Typeface = ui.ly.Theme.Fonts["Secondary"]
							txt.Color = ui.ly.Theme.Colors["Light"]
							txt.Layout(ui.ly.Context)
						})
					}),
				)
			}),
			layout.Flexed(1, component.DuoUIlogger(ui.rc, ui.ly.Context, ui.ly.Theme)),
		)
	})
}

// Main wallet screen
func (ui *DuoUI) DuoUImainScreen() {
	ui.ly.Theme.DuoUIcontainer(0, ui.ly.Theme.Colors["Dark"]).Layout(ui.ly.Context, layout.Center, func() {
		layout.Flex{Axis: layout.Vertical}.Layout(ui.ly.Context,
			layout.Rigid(ui.DuoUIheader()),
			layout.Flexed(1, ui.DuoUIbody()),
			layout.Rigid(ui.DuoUIfooter()),
		)
	})
}
