package pages

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/gui/rcd"
)

func txPage(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, tx string) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title:       "BLOCK",
		TxColor:     "",
		Command:     rc.GetSingleTx(tx),
		Border:      4,
		BorderColor: th.Colors["Light"],
		//Header:        func(gtx layout.Context)layout.Dimensions{},
		HeaderBgColor: "",
		HeaderPadding: 0,
		//Body:          singleTxBody(rc, gtx, th, rc.History.SingleTx),
		BodyBgColor: th.Colors["Light"],
		BodyPadding: 0,
		//Footer:        func(gtx layout.Context)layout.Dimensions{},
		FooterBgColor: "",
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

//func singleTxBody(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, tx btcjson.GetTransactionResult) func(gtx layout.Context)layout.Dimensions{
//	return func(gtx layout.Context)layout.Dimensions{

//duo := layout.Horizontal
//if gtx.Constraints.Width.Max < 1280 {
//	duo = layout.Vertical
//}
//trio := layout.Horizontal
//if gtx.Constraints.Width.Max < 780 {
//	trio = layout.Vertical
//}

//blockJSON, _ := json.MarshalIndent(block, "", "  ")
//blockText := string(blockJSON)
//widgets := []func(){
//
//	func(gtx layout.Context)layout.Dimensions{
//		th.H6(tx.TxID).Layout(gtx)
//	},
//component.UnoField(gtx, component.ContentLabeledField(gtx, th, layout.Vertical, 4, 12, 14, "Hash", "Dark", "LightGrayII", "Dark", "LightGrayI", fmt.Sprint(block.Hash))),
//component.DuoFields(gtx, duo,
//	component.TrioFields(gtx, th, trio, 12, 16,
//		"Height", fmt.Sprint(block.Height), "Dark", "LightGrayII", "Dark", "LightGrayI",
//		"Confirmations", fmt.Sprint(block.Confirmations), "Dark", "LightGrayII", "Dark", "LightGrayI",
//		"Time", fmt.Sprint(time.Unix(block.Time, 0).Format("2006-01-02 15:04:05")), "LightGrayII", "Dark", "Dark", "LightGrayI",
//	),
//	component.TrioFields(gtx, th, trio, 12, 16,
//		"PowAlgo", fmt.Sprint(block.PowAlgo), algoHeadColor, algoHeadBgColor, algoColor, algoBgColor,
//		"Difficulty", fmt.Sprint(block.Difficulty), "Dark", "LightGrayII", "Dark", "LightGrayI",
//		"Nonce", fmt.Sprint(block.Nonce), "LightGrayII", "Dark", "Dark", "LightGrayI",
//	),
//),
//component.DuoFields(gtx, duo,
//	component.ContentLabeledField(gtx, th, layout.Vertical, 4, 12, 12, "MerkleRoot", "Dark", "LightGrayII", "Dark", "LightGrayI", block.MerkleRoot),
//	component.ContentLabeledField(gtx, th, layout.Vertical, 4, 12, 12, "PowHash", "Dark", "LightGrayII", "Dark", "LightGrayI", fmt.Sprint(block.PowHash)),
//),
//component.DuoFields(gtx, duo,
//	component.TrioFields(gtx, th, trio, 12, 16,
//		"Size", fmt.Sprint(block.Size), "Dark", "LightGrayII", "Dark", "LightGrayI",
//		"Weight", fmt.Sprint(block.Weight), "Dark", "LightGrayII", "Dark", "LightGrayI",
//		"Bits", fmt.Sprint(block.Bits), "LightGrayII", "Dark", "Dark", "LightGrayI",
//	),
//	component.TrioFields(gtx, th, trio, 12, 16,
//		"TxNum", fmt.Sprint(block.TxNum), "Dark", "LightGrayII", "Dark", "LightGrayI",
//		"StrippedSize", fmt.Sprint(block.StrippedSize), "Dark", "LightGrayII", "Dark", "LightGrayI",
//		"Version", fmt.Sprint(block.Version), "LightGrayII", "Dark", "Dark", "LightGrayI",
//	),
//),
//component.UnoField(gtx, component.ContentLabeledField(gtx, th, layout.Vertical, 4, 12, 12, "Tx", "Dark", "LightGrayII", "Dark", "LightGrayI", fmt.Sprint(block.Tx))),
//component.UnoField(gtx, component.ContentLabeledField(gtx, th, layout.Vertical, 4, 12, 12, "RawTx", "Dark", "LightGrayII", "Dark", "LightGrayI", fmt.Sprint(blockText))),
//component.PageNavButtons(rc, gtx, th, block.PreviousHash, block.NextHash, blockPage(rc, gtx, th, block.PreviousHash), blockPage(rc, gtx, th, block.NextHash)),
//}
//layautList.Layout(gtx, len(widgets), func(i int) {
//	layout.UniformInset(unit.Dp(0)).Layout(gtx, widgets[i])
//})
//}
//}
