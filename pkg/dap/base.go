// SPDX-License-Identifier: Unlicense OR MIT

package dap

import (
	"gioui.org/io/system"
	"gioui.org/layout"
	gwallet "github.com/gioapp/wallet/app"
	"github.com/gioapp/wallet/pkg/lyt"
)

func (d *dap) DAppP() error {
	for {
		select {
		case e := <-d.boot.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&d.boot.UI.Ops, e)
				d.BeforeMain()
				lyt.Format(gtx, "max(inset(0dp0dp0dp0dp,_))", d.Main())
				//d.AfterMain()
				e.Frame(gtx.Ops)
			}
			d.boot.UI.Window.Invalidate()
		}
	}
}

func (d *dap) BeforeMain() {

}

func (d *dap) AfterMain() {

}

func Tik(d *dap) func() {
	return func() {
		gw := d.boot.Apps["ParallelCoinWallet"].(gwallet.GioWallet)
		gw.GetBalances()
		gw.GetLatestTransactions()
	}
}

func (d *dap) Main() W {
	return func(gtx C) D {
		return lyt.Format(gtx, d.boot.UI.R.Mod["Main"].(string),
			func(gtx C) D {
				return d.boot.UI.N.Nav(d.boot.UI.Theme, gtx)
			},
			func(gtx C) D {
				return lyt.Format(gtx, d.boot.UI.R.Mod["Content"].(string),
					noReturn,
					noReturn,
					//Page(d.boot.UI.Theme, d.boot.UI.P[d.boot.UI.N.CurrentPage]),
					noReturn,
				)
			})
	}
}
