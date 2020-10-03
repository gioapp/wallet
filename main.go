package main

import (
	"fmt"
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	"gioui.org/io/system"
	"gioui.org/layout"
	gwallet "github.com/gioapp/wallet/app"
	"github.com/gioapp/wallet/cfg"
	in "github.com/gioapp/wallet/cfg/ini"
	"log"
	"os"
	"time"
)

func main() {

	g := gwallet.NewGioWallet("bitcoin")

	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	in.Init(g.Settings.File)
	ticker(g.Tik())

	go func() {
		defer os.Exit(0)
		if err := loop(g); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(g *gwallet.GioWallet) error {
	for {
		select {
		case e := <-g.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&g.UI.Ops, e)
				g.BeforeMain()
				//if !g.API.OK {
				//g.GreskaEkran()
				//} else {
				g.AppMain(gtx)
				//}
				g.AfterMain()

				e.Frame(gtx.Ops)
			}
			g.UI.Window.Invalidate()
		}
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
