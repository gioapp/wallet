package pages

import (
	"gioui.org/layout"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"github.com/gioapp/wallet/gui/rcd"
)

func LoadPages(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) (p map[string]*page.DuoUIpage) {
	p = make(map[string]*page.DuoUIpage)
	p["OVERVIEW"] = Overview(rc, gtx, th)
	p["SEND"] = Send(rc, gtx, th)
	// p["RECEIVE"] = th.DuoUIpage("RECEIVE", 10, func(gtx layout.Context)layout.Dimensions{}, func(gtx layout.Context)layout.Dimensions{}, func(gtx layout.Context)layout.Dimensions{ th.H5("receive :").Layout(gtx) }, func(gtx layout.Context)layout.Dimensions{})
	p["ADDRESSBOOK"] = DuoUIaddressBook(rc, gtx, th)
	p["SETTINGS"] = Settings(rc, gtx, th)
	p["NETWORK"] = Network(rc, gtx, th)
	// p["BLOCK"] = th.DuoUIpage("BLOCK", 0, func(gtx layout.Context)layout.Dimensions{}, func(gtx layout.Context)layout.Dimensions{}, func(gtx layout.Context)layout.Dimensions{ th.H5("block :").Layout(gtx) }, func(gtx layout.Context)layout.Dimensions{})
	p["HISTORY"] = History(rc, gtx, th)
	p["EXPLORER"] = DuoUIexplorer(rc, gtx, th)
	p["MINER"] = Miner(rc, gtx, th)
	p["CONSOLE"] = Console(rc, gtx, th)
	p["LOG"] = Logger(rc, gtx, th)
	return
}
