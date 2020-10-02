package component

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/gui/rcd"
)

var (
	itemsList = &layout.List{
		Axis: layout.Vertical,
	}
	singleItem = &layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}
)

func DuoUIstatus(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//th.DuoUIcontainer(8, th.Colors["Light"]).Layout(gtx, layout.NW, func(gtx layout.Context)layout.Dimensions{
		//	bigStatus := []func(){
		//		listItem(gtx, th, 22, 6, "EditorMonetizationOn", "BALANCE :", rc.Status.Wallet.Balance.Load()+" "+rc.Settings.Abbrevation),
		//		th.DuoUIline(gtx, 8, 0, 1, th.Colors["LightGray"]),
		//		listItem(gtx, th, 22, 6, "MapsLayersClear", "UNCONFIRMED :", rc.Status.Wallet.Unconfirmed.Load()+" "+
		//			rc.Settings.Abbrevation),
		//		th.DuoUIline(gtx, 8, 0, 1, th.Colors["LightGray"]),
		//		listItem(gtx, th, 22, 6, "CommunicationImportExport", "TRANSACTIONS :", fmt.Sprint(rc.Status.Wallet.TxsNumber.Load())),
		//
		//		th.DuoUIline(gtx, 8, 0, 1, th.Colors["LightGray"]),
		//		listItem(gtx, th, 16, 4, "DeviceWidgets", "Block Count :", fmt.Sprint(rc.Status.Node.BlockCount.Load())),
		//
		//		th.DuoUIline(gtx, 4, 0, 1, th.Colors["LightGray"]),
		//		listItem(gtx, th, 16, 4, "ImageTimer", "Difficulty :", fmt.Sprint(rc.Status.Node.Difficulty.Load())),
		//
		//		th.DuoUIline(gtx, 4, 0, 1, th.Colors["LightGray"]),
		//		listItem(gtx, th, 16, 4, "NotificationVPNLock", "Connections :", fmt.Sprint(rc.Status.Node.ConnectionCount.Load())),
		//	}
		//	itemsList.Layout(gtx, len(bigStatus), func(i int) {
		//		layout.UniformInset(unit.Dp(0)).Layout(gtx, bigStatus[i])
		//	})
		//})
		return layout.Dimensions{}
	}
}

func listItem(gtx *layout.Context, th *theme.DuoUItheme, size, top int, iconName, name, value string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//icon := th.Icons[iconName]
		//layout.Flex{
		//	Axis:    layout.Horizontal,
		//	Spacing: layout.SpaceBetween,
		//}.Layout(gtx,
		//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//		layout.Flex{}.Layout(gtx,
		//			layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//				layout.Inset{Top: unit.Dp(float32(top)), Bottom: unit.Dp(0), Left: unit.Dp(0), Right: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context)layout.Dimensions{
		//					if icon != nil {
		//						icon.Color = gelook.HexARGB(th.Colors["Dark"])
		//						icon.Layout(gtx, unit.Px(float32(size)))
		//					}
		//					gtx.Dimensions = layout.Dimensions{
		//						Size: image.Point{X: size, Y: size},
		//					}
		//				})
		//			}),
		//			layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//				txt := th.DuoUIlabel(unit.Dp(float32(size)), name)
		//				txt.Font.Typeface = th.Fonts["Primary"]
		//				txt.Color = th.Colors["Primary"]
		//				txt.Layout(gtx)
		//			}),
		//		)
		//	}),
		//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//		value := th.H5(value)
		//		value.TextSize = unit.Dp(float32(size))
		//		value.Font.Typeface = th.Fonts["Primary"]
		//		value.Color = th.Colors["Dark"]
		//		value.Alignment = text.End
		//		value.Layout(gtx)
		//	}),
		//)
		return layout.Dimensions{}
	}
}
