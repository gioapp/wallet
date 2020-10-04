package gwallet

import (
	"math"

	"strconv"
)

var (
	noReturn = func(gtx C) D { return D{} }

	//contentList = &layout.List{
	//	Axis: layout.Vertical,
	//}
	suffixes = [7]string{
		0: "B",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
		5: "PB",
		6: "EB",
	}
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func formatByteSize(size float64) string {
	//size := sizeInMB * 1024 * 1024
	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	//fmt.Println("dole", strconv.FormatFloat(getSize, 'f', -1, 64)+" "+string(getSuffix))
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
