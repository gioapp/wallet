package model

import (
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/gel/navigation"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"
	"sync"

	"gioui.org/app"
	"gioui.org/layout"
	"go.uber.org/atomic"
)

type DuOScomponent struct {
	Name       string
	Version    string
	Model      interface{}
	View       func()
	Controller func()
}

type DuoUI struct {
	Window     *app.Window
	Ops        op.Ops
	Context    layout.Context
	Theme      *theme.DuoUItheme
	Pages      *DuoUIpages
	Navigation *DuoUInav
	// Configuration *DuoUIconfiguration
	Viewport int
	IsReady  bool
}

type DuoUIpages struct {
	CurrentPage *page.DuoUIpage
	Controller  map[string]*page.DuoUIpage
	Theme       map[string]*page.DuoUIpage
}
type DuoUIlog struct {
	Mx          sync.Mutex
	LogMessages atomic.Value // []log.Entry
	//LogChan     chan log.Entry
	StopLogger chan struct{}
}

// type DuoUIconfiguration struct {
//	Abbrevation        string
//	PrimaryTextColor   color.RGBA
//	SecondaryTextColor color.RGBA
//	PrimaryBgColor     color.RGBA
//	SecondaryBgColor   color.RGBA
//	Navigations        map[string]*view.DuoUIthemeNav
// }

type DuoUIconfTabs struct {
	Current  string
	TabsList map[string]*widget.Clickable
}

// type DuoUIalert struct {
//	Time      time.Time   `json:"time"`
//	Title     string      `json:"title"`
//	Message   interface{} `json:"message"`
//	AlertType string      `json:"type"`
// }

type DuoUIsettings struct {
	Abbrevation string
	Tabs        *DuoUIconfTabs
	Daemon      *DaemonConfig `json:"daemon"`
}

type DaemonConfig struct {
	// Config  *pod.Config `json:"config"`
	Config map[string]interface{} `json:"config"`
	//Schema  pod.Schema             `json:"schema"`
	Widgets map[string]interface{}
}

type DuoUIblock struct {
	Height        int64   `json:"height"`
	BlockHash     string  `json:"hash"`
	PowAlgoID     uint32  `json:"pow"`
	Difficulty    float64 `json:"diff"`
	Amount        float64 `json:"amount"`
	TxNum         int     `json:"txnum"`
	Confirmations int64
	Time          int64 `json:"time"`
	Link          *widget.Clickable
}

type DuoUItoast struct {
	Title   string
	Message string
}

type DuoUInav struct {
	Items             map[string]*navigation.DuoUIthemeNav
	Width             int
	Height            int
	TextSize          int
	IconSize          int
	PaddingVertical   int
	PaddingHorizontal int
}
