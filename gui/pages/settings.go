package pages

import (
	"gioui.org/widget"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"gioui.org/layout"

	"github.com/gioapp/wallet/gui/component"
	"github.com/gioapp/wallet/gui/rcd"
	"github.com/gioapp/wallet/pkg/gel"
)

var (
	settingsPanelElement  = gel.NewPanel()
	buttonSettingsRestart = new(widget.Clickable)
)

func Settings(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title:   "SETTINGS",
		TxColor: "",
		//Command:       func(gtx layout.Context)layout.Dimensions{},
		Border:        4,
		BorderColor:   th.Colors["Light"],
		Header:        SettingsHeader(rc, gtx, th),
		HeaderBgColor: "",
		HeaderPadding: 4,
		Body:          SettingsBody(rc, gtx, th),
		BodyBgColor:   th.Colors["Light"],
		BodyPadding:   0,
		//Footer:        func(gtx layout.Context)layout.Dimensions{},
		FooterBgColor: "",
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)

	// return th.DuoUIpage("SETTINGS", 0, func(gtx layout.Context)layout.Dimensions{}, component.ContentHeader(gtx, th, SettingsHeader(rc, gtx, th)), SettingsBody(rc, gtx, th), func(gtx layout.Context)layout.Dimensions{
	// var msg string
	// if rc.Settings.Daemon.Config["DisableBanning"].(*bool) != true{
	//	msg = "ima"
	// }else{
	//	msg = "nema"
	// //}
	// ttt := th.H6(fmt.Sprint(rc.Settings.Daemon.Config))
	// ttt.Color = gelook.HexARGB("ffcfcfcf")
	// ttt.Layout(gtx)
	// })
}

func SettingsHeader(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
			//layout.Rigid(component.SettingsTabs(rc, gtx, th)),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				// var settingsRestartButton gelook.DuoUIbutton
				// settingsRestartButton = th.DuoUIbutton(th.Fonts["Secondary"],
				// 	"restart",
				// 	th.Colors["Light"],
				// 	th.Colors["Dark"],
				// 	th.Colors["Dark"],
				// 	th.Colors["Light"],
				// 	"",
				// 	th.Colors["Light"],
				// 	23, 0, 80, 48, 4, 4)
				// for buttonSettingsRestart.Clicked() {
				// 	rc.SaveDaemonCfg()
				// }
				// settingsRestartButton.Layout(gtx, buttonSettingsRestart)
				return layout.Dimensions{}
			}),
		)
		return layout.Dimensions{}
	}
}

func SettingsBody(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		// layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context)layout.Dimensions{
		container.DuoUIcontainer(th, 16, th.Colors["Light"]).Layout(gtx, layout.N, func(gtx layout.Context) layout.Dimensions {
			//for _, fields := range rc.Settings.Daemon.Schema.Groups {
			//
			//	if fmt.Sprint(fields.Legend) == rc.Settings.Tabs.Current {
			//		settingsPanel := th.DuoUIpanel()
			//		settingsPanel.PanelObject = fields.Fields
			//		settingsPanel.ScrollBar = th.ScrollBar(16)
			//		settingsPanelElement.PanelObjectsNumber = len(fields.Fields)
			//		settingsPanel.Layout(gtx, settingsPanelElement, func(i int, in interface{}) {
			//			//settings := in.(pod.Fields)
			//			//t := settings[i]
			//
			//			//fieldsList.Layout(gtx, len(fields.Fields), func(il int) {
			//			i = settingsPanelElement.PanelObjectsNumber - 1 - i
			//			tl := component.Field{
			//				//Field: &settings[i],
			//			}
			//			layout.Flex{
			//				Axis: layout.Vertical,
			//			}.Layout(gtx,
			//				layout.Rigid(SettingsItemRow(rc, gtx, th, &tl)),
			//				layout.Rigid(th.DuoUIline(gtx, 4, 0, 1, th.Colors["LightGray"])))
			//		})
			//	}
			//}
			return layout.Dimensions{}
		})
		// })
		return layout.Dimensions{}
	}
}

func SettingsItemRow(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, f *component.Field) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{
			Axis:      layout.Horizontal,
			Alignment: layout.Middle,
		}.Layout(gtx)// layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//	gelook.DuoUIdrawRectangle(gtx, 30, 3, th.Colors["Light"], [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
		//// }),
		//layout.Flexed(0.62, func(gtx layout.Context)layout.Dimensions{
		//	layout.Flex{
		//		Axis:    layout.Vertical,
		//		Spacing: 4,
		//	}.Layout(gtx,
		//		layout.Rigid(component.SettingsFieldLabel(gtx, th, f)),
		//		layout.Rigid(component.SettingsFieldDescription(gtx, th, f)),
		//	)
		//}),
		//layout.Flexed(1, component.DuoUIinputField(rc, gtx, th, f)),

		return layout.Dimensions{}
	}
}
