package duoui

import (
	"gioui.org/layout"

	"github.com/gioapp/wallet/gui/component"
)

var (
	footerNav = &layout.List{
		Axis: layout.Horizontal,
	}
)

func (ui *DuoUI) DuoUIfooter() func() {
	return func() {
		footer := ui.ly.Theme.DuoUIcontainer(0, ui.ly.Theme.Colors["Dark"])
		footer.FullWidth = true
		footer.Layout(ui.ly.Context, layout.N, func() {
			layout.Flex{Spacing: layout.SpaceBetween}.Layout(ui.ly.Context,
				layout.Rigid(component.FooterLeftMenu(ui.rc, ui.ly.Context, ui.ly.Theme, ui.ly.Pages)),
				layout.Flexed(1, func() {}),
				layout.Rigid(component.FooterRightMenu(ui.rc, ui.ly.Context, ui.ly.Theme, ui.ly.Pages)),
			)
		})
	}
}
