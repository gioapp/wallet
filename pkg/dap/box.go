package box

import (
	"gioui.org/layout"
	"github.com/gioapp/wallet/pkg/container"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
//noReturn = func(gtx C) D { return D{} }
)

func BoxBase(color string, content layout.Widget) layout.Widget {
	return container.C().
		OutsideColor(color).
		BorderColor(color).
		InsideColor(color).
		Margin(0).
		Border(0).
		Padding(0).
		Layout(content)
}
func BoxPanel(th *theme.Theme, content layout.Widget) layout.Widget {
	return container.C().
		OutsideColor(th.Colors["PanelBg"]).
		BorderColor(th.Colors["Border"]).
		InsideColor(th.Colors["PanelBg"]).
		Margin(4).
		Border(1).
		Padding(8).
		Layout(content)
}
func BoxEditor(th *theme.Theme, content layout.Widget) layout.Widget {
	return container.C().
		OutsideColor(th.Colors["White"]).
		BorderColor(th.Colors["White"]).
		InsideColor(th.Colors["White"]).
		Margin(8).
		Border(8).
		Padding(8).
		Layout(content)
}
