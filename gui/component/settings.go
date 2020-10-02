// SPDX-License-Identifier: Unlicense OR MIT
package component

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/gui/rcd"
)

var (
	groupsList = &layout.List{
		Axis:      layout.Horizontal,
		Alignment: layout.Start,
	}
)

type Field struct {
	//Field *pod.Field
}

func SettingsTabs(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//groupsNumber := len(rc.Settings.Daemon.Schema.Groups)
		//groupsList.Layout(gtx, groupsNumber, func(i int) {
		//	layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//		color := th.Colors["Light"]
		//		bgColor := th.Colors["Dark"]
		//		i = groupsNumber - 1 - i
		//		t := rc.Settings.Daemon.Schema.Groups[i]
		//		txt := fmt.Sprint(t.Legend)
		//		for rc.Settings.Tabs.TabsList[txt].Clicked() {
		//			rc.Settings.Tabs.Current = txt
		//		}
		//		if rc.Settings.Tabs.Current == txt {
		//			color = th.Colors["Dark"]
		//			bgColor = th.Colors["Light"]
		//		}
		//		th.DuoUIbutton(th.Fonts["Primary"],
		//			txt, color, bgColor, "", "", "", "",
		//			16, 0, 80, 32, 4, 4, 4, 4).Layout(gtx, rc.Settings.Tabs.TabsList[txt])
		//	})
		//})
		return layout.Dimensions{}
	}
}

func DuoUIinputField(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, f *Field) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//switch f.Field.Type {
		//case "stringSlice":
		//	switch f.Field.InputType {
		//	case "text":
		//		if f.Field.Model != "MiningAddrs" {
		//			StringsArrayEditor(gtx, th, (rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor),
		//				(rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor).Text(),
		//				func(e widget.EditorEvent) {
		//					rc.Settings.Daemon.Config[f.Field.Model] = strings.Fields((rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor).Text())
		//					Debug()
		//					if e != nil {
		//						rc.SaveDaemonCfg()
		//					}
		//				})()
		//		}
		//	default:
		//
		//	}
		//case "input":
		//	switch f.Field.InputType {
		//	case "text":
		//		Editor(gtx, th, (rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor),
		//			(rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor).Text(),
		//			func(e widget.EditorEvent) {
		//				txt := rc.Settings.Daemon.Widgets[f.Field.Model].(*widget.Editor).Text()
		//				rc.Settings.Daemon.Config[f.Field.Model] = txt
		//				if e != nil {
		//					rc.SaveDaemonCfg()
		//				}
		//			})()
		//	case "number":
		//		Editor(gtx, th, (rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor), (rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor).Text(),
		//			func(e widget.EditorEvent) {
		//				number, err := strconv.Atoi((rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor).Text())
		//				if err == nil {
		//				}
		//				rc.Settings.Daemon.Config[f.Field.Model] = number
		//				if e != nil {
		//					rc.SaveDaemonCfg()
		//				}
		//			})()
		//	case "decimal":
		//		Editor(gtx, th, (rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor), (rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor).Text(),
		//			func(e widget.EditorEvent) {
		//				decimal, err := strconv.ParseFloat((rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor).Text(), 64)
		//				if err != nil {
		//				}
		//				rc.Settings.Daemon.Config[f.Field.Model] = decimal
		//				if e != nil {
		//					rc.SaveDaemonCfg()
		//				}
		//			})()
		//	case "password":
		//		e := th.DuoUIeditor(f.Field.Label, "DocText", "DocBg", 32)
		//		e.Font.Typeface = th.Fonts["Primary"]
		//		e.Font.Style = text.Italic
		//		e.Layout(gtx, (rc.Settings.Daemon.Widgets[f.Field.Model]).(*widget.Editor))
		//	default:
		//	}
		//case "switch":
		//	sw := th.DuoUIcheckBox(f.Field.Label, th.Colors["Primary"],
		//		th.Colors["Primary"])
		//	sw.PillColor = th.Colors["LightGray"]
		//	sw.PillColorChecked = th.Colors["LightGrayI"]
		//	sw.CircleColor = th.Colors["LightGrayII"]
		//	sw.CircleColorChecked = th.Colors["Primary"]
		//	sw.DrawLayout(gtx,
		//		(rc.Settings.Daemon.Widgets[f.Field.Model]).(*gel.CheckBox))
		//	if (rc.Settings.Daemon.Widgets[f.Field.Model]).(*gel.CheckBox).Checked(gtx) {
		//		if !*rc.Settings.Daemon.Config[f.Field.Model].(*bool) {
		//			tt := true
		//			rc.Settings.Daemon.Config[f.Field.Model] = &tt
		//			rc.SaveDaemonCfg()
		//		}
		//	} else {
		//		if *rc.Settings.Daemon.Config[f.Field.Model].(*bool) {
		//			ff := false
		//			rc.Settings.Daemon.Config[f.Field.Model] = &ff
		//			rc.SaveDaemonCfg()
		//		}
		//	}
		//case "radio":
		//	// radioButtonsGroup := (duo.Configuration.Settings.Daemon.Widgets[fieldName]).(*widget.Enum)
		//	// layout.Flex{}.Layout(gtx,
		//	//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//	//		duo.Theme.RadioButton("r1", "RadioButton1").Layout(gtx,  radioButtonsGroup)
		//	//
		//	//	}),
		//	//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//	//		duo.Theme.RadioButton("r2", "RadioButton2").Layout(gtx, radioButtonsGroup)
		//	//
		//	//	}),
		//	//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//	//		duo.Theme.RadioButton("r3", "RadioButton3").Layout(gtx, radioButtonsGroup)
		//	//
		//	//	}))
		//default:
		//	// duo.Theme.CheckBox("Checkbox").Layout(gtx, (duo.Configuration.Settings.Daemon.Widgets[fieldName]).(*widget.CheckBox))
		//
		//}
		return layout.Dimensions{}
	}
}

//
// var typeRegistry = make(map[string]reflect.Type)
//
// func makeInstance(name string) interface{} {
//	v := reflect.New(typeRegistry["cx.DuoUIconfigurationig."+name]).Elem()
//	// Maybe fill in fields here if necessary
//	return v.Interface()
// }

func SettingsFieldLabel(gtx *layout.Context, th *theme.DuoUItheme, f *Field) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//	name := th.H6(fmt.Sprint(f.Field.Label))
		//	name.Color = th.Colors["Dark"]
		//	name.Font.Typeface = th.Fonts["Primary"]
		//	name.Layout(gtx)
		//})
		return layout.Dimensions{}
	}
}

func SettingsFieldDescription(gtx *layout.Context, th *theme.DuoUItheme, f *Field) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//	desc := th.Body2(fmt.Sprint(f.Field.Description))
		//	desc.Font.Typeface = th.Fonts["Primary"]
		//	desc.Color = th.Colors["Dark"]
		//	desc.Layout(gtx)
		//})
		return layout.Dimensions{}
	}
}
