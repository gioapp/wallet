package component

import (
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/panel"
	"github.com/gioapp/gel/theme"
	"github.com/gioapp/wallet/pkg/gel"

	"gioui.org/layout"
	"gioui.org/unit"

	"github.com/gioapp/wallet/gui/rcd"
)

var (
	latestTxsPanelElement = gel.NewPanel()
)

func DuoUIlatestTransactions(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//width := gtx.Constraints.Width.Max
		//cs := gtx.Constraints
		container.DuoUIcontainer(th, 0, th.Colors["DarkGray"]).Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					container.DuoUIcontainer(th, 8, th.Colors["Primary"]).Layout(gtx, layout.N, func(gtx layout.Context) layout.Dimensions {
						//gtx.Constraints.Width.Min = width
						//latestx := th.H5("LATEST TRANSACTIONS")
						//latestx.Color = th.Colors["Light"]
						//latestx.Alignment = text.Start
						//latestx.Layout(gtx)
						return layout.Dimensions{}
					})
					return layout.Dimensions{}
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {

								latestTxsBookPanel := panel.NewPanel()
								latestTxsBookPanel.PanelObject = rc.Status.Wallet.LastTxs.Txs
								//latestTxsBookPanel.ScrollBar = th.ScrollBar(0)
								latestTxsPanelElement.PanelObjectsNumber = len(rc.Status.Wallet.LastTxs.Txs)
								//latestTxsBookPanel.Layout(gtx, latestTxsPanelElement, func(i int, in interface{}) {
								//	//txs := in.([]model.DuoUItransactionExcerpt)
								//	//t := txs[i]
								//	container.DuoUIcontainer(16, th.Colors["Dark"]).Layout(gtx, layout.NW, func(gtx layout.Context) layout.Dimensions {
								//		//width := gtx.Constraints.Width.Max
								//		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								//			//layout.Rigid(lTtxid(gtx, th, t.TxID)),
								//			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								//				//gtx.Constraints.Width.Min = width
								//				layout.Flex{
								//					Spacing: layout.SpaceBetween,
								//				}.Layout(gtx,
								//					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								//						//layout.Flex{
								//						//	Axis: layout.Vertical,
								//						//}.Layout(gtx,
								//						//	layout.Rigid(lTcategory(gtx, th, t.Category)),
								//						//	layout.Rigid(lTtime(gtx, th, t.Time)),
								//						//)
								//						return layout.Dimensions{}
								//					}),
								//					//layout.Rigid(lTamount(gtx, th, t.Amount)),
								//				)
								//				return layout.Dimensions{}
								//			}),
								//			//layout.Rigid(th.DuoUIline(gtx, 0, 0, 1, th.Colors["Hint"])),
								//		)
								//		return layout.Dimensions{}
								//	})
								//})
								return layout.Dimensions{}
							}))
						return layout.Dimensions{}
					})
					return layout.Dimensions{}
				}),
			)
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}

func lTtxid(gtx *layout.Context, th *theme.DuoUItheme, v string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//tim := th.Caption(v)
		//tim.Font.Typeface = th.Fonts["Mono"]
		//tim.Color = th.Colors["Light"]
		//tim.Layout(gtx)
		return layout.Dimensions{}
	}
}

func lTcategory(gtx *layout.Context, th *theme.DuoUItheme, v string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//sat := th.Body1(v)
		//sat.Color = th.Colors["Light"]
		//sat.Font.Typeface = "bariol"
		//sat.Layout(gtx)
		return layout.Dimensions{}
	}
}

func lTtime(gtx *layout.Context, th *theme.DuoUItheme, v string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//l := th.Body1(v)
		//l.Font.Typeface = "bariol"
		//l.Color = th.Colors["Light"]
		//l.Color = th.Colors["Hint"]
		//l.Layout(gtx)
		return layout.Dimensions{}
	}
}

func lTamount(gtx *layout.Context, th *theme.DuoUItheme, v float64) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//sat := th.Body1(fmt.Sprintf("%0.8f", v))
			//sat.Font.Typeface = "bariol"
			//sat.Color = th.Colors["Light"]
			//sat.Layout(gtx)
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}
