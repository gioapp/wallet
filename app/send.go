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
	sendAddresses     []*SendAddress
	sendAddressesList = &layout.List{
		Axis: layout.Vertical,
	}
	sendChooseFeeBtn = new(widget.Clickable)
	sendClearAllBtn  = new(widget.Clickable)
	sendBtn          = new(widget.Clickable)
	addRecipientBtn  = new(widget.Clickable)
)

type SendAddress struct {
	sendAddressInput      *widget.Editor
	sendLabelInput        *widget.Editor
	sendAmountInput       *widget.Editor
	sendAddressBookBtn    *widget.Clickable
	sendPasteClipboardBtn *widget.Clickable
	sendClearBtn          *widget.Clickable
	sendAllAvailableBtn   *widget.Clickable
}

func (g *GioWallet) GetSend() {

}

func CreateSendAddressItem() func() {
	return func() {
		sendAddresses = append(sendAddresses,
			&SendAddress{
				sendAddressInput: &widget.Editor{
					SingleLine: true,
					Submit:     true,
				},
				sendLabelInput: &widget.Editor{
					SingleLine: true,
					Submit:     true,
				},
				sendAmountInput: &widget.Editor{
					SingleLine: true,
					Submit:     true,
				},
				sendAddressBookBtn:    new(widget.Clickable),
				sendPasteClipboardBtn: new(widget.Clickable),
				sendClearBtn:          new(widget.Clickable),
				sendAllAvailableBtn:   new(widget.Clickable),
			})
	}
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
		return lyt.Format(gtx, "hmax(hflex(middle,r(_),r(inset(0dp8dp0dp8dp,_)),r(_),f(1,inset(8dp8dp8dp8dp,_))))",
			sendButton(th, sendBtn, "Send", "hflex(r(_),r(_)))", "Send", func() {}),
			sendButton(th, sendClearAllBtn, "Close", "hflex(r(_),r(_)))", "Clear All", func() {}),
			sendButton(th, addRecipientBtn, "CounterPlus", "hflex(r(_),r(_)))", "Add Recipient", CreateSendAddressItem()),
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
		return lyt.Format(gtx, "hflex(middle,r(_),r(inset(0dp8dp0dp8dp,_)),r(_),f(1,inset(8dp8dp8dp8dp,_)))",
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
		return sendAddressesList.Layout(gtx, len(sendAddresses), singleAddress(g.ui.Theme))
	})
}

func singleAddress(th *theme.Theme) func(gtx C, i int) D {
	return func(gtx C, i int) D {
		return lyt.Format(gtx, "vflex(start,r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_)),r(inset(5dp0dp5dp0dp,_))))",
			labeledRow(th, "Pay to:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,f(1,inset(8dp8dp8dp8dp,_)),r(_),r(_),r(_))",
						box.BoxEditor(th, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							e := material.Editor(th.T, sendAddresses[i].sendAddressInput, "Enter a ParallelCoin address (e.g. 9ef0sdjifvmlkdsfnsdlkg)")
							return e.Layout(gtx)
						}),
						sendButton(th, sendAddresses[i].sendAddressBookBtn, "AddressBook", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
						sendButton(th, sendAddresses[i].sendPasteClipboardBtn, "Paste", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
						sendButton(th, sendAddresses[i].sendClearBtn, "Close", "max(inset(0dp0dp0dp0dp,_))", "tetet", func() {}),
					)
				}),
			labeledRow(th, "Label:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,f(1,inset(8dp8dp8dp8dp,_)))",
						box.BoxEditor(th, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							e := material.Editor(th.T, sendAddresses[i].sendLabelInput, "Enter a label for this address to add it to the list of used addresses")
							return e.Layout(gtx)
						}),
					)
				}),
			labeledRow(th, "Amount:",
				func(gtx C) D {
					return lyt.Format(gtx, "hflex(middle,r(_),r(_),r(_),r(_))",
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
							btn := material.IconButton(th.T, connectionsBtn, th.Icons["networkIcon"])
							btn.Inset = layout.Inset{unit.Dp(2), unit.Dp(2), unit.Dp(2), unit.Dp(2)}
							btn.Size = unit.Dp(21)
							btn.Background = helper.HexARGB(th.Colors["Secondary"])
							for connectionsBtn.Clicked() {
								//ui.N.CurrentPage = "Welcome"
							}
							return btn.Layout(gtx)
						},
					)
				}),
			helper.DuoUIline(false, 1, 0, 1, th.Colors["Border"]),
		)
	}
}

func sendButton(th *theme.Theme, c *widget.Clickable, icon, lay, label string, onClick func()) func(gtx C) D {
	return box.BoxButton(th, func(gtx C) D {
		noLabel := true
		if label != "" {
			noLabel = true
		}
		b := btn.IconTextBtn(th, c, lay, noLabel, label)
		b.TextSize = unit.Dp(12)
		b.Background = th.Colors["ButtonBg"]
		b.Icon = th.Icons[icon]
		b.IconSize = unit.Dp(16)
		b.CornerRadius = unit.Dp(0)
		//b.Background = th.Colors["NavBg"]
		b.TextColor = th.Colors["ButtonText"]
		for c.Clicked() {
			//n.CurrentPage = item.Page
			onClick()
		}
		return b.Layout(gtx)
	})
}
