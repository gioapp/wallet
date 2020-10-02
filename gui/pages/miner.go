package pages

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/rcd"
)

func Miner(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title:   "MINER",
		TxColor: "",
		//Command:       func(gtx layout.Context)layout.Dimensions{},
		Border:      4,
		BorderColor: th.Colors["Light"],
		//Header:        func(gtx layout.Context)layout.Dimensions{},
		HeaderBgColor: "",
		HeaderPadding: 0,
		Body:          DuoUIminer(rc, gtx, th),
		BodyBgColor:   th.Colors["Dark"],
		BodyPadding:   0,
		//Footer:        func(gtx layout.Context)layout.Dimensions{},
		FooterBgColor: "",
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

func DuoUIminer(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		rc.GetDuoUIhashesPerSecList()
		layout.Flex{}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceAround,
					}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							//	consoleOutputList.Layout(gtx, rc.Status.Kopach.Hps.Len(), func(i int) {
							//		t := rc.Status.Kopach.Hps.Get(i)
							//		layout.Flex{
							//			Axis:      layout.Vertical,
							//			Alignment: layout.End,
							//		}.Layout(gtx,
							//			layout.Rigid(func(gtx layout.Context)layout.Dimensions {
							//				sat := th.Body1(fmt.Sprint(t))
							//				sat.Font.Typeface = th.Fonts["Mono"]
							//				sat.Color = gelook.HexARGB(th.Colors["Dark"])
							//				sat.Layout(gtx)
							//			}),
							//		)
							//	})
							return layout.Dimensions{}
						}),
					)
					return layout.Dimensions{}
				})
				return layout.Dimensions{}
			}),
		)
		return layout.Dimensions{}
	}
}
