package main

import (
	"fmt"
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	gwallet "github.com/gioapp/wallet/app"
	"github.com/gioapp/wallet/cfg"
	"github.com/gioapp/wallet/pkg/dap"
	"log"
	"os"
)

func main() {
	d := dap.NewDap("Duo App Plan9")
	d.NewSap(gwallet.NewGioWallet())
	if cfg.Initial {
		fmt.Println("running initial setup")
	}
	go func() {
		defer os.Exit(0)
		if err := d.DAppP(); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}
