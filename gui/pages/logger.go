package pages

import (
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"
	"time"

	"github.com/gioapp/wallet/gui/rcd"

	"gioui.org/layout"
)

var (
	logOutputList = &layout.List{
		Axis:        layout.Vertical,
		ScrollToEnd: true,
	}
)

var StartupTime = time.Now()

func Logger(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme) *page.DuoUIpage {
	p := page.DuoUIpage{
		Title: "LOG",
		//Command:       func(gtx layout.Context)layout.Dimensions{},
		Border:      4,
		BorderColor: th.Colors["Dark"],
		//Header:        func(gtx layout.Context)layout.Dimensions{},
		HeaderBgColor: "",
		//Body:          component.DuoUIlogger(rc, gtx, th),
		BodyBgColor: th.Colors["Dark"],
		//Footer:        func(gtx layout.Context)layout.Dimensions{},
		FooterBgColor: "",
	}
	return page.NewDuoUIpage(th, p)
}
