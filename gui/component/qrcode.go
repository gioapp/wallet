package component

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"github.com/nfnt/resize"

	"github.com/gioapp/wallet/gui/model"
	"github.com/gioapp/wallet/gui/rcd"
	"github.com/gioapp/wallet/pkg/qrcode"
)

// func DuoUIqrCode(pubAddr string) {
//	//qr, err := qrcode.New(strings.ToUpper(pubAddr), qrcode.Medium)
//	//if err != nil {
//	//	Fatal(err)
//	//}
//	//qr.BackgroundColor = rgb(0xe8f5e9)
//	//addrQR := paint.NewImageOp(qr.Image(256))
//	return
// }

// func NewQrCode(pubAddr string) *model.DuoUIqrCode {
//	//qr, err := qrcode.New(strings.ToUpper(pubAddr), qrcode.Medium)
//	//if err != nil {
//	//	Fatal(err)
//	//}
//	//Info(pubAddr)
//	//qr.BackgroundColor = theme.HexARGB("ff3030cf")
//	//return &model.DuoUIqrCode{
//	//	AddrQR:  paint.NewImageOp(qr.Image(256)),
//	//	PubAddr: pubAddr,
//	//}
// }

func DuoUIqrCode(gtx *layout.Context, hash string, size uint) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		qr, err := qrcode.Encode(hash, 3, qrcode.ECLevelM)
		if err != nil {
		}
		qrResize := resize.Resize(size, 0, qr, resize.NearestNeighbor)
		addrQR := paint.NewImageOp(qrResize)

		//sz := gtx.Constraints.Width.Constrain(gtx.Px(unit.Dp(float32(size))))
		addrQR.Add(gtx.Ops)
		paint.PaintOp{
			Rect: f32.Rectangle{
				Max: f32.Point{
					//X: float32(sz), Y: float32(sz),
				},
			},
		}.Add(gtx.Ops)
		//gtx.Dimensions.Size = image.Point{X: sz, Y: sz}
		return layout.Dimensions{}
	}
}

func QrDialog(rc *rcd.RcVar, gtx *layout.Context, address string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		// clipboard.Set(t.Address)
		rc.Dialog.Show = true
		rc.Dialog = &model.DuoUIdialog{
			Show: true,
			//CustomField: DuoUIqrCode(gtx, address, 256),
			//Orange:      func(gtx layout.Context)layout.Dimensions{ rc.Dialog.Show = false },
			OrangeLabel: "CLOSE",
			Title:       "ParallelCoin address",
			Text:        address,
		}
		return layout.Dimensions{}
	}
}
