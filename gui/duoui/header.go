package duoui

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/gioapp/wallet/gui/component"
	"github.com/gioapp/wallet/pkg/gel"
	"github.com/gioapp/wallet/pkg/gelook"
)

var (
	logoButton = new(gel.Button)
	headerList = &layout.List{
		Axis:      layout.Horizontal,
		Alignment: layout.Start,
	}
)

func (ui *DuoUI) DuoUIheader() func() {
	return func() {
		iconSize := 32
		iconWidth := 48
		iconHeight := 48
		iconPaddingVertical := 3
		iconPaddingHorizontal := 3
		if ui.ly.Viewport > 740 {
			iconSize = 64
			iconWidth = 96
			iconHeight = 96
			iconPaddingVertical = 6
			iconPaddingHorizontal = 6
		}
		ui.ly.Theme.DuoUIcontainer(0, ui.ly.Theme.Colors["Dark"]).Layout(ui.ly.Context, layout.NW, func() {
			layout.Flex{
				Axis:      layout.Horizontal,
				Spacing:   layout.SpaceBetween,
				Alignment: layout.Middle,
			}.Layout(ui.ly.Context,
				layout.Rigid(func() {
					var logoMeniItem gelook.DuoUIbutton
					logoMeniItem = ui.ly.Theme.DuoUIbutton("", "", "", ui.ly.Theme.Colors["Dark"], "", "", "logo", ui.ly.Theme.Colors["Light"], 0, iconSize, iconWidth, iconHeight, iconPaddingVertical, iconPaddingHorizontal, iconPaddingVertical, iconPaddingHorizontal)
					for logoButton.Clicked(ui.ly.Context) {
						ui.ly.Theme.ChangeLightDark()
					}
					logoMeniItem.IconLayout(ui.ly.Context, logoButton)
				}),
				layout.Flexed(1, component.HeaderMenu(ui.rc, ui.ly.Context, ui.ly.Theme, ui.ly.Pages)),
				layout.Rigid(component.Label(ui.ly.Context, ui.ly.Theme, ui.ly.Theme.Fonts["Primary"], 12, ui.ly.Theme.Colors["Light"], ui.rc.Status.Wallet.Balance.Load()+" "+ui.rc.Settings.Abbrevation)),
				layout.Rigid(component.Label(ui.ly.Context, ui.ly.Theme, ui.ly.Theme.Fonts["Primary"], 12, ui.ly.Theme.Colors["Light"], fmt.Sprint(ui.ly.Viewport))),
			)
		})
	}
}

func renderIcon(gtx *layout.Context, icon *gelook.DuoUIicon) func() {
	return func() {
		icon.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0x55, B: 0x30}
		icon.Layout(gtx, unit.Dp(float32(48)))
		pointer.Rect(image.Rectangle{Max: image.Point{
			X: 64,
			Y: 64,
		}}).Add(gtx.Ops)
	}
}
