package rcd

import (
	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/pkg/gel"
	"github.com/p9c/pod/pkg/rpc/btcjson"
	"github.com/p9c/pod/pkg/rpc/chainrpc"
)

func (r *RcVar) GetSingleBlock(hash string) func() {
	return func() {
		r.Explorer.SingleBlock = r.GetBlock(hash)
	}
}

func (r *RcVar) GetBlock(hash string) btcjson.GetBlockVerboseResult {
	verbose, verbosetx := true, true
	bcmd := btcjson.GetBlockCmd{
		Hash:      hash,
		Verbose:   &verbose,
		VerboseTx: &verbosetx,
	}
	bl, err := chainrpc.HandleGetBlock(r.cx.RPCServer, &bcmd, nil)
	if err != nil {
		// dv.PushDuoVUEalert("BTCJSONError", err.BTCJSONError(), "error")
	}
	gbvr, ok := bl.(btcjson.GetBlockVerboseResult)
	if ok {
		return gbvr
	}
	return btcjson.GetBlockVerboseResult{}
}

func (r *RcVar) GetNetworkLastBlock() (out int32) {
	for _, g := range r.cx.RPCServer.Cfg.ConnMgr.ConnectedPeers() {
		out := g.ToPeer().StatsSnapshot().LastBlock
		if out > r.Status.Node.NetworkLastBlock.Load() {
			r.Status.Node.NetworkLastBlock.Store(out)
		}
	}
	return
}

func (r *RcVar) GetBlockExcerpt(height int) (b model.DuoUIblock) {
	b = *new(model.DuoUIblock)
	hashHeight, err := r.cx.RPCServer.Cfg.Chain.BlockHashByHeight(int32(height))
	if err != nil {
		Error("Block Hash By Height:", err)
	}

	verbose, verbosetx := true, true
	bcmd := btcjson.GetBlockCmd{
		Hash:      hashHeight.String(),
		Verbose:   &verbose,
		VerboseTx: &verbosetx,
	}
	bl, err := chainrpc.HandleGetBlock(r.cx.RPCServer, &bcmd, nil)
	if err != nil {
		// dv.PushDuoVUEalert("BTCJSONError", err.BTCJSONError(), "error")
	}
	block := bl.(btcjson.GetBlockVerboseResult)
	b.Height = block.Height
	b.BlockHash = block.Hash
	b.Confirmations = block.Confirmations
	b.TxNum = block.TxNum

	// t := time.Unix(0, block.Time)
	// b.Time = t.Format("02/01/2006, 15:04:05")
	b.Time = block.Time

	b.Link = &gel.Button{}
	return
}

func (r *RcVar) GetBlocksExcerpts() func() {
	return func() {
		r.Explorer.Page.To = int(r.Status.Node.BlockCount.Load()) / r.Explorer.PerPage.Value
		startBlock := r.Explorer.Page.Value * r.Explorer.PerPage.Value
		endBlock := r.Explorer.Page.Value*r.Explorer.PerPage.Value + r.Explorer.PerPage.Value
		height := int(r.cx.RPCServer.Cfg.Chain.BestSnapshot().Height)
		Debug("GetBlocksExcerpts", startBlock, endBlock, height)
		if endBlock > height {
			endBlock = height
		}
		blocks := *new([]model.DuoUIblock)
		for i := startBlock; i < endBlock; i++ {
			blocks = append(blocks, r.GetBlockExcerpt(i))
			// Info("trazo")
			// Info(r.Status.Node.BlockHeight)
		}
		r.Explorer.Blocks = blocks
		return
	}
}

func (r *RcVar) GetBlockCount() {
	getBlockCount, err := chainrpc.HandleGetBlockCount(r.cx.RPCServer, nil, nil)
	if err != nil {
		// dv.PushDuoVUEalert("BTCJSONError", err.BTCJSONError(), "error")
	}
	r.Status.Node.BlockCount.Store(uint64(getBlockCount.(int64)))
	return
}
func (r *RcVar) GetBlockHash(blockHeight int) string {
	hcmd := btcjson.GetBlockHashCmd{
		Index: int64(blockHeight),
	}
	hash, err := chainrpc.HandleGetBlockHash(r.cx.RPCServer, &hcmd, nil)
	if err != nil {
		// dv.PushDuoVUEalert("BTCJSONError", err.BTCJSONError(), "error")
	}
	return hash.(string)
}

func (r *RcVar) GetConnectionCount() {
	r.Status.Node.ConnectionCount.Store(r.cx.RealNode.ConnectedCount())
	return
}

func (r *RcVar) GetDifficulty() {
	c := btcjson.GetDifficultyCmd{}
	diff, err := chainrpc.HandleGetDifficulty(r.cx.RPCServer, c, nil)
	if err != nil {
		// dv.PushDuoVUEalert("BTCJSONError", err.BTCJSONError(), "error")
	}
	r.Status.Node.Difficulty.Store(diff.(float64))
	return
}

// func (v *DuoVUEnode) Gethashespersec() {
// 	r, err := v.r.cx.RPCServer.HandleGetHashesPerSec(v.r.cx.RPCServer, a, nil)
// 	r = int64(0)
// 	return
// }
// func (v *DuoVUEnode) Getheaders(a *btcjson.GetHeadersCmd) {
// 	r, err := v.r.cx.RPCServer.HandleGetHeaders(v.r.cx.RPCServer, a, nil)
// 	r = []string{}
// 	return
// }
// func (v *DuoVUEnode) Getinfo() {
// 	r, err := v.r.cx.RPCServer.HandleGetInfo(v.r.cx.RPCServer, a, nil)
// 	r = btcjson.InfoChainResult{}
// 	return
// }
// func (v *DuoVUEnode) Getmempoolinfo() {
// 	r, err := v.r.cx.RPCServer.HandleGetMempoolInfo(v.r.cx.RPCServer, a, nil)
// 	r = btcjson.GetMempoolInfoResult{}
// 	return
// }
// func (v *DuoVUEnode) Getmininginfo() {
// 	r, err := v.r.cx.RPCServer.HandleGetMiningInfo(v.r.cx.RPCServer, a, nil)
// 	r = btcjson.GetMiningInfoResult{}
// 	return
// }
// func (v *DuoVUEnode) Getnettotals() {
// 	r, err := v.r.cx.RPCServer.HandleGetNetTotals(v.r.cx.RPCServer, a, nil)
// 	r = btcjson.GetNetTotalsResult{}
// 	return
// }
// func (v *DuoVUEnode) Getnetworkhashps(a *btcjson.GetNetworkHashPSCmd) {
// 	r, err := v.r.cx.RPCServer.HandleGetNetworkHashPS(v.r.cx.RPCServer, a, nil)
// 	r = int64(0)
// 	return
// }

// func (v *DuoVUEnode) Stop() {
// 	r, err := v.r.cx.RPCServer.HandleStop(v.r.cx.RPCServer, a, nil)
// 	r = ""
// 	return
// }
func (r *RcVar) GetUptime() {
	rRaw, err := chainrpc.HandleUptime(r.cx.RPCServer, nil, nil)
	if err != nil {
	}
	// rRaw = int64(0)
	r.Uptime = rRaw.(int)
	return
}

// func (v *DuoVUEnode) Validateaddress(a *btcjson.ValidateAddressCmd) {
// 	r, err := v.r.cx.RPCServer.HandleValidateAddress(v.r.cx.RPCServer, a, nil)
// 	r = btcjson.ValidateAddressChainResult{}
// 	return
// }
// func (v *DuoVUEnode) Verifychain(a *btcjson.VerifyChainCmd) {
// 	r, err := v.r.cx.RPCServer.HandleVerifyChain(v.r.cx.RPCServer, a, nil)
// }
// func (v *DuoVUEnode) Verifymessage(a *btcjson.VerifyMessageCmd) {
// 	r, err := v.r.cx.RPCServer.HandleVerifyMessage(v.r.cx.RPCServer, a, nil)
// 	r = ""
// 	return
// }
func (r *RcVar) GetWalletVersion() map[string]btcjson.VersionResult {
	v, err := chainrpc.HandleVersion(r.cx.RPCServer, nil, nil)
	if err != nil {
	}
	return v.(map[string]btcjson.VersionResult)
}
