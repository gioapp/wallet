package duoui

import (
	"gioui.org/layout"
)

var (
	footerNav = &layout.List{
		Axis: layout.Horizontal,
	}
)

func (ui *DuoUI) DuoUIfooter() func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//footer := ui.ly.Theme.DuoUIcontainer(0, ui.ly.Theme.Colors["Dark"])
		//footer.FullWidth = true
		//footer.Layout(gtx, layout.N, func(gtx layout.Context)layout.Dimensions{
		//	layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
		//		layout.Rigid(component.FooterLeftMenu(ui.rc, gtx, ui.ly.Theme, ui.ly.Pages)),
		//		layout.Flexed(1, func(gtx layout.Context)layout.Dimensions{}),
		//		layout.Rigid(component.FooterRightMenu(ui.rc, gtx, ui.ly.Theme, ui.ly.Pages)),
		//	)
		//})
		return layout.Dimensions{}
	}
}
