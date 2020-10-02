package duoui

import (
	"gioui.org/layout"

	"github.com/gioapp/wallet/gui/component"
)

func (ui *DuoUI) DuoUImenu() func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		ui.ly.Navigation.Width = 48
		ui.ly.Navigation.Height = 48
		ui.ly.Navigation.TextSize = 0
		ui.ly.Navigation.IconSize = 24
		ui.ly.Navigation.PaddingVertical = 4
		ui.ly.Navigation.PaddingHorizontal = 0
		if ui.ly.Viewport > 740 {
			ui.ly.Navigation.Width = 96
			ui.ly.Navigation.Height = 72
			ui.ly.Navigation.TextSize = 48
			ui.ly.Navigation.IconSize = 36
			ui.ly.Navigation.PaddingVertical = 8
			ui.ly.Navigation.PaddingHorizontal = 0
		}

		layout.Flex{
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
			Spacing:   layout.SpaceEvenly}.Layout(gtx,
			layout.Rigid(component.MainNavigation(ui.rc, gtx, ui.ly.Theme, ui.ly.Pages, ui.ly.Navigation)),
		)
		return layout.Dimensions{}
	}
}
