package duoui

import (
	"fmt"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/gui/component"
	"image"
	"image/color"

	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
)

var (
	logoButton = new(widget.Clickable)
	headerList = &layout.List{
		Axis:      layout.Horizontal,
		Alignment: layout.Start,
	}
)

func (ui *DuoUI) DuoUIheader() func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//iconSize := 32
		//iconWidth := 48
		//iconHeight := 48
		//iconPaddingVertical := 3
		//iconPaddingHorizontal := 3
		//if ui.ly.Viewport > 740 {
		//	iconSize = 64
		//	iconWidth = 96
		//	iconHeight = 96
		//	iconPaddingVertical = 6
		//	iconPaddingHorizontal = 6
		//}
		container.DuoUIcontainer(ui.ly.Theme, 0, ui.ly.Theme.Colors["Dark"]).Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:      layout.Horizontal,
				Spacing:   layout.SpaceBetween,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					logoMeniItem := material.IconButton(ui.ly.Theme.T, logoButton, ui.ly.Theme.Icons["logo"])
					logoMeniItem.Inset = layout.Inset{unit.Dp(5), unit.Dp(3), unit.Dp(5), unit.Dp(5)}
					logoMeniItem.Color = helper.HexARGB(ui.ly.Theme.Colors["Danger"])
					logoMeniItem.Size = unit.Dp(16)
					logoMeniItem.Background = helper.HexARGB(ui.ly.Theme.Colors["White"])
					//logoMeniItem = ui.ly.Theme.IconButton("", "", "", ui.ly.Theme.Colors["Dark"], "", "", "logo", ui.ly.Theme.Colors["Light"], 0, iconSize, iconWidth, iconHeight, iconPaddingVertical, iconPaddingHorizontal, iconPaddingVertical, iconPaddingHorizontal)
					//logoMeniItem = material.IconButton{}( ui.ly.Theme.Icons["logo"])
					for logoButton.Clicked() {
						ui.ly.Theme.ChangeLightDark()
					}
					return logoMeniItem.Layout(gtx)
				}),
				layout.Flexed(1, component.HeaderMenu(ui.rc, gtx, ui.ly.Theme, ui.ly.Pages)),
				layout.Rigid(component.Label(gtx, ui.ly.Theme, ui.ly.Theme.Fonts["Primary"], 12, ui.ly.Theme.Colors["Light"], ui.rc.Status.Wallet.Balance.Load()+" "+ui.rc.Settings.Abbrevation)),
				layout.Rigid(component.Label(gtx, ui.ly.Theme, ui.ly.Theme.Fonts["Primary"], 12, ui.ly.Theme.Colors["Light"], fmt.Sprint(ui.ly.Viewport))),
			)
		})
		return layout.Dimensions{}
	}
}

func renderIcon(gtx *layout.Context, icon *widget.Icon) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		icon.Color = color.RGBA{A: 0xff, R: 0xcf, G: 0x55, B: 0x30}
		icon.Layout(gtx, unit.Dp(float32(48)))
		pointer.Rect(image.Rectangle{Max: image.Point{
			X: 64,
			Y: 64,
		}}).Add(gtx.Ops)
		return layout.Dimensions{}
	}
}
