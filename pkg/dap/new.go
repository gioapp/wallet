package dap

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/wallet/cfg"
	"github.com/gioapp/wallet/pkg/appdata"
	"github.com/gioapp/wallet/pkg/dap/mod"
	"github.com/gioapp/wallet/pkg/dap/res"
	"github.com/gioapp/wallet/pkg/nav"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	noReturn = func(gtx C) D { return D{} }
)

type (
	D = layout.Dimensions
	C = layout.Context
	W = layout.Widget
)
type dap struct {
	boot mod.Dap
}

func NewDap(title string) dap {
	if cfg.Initial {
		fmt.Println("running initial setup")
	}
	d := mod.Dap{
		Ctx:  context.Background(),
		Apps: make(map[string]mod.Sap),
	}

	d.UI = &mod.UserInterface{
		Theme: theme.NewTheme(),
		//mob:   make(chan bool),
	}

	d.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1024), unit.Dp(800)),
		app.Title(title),
	)

	n := &nav.Navigation{
		Name:         "Navigacion",
		Bg:           d.UI.Theme.Colors["NavBg"],
		ItemIconSize: unit.Px(24),
	}
	d.UI.N = n

	s := &mod.Settings{
		Dir: appdata.Dir("dap", false),
	}
	d.S = s

	d.UI.R = res.Resposnsivity(0, 0)

	return dap{boot: d}
}

func (d *dap) NewSap(s mod.Sap) {
	d.boot.Apps[s.Title] = s
	return
}
func (d *dap) BOOT() *mod.Dap {
	return &d.boot
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
