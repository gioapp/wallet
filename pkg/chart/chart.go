package chart

import (
	"gioui.org/layout"
)

var (
	chartList = &layout.List{
		Axis: layout.Horizontal,
	}
)

//
//func Chart(values []float64) func(gtx layout.Context) layout.Dimensions {
//	return func(gtx layout.Context) layout.Dimensions {
//		return chartList.Layout(gtx, len(values), func(gtx C, i int) D {
//			value := values[i]
//
//		})
//	}
//}

//func verticalPx(topSize, bottomSize float64) func(gtx layout.Context) layout.Dimensions {
//	lF := "vflexb(start,f(" + fmt.Sprintf("%.2f", topSize) + ",_)),f(" + fmt.Sprintf("%.2f", bottomSize) + ",_))"
//	top := "ff664825"
//	bottom := "ff994672"
//	return func(gtx layout.Context) layout.Dimensions {
//		return lyt.Format(gtx, lF,
//			verticalPart(top),
//			verticalPart(bottom),
//		)
//	}
//}
//
//func verticalPart(c string) func(gtx layout.Context) layout.Dimensions {
//	return func(gtx layout.Context) layout.Dimensions {
//		return helper.Fill(gtx, helper.HexARGB(c))
//	}
//}
