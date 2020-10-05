package main

import (
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	gwallet "github.com/gioapp/wallet/app"
	"github.com/gioapp/wallet/pkg/dap"
)

func main() {
	d := dap.NewDap("Duo App Plan9")
	d.NewSap(gwallet.NewGioWallet(d.BOOT()))
	go d.DAP()
	app.Main()
}
