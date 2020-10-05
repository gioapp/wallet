// SPDX-License-Identifier: Unlicense OR MIT

package mod

import (
	"context"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/gioapp/wallet/pkg/dap/res"
	"github.com/gioapp/wallet/pkg/nav"
	"github.com/gioapp/wallet/pkg/theme"
)

//Duo App Platform
type Dap struct {
	BeforeMain map[int]func()
	Main       func()
	AfterMain  map[int]func()
	Ctx        context.Context
	Tik        map[int]func()
	UI         UserInterface
	S          *Settings
	Apps       map[string]*Sap
}

type Sap struct {
	Title string
	App   interface{}
	UI    *UserInterface
}

type Settings struct {
	Dir  string
	File string
}

type UserInterface struct {
	Device string
	Window *app.Window
	Theme  *theme.Theme
	//Ekran   func(gtx C) D
	FontSize float32
	R        res.Responsivity
	//P        Pages
	N   *nav.Navigation
	G   layout.Context
	Ops op.Ops
}

//type Pages map[string]Page
