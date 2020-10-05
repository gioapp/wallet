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
	boot *mod.Dap
}

func NewDap(title string) *dap {
	d := &mod.Dap{
		Ctx:  context.Background(),
		Apps: make(map[string]interface{}),
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
		app.Title(title),
	)
	//d.UI.N.Items = g.getMenuItems()

	d.UI.R.Mode = "mobile"

	n := nav.Navigation{
		Name:         "Navigacion",
		Bg:           d.UI.Theme.Colors["NavBg"],
		Items:        make(map[int]nav.Item),
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

	return &dap{boot: d}
}

func (d *dap) NewSap(title string, sap interface{}) {
	d.boot.Apps[title] = sap
	return
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
}
