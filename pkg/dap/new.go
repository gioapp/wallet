package dap

import (
	"context"
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/appdata"
	"github.com/gioapp/wallet/pkg/dap/mod"
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
	boot *mod.Dap
}

func NewDap(apps map[string]interface{}) *dap {
	d := &mod.Dap{
		Ctx:  context.Background(),
		Apps: apps,
	}

	d.UI = &mod.UserInterface{
		Theme: theme.NewTheme(),
		//mob:   make(chan bool),
	}
	//d.UI.P = g.getPages()
	d.UI.Theme.T.Color.Primary = helper.HexARGB(d.UI.Theme.Colors["Primary"])
	d.UI.Theme.T.Color.Text = helper.HexARGB(d.UI.Theme.Colors["Charcoal"])
	d.UI.Theme.T.Color.Hint = helper.HexARGB(d.UI.Theme.Colors["Silver"])
	d.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1024), unit.Dp(800)),
		app.Title("ParallelCoin - DUO - True Story"),
	)
	//d.UI.N.Items = g.getMenuItems()

	d.UI.R.Mode = "mobile"

	n := nav.Navigation{
		Name: "Navigacion",
		Bg:   d.UI.Theme.Colors["NavBg"],
		//Items:        g.menuItems,
		ItemIconSize: unit.Px(24),
		CurrentPage:  "Overview",
	}
	d.UI.N = n

	s := &mod.Settings{
		Dir: appdata.Dir("dap", false),
	}
	d.S = s

	return &dap{boot: d}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
