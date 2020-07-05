package pages

import (
	"time"

	"github.com/gioapp/wallet/gui/component"
	"github.com/gioapp/wallet/gui/rcd"

	"gioui.org/layout"

	"github.com/gioapp/wallet/pkg/gelook"
)

var (
	logOutputList = &layout.List{
		Axis:        layout.Vertical,
		ScrollToEnd: true,
	}
)

var StartupTime = time.Now()

func Logger(rc *rcd.RcVar, gtx *layout.Context, th *gelook.DuoUItheme) *gelook.DuoUIpage {
	page := gelook.DuoUIpage{
		Title:         "LOG",
		Command:       func() {},
		Border:        4,
		BorderColor:   th.Colors["Dark"],
		Header:        func() {},
		HeaderBgColor: "",
		Body:          component.DuoUIlogger(rc, gtx, th),
		BodyBgColor:   th.Colors["Dark"],
		Footer:        func() {},
		FooterBgColor: "",
	}
	return th.DuoUIpage(page)
}
