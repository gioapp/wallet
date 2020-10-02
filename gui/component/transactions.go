package component

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/gui/rcd"
)

//func TransactionsList(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) func(gtx layout.Context)layout.Dimensions{
//	return func(gtx layout.Context)layout.Dimensions{
//		transactionsPanel := th.DuoUIpanel()
//		transactionsPanel.PanelObject = rc.History.Txs.Txs
//		transactionsPanel.ScrollBar = th.ScrollBar()
//		transactionsPanelElement.PanelObjectsNumber = len(rc.History.Txs.Txs)
//		transactionsPanel.Layout(gtx, transactionsPanelElement, func(i int, in interface{}) {
//			txs := in.([]model.DuoUItransactionExcerpt)
//			t := txs[i]
//			th.DuoUIline(gtx, 0, 0, 1, th.Colors["Hint"])()
//			for t.Link.Clicked() {
//				rc.ShowPage = fmt.Sprintf("TRANSACTION %s", t.TxID)
//				rc.GetSingleTx(t.TxID)()
//				//SetPage(rc, txPage(rc, gtx, th, t.TxID))
//			}
//			width := gtx.Constraints.Width.Max
//			button := th.DuoUIbutton("", "", "", "", "", "", "", "", 0, 0, 0, 0, 0, 0, 0, 0)
//			button.InsideLayout(gtx, t.Link, func(gtx layout.Context)layout.Dimensions{
//				gtx.Constraints.Width.Min = width
//				layout.Flex{
//					Spacing: layout.SpaceBetween,
//				}.Layout(gtx,
//					layout.Rigid(txsDetails(gtx, th, i, &t)),
//					layout.Rigid(Label(gtx, th, th.Fonts["Mono"], 12, th.Colors["Secondary"], fmt.Sprintf("%0.8f", t.Amount))))
//			})
//		})
//	}
//}

func TxsDetails(gtx *layout.Context, th *theme.DuoUItheme, i int, t *model.DuoUItransactionExcerpt) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Flex{
		//	Axis: layout.Vertical,
		//}.Layout(gtx,
		//	layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(i))),
		//	layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.TxID)),
		//	layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprintf("%0.8f", t.Amount))),
		//	layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.Category)),
		//	layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.Time)),
		//)
		return layout.Dimensions{}
	}
}

func TransactionsFilter(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Flex{}.Layout(gtx,
		//	layout.Rigid(txsFilterItem(gtx, th, "ALL", rc.History.Categories.AllTxs)),
		//	layout.Rigid(txsFilterItem(gtx, th, "MINTED", rc.History.Categories.MintedTxs)),
		//	layout.Rigid(txsFilterItem(gtx, th, "IMATURE", rc.History.Categories.ImmatureTxs)),
		//	layout.Rigid(txsFilterItem(gtx, th, "SENT", rc.History.Categories.SentTxs)),
		//	layout.Rigid(txsFilterItem(gtx, th, "RECEIVED", rc.History.Categories.ReceivedTxs)))
		//switch c := true; c {
		//case rc.History.Categories.AllTxs.Checked(gtx):
		//	rc.History.Category = "all"
		//case rc.History.Categories.MintedTxs.Checked(gtx):
		//	rc.History.Category = "generate"
		//case rc.History.Categories.ImmatureTxs.Checked(gtx):
		//	rc.History.Category = "immature"
		//case rc.History.Categories.SentTxs.Checked(gtx):
		//	rc.History.Category = "sent"
		//case rc.History.Categories.ReceivedTxs.Checked(gtx):
		//	rc.History.Category = "received"
		//}
		return layout.Dimensions{}
	}
}

func txsFilterItem(gtx *layout.Context, th *theme.DuoUItheme, id string, c *widget.Bool) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//th.DuoUIcheckBox(id, th.Colors["Light"], th.Colors["Light"]).Layout(gtx, c)
		return layout.Dimensions{}
	}
}
