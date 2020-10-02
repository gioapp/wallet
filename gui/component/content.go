package component

import (
	"fmt"
	"gioui.org/widget"
	"github.com/gioapp/gel/page"
	"github.com/gioapp/gel/theme"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"

	"github.com/gioapp/wallet/gui/rcd"
)

var (
	previousBlockHashButton = new(widget.Clickable)
	nextBlockHashButton     = new(widget.Clickable)
	list                    = &layout.List{
		Axis: layout.Vertical,
	}
)

func UnoField(gtx *layout.Context, field func()) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//layout.Flex{
		//	Axis:    layout.Horizontal,
		//	Spacing: layout.SpaceAround,
		//}.Layout(gtx,
		//	layout.Flexed(1, field),
		//)
		return layout.Dimensions{}
	}
}
func DuoFields(gtx *layout.Context, axis layout.Axis, left, right func()) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{
			Axis:    axis,
			Spacing: layout.SpaceAround,
		}.Layout(gtx,
			fieldAxis(axis, left, 0.5),
			fieldAxis(axis, right, 0.5),
		)
		return layout.Dimensions{}
	}
}

func TrioFields(gtx *layout.Context, th *theme.DuoUItheme, axis layout.Axis, labelTextSize, valueTextSize float32, unoLabel, unoValue, unoHeadcolor, unoHeadbgColor, unoColor, unoBgColor, duoLabel, duoValue, duoHeadcolor, duoHeadbgColor, duoColor, duoBgColor, treLabel, treValue, treHeadcolor, treHeadbgColor, treColor, treBgColor string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{
			Axis:    axis,
			Spacing: layout.SpaceAround,
		}.Layout(gtx)//fieldAxis(axis, ContentLabeledField(gtx, th, layout.Vertical, 4, labelTextSize, valueTextSize, unoLabel, unoHeadcolor, unoHeadbgColor, unoColor, unoBgColor, fmt.Sprint(unoValue)), 0.3333),
		//fieldAxis(axis, ContentLabeledField(gtx, th, layout.Vertical, 4, labelTextSize, valueTextSize, duoLabel, duoHeadcolor, duoHeadbgColor, duoColor, duoBgColor, fmt.Sprint(duoValue)), 0.3333),
		//fieldAxis(axis, ContentLabeledField(gtx, th, layout.Vertical, 4, labelTextSize, valueTextSize, treLabel, treHeadbgColor, treHeadcolor, treColor, treBgColor, fmt.Sprint(treValue)), 0.3333),

		return layout.Dimensions{}
	}
}

func fieldAxis(axis layout.Axis, field func(), size float32) layout.FlexChild {
	var f layout.FlexChild
	switch axis {
	case layout.Horizontal:
		//f = layout.Flexed(size, field)
	case layout.Vertical:
		//f = layout.Rigid(field)
	}
	return f
}

func ContentLabeledField(gtx layout.Context, th *theme.DuoUItheme, axis layout.Axis, margin, labelTextSize, valueTextSize float32, label, headcolor, headbgColor, color, bgColor, value string) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.UniformInset(unit.Dp(margin)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			layout.Flex{
				Axis: axis,
			}.Layout(gtx,
				layout.Rigid(contentField(gtx, th, label, th.Colors[headcolor], th.Colors[headbgColor], th.Fonts["Primary"], 4, labelTextSize)),
				layout.Rigid(contentField(gtx, th, value, th.Colors[color], th.Colors[bgColor], th.Fonts["Mono"], 4, valueTextSize)))
			return layout.Dimensions{}
		})
		return layout.Dimensions{}
	}
}

func PageNavButtons(rc *rcd.RcVar, gtx *layout.Context, th *theme.DuoUItheme, previousBlockHash, nextBlockHash string, prevPage, nextPage *page.DuoUIpage) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		layout.Flex{}.Layout(gtx,
			layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
				//eh := chainhash.Hash{}
				//if previousBlockHash != eh.String() {
				//	pageNavButton(rc, gtx, th, nextPage, previousBlockHashButton, "Previous Block", previousBlockHash)
				//}
				return layout.Dimensions{}
			}),
			layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
				if nextBlockHash != "" {
					pageNavButton(rc, gtx, th, nextPage, nextBlockHashButton, "Next Block", nextBlockHash)
				}
				return layout.Dimensions{}
			}))
		return layout.Dimensions{}
	}
}

func pageNavButton(rc *rcd.RcVar, gtx layout.Context, th *theme.DuoUItheme, page *page.DuoUIpage, b *widget.Clickable, label, hash string) {
	layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//var blockButton gelook.DuoUIbutton
		//blockButton = th.DuoUIbutton(th.Fonts["Mono"], label+" "+hash, th.Colors["Light"], th.Colors["Info"], th.Colors["Info"], th.Colors["Light"], "", th.Colors["Light"], 16, 0, 60, 24, 0, 0, 0, 0)
		for b.Clicked() {
			rc.ShowPage = fmt.Sprintf("BLOCK %s", hash)
			//rc.GetSingleBlock(hash)()
			SetPage(rc, page)
		}
		//blockButton.Layout(gtx, b)
		return layout.Dimensions{}
	})
}

func contentField(gtx layout.Context, th *theme.DuoUItheme, text, color, bgColor string, font text.Typeface, padding, textSize float32) func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//hmin := gtx.Constraints.Width.Min
		//vmin := gtx.Constraints.Height.Min
		layout.Stack{Alignment: layout.W}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {

				//fill(gtx, gelook.HexARGB(bgColor))
				return layout.Dimensions{}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				//gtx.Constraints.Width.Min = hmin
				//gtx.Constraints.Height.Min = vmin
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					layout.UniformInset(unit.Dp(padding)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						//l := th.DuoUIlabel(unit.Dp(textSize), text)
						//l.Font.Typeface = font
						//l.Color = color
						//l.Layout(gtx)
						return layout.Dimensions{}
					})
					return layout.Dimensions{}
				})
				return layout.Dimensions{}
			}),
		)
		return layout.Dimensions{}
	}
}
