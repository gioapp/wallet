package pages

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/rcd"
)

var (
	consoleInputField = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	consoleOutputList = &layout.List{
		Axis:        layout.Vertical,
		ScrollToEnd: true,
	}
)

func Console(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title: "CONSOLE",
		//Command:       func(gtx layout.Context) layout.Dimensions {},
		Border:      4,
		BorderColor: "ff000000",
		//Header:        func(gtx layout.Context) layout.Dimensions {},
		HeaderPadding: 0,
		Body:          consoleBody(rc, gtx, th),
		BodyBgColor:   "ff000000",
		BodyPadding:   4,
		//Footer:        func(gtx layout.Context) layout.Dimensions {},
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

func consoleBody(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//	th.DuoUIcontainer(0, "ff000000").Layout(gtx, layout.N, func(gtx layout.Context) layout.Dimensions {
		//		layout.Flex{}.Layout(gtx, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		//layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//	layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceAround}.Layout(gtx,
		//		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		//consoleOutputList.Layout(gtx,
		//len(rc.ConsoleHistory.Commands), func(i int) {
		//	t := rc.ConsoleHistory.Commands[i]
		//	layout.Flex{
		//		Axis:      layout.Vertical,
		//		Alignment: layout.End,
		//	}.Layout(gtx,
		//		layout.Rigid(component.Label(
		//			gtx, th, th.Fonts["Mono"],
		//			12, th.Colors["Light"], "ds://"+t.ComID)),
		//		layout.Rigid(component.Label(
		//			gtx, th, th.Fonts["Mono"],
		//			12, th.Colors["Light"], t.Out)),
		//	)
		//})
		//return layout.Dimensions{}
		//}),
		//layout.Rigid(
		//component.ConsoleInput(gtx, th,
		//	consoleInputField, "Run command",
		//	func(e gel.SubmitEvent) {
		//		rc.ConsoleHistory.Commands = append(
		//			rc.ConsoleHistory.Commands,
		//			model.DuoUIconsoleCommand{
		//				ComID: e.Text,
		//				Time:  time.Time{},
		//				Out:   rc.ConsoleCmd(e.Text),
		//			})
		//return layout.Dimensions{}
		//})),
		//)
		//	return layout.Dimensions{}
		//})
		//return layout.Dimensions{}
		//}),
		//)
		//	return layout.Dimensions{}
		//})
		return layout.Dimensions{}
	}
}
