package rcd

import (
	"gioui.org/widget"
	"github.com/gioapp/gel/page"
	"time"

	"gioui.org/layout"
	"gioui.org/text"

	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/pkg/gel"
)

type RcVar struct {
	//cx             *conte.Xt
	db             *DuoUIdb
	Boot           *Boot
	Events         chan Event
	UpdateTrigger  chan struct{}
	Status         *model.DuoUIstatus
	Dialog         *model.DuoUIdialog
	Log            *model.DuoUIlog
	ConsoleHistory *model.DuoUIconsoleHistory

	Commands *DuoUIcommands

	Settings  *model.DuoUIsettings
	Sent      bool
	Toasts    []model.DuoUItoast
	Localhost model.DuoUIlocalHost
	Uptime    int

	AddressBook *model.DuoUIaddressBook
	ShowPage    string
	CurrentPage *page.DuoUIpage
	// NodeChan   chan *rpc.Server
	// WalletChan chan *wallet.Wallet
	Explorer *model.DuoUIexplorer
	History  *model.DuoUIhistory
	Network  *model.DuoUInetwork
	Quit     chan struct{}
	Ready    chan struct{}
	IsReady  bool
}

type Boot struct {
	IsBoot     bool   `json:"boot"`
	IsFirstRun bool   `json:"firstrun"`
	IsBootMenu bool   `json:"menu"`
	IsBootLogo bool   `json:"logo"`
	IsLoading  bool   `json:"loading"`
	IsScreen   string `json:"screen"`
}

// type rcVar interface {
//	GetDuoUItransactions(sfrom, count int, cat string)
//	GetDuoUIbalance()
//	GetDuoUItransactionsExcerpts()
//	DuoSend(wp string, ad string, am float64)
//	GetDuoUItatus()
//	PushDuoUIalert(t string, m interface{}, at string)
//	GetDuoUIblockHeight()
//	GetDuoUIblockCount()
//	GetDuoUIconnectionCount()
// }

func RcInit() (r *RcVar) {
	b := Boot{
		IsBoot:     true,
		IsFirstRun: false,
		IsBootMenu: false,
		IsBootLogo: false,
		IsLoading:  false,
		IsScreen:   "",
	}
	// d := models.DuoUIdialog{
	//	Show:   true,
	//	Ok:     func(gtx layout.Context)layout.Dimensions{ r.Dialog.Show = false },
	//	Cancel: func(gtx layout.Context)layout.Dimensions{ r.Dialog.Show = false },
	//	Title:  "Dialog!",
	//	Text:   "Dialog text",
	// }
	l := new(model.DuoUIlog)

	r = &RcVar{
		//cx:          cx,
		db:          new(DuoUIdb),
		Boot:        &b,
		AddressBook: new(model.DuoUIaddressBook),
		Status: &model.DuoUIstatus{
			Node: &model.NodeStatus{},
			Wallet: &model.WalletStatus{
				//WalletVersion: make(map[string]btcjson.VersionResult),
				LastTxs: &model.DuoUItransactionsExcerpts{},
			},
			Kopach: &model.KopachStatus{},
		},
		Dialog: &model.DuoUIdialog{},
		//Settings: settings(cx),
		Log:      l,
		Commands: new(DuoUIcommands),
		ConsoleHistory: &model.DuoUIconsoleHistory{
			Commands: []model.DuoUIconsoleCommand{
				{
					ComID:    "input",
					Category: "input",
					Time:     time.Now(),

					// Out: input(duo),
				},
			},
			CommandsNumber: 1,
		},
		Sent:      false,
		Localhost: model.DuoUIlocalHost{},
		ShowPage:  "OVERVIEW",
		Explorer: &model.DuoUIexplorer{
			PerPage: &gel.DuoUIcounter{
				Value:        20,
				OperateValue: 1,
				From:         0,
				To:           50,
				CounterInput: &widget.Editor{
					Alignment:  text.Middle,
					SingleLine: true,
				},
				CounterIncrease: new(widget.Clickable),
				CounterDecrease: new(widget.Clickable),
				CounterReset:    new(widget.Clickable),
			},
			Page: &gel.DuoUIcounter{
				Value:        0,
				OperateValue: 1,
				From:         0,
				To:           50,
				CounterInput: &widget.Editor{
					Alignment:  text.Middle,
					SingleLine: true,
				},
				CounterIncrease: new(widget.Clickable),
				CounterDecrease: new(widget.Clickable),
				CounterReset:    new(widget.Clickable),
			},
			Blocks: []model.DuoUIblock{},
			//SingleBlock: btcjson.GetBlockVerboseResult{},
		},

		Network: &model.DuoUInetwork{
			PerPage: &gel.DuoUIcounter{
				Value:        20,
				OperateValue: 1,
				From:         0,
				To:           50,
				CounterInput: &widget.Editor{
					Alignment:  text.Middle,
					SingleLine: true,
				},
				CounterIncrease: new(widget.Clickable),
				CounterDecrease: new(widget.Clickable),
				CounterReset:    new(widget.Clickable),
			},
			Page: &gel.DuoUIcounter{
				Value:        0,
				OperateValue: 1,
				From:         0,
				To:           50,
				CounterInput: &widget.Editor{
					Alignment:  text.Middle,
					SingleLine: true,
				},
				CounterIncrease: new(widget.Clickable),
				CounterDecrease: new(widget.Clickable),
				CounterReset:    new(widget.Clickable),
			},
			PeersList: &layout.List{
				Axis: layout.Vertical,
			},
			// Peers:     []*btcjson.GetPeerInfoResult
		},

		History: &model.DuoUIhistory{
			PerPage: &gel.DuoUIcounter{
				Value:        20,
				OperateValue: 1,
				From:         1,
				To:           50,
				CounterInput: &widget.Editor{
					Alignment:  text.Middle,
					SingleLine: true,
				},
				CounterIncrease: new(widget.Clickable),
				CounterDecrease: new(widget.Clickable),
				CounterReset:    new(widget.Clickable),
			},
			Page: &gel.DuoUIcounter{
				Value:        0,
				OperateValue: 1,
				From:         0,
				To:           50,
				CounterInput: &widget.Editor{
					Alignment:  text.Middle,
					SingleLine: true,
				},
				CounterIncrease: new(widget.Clickable),
				CounterDecrease: new(widget.Clickable),
				CounterReset:    new(widget.Clickable),
			},
			TransList: &layout.List{
				Axis: layout.Vertical,
			},
			Categories: &model.DuoUIhistoryCategories{
				AllTxs:      new(widget.Bool),
				MintedTxs:   new(widget.Bool),
				ImmatureTxs: new(widget.Bool),
				SentTxs:     new(widget.Bool),
				ReceivedTxs: new(widget.Bool),
			},
			Txs: &model.DuoUItransactionsExcerpts{
				ModelTxsListNumber: 0,
				TxsListNumber:      0,
				Txs:                []model.DuoUItransactionExcerpt{},
				TxsNumber:          0,
				Balance:            0,
				BalanceHeight:      0,
			},
			//SingleTx: btcjson.GetTransactionResult{},
		},
		Quit:  make(chan struct{}),
		Ready: make(chan struct{}),
	}
	//r.db.DuoUIdbInit(r.cx.DataDir)
	return
}
