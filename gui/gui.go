package gui

import (
	"github.com/gioapp/wallet/gui/duoui"
	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/gui/rcd"
)

func WalletGUI(duo *model.DuoUI, rc *rcd.RcVar) (err error) {
	go func() {
		//Debug("starting UI main loop")
		if rc.IsReady != false {
		}
		if err := duoui.DuoUImainLoop(duo, rc); err != nil {
			//Fatal(err.Error(), "- shutting down")
		}
	}()
	//Debug("starting up gio app main")
	//Debug("GUI shut down")
	return
}
