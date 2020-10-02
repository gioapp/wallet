package pages

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/pkg/gel"

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
	transactionsPanelElement = gel.NewPanel()
)

func History(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title: "HISTORY",
		//Command:       rc.GetDuoUItransactions(),
		Border:        4,
		BorderColor:   th.Colors["Light"],
		Header:        historyHeader(rc, gtx, th),
		HeaderPadding: 4,
		Body:          historyBody(rc, gtx, th),
		BodyBgColor:   th.Colors["Light"],
		BodyPadding:   0,
		//Footer:        func(gtx layout.Context)layout.Dimensions{},
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

func historyHeader(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Flex{
		//	Spacing: layout.SpaceBetween,
		//}.Layout(gtx,
		//	layout.Rigid(component.TransactionsFilter(rc, gtx, th)),
		//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//		th.DuoUIcounter(rc.GetDuoUItransactions()).Layout(gtx, rc.History.PerPage, "TxNum per page: ", fmt.Sprint(rc.History.PerPage.Value))
		//	}),
		//	layout.Rigid(func(gtx layout.Context)layout.Dimensions {
		//		th.DuoUIcounter(rc.GetDuoUItransactions()).Layout(gtx, rc.History.Page, "TxNum page: ", fmt.Sprint(rc.History.Page.Value))
		//	}),
		//)
		return layout.Dimensions{}
	}
}

func historyBody(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//layout.Flex{
			//	Axis: layout.Vertical,
			//}.Layout(gtx,
			//	layout.Rigid(
			//func(gtx layout.Context)layout.Dimensions{
			//	transactionsPanel := th.DuoUIpanel()
			//	transactionsPanel.PanelObject = rc.History.Txs.Txs
			//	transactionsPanel.ScrollBar = th.ScrollBar(16)
			//	transactionsPanelElement.PanelObjectsNumber = len(rc.History.Txs.Txs)
			//	transactionsPanel.Layout(gtx, transactionsPanelElement, func(i int, in interface{}) {
			//		txs := in.([]model.DuoUItransactionExcerpt)
			//		t := txs[i]
			//		th.DuoUIline(gtx, 0, 0, 1, th.Colors["Hint"])()
			//		for t.Link.Clicked() {
			//			rc.ShowPage = fmt.Sprintf("TRANSACTION %s", t.TxID)
			//			rc.GetSingleTx(t.TxID)()
			//			component.SetPage(rc, txPage(rc, gtx, th, t.TxID))
			//		}
			//		width := gtx.Constraints.Width.Max
			//		button := th.DuoUIbutton("", "", "", "", "", "", "", "", 0, 0, 0, 0, 0, 0, 0, 0)
			//		button.InsideLayout(gtx, t.Link, func(gtx layout.Context)layout.Dimensions{
			//			gtx.Constraints.Width.Min = width
			//			layout.Flex{
			//				Spacing: layout.SpaceBetween,
			//			}.Layout(gtx,
			//				layout.Rigid(component.TxsDetails(gtx, th, i, &t)),
			//				layout.Rigid(component.Label(gtx, th, th.Fonts["Mono"], 12, th.Colors["Secondary"], fmt.Sprintf("%0.8f", t.Amount))))
			//		})
			//	})
			//}))
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}
