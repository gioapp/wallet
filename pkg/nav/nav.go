package nav

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/btn"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	navList = &layout.List{}
)

type Navigation struct {
	Name         string
	Bg           string
	Logo         Logo
	Items        []Item
	wide         bool
	mode         string
	navLayout    string
	itemLayout   string
	ItemIconSize unit.Value
	axis         layout.Axis
	size         int
	noContent    bool
	logo         layout.Widget
	menuItems    []Item
	CurrentPage  string
}
type Item struct {
	Title string
	Bg    string
	Icon  *widget.Icon
	Btn   *widget.Clickable
}

func (n *Navigation) Nav(th *theme.Theme, gtx C) D {
	n.size = 128
	n.noContent = true
	if n.wide {
		n.size = 64
		n.noContent = false
	}
	switch n.mode {
	case "mobile":
		gtx.Constraints.Min.Y = n.size
		gtx.Constraints.Max.Y = n.size
		//gtx.Constraints.Min.X = gtx.Constraints.Max.X
	case "tablet":
		gtx.Constraints.Min.X = n.size
		gtx.Constraints.Max.X = n.size
		//gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
	}
	helper.Fill(gtx, helper.HexARGB(n.Bg))
	navList.Axis = n.axis

	return lyt.Format(gtx, n.navLayout,
		n.logo,
		func(gtx C) D {
			return navList.Layout(gtx, len(n.Items), func(gtx C, i int) D {
				item := n.Items[i]
				btn := btn.IconTextBtn(th, item.Btn, n.itemLayout, n.noContent, item.Title)
				btn.TextSize = unit.Dp(12)
				btn.Icon = item.Icon
				btn.IconSize = n.ItemIconSize
				btn.CornerRadius = unit.Dp(0)
				btn.Background = th.Colors["NavBg"]
				btn.TextColor = th.Colors["NavItem"]
				for item.Btn.Clicked() {
					n.CurrentPage = item.Title
				}
				return btn.Layout(gtx)
			})
		})
}