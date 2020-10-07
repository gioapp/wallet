// SPDX-License-Identifier: Unlicense OR MIT

package dap

import (
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/gioapp/wallet/pkg/dap/box"
	"github.com/gioapp/wallet/pkg/dap/res"
	"github.com/gioapp/wallet/pkg/lyt"
	"log"
	"os"
	"time"
)

func (d *dap) DAP() {
	defer os.Exit(0)
	if err := d.DAppP(); err != nil {
		log.Fatal(err)
	}
}

func (d *dap) DAppP() error {
	for {
		select {
		case e := <-d.boot.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&d.boot.UI.Ops, e)
				d.BeforeMain(&gtx)
				lyt.Format(gtx, "max(inset(0dp0dp0dp0dp,_))", d.Main())
				//d.AfterMain()
				e.Frame(gtx.Ops)
			}
			d.boot.UI.Window.Invalidate()
		}
	}
}

func (d *dap) BeforeMain(gtx *C) {
	d.boot.UI.R = res.Resposnsivity(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
	d.boot.UI.N.Mode = d.boot.UI.R.Mode
	d.boot.UI.N.NavLayout = d.boot.UI.R.Mod["Nav"].(string)
	d.boot.UI.N.ItemLayout = d.boot.UI.R.Mod["NavIconAndLabel"].(string)
	d.boot.UI.N.Axis = d.boot.UI.R.Mod["NavItemsAxis"].(layout.Axis)
	d.boot.UI.N.Size = d.boot.UI.R.Mod["NavSize"].(int)
	d.boot.UI.N.NoContent = d.boot.UI.N.Wide
	d.boot.UI.N.LogoWidget = d.boot.UI.N.LogoLayout(d.boot.UI.Theme)
}

func (d *dap) AfterMain() {

}

func Tik(d *dap) func() {
	return func() {
		//gw := d.boot.Apps["ParallelCoinWallet"].App.(gwallet.GioWallet)
		//gw.GetBalances()
		//gw.GetLatestTransactions()
	}
}

func ticker(f func()) {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				f()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (d *dap) Main() W {
	return func(gtx C) D {
		return lyt.Format(gtx, d.boot.UI.R.Mod["Container"].(string),
			box.BoxBase(d.boot.UI.Theme.Colors["NavBg"], d.boot.UI.N.Nav(d.boot.UI.Theme, gtx)),
			func(gtx C) D {
				return lyt.Format(gtx, d.boot.UI.R.Mod["Main"].(string),
					d.boot.UI.N.CurrentPage.P(d.boot.UI.Theme, d.boot.UI.R.Mod["Page"].(string)),
					d.boot.UI.F,
				)
			})
	}
}
