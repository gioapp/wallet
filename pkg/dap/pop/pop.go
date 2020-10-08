package pop

import (
	"gioui.org/layout"
	"github.com/gioapp/wallet/pkg/container"
	"github.com/gioapp/wallet/pkg/theme"
)

func Popup(th *theme.Theme, content layout.Widget) {
	container.C().
		OutsideColor(th.Colors["White"]).
		BorderColor(th.Colors["Border"]).
		InsideColor(th.Colors["White"]).
		Margin(0).
		Border(1).
		Padding(4).
		Layout(content)

}
