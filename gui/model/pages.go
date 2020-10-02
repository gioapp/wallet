package model

import (
	"gioui.org/layout"
	"gioui.org/widget"

	"github.com/gioapp/wallet/pkg/gel"
)

type DuoUIexplorer struct {
	Page    *gel.DuoUIcounter
	PerPage *gel.DuoUIcounter
	Blocks  []DuoUIblock
	//SingleBlock btcjson.GetBlockVerboseResult
}

type DuoUIhistory struct {
	TransList  *layout.List
	Category   string
	Categories *DuoUIhistoryCategories
	Page       *gel.DuoUIcounter
	PerPage    *gel.DuoUIcounter
	Txs        *DuoUItransactionsExcerpts
	//SingleTx   btcjson.GetTransactionResult
}
type DuoUIhistoryCategories struct {
	AllTxs      *widget.Bool
	MintedTxs   *widget.Bool
	ImmatureTxs *widget.Bool
	SentTxs     *widget.Bool
	ReceivedTxs *widget.Bool
}

type DuoUInetwork struct {
	PeersList *layout.List
	Page      *gel.DuoUIcounter
	PerPage   *gel.DuoUIcounter
	//Peers     []*btcjson.GetPeerInfoResult `json:"peers"`
}
