package pages

import (
	"gioui.org/widget"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"gioui.org/layout"
	"github.com/gioapp/wallet/gui/rcd"
	"github.com/gioapp/wallet/pkg/gel"
)

var (
	addressBookPanelElement = gel.NewPanel()
	showMiningAddresses     = &widget.Bool{}
	buttonNewAddress        = new(widget.Clickable)
	address                 string
)

func DuoUIaddressBook(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title:       "ADDRESSBOOK",
		Command:     rc.GetAddressBook(),
		Border:      4,
		BorderColor: th.Colors["Light"],
		//Header:        addressBookHeader(rc, gtx, th, rc.GetAddressBook()),
		HeaderPadding: 4,
		Body:          addressBookBody(rc, gtx, th),
		BodyBgColor:   th.Colors["Light"],
		BodyPadding:   4,
		//Footer:        func(gtx layout.Context) layout.Dimensions{},
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

func addressBookBody(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				//layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context)layout.Dimensions{
				//	layout.Flex{
				//		Axis:    layout.Vertical,
				//		Spacing: layout.SpaceAround,
				//	}.Layout(gtx,
				//		layout.Flexed(1, addressBookContent(rc, gtx, th)))
				//})
				return layout.Dimensions{}
			}))
		return layout.Dimensions{}
	}
}

func addressBookHeader(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, pageFunc func()) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{
			Spacing:   layout.SpaceBetween,
			Axis:      layout.Horizontal,
			Alignment: layout.Middle,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				//if showMiningAddresses.Checked(gtx) {
				//	rc.AddressBook.ShowMiningAddresses = true
				// rc.GetAddressBook()()
				//} else {
				// rc.GetAddressBook()()
				//}
				//th.DuoUIcheckBox("SHOW MINING ADDRESSES", th.Colors["Light"], th.Colors["Light"]).Layout(gtx, showMiningAddresses)
				return layout.Dimensions{}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				// th.DuoUIcounter(rc.GetBlocksExcerpts()).Layout(gtx, rc.Explorer.Page, "PAGE", fmt.Sprint(rc.Explorer.Page.Value))
				return layout.Dimensions{}
			}))
		// layout.Rigid(component.Button(gtx, th, buttonNewAddress, th.Fonts["Secondary"], 12, th.Colors["ButtonText"], th.Colors["Dark"], "NEW ADDRESS", component.QrDialog(rc, gtx, rc.CreateNewAddress("")))))
		//layout.Rigid(component.MonoButton(gtx, th, buttonNewAddress, 12, "Primary", "Light", "Secondary", "NEW ADDRESS", func(gtx layout.Context) layout.Dimensions {
		//	rc.Dialog.Show = true
		//	rc.Dialog = &model.DuoUIdialog{
		//		Show: true,
		//Orange: func(gtx layout.Context)layout.Dimensions{
		//	rc.Dialog.Show = false
		//},
		//CustomField: component.DuoUIqrCode(gtx, address, 256),
		//Title: "Copy address",
		//Text:        rc.CreateNewAddress(""),
		//}
		//pageFunc()
		//return layout.Dimensions{}
		//})))
		return layout.Dimensions{}
	}
}

func addressBookContent(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		addressBookPanelElement.PanelObject = rc.AddressBook.Addresses
		addressBookPanelElement.PanelObjectsNumber = len(rc.AddressBook.Addresses)
		//addressBookPanel := panel.NewPanel()
		//addressBookPanel.ScrollBar = th.ScrollBar(16)
		//addressBookPanel.Layout(gtx, addressBookPanelElement, func(i int, in interface{}) {
		//if in != nil {
		//addresses := in.([]model.DuoUIaddress)
		//t := rc.AddressBook.Addresses[i]
		//layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//		layout.Flex{
		//			Alignment: layout.Middle,
		//		}.Layout(gtx,
		//			layout.Flexed(0.2, component.Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.Index))),
		//			layout.Flexed(0.2, component.Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.Account)),
		//			layout.Rigid(component.MonoButton(gtx, th, t.Copy, 12, "", "", "Mono", t.Address, func(gtx layout.Context)layout.Dimensions{ clipboard.Set(t.Address) })),
		//			layout.Flexed(0.4, component.Label(gtx, th, th.Fonts["Primary"], 14, th.Colors["Dark"], t.Label)),
		//			layout.Flexed(0.2, component.Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.Amount))),
		//			layout.Rigid(component.MonoButton(gtx, th, t.QrCode, 12, "", "", "Secondary", "QR", component.QrDialog(rc, gtx, t.Address))),
		//		)
		//		return layout.Dimensions{}
		//	}),
		//	layout.Rigid(th.DuoUIline(gtx, 1, 0, 1, th.Colors["Gray"])),
		//)
		//}
		//})
		// }).Layout(gtx, addressBookPanel)
		return layout.Dimensions{}
	}
}
