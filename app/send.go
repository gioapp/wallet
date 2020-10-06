package gwallet

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/wallet/pkg/btn"
	"github.com/gioapp/wallet/pkg/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	addressBookBtn    = new(widget.Clickable)
	pasteClipboardBtn = new(widget.Clickable)
	clearBtn          = new(widget.Clickable)
	allAvailableBtn   = new(widget.Clickable)
	chooseFeeBtn      = new(widget.Clickable)
	clearAllBtn       = new(widget.Clickable)
	sendBtn           = new(widget.Clickable)
	addressInput      = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	labelInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	amountInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func (g *GioWallet) GetSend() {

}

func (g *GioWallet) sendHeader() func(gtx C) D {
	return boxBase(g.ui.Theme.Colors["PanelBg"], func(gtx C) D {
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
	return boxBase(g.ui.Theme.Colors["PanelBg"], func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.ui.Theme.Colors["PanelBg"]))
		return lyt.Format(gtx, "vflex(middle,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			func(gtx C) D {
				return lyt.Format(gtx, "hflex(middle,r(_),r(_),r(_),f(1,inset(8dp8dp8dp8dp,_)))",
					func(gtx C) D {
						title := theme.H6(g.ui.Theme, "Transaction Fee:")
						title.Alignment = text.Start
						return title.Layout(gtx)
					},
					func(gtx C) D {
						title := theme.H6(g.ui.Theme, "0.00000 DUO/kb")
						title.Alignment = text.Start
						return title.Layout(gtx)
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
						title := theme.H6(g.ui.Theme, "Warning: Fee estimation is currently possible.")
						title.Alignment = text.Start
						return title.Layout(gtx)
					},
				)
			},
			func(gtx C) D {
				return lyt.Format(gtx, "hmax(hflex(middle,r(_),r(_),r(_),f(1,inset(8dp8dp8dp8dp,_))))",
					sendButton(g.ui.Theme, sendBtn, "Send", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
					sendButton(g.ui.Theme, clearAllBtn, "Close", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
					sendButton(g.ui.Theme, clearAllBtn, "CounterPlus", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
					func(gtx C) D {
						title := theme.H6(g.ui.Theme, "Balance: 15.26656.5664654 DUO")
						title.Alignment = text.Start
						return title.Layout(gtx)
					},
				)
			})
	})
}

func (g *GioWallet) sendBody() func(gtx C) D {
	return func(gtx C) D {
		//gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflex(start,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			sendRow(g.ui.Theme, "Pay to:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,f(1,inset(8dp8dp8dp8dp,_)),r(_),r(_),r(_))",
						boxBase(g.ui.Theme.Colors["PanelBg"], func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							e := material.Editor(g.ui.Theme.T, addressInput, "Enter a ParallelCoin address (e.g. 9ef0sdjifvmlkdsfnsdlkg)")
							return e.Layout(gtx)
						}),
						sendButton(g.ui.Theme, addressBookBtn, "AddressBook", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
						sendButton(g.ui.Theme, pasteClipboardBtn, "Paste", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
						sendButton(g.ui.Theme, clearBtn, "Close", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
					)
				}),
			sendRow(g.ui.Theme, "Label:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,f(1,inset(8dp8dp8dp8dp,_)))",
						boxBase(g.ui.Theme.Colors["PanelBg"], func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							e := material.Editor(g.ui.Theme.T, labelInput, "Enter a label for this address to add it to the list of used addresses")
							return e.Layout(gtx)
						}),
					)
				}),
			sendRow(g.ui.Theme, "Amount:",
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
	}
}

func sendRow(th *theme.Theme, label string, content func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflex(start,r(inset(0dp0dp0dp0dp,_)),f(1,_))",
			func(gtx C) D {
				gtx.Constraints.Min.X = 80
				gtx.Constraints.Min.X = 80
				title := theme.Body(th, label)
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			content,
		)
	}
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
