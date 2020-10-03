package gwallet

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
	"github.com/gioapp/wallet/pkg/icontextbtn"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	navList = &layout.List{}
)

func (n *Navigation) Nav(th *theme.Theme, gtx C) D {
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
				btn := icontextbtn.IconTextBtn(th, item.Btn, item.Icon, unit.Dp(48), th.Colors["NavItem"], item.Title, n.itemLayout, n.noText)
				btn.TextSize = unit.Dp(12)
				btn.CornerRadius = unit.Dp(0)
				btn.Background = helper.HexARGB(th.Colors["NavBg"])
				for item.Btn.Clicked() {
					currentPage = item.Title
				}
				return btn.Layout(gtx)
			})
		})
}
