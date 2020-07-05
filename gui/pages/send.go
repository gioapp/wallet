package pages

import (
	"strconv"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/gioapp/wallet/gui/component"
	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/gui/rcd"
	"github.com/gioapp/wallet/pkg/clipboard"
	"github.com/gioapp/wallet/pkg/gel"
	"github.com/gioapp/wallet/pkg/gelook"
)

var (
	layautList = &layout.List{
		Axis: layout.Vertical,
	}
	addressLineEditor = &gel.Editor{
		SingleLine: true,
	}
	amountLineEditor = &gel.Editor{
		SingleLine: true,
	}
	passLineEditor = &gel.Editor{
		SingleLine: true,
	}
	buttonPasteAddress = new(gel.Button)
	buttonPasteAmount  = new(gel.Button)
	buttonSend         = new(gel.Button)
	sendStruct         = new(send)
)

type send struct {
	address    string
	amount     float64
	passPharse string
}

func Send(rc *rcd.RcVar, gtx *layout.Context, th *gelook.DuoUItheme) *gelook.DuoUIpage {
	page := gelook.DuoUIpage{
		Title:         "SEND",
		TxColor:       "",
		Command:       func() {},
		Border:        4,
		BorderColor:   th.Colors["Light"],
		Header:        func() {},
		HeaderBgColor: "",
		HeaderPadding: 0,
		Body:          sendBody(rc, gtx, th),
		BodyBgColor:   th.Colors["Light"],
		BodyPadding:   0,
		Footer:        func() {},
		FooterBgColor: "",
		FooterPadding: 0,
	}
	return th.DuoUIpage(page)
}

func sendBody(rc *rcd.RcVar, gtx *layout.Context, th *gelook.DuoUItheme) func() {
	return func() {
		layout.Flex{}.Layout(gtx,
			layout.Rigid(func() {
				widgets := []func(){
					func() {
						th.DuoUIcontainer(1, th.Colors["Gray"]).Layout(gtx, layout.Center, func() {
							layout.Flex{}.Layout(gtx,
								layout.Flexed(1, component.Editor(gtx, th, addressLineEditor, "DUO address",
									func(e gel.EditorEvent) {
										sendStruct.address = addressLineEditor.Text()
									})),
								layout.Rigid(component.Button(gtx, th, buttonPasteAddress, th.Fonts["Primary"], 10, 13, 8, 12, 8,
									th.Colors["ButtonText"], th.Colors["ButtonBg"], "PASTE ADDRESS", func() {
										addressLineEditor.SetText(clipboard.Get())
									})))
						})
					},
					func() {
						th.DuoUIcontainer(1, th.Colors["Gray"]).Layout(gtx, layout.Center, func() {
							layout.Flex{}.Layout(gtx,
								layout.Flexed(1, component.Editor(gtx, th, amountLineEditor,
									"DUO Amount", func(e gel.EditorEvent) {
										f, err := strconv.ParseFloat(amountLineEditor.Text(), 64)
										if err != nil {
										}
										sendStruct.amount = f
									})),
								layout.Rigid(component.Button(gtx, th, buttonPasteAmount, th.Fonts["Primary"], 10, 13, 8, 12, 8,
									th.Colors["ButtonText"], th.Colors["ButtonBg"], "PASTE AMOUNT", func() {
										amountLineEditor.SetText(clipboard.Get())
									})))
						})
					},
					func() {
						layout.Flex{}.Layout(gtx,
							layout.Rigid(component.Button(gtx, th, buttonSend, th.Fonts["Primary"], 14, 10, 10, 9, 10,
								th.Colors["ButtonText"], th.Colors["ButtonBg"], "SEND", func() {
									rc.Dialog.Show = true
									rc.Dialog = &model.DuoUIdialog{
										Show:       true,
										Green:      rc.DuoSend(sendStruct.passPharse, sendStruct.address, 11),
										GreenLabel: "SEND",
										CustomField: func() {
											layout.Flex{}.Layout(gtx,
												layout.Flexed(1, component.Editor(gtx, th, passLineEditor, "Enter your password",
													func(e gel.EditorEvent) {
														sendStruct.passPharse = passLineEditor.Text()
													})))
										},
										Red:      func() { rc.Dialog.Show = false },
										RedLabel: "CANCEL",
										Title:    "Are you sure?",
										Text:     "Confirm ParallelCoin send",
									}
								})))
					},
				}
				layautList.Layout(gtx, len(widgets), func(i int) {
					layout.UniformInset(unit.Dp(8)).Layout(gtx, widgets[i])
				})
			}))
		//Info("passPharse:" + sendStruct.passPharse)
	}
}
