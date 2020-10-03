package gwallet

import (
	"context"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/wallet/pkg/theme"
	rpcclient "github.com/p9c/pod/pkg/rpc/client"
)

var (
	selected int
	pwd      []string
)

type (
	D = layout.Dimensions
	C = layout.Context
)

type GioWallet struct {
	rpc *rpcclient.Client

	dataDir string

	ctx       context.Context
	UI        walletUI
	menuItems []Item
	Settings  walletSettings
	ItemsList []*folderListItem
	Status    Status
}

type folderListItem struct {
	Name  string
	Hash  string
	Size  uint64
	Type  uint8
	btn   *widget.Clickable
	check *widget.Bool
}

type walletUI struct {
	Device string
	Window *app.Window
	Theme  *theme.Theme
	//Ekran   func(gtx C) D
	FontSize float32

	res   walletResponsivity
	pages pages
	nav   Navigation
	Ops   op.Ops
}

type walletResponsivity struct {
	Mode string
	mod  map[string]interface{}
}

type walletSettings struct {
	Dir  string
	File string
}

type walletPage struct {
	Title  string
	Header func(gtx C) D
	Body   func(gtx C) D
}

type pages map[string]walletPage

type Navigation struct {
	Name         string
	Bg           string
	Logo         Logo
	Items        []Item
	wide         bool
	mode         string
	navLayout    string
	itemLayout   string
	itemIconSize unit.Value
	axis         layout.Axis
	size         int
	noText       bool
	logo         func(gtx C) D
}
type Item struct {
	Title string
	Bg    string
	Icon  *widget.Icon
	Btn   *widget.Clickable
	//Page *walletPage
}

type Logo struct {
	Title string
	Logo  string
}

type Status struct {
	bal *Balances
	txs []Tx
}

type Balances struct {
	available string
	pending   string
	immature  string
	total     string
}
type Tx struct {
	time    string
	address string
	amount  string
}
