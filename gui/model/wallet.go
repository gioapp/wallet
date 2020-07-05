package model

import (
	"gioui.org/op/paint"

	"github.com/gioapp/wallet/pkg/gel"
)

type DuoUIbalance struct {
	Balance string `json:"balance"`
}
type DuoUIunconfirmed struct {
	Unconfirmed string `json:"unconfirmed"`
}

type DuoUItransactionsNumber struct {
	TxsNumber int `json:"txsnumber"`
}
type DuoUItransactionsExcerpts struct {
	ModelTxsListNumber int
	TxsListNumber      int
	Txs                []DuoUItransactionExcerpt `json:"txs"`
	TxsNumber          int                       `json:"txsnumber"`
	Balance            float64                   `json:"balance"`
	BalanceHeight      float64                   `json:"balanceheight"`
}
type DuoUItransactionExcerpt struct {
	Balance       float64 `json:"balance"`
	Amount        float64 `json:"amount"`
	Category      string  `json:"category"`
	Confirmations int64   `json:"confirmations"`
	Time          string  `json:"time"`
	TxID          string  `json:"txid"`
	Comment       string  `json:"comment,omitempty"`
	Link          *gel.Button
}

type DuoUIaddress struct {
	Index   int     `json:"num"`
	Label   string  `json:"label"`
	Account string  `json:"account"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
	Copy    *gel.Button
	QrCode  *gel.Button
}
type DuoUIaddressBook struct {
	ShowMiningAddresses bool
	Num                 int            `json:"num"`
	Addresses           []DuoUIaddress `json:"addresses"`
}

type DuoUIqrCode struct {
	AddrQR  paint.ImageOp
	PubAddr string
}
