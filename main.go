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
	"time"
)

func main() {
	apps := make(map[string]interface{})
	apps["ParallelCoinWallet"] = gwallet.NewGioWallet("parallelcoin")
	d := dap.NewDap(apps)

	if cfg.Initial {
		fmt.Println("running initial sync")
	}

	//ticker(d.Apps["ParallelcoinWallet"].(gwallet.GioWallet)).Tik()

	go func() {
		defer os.Exit(0)
		if err := d.DAppP(); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
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
