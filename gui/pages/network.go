package pages

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/gui/rcd"
)

var (
// itemValue = &gel.DuoUIcounter{
//	Value:        11,
//	OperateValue: 1,
//	From:         0,
//	To:           15,
//	CounterInput: &widget.Editor{
//		Alignment:  text.Middle,
//		SingleLine: true,
//	},
//	CounterIncrease: new(widget.Clickable),
//	//CounterDecrease: new(controller.Button),
//	CounterReset: new(widget.Clickable),
// }
)

func Network(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title:         "NETWORK",
		TxColor:       "",
		Command:       rc.GetPeerInfo(),
		Border:        4,
		BorderColor:   th.Colors["Light"],
		Header:        networkHeader(rc, gtx, th),
		HeaderBgColor: "",
		HeaderPadding: 0,
		Body:          networkBody(rc, gtx, th),
		BodyBgColor:   th.Colors["Light"],
		BodyPadding:   0,
		//Footer:        func(gtx layout.Context)layout.Dimensions{},
		FooterBgColor: "",
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

func networkHeader(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Flex{
		//	Spacing: layout.SpaceBetween,
		//}.Layout(gtx,
		//	// layout.Rigid(component.TransactionsFilter(rc, gtx, th)),
		//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//		th.DuoUIcounter(rc.GetPeerInfo()).Layout(gtx, rc.Network.PerPage, "Peers per page: ", fmt.Sprint(rc.Network.PerPage.Value))
		//	}),
		//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//		th.DuoUIcounter(rc.GetPeerInfo()).Layout(gtx, rc.Network.Page, "Peers page: ", fmt.Sprint(rc.Network.Page.Value))
		//	}),
		//)
		return layout.Dimensions{}
	}
}

func networkBody(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context)layout.Dimensions{
		//	layout.Flex{
		//		Axis: layout.Vertical,
		//	}.Layout(gtx,
		//		layout.Rigid(component.PeersList(rc, gtx, th)))
		//})
		return layout.Dimensions{}
	}
}
