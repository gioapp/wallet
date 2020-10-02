// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"time"

	"gioui.org/layout"
)

type Command struct {
	Com      interface{}
	ComID    string
	Category string
	Out      func()
	Time     time.Time
}

func RunCommand(name string) *Command {
	return &Command{
		ComID: name,
	}
}

func (p Command) Layout(gtx layout.Context) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Flex{}.Layout(gtx,
		//	layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		//in := layout.UniformInset(unit.Dp(0))
		//return in.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//cs := gtx.Constraints
		//DuoUIdrawRectangle(gtx, cs.Width.Max, cs.Height.Max, "ffacacac", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
		//layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//cs := gtx.Constraints
		//DuoUIdrawRectangle(gtx, cs.Width.Max, cs.Height.Max, "ffcfcfcf", [4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
		//f()
		//})
		//})
		//}),
		//)
		//}
		return layout.Dimensions{}
	}
}
