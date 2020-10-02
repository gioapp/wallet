package duoui

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

var (
	passPhrase        string
	confirmPassPhrase string
	passEditor        = &widget.Editor{
		SingleLine: true,
		// Submit:     true,
	}
	confirmPassEditor = &widget.Editor{
		SingleLine: true,
		// Submit:     true,
	}
	listWallet = &layout.List{
		Axis: layout.Vertical,
	}
	encryption         = new(widget.Bool)
	seed               = new(widget.Bool)
	testnet            = new(widget.Bool)
	buttonCreateWallet = new(widget.Clickable)
)

func (ui *DuoUI) DuoUIloaderCreateWallet() func(gtx layout.Context) layout.Dimensions {
	return func(gtx layout.Context) layout.Dimensions {
		//cs := gtx.Constraints
		//gelook.DuoUIdrawRectangle(gtx, cs.Width.Max,
		//	cs.Height.Max, ui.ly.Theme.Colors["Light"],
		//	[4]float32{0, 0, 0, 0}, [4]float32{0, 0, 0, 0})
		//return	layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//	controllers := []func(gtx layout.Context) layout.Dimensions{
		//		func(gtx layout.Context) layout.Dimensions {
		//			bal := ui.ly.Theme.H5("Enter the private passphrase for your new wallet:")
		//			bal.Font.Typeface = ui.ly.Theme.Fonts["Primary"]
		//			bal.Color = ui.ly.Theme.Colors["Dark"]
		//			bal.Layout(gtx)
		//		},
		//		func(gtx layout.Context) layout.Dimensions {
		//			layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//				e := ui.ly.Theme.DuoUIeditor("Enter Passphrase", "Dark", "Light", 32)
		//				e.Font.Typeface = ui.ly.Theme.Fonts["Primary"]
		//				e.Font.Style = text.Regular
		//				e.Layout(gtx, passEditor)
		//				for _, e := range passEditor.Events(gtx) {
		//					switch e.(type) {
		//					case gel.ChangeEvent:
		//						passPhrase = passEditor.Text()
		//					}
		//				}
		//			})
		//		},
		//		func(gtx layout.Context) layout.Dimensions {
		//			layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		//				e := ui.ly.Theme.DuoUIeditor("Repeat Passphrase", "Dark", "Light", 32)
		//				e.Font.Typeface = ui.ly.Theme.Fonts["Primary"]
		//				e.Font.Style = text.Regular
		//				e.Layout(gtx, confirmPassEditor)
		//				for _, e := range confirmPassEditor.Events(gtx) {
		//					switch e.(type) {
		//					case gel.ChangeEvent:
		//						confirmPassPhrase = confirmPassEditor.Text()
		//					}
		//				}
		//			})
		//		},
		//		func(gtx layout.Context) layout.Dimensions {
		//			encryptionCheckBox := ui.ly.Theme.DuoUIcheckBox(
		//				"Do you want to add an additional layer of encryption"+
		//					" for public data?", ui.ly.Theme.Colors["Dark"], ui.ly.Theme.Colors["Dark"])
		//			encryptionCheckBox.Font.Typeface = ui.ly.Theme.Fonts["Primary"]
		//			encryptionCheckBox.Color = gelook.HexARGB(ui.ly.Theme.Colors["Dark"])
		//			encryptionCheckBox.Layout(gtx, encryption)
		//		},
		//		func(gtx layout.Context) layout.Dimensions {
		//			seedCheckBox := ui.ly.Theme.DuoUIcheckBox(
		//				"Do you have an existing wallet seed you want to use?",
		//				ui.ly.Theme.Colors["Dark"], ui.ly.Theme.Colors["Dark"])
		//			seedCheckBox.Font.Typeface = ui.ly.Theme.Fonts["Primary"]
		//			seedCheckBox.Color = gelook.HexARGB(ui.ly.Theme.Colors["Dark"])
		//			seedCheckBox.Layout(gtx, seed)
		//		},
		//		func(gtx layout.Context) layout.Dimensions {
		//			testnetCheckBox := ui.ly.Theme.DuoUIcheckBox(
		//				"Use testnet?", ui.ly.Theme.Colors["Dark"], ui.ly.Theme.Colors["Dark"])
		//			testnetCheckBox.Font.Typeface = ui.ly.Theme.Fonts["Primary"]
		//			testnetCheckBox.Color = gelook.HexARGB(ui.ly.Theme.Colors["Dark"])
		//			testnetCheckBox.Layout(gtx, testnet)
		//		},
		//		func(gtx layout.Context) layout.Dimensions {
		//			var createWalletbuttonComp gelook.DuoUIbutton
		//			createWalletbuttonComp = ui.ly.Theme.DuoUIbutton(ui.ly.Theme.Fonts["Secondary"], "CREATE WALLET", ui.ly.Theme.Colors["Dark"], ui.ly.Theme.Colors["Light"], ui.ly.Theme.Colors["Light"], ui.ly.Theme.Colors["Dark"], "", ui.ly.Theme.Colors["Dark"], 16, 0, 125, 32, 4, 4, 4, 4)
		//			for buttonCreateWallet.Clicked() {
		//				if passPhrase != "" && passPhrase == confirmPassPhrase {
		//					if testnet.Checked(gtx) {
		//						ui.rc.UseTestnet()
		//					}
		//					ui.rc.CreateWallet(passPhrase, "", "", "")
		//					if testnet.Checked(gtx) {
		//						//interrupt.RequestRestart()
		//					}
		//				}
		//			}
		//			createWalletbuttonComp.Layout(gtx, buttonCreateWallet)
		//		},
		//	}
		//	listWallet.Layout(gtx, len(controllers), func(i int, gtx layout.Context) layout.Dimensions)
		//	{
		//		layout.UniformInset(unit.Dp(10)).Layout(gtx, controllers[i])
		//	})
		//})
		return layout.Dimensions{}
	}
}
