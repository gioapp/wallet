package duoui

import (
	"gioui.org/layout"
)

func (ui *DuoUI) DuoUIbody() func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(ui.DuoUIsidebar()),
			layout.Flexed(1, ui.DuoUIcontent()),
		)
		return layout.Dimensions{}
	}
}
