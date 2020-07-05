package main

import (
	"fmt"
	"gioui.org/app"
	"git.parallelcoin.io/dev/9/pkg/util/interrupt"
	"github.com/gioapp/wallet/cfg"
	"github.com/gioapp/wallet/cfg/ini"
	"github.com/gioapp/wallet/gui"
	"github.com/gioapp/wallet/gui/duoui"
	"github.com/gioapp/wallet/gui/rcd"
	"time"
)

func main() {
	//coins := c.Coins{}
	in.Init()
	if cfg.Initial {
		fmt.Println("running initial sync")
		time.Sleep(time.Second * 2)
	}

	go func() {
		rc := rcd.RcInit()
		//if !apputil.FileExists(*cx.Config.WalletFile) {
		//	rc.Boot.IsFirstRun = true
		//}
		duo, err := duoui.DuOuI(rc)
		rc.DuoUIloggerController()
		interrupt.AddHandler(func() {
			//Debug("guiHandle interrupt")
			close(rc.Quit)
		})
		//Info("IsFirstRun? ", rc.Boot.IsFirstRun)
		// signal the GUI that the back end is ready
		//Debug("sending ready signal")
		// we can do this without blocking because the channel has 1 buffer this way it falls immediately the GUI starts
		if !rc.Boot.IsFirstRun {
			go rc.StartServices()
		}
		// Start up GUI
		//Debug("starting up GUI")
		err = gui.WalletGUI(duo, rc)
		if err != nil {
			//Error(err)
		}

	}()
	app.Main()

	// exp.SrcNode().GetAddrs()

}
