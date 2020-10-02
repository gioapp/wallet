package component

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/gioapp/gel/theme"

	"gioui.org/layout"

	"github.com/gioapp/wallet/gui/rcd"
)

func PeersList(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//rc.Network.PeersList.Layout(gtx, len(rc.Network.Peers), func(i int,gtx layout.Context) layout.Dimensions {
		//	t := rc.Network.Peers[i]
		//	th.DuoUIline(gtx, 0, 0, 1, th.Colors["Hint"])()
		//	layout.Flex{
		//		Spacing: layout.SpaceBetween,
		//	}.Layout(gtx,
		//		layout.Rigid(Label(gtx, th, th.Fonts["Mono"], 14, th.Colors["Dark"], fmt.Sprint(t.ID))),
		//		layout.Flexed(1, peerDetails(gtx, th, i, t)),
		//		layout.Rigid(Label(gtx, th, th.Fonts["Mono"], 14, th.Colors["Dark"], t.Addr)))
		//})
		return layout.Dimensions{}
	}
}

func peerDetails(gtx *layout.Context, th *theme.DuoUItheme, i int, t *btcjson.GetPeerInfoResult) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Flex{
		//	Axis:    layout.Horizontal,
		//	Spacing: layout.SpaceAround,
		//}.Layout(gtx,
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.AddrLocal)),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.Services)),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.RelayTxes))),
		//// layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.LastSend))),
		//// layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.LastRecv))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.BytesSent))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.BytesRecv))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.ConnTime))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.TimeOffset))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.PingTime))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.PingWait))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.Version))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], t.SubVer)),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.Inbound))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.StartingHeight))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.CurrentHeight))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.BanScore))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.FeeFilter))),
		//layout.Rigid(Label(gtx, th, th.Fonts["Primary"], 12, th.Colors["Dark"], fmt.Sprint(t.SyncNode))))
		return layout.Dimensions{}
	}
}
