package gwallet

import (
	"gioui.org/layout"
	"gioui.org/widget"
	rpcclient "github.com/p9c/pod/pkg/rpc/client"
)

var (
	selected int
	pwd      []string
)

type (
	D = layout.Dimensions
	C = layout.Context
	W = layout.Widget
)

type GioWallet struct {
	rpc    *rpcclient.Client
	Status Status
}

type folderListItem struct {
	Name  string
	Hash  string
	Size  uint64
	Type  uint8
	btn   *widget.Clickable
	check *widget.Bool
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
	Id      string
	Time    string
	Address string
	Amount  string
	Btn     *widget.Clickable
}
