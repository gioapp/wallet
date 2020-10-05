package nav

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/btn"
	"github.com/gioapp/wallet/pkg/dap/page"
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
	Items        map[int]Item
	Wide         bool
	Mode         string
	NavLayout    string
	ItemLayout   string
	ItemIconSize unit.Value
	Axis         layout.Axis
	Size         int
	NoContent    bool
	LogoWidget   layout.Widget
	MenuItems    []Item
	CurrentPage  page.Page
}
type Item struct {
	Title string
	Page  page.Page
	Bg    string
	Icon  *widget.Icon
	Btn   *widget.Clickable
}

func (n *Navigation) Nav(th *theme.Theme, gtx C) D {
	n.Size = 128
	n.NoContent = true
	if n.Wide {
		n.Size = 64
		n.NoContent = false
	}
	//switch n.Mode {
	//case "mobile":
	//	gtx.Constraints.Min.Y = n.Size
	//	gtx.Constraints.Max.Y = n.Size
	//	//gtx.Constraints.Min.X = gtx.Constraints.Max.X
	//case "tablet":
	//	gtx.Constraints.Min.X = n.Size
	//	gtx.Constraints.Max.X = n.Size
	//	//gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
	//}
	helper.Fill(gtx, helper.HexARGB(n.Bg))
	navList.Axis = n.Axis

	return lyt.Format(gtx, n.NavLayout,
		n.LogoWidget,
		//func(gtx C)D{ return D{}},
		func(gtx C) D {
			return navList.Layout(gtx, len(n.Items), func(gtx C, i int) D {
				item := n.Items[i]
				btn := btn.IconTextBtn(th, item.Btn, n.ItemLayout, n.NoContent, item.Title)
				btn.TextSize = unit.Dp(12)
				btn.Icon = item.Icon
				btn.IconSize = n.ItemIconSize
				btn.CornerRadius = unit.Dp(0)
				btn.Background = th.Colors["NavBg"]
				btn.TextColor = th.Colors["NavItem"]
				for item.Btn.Clicked() {
					n.CurrentPage = item.Page
				}
				return btn.Layout(gtx)
			})
			//func(gtx C)D{ return D{}})
		})
}
