package component

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/gui/rcd"
)

var (
	buttonDialogCancel = new(widget.Clickable)
	buttonDialogOK     = new(widget.Clickable)
	buttonDialogClose  = new(widget.Clickable)
)

func DuoUIdialog(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) {
	// cs := gtx.Constraints
	container.DuoUIcontainer(th, 0, "ee000000").Layout(gtx, layout.Center, func(gtx layout.Context) layout.Dimensions {
		//cs := gtx.Constraints
		layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				//rr := float32(gtx.Px(unit.Dp(0)))
				//clip.Rect{
				//	Rect: f32.Rectangle{Max: f32.Point{
				//		//X: float32(cs.Width.Max),
				//		//Y: float32(cs.Height.Max),
				//	}},
				//	NE: rr, NW: rr, SE: rr, SW: rr,
				//}.Op(gtx.Ops).Add(gtx.Ops)
				//fill(gtx, gelook.HexARGB("ff888888"))
				//pointer.Rect(image.Rectangle{Max: gtx.Dimensions.Size}).Add(gtx.Ops)
				return layout.Dimensions{}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				//if gtx.Constraints.Width.Max > 500 {
				//	gtx.Constraints.Width.Max = 500
				//}
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					layout.Inset{Top: unit.Dp(16), Bottom: unit.Dp(16), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						//layout.Flex{
						//	Axis:      layout.Vertical,
						//	Alignment: layout.Middle,
						//}.Layout(gtx,
						//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
						//layout.Flex{
						//	Axis:      layout.Horizontal,
						//	Alignment: layout.Middle,
						//}.Layout(gtx,
						//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
						//layout.Inset{Top: unit.Dp(0), Bottom: unit.Dp(8), Left: unit.Dp(4), Right: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context)layout.Dimensions{
						//cur := th.DuoUIlabel(unit.Dp(14), rc.Dialog.Text)
						//cur.Font.Typeface = th.Fonts["Primary"]
						//cur.Color = th.Colors["Dark"]
						//cur.Alignment = text.Start
						//cur.Layout(gtx)
						//})
						//}),
						//)
						//}),
						//layout.Rigid(func(gtx layout.Context)layout.Dimensions {
						//
						//	rc.Dialog.CustomField()
						//
						//}),
						//layout.Rigid(func(gtx layout.Context)layout.Dimensions {
						//	layout.Flex{
						//		Axis:      layout.Horizontal,
						//		Alignment: layout.Middle,
						//	}.Layout(gtx,
						//		layout.Rigid(
						//			dialogButon(gtx, th,
						//				rc.Dialog.Red, rc.Dialog.RedLabel,
						//				"ffcf3030", "iconCancel", "ffcf8080",
						//				buttonDialogCancel)),
						//		layout.Rigid(dialogButon(gtx, th, rc.Dialog.Green, rc.Dialog.GreenLabel, "ff30cf30", "iconOK", "ff80cf80", buttonDialogOK)),
						//		layout.Rigid(dialogButon(gtx, th, rc.Dialog.Orange, rc.Dialog.OrangeLabel, "ffcf8030", "iconClose", "ffcfa880", buttonDialogClose)),
						//	)
						//}),
						//)
						return layout.Dimensions{}
					})
					return layout.Dimensions{}
				})
				return layout.Dimensions{}
			}))
		return layout.Dimensions{}
	})

}

func dialogButon(gtx *layout.Context, th *theme.DuoUItheme, f func(), t, bgColor, icon, iconColor string, button *widget.Clickable) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		if f != nil {
			//var b gelook.DuoUIbutton
			//layout.Inset{Top: unit.Dp(8), Bottom: unit.Dp(8), Left: unit.Dp(8), Right: unit.Dp(8)}.Layout(gtx, func(gtx layout.Context)layout.Dimensions{
			//b = th.DuoUIbutton(th.Fonts["Primary"], t, th.Colors["Dark"], bgColor, th.Colors["Info"], bgColor, icon, iconColor, 16, 32, 120, 64, 0, 0, 0, 0)
			//for button.Clicked() {
			//	f()
			//}
			//b.MenuLayout(gtx, button)
			//})
		}
		return layout.Dimensions{}
	}
}
