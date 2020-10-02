package duoui

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/container"
)

func (ui *DuoUI) DuoUIsidebar() func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		container.DuoUIcontainer(ui.ly.Theme, 0, ui.ly.Theme.Colors["Dark"]).Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(ui.DuoUImenu()),
			)
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}
