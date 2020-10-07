package gwallet

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/btn"
	"github.com/gioapp/wallet/pkg/dap/box"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	sendAddressBookBtn    = new(widget.Clickable)
	sendPasteClipboardBtn = new(widget.Clickable)
	sendClearBtn          = new(widget.Clickable)
	sendAllAvailableBtn   = new(widget.Clickable)
	sendChooseFeeBtn      = new(widget.Clickable)
	sendClearAllBtn       = new(widget.Clickable)
	sendBtn               = new(widget.Clickable)
	sendAddressInput      = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	sendLabelInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	sendAmountInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func (g *GioWallet) GetSend() {

}

func (g *GioWallet) sendHeader() func(gtx C) D {
	return box.BoxPanel(g.ui.Theme, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.ui.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflex(middle,r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				gtx.Constraints.Min.X = gtx.Constraints.Max.X
				title := theme.H6(g.ui.Theme, "Send Header")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}

func (g *GioWallet) sendFooter() func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflex(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			sendFooterTop(g.ui.Theme),
			sendFooterBottom(g.ui.Theme))
	}
}

func sendFooterBottom(th *theme.Theme) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hmax(hflex(middle,r(_),r(_),r(_),f(1,inset(8dp8dp8dp8dp,_))))",
			sendButton(th, sendBtn, "Send", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
			sendButton(th, sendClearAllBtn, "Close", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
			sendButton(th, sendClearAllBtn, "CounterPlus", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
			func(gtx C) D {
				title := theme.H6(th, "Balance: 15.26656.5664654 DUO")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
		)
	}
}

func sendFooterTop(th *theme.Theme) func(gtx C) D {
	return box.BoxPanel(th, func(gtx C) D {
		return lyt.Format(gtx, "hflex(middle,r(_),r(_),r(_),f(1,inset(8dp8dp8dp8dp,_)))",
			func(gtx C) D {
				title := theme.H6(th, "Transaction Fee:")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				title := theme.H6(th, "0.00000 DUO/kb")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				btn := material.IconButton(th.T, connectionsBtn, th.Icons["networkIcon"])
				btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
				btn.Size = unit.Dp(21)
				btn.Background = helper.HexARGB(th.Colors["Secondary"])
				for connectionsBtn.Clicked() {
					//ui.N.CurrentPage = "Welcome"
				}
				return btn.Layout(gtx)
			},
			func(gtx C) D {
				title := theme.H6(th, "Warning: Fee estimation is currently possible.")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
		)
	})
}

func (g *GioWallet) sendBody() func(gtx C) D {
	return box.BoxPanel(g.ui.Theme, func(gtx C) D {
		return lyt.Format(gtx, "vflex(start,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			labeledRow(g.ui.Theme, "Pay to:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,f(1,inset(8dp8dp8dp8dp,_)),r(_),r(_),r(_))",
						box.BoxEditor(g.ui.Theme, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							e := material.Editor(g.ui.Theme.T, sendAddressInput, "Enter a ParallelCoin address (e.g. 9ef0sdjifvmlkdsfnsdlkg)")
							return e.Layout(gtx)
						}),
						sendButton(g.ui.Theme, sendAddressBookBtn, "AddressBook", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
						sendButton(g.ui.Theme, sendPasteClipboardBtn, "Paste", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
						sendButton(g.ui.Theme, sendClearBtn, "Close", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
					)
				}),
			labeledRow(g.ui.Theme, "Label:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,f(1,inset(8dp8dp8dp8dp,_)))",
						box.BoxEditor(g.ui.Theme, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							e := material.Editor(g.ui.Theme.T, sendLabelInput, "Enter a label for this address to add it to the list of used addresses")
							return e.Layout(gtx)
						}),
					)
				}),
			labeledRow(g.ui.Theme, "Amount:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,r(_),r(_),r(_),r(_))",
						func(gtx C) D {
							btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
							btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
							btn.Size = unit.Dp(21)
							btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
							for connectionsBtn.Clicked() {
								//ui.N.CurrentPage = "Welcome"
							}
							return btn.Layout(gtx)
						},
						func(gtx C) D {
							btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
							btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
							btn.Size = unit.Dp(21)
							btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
							for connectionsBtn.Clicked() {
								//ui.N.CurrentPage = "Welcome"
							}
							return btn.Layout(gtx)
						},
						func(gtx C) D {
							btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
							btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
							btn.Size = unit.Dp(21)
							btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
							for connectionsBtn.Clicked() {
								//ui.N.CurrentPage = "Welcome"
							}
							return btn.Layout(gtx)
						},
						func(gtx C) D {
							btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
							btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
							btn.Size = unit.Dp(21)
							btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
							for connectionsBtn.Clicked() {
								//ui.N.CurrentPage = "Welcome"
							}
							return btn.Layout(gtx)
						},
						func(gtx C) D {
							btn := material.IconButton(g.ui.Theme.T, connectionsBtn, g.ui.Theme.Icons["networkIcon"])
							btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
							btn.Size = unit.Dp(21)
							btn.Background = helper.HexARGB(g.ui.Theme.Colors["Secondary"])
							for connectionsBtn.Clicked() {
								//ui.N.CurrentPage = "Welcome"
							}
							return btn.Layout(gtx)
						},
					)
				}),
		)
	})
}

func sendButton(th *theme.Theme, c *widget.Clickable, icon, lay, label string, onClick func()) func(gtx C) D {
	return func(gtx C) D {
		noLabel := true
		if label != "" {
			noLabel = false
		}
		b := btn.IconTextBtn(th, c, lay, noLabel, label)
		b.TextSize = unit.Dp(12)
		b.Icon = th.Icons[icon]
		b.IconSize = unit.Dp(16)
		b.CornerRadius = unit.Dp(0)
		b.Background = th.Colors["NavBg"]
		b.TextColor = th.Colors["NavItem"]
		for c.Clicked() {
			//n.CurrentPage = item.Page
			onClick()
		}
		return b.Layout(gtx)
	}
}
