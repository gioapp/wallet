package duoui

import "gioui.org/layout"

func (ui *DuoUI) DuoUIcontent() func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		ui.rc.CurrentPage.Layout(gtx)
		return layout.Dimensions{}
	}
}
