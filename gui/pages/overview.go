package pages

import (
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/component"
	"github.com/gioapp/wallet/gui/rcd"
)

func Overview(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title:  "OVERVIEW",
		Border: 0,
		//Command:       func(gtx layout.Context)layout.Dimensions{},
		//Header:        func(gtx layout.Context)layout.Dimensions{},
		HeaderBgColor: "",
		HeaderPadding: 0,
		Body:          overviewBody(rc, gtx, th),
		BodyBgColor:   th.Colors["Light"],
		BodyPadding:   0,
		//Footer:        func(gtx layout.Context)layout.Dimensions{},
		FooterBgColor: "",
		FooterPadding: 0,
	}
	return page.NewDuoUIpage(th, p)
}

func overviewBody(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		viewport := layout.Flex{Axis: layout.Horizontal}
		//if gtx.Constraints.Width.Max < 780 {
		//	viewport = layout.Flex{Axis: layout.Vertical}
		//}
		viewport.Layout(gtx,
			//layout.Flexed(0.5, component.DuoUIstatus(rc, gtx, th)),
			layout.Flexed(0.5, component.DuoUIlatestTransactions(rc, gtx, th)),
		)
		op.InvalidateOp{}.Add(gtx.Ops)
		return layout.Dimensions{}
	}
}
