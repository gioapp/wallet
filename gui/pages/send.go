package pages

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/gui/rcd"
)

var (
	layautList = &layout.List{
		Axis: layout.Vertical,
	}
	addressLineEditor = &widget.Editor{
		SingleLine: true,
	}
	amountLineEditor = &widget.Editor{
		SingleLine: true,
	}
	passLineEditor = &widget.Editor{
		SingleLine: true,
	}
	buttonPasteAddress = new(widget.Clickable)
	buttonPasteAmount  = new(widget.Clickable)
	buttonSend         = new(widget.Clickable)
	sendStruct         = new(send)
)

type send struct {
	address    string
	amount     float64
	passPharse string
}

func Send(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title:   "SEND",
		TxColor: "",
		//Command:       func(gtx layout.Context) layout.Dimensions {},
		Border:      4,
		BorderColor: th.Colors["Light"],
		//Header:        func(gtx layout.Context) layout.Dimensions {},
		HeaderBgColor: "",
		HeaderPadding: 0,
		//Body:          sendBody(rc, gtx, th),
		BodyBgColor: th.Colors["Light"],
		BodyPadding: 0,
		//Footer:        func(gtx layout.Context) layout.Dimensions {},
		FooterBgColor: "",
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

func sendBody(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				//widgets := []func(gtx layout.Context) layout.Dimensions{
				//	func(gtx layout.Context) layout.Dimensions {
				//		th.DuoUIcontainer(1, th.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx layout.Context) layout.Dimensions {
				//			//layout.Flex{}.Layout(gtx,
				//			//	layout.Flexed(1, component.Editor(gtx, th, addressLineEditor, "DUO address",
				//			//		func(e widget.EditorEvent) {
				//			//			sendStruct.address = addressLineEditor.Text()
				//			//		})),
				//			//	layout.Rigid(component.Button(gtx, th, buttonPasteAddress, th.Fonts["Primary"], 10, 13, 8, 12, 8,
				//			//		th.Colors["ButtonText"], th.Colors["ButtonBg"], "PASTE ADDRESS", func(gtx layout.Context) layout.Dimensions {
				//			//			addressLineEditor.SetText(clipboard.Get())
				//			//		})))
				//			return layout.Dimensions{}
				//		})
				//		return layout.Dimensions{}
				//	},
				//	func(gtx layout.Context) layout.Dimensions {
				//		th.DuoUIcontainer(1, th.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx layout.Context) layout.Dimensions {
				//			//layout.Flex{}.Layout(gtx,
				//			//	layout.Flexed(1, component.Editor(gtx, th, amountLineEditor,
				//			//		"DUO Amount", func(e widget.EditorEvent) {
				//			//			f, err := strconv.ParseFloat(amountLineEditor.Text(), 64)
				//			//			if err != nil {
				//			//			}
				//			//			sendStruct.amount = f
				//			//		})),
				//			//	layout.Rigid(component.Button(gtx, th, buttonPasteAmount, th.Fonts["Primary"], 10, 13, 8, 12, 8,
				//			//		th.Colors["ButtonText"], th.Colors["ButtonBg"], "PASTE AMOUNT", func(gtx layout.Context) layout.Dimensions {
				//			//			amountLineEditor.SetText(clipboard.Get())
				//			//		})))
				//			return layout.Dimensions{}
				//		})
				//		return layout.Dimensions{}
				//	},
				//	func(gtx layout.Context) layout.Dimensions {
				//		//layout.Flex{}.Layout(gtx,
				//			//layout.Rigid(component.Button(gtx, th, buttonSend, th.Fonts["Primary"], 14, 10, 10, 9, 10,
				//			//	th.Colors["ButtonText"], th.Colors["ButtonBg"], "SEND", func(gtx layout.Context) layout.Dimensions {
				//			//		rc.Dialog.Show = true
				//			//		rc.Dialog = &model.DuoUIdialog{
				//			//			Show:       true,
				//			//			Green:      rc.DuoSend(sendStruct.passPharse, sendStruct.address, 11),
				//			//			GreenLabel: "SEND",
				//			//			CustomField: func(gtx layout.Context) layout.Dimensions {
				//			//				layout.Flex{}.Layout(gtx,
				//			//					layout.Flexed(1, component.Editor(gtx, th, passLineEditor, "Enter your password",
				//			//						func(e widget.EditorEvent) {
				//			//							sendStruct.passPharse = passLineEditor.Text()
				//			//						})))
				//			//			},
				//			//			Red:      func(gtx layout.Context) layout.Dimensions { rc.Dialog.Show = false },
				//			//			RedLabel: "CANCEL",
				//			//			Title:    "Are you sure?",
				//			//			Text:     "Confirm ParallelCoin send",
				//			//		}
				//			//	})))
				//		return layout.Dimensions{}
				//	},
				//}
				//layautList.Layout(gtx, len(widgets), func(i int) {
				//	layout.UniformInset(unit.Dp(8)).Layout(gtx, widgets[i])
				//})
				return layout.Dimensions{}
			}))
		//Info("passPharse:" + sendStruct.passPharse)
		return layout.Dimensions{}
	}
}
