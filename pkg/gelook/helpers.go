package gelook

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func DuoUIdrawRectangle(gtx *layout.Context, w, h int, color string, borderRadius [4]float32, padding [4]float32) {
	in := layout.Inset{
		Top:    unit.Dp(padding[0]),
		Right:  unit.Dp(padding[1]),
		Bottom: unit.Dp(padding[2]),
		Left:   unit.Dp(padding[3]),
	}
	in.Layout(gtx, func() {
		square := f32.Rectangle{
			Max: f32.Point{
				X: float32(w),
				Y: float32(h),
			},
		}
		paint.ColorOp{Color: HexARGB(color)}.Add(gtx.Ops)
		clip.Rect{Rect: square,
			NE: borderRadius[0], NW: borderRadius[1], SE: borderRadius[2], SW: borderRadius[3]}.Op(gtx.Ops).Add(gtx.Ops) // HLdraw
		paint.PaintOp{Rect: square}.Add(gtx.Ops)
		gtx.Dimensions = layout.Dimensions{Size: image.Point{X: w, Y: h}}
	})
}

func HexARGB(s string) (c color.RGBA) {
	_, _ = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.A, &c.R, &c.G, &c.B)
	return
}

func DuoUIfill(gtx *layout.Context, col string) {
	cs := gtx.Constraints
	d := image.Point{X: cs.Width.Min, Y: cs.Height.Min}
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: HexARGB(col)}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = layout.Dimensions{Size: d}
}

func (t *DuoUItheme) DuoUIline(gtx *layout.Context, verticalPadding, horizontalPadding float32, size int, color string) func() {
	return func() {
		layout.Inset{
			Top:    unit.Dp(verticalPadding),
			Right:  unit.Dp(horizontalPadding),
			Bottom: unit.Dp(verticalPadding),
			Left:   unit.Dp(horizontalPadding),
		}.Layout(gtx, func() {
			DuoUIdrawRectangle(gtx, gtx.Constraints.Width.Max, size, color, [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
		})
	}
}
