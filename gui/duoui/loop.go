package duoui

import (
	"errors"

	"github.com/gioapp/wallet/gui/component"

	"gioui.org/io/system"

	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/gui/rcd"
	"github.com/p9c/pod/pkg/util/interrupt"
)

func DuoUImainLoop(d *model.DuoUI, r *rcd.RcVar) error {
	Debug("starting up duo ui main loop")
	ui := new(DuoUI)
	ui = &DuoUI{
		ly: d,
		rc: r,
	}
	for {
		select {
		case <-ui.rc.Ready:
			updateTrigger := make(chan struct{}, 1)
			go func() {
			quitTrigger:
				for {
					select {
					case <-updateTrigger:
						Trace("repaint forced")
						// ui.ly.Window.Invalidate()
					case <-ui.rc.Quit:
						break quitTrigger
					}
				}
			}()
			go ui.rc.ListenInit(updateTrigger)
			ui.rc.IsReady = true
			r.Boot.IsBoot = false
		case <-ui.rc.Quit:
			Debug("quit signal received")
			if !interrupt.Requested() {
				interrupt.Request()
			}
			// This case is for handling when some external application is
			//controlling the GUI and to gracefully handle the back-end
			//servers being shut down by the interrupt library receiving an
			//interrupt signal  Probably nothing needs to be run between
			//starting it and shutting down
			<-interrupt.HandlersDone
			Debug("closing GUI from interrupt/quit signal")
			return errors.New("shutdown triggered from back end")
			// TODO events of gui
		case e := <-ui.rc.Commands.Events:
			switch e := e.(type) {
			case rcd.DuoUIcommandEvent:
				ui.rc.Commands.History = append(ui.rc.Commands.History, e.Command)
				ui.ly.Window.Invalidate()
			}
		case e := <-ui.ly.Window.Events():
			ui.ly.Viewport = ui.ly.Context.Constraints.Width.Max
			switch e := e.(type) {
			case system.DestroyEvent:
				Debug("destroy event received")
				interrupt.Request()
				// Here do cleanup like are you sure (
				//optional) modal or shutting down indefinite spinner
				<-interrupt.HandlersDone
				return e.Err
			case system.FrameEvent:
				ui.ly.Context.Reset(e.Config, e.Size)
				if ui.rc.Boot.IsBoot {
					if ui.rc.Boot.IsFirstRun {
						ui.DuoUIloaderCreateWallet()
					} else {
						ui.DuoUIsplashScreen()
					}
					e.Frame(ui.ly.Context.Ops)
				} else {
					ui.DuoUImainScreen()
					if ui.rc.Dialog.Show {
						component.DuoUIdialog(ui.rc, ui.ly.Context, ui.ly.Theme)
						// ui.DuoUItoastSys()
					}
					e.Frame(ui.ly.Context.Ops)
				}
				ui.ly.Window.Invalidate()
			}
		}
	}
}
