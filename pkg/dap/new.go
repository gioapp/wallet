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
	"github.com/gioapp/wallet/pkg/dap/page"
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

	d.UI = mod.UserInterface{
		Theme: theme.NewTheme(),
		//mob:   make(chan bool),
	}
	//d.UI.P = g.getPages()

	d.UI.Window = app.NewWindow(
		app.Size(unit.Dp(1024), unit.Dp(800)),
		app.Title(title),
	)
	//d.UI.N.Items = g.getMenuItems()

	d.UI.R.Mode = "mobile"

	n := &nav.Navigation{
		Name: "Navigacion",
		Bg:   d.UI.Theme.Colors["NavBg"],
		//Items:        make(map[int]nav.Item),
		ItemIconSize: unit.Px(24),
		//CurrentPage:  "Overview",
		CurrentPage: page.Page{
			Title:  "",
			Header: noReturn,
			Body:   noReturn,
			Footer: noReturn,
		},
	}
	d.UI.N = n

	s := &mod.Settings{
		Dir: appdata.Dir("dap", false),
	}
	d.S = s

	r := res.Responsivity{
		Mode: "mobile",
		Mod: map[string]interface{}{
			"Screen":          "max(inset(80dp0dp80dp0dp,_))",
			"Main":            "max(hflex(start,r(_),f(1,_)))",
			"Content":         "max(inset(80dp0dp80dp0dp,_))",
			"Nav":             "hflex(start,r(_),f(1,_))",
			"NavSize":         128,
			"NavIconSize":     128,
			"NavIconAndLabel": "hflex(r(_),r(_))",
			"Logo":            d.UI.Theme.Icons["Logo"],
			"NavItemsAxis":    layout.Horizontal,
			"TwoEqual":        "vflex(start,r(inset(0dp,_)),f(1,inset(0dp,_))))",

			"MainLayout":    "vflex(start,r(_),f(1,_))",
			"ContentLayout": "vflex(start,f(1,_),r(_))",
			"NavLayout":     "hflex(start,r(_),f(1,_))",

			"ContentBodyLayout": "vflex(start,r(inset(0dp,_)),f(1,inset(0dp,_))))",
		},
	}
	d.UI.R = r

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
