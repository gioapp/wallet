package ico

import (
	"gioui.org/widget"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

func NewDuoUIicons() (i map[string]*widget.Icon) {
	i = make(map[string]*widget.Icon)
	i["Console"] = mustIcon(widget.NewIcon(icons.ActionInput))
	i["ActionTimeline"] = mustIcon(widget.NewIcon(icons.ActionTimeline))
	i["AddressBook"] = mustIcon(widget.NewIcon(icons.ActionBook))
	i["Explore"] = mustIcon(widget.NewIcon(icons.ActionExplore))
	i["Build"] = mustIcon(widget.NewIcon(icons.ActionBuild))
	i["Checked"] = mustIcon(widget.NewIcon(icons.ToggleCheckBox))
	i["CHK"] = mustIcon(widget.NewIcon(icons.AlertWarning))
	i["Close"] = mustIcon(widget.NewIcon(icons.NavigationClose))
	i["CommunicationImportExport"] = mustIcon(widget.NewIcon(icons.CommunicationImportExport))
	i["Console"] = mustIcon(widget.NewIcon(icons.HardwareKeyboardHide))
	i["Copy"] = mustIcon(widget.NewIcon(icons.ContentContentCopy))
	i["CounterMinus"] = mustIcon(widget.NewIcon(icons.ImageExposureNeg1))
	i["CounterPlus"] = mustIcon(widget.NewIcon(icons.ImageExposurePlus1))
	i["DBG"] = mustIcon(widget.NewIcon(icons.ActionBugReport))
	i["Delete"] = mustIcon(widget.NewIcon(icons.ActionDelete))
	i["DeviceSignalCellular0Bar"] = mustIcon(widget.NewIcon(icons.DeviceSignalCellular0Bar))
	i["DeviceWidgets"] = mustIcon(widget.NewIcon(icons.DeviceWidgets))
	i["Down"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["EditorMonetizationOn"] = mustIcon(widget.NewIcon(icons.EditorMonetizationOn))
	i["ERR"] = mustIcon(widget.NewIcon(icons.AlertError))
	i["Filter"] = mustIcon(widget.NewIcon(icons.ContentFilterList))
	i["FilterAll"] = mustIcon(widget.NewIcon(icons.ActionDoneAll))
	i["FilterNone"] = mustIcon(widget.NewIcon(icons.ContentBlock))
	i["Folded"] = mustIcon(widget.NewIcon(icons.NavigationChevronRight))
	i["FoldIn"] = mustIcon(widget.NewIcon(icons.ContentRemove))
	i["FTL"] = mustIcon(widget.NewIcon(icons.ImageFlashOn))
	i["Grab"] = mustIcon(widget.NewIcon(icons.NavigationMenu))
	i["HardwareWatch"] = mustIcon(widget.NewIcon(icons.HardwareWatch))
	i["Help"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["HideAll"] = mustIcon(widget.NewIcon(icons.NavigationUnfoldLess))
	i["HideItem"] = mustIcon(widget.NewIcon(icons.ActionVisibilityOff))
	i["Transactions"] = mustIcon(widget.NewIcon(icons.ActionHistory))
	i["Cancel"] = mustIcon(widget.NewIcon(icons.NavigationCancel))
	i["Close"] = mustIcon(widget.NewIcon(icons.NavigationClose))
	i["Down"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["Grab"] = mustIcon(widget.NewIcon(icons.NavigationMenu))
	i["OK"] = mustIcon(widget.NewIcon(icons.NavigationCheck))
	i["Up"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropUp))
	i["ImageRemoveRedEye"] = mustIcon(widget.NewIcon(icons.ImageRemoveRedEye))
	i["ImageTimer"] = mustIcon(widget.NewIcon(icons.ImageTimer))
	i["INF"] = mustIcon(widget.NewIcon(icons.ActionInfo))
	i["Kill"] = mustIcon(widget.NewIcon(icons.NavigationCancel))
	i["Logo"] = mustIcon(widget.NewIcon(ParallelCoin))
	i["MapsLayers"] = mustIcon(widget.NewIcon(icons.MapsLayers))
	i["MapsLayersClear"] = mustIcon(widget.NewIcon(icons.MapsLayersClear))
	i["Minimize"] = mustIcon(widget.NewIcon(icons.NavigationExpandMore))
	i["Peers"] = mustIcon(widget.NewIcon(icons.NotificationWiFi))
	i["networkIcon"] = mustIcon(widget.NewIcon(icons.ActionFingerprint))
	i["NotificationNetworkCheck"] = mustIcon(widget.NewIcon(icons.NotificationNetworkCheck))
	i["NotificationSync"] = mustIcon(widget.NewIcon(icons.NotificationSync))
	i["NotificationSyncDisabled"] = mustIcon(widget.NewIcon(icons.NotificationSyncDisabled))
	i["NotificationSyncProblem"] = mustIcon(widget.NewIcon(icons.NotificationSyncProblem))
	i["NotificationVPNLock"] = mustIcon(widget.NewIcon(icons.NotificationVPNLock))
	i["Overview"] = mustIcon(widget.NewIcon(icons.ActionHome))
	i["Paste"] = mustIcon(widget.NewIcon(icons.ContentContentPaste))
	i["Pause"] = mustIcon(widget.NewIcon(icons.AVPause))
	i["RadioChecked"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonChecked))
	i["RadioUnchecked"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonUnchecked))
	i["Receive"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropDown))
	i["Restart"] = mustIcon(widget.NewIcon(icons.NavigationRefresh))
	i["Run"] = mustIcon(widget.NewIcon(icons.AVPlayArrow))
	i["Screenshot"] = mustIcon(widget.NewIcon(icons.ContentSelectAll))
	i["Send"] = mustIcon(widget.NewIcon(icons.ContentSend))
	i["Settings"] = mustIcon(widget.NewIcon(icons.ActionSettings))
	i["ShowAll"] = mustIcon(widget.NewIcon(icons.NavigationUnfoldMore))
	i["ShowItem"] = mustIcon(widget.NewIcon(icons.ActionVisibility))
	i["Sidebar"] = mustIcon(widget.NewIcon(icons.ActionChromeReaderMode))
	i["Stop"] = mustIcon(widget.NewIcon(icons.AVStop))
	i["ToggleOff"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonUnchecked))
	i["ToggleOn"] = mustIcon(widget.NewIcon(icons.ToggleRadioButtonChecked))
	i["Trace"] = mustIcon(widget.NewIcon(icons.ActionTrackChanges))
	i["TRC"] = mustIcon(widget.NewIcon(icons.ActionSearch))
	i["Unchecked"] = mustIcon(widget.NewIcon(icons.ToggleCheckBoxOutlineBlank))
	i["Unfolded"] = mustIcon(widget.NewIcon(icons.NavigationExpandMore))
	i["Up"] = mustIcon(widget.NewIcon(icons.NavigationArrowDropUp))
	i["WRN"] = mustIcon(widget.NewIcon(icons.ActionHelp))
	i["Zoom"] = mustIcon(widget.NewIcon(icons.NavigationExpandLess))
	//i["AddressBook"] = mustIcon(widget.NewIcon(AddressBook))
	//i["Balance"] = mustIcon(widget.NewIcon(Balance))
	//i["Blocks"] = mustIcon(widget.NewIcon(Blocks))
	//i["Copy"] = mustIcon(widget.NewIcon(Copy))
	//i["Help"] = mustIcon(widget.NewIcon(Help))
	//i["History"] = mustIcon(widget.NewIcon(History))
	//i["Info"] = mustIcon(widget.NewIcon(Info))
	//i["Light"] = mustIcon(widget.NewIcon(Light))
	//i["Link"] = mustIcon(widget.NewIcon(Link))
	//i["LinkOff"] = mustIcon(widget.NewIcon(LinkOff))
	//i["Loading"] = mustIcon(widget.NewIcon(Loading))
	//i["NoLight"] = mustIcon(widget.NewIcon(NoLight))
	//i["Ok"] = mustIcon(widget.NewIcon(Ok))
	//i["Overview"] = mustIcon(widget.NewIcon(Overview))
	//i["ParallelCoin"] = mustIcon(widget.NewIcon(ParallelCoin))
	//i["ParallelCoinRound"] = mustIcon(widget.NewIcon(ParallelCoinRound))
	//i["Peers"] = mustIcon(widget.NewIcon(Peers))
	//i["Receive"] = mustIcon(widget.NewIcon(Receive))
	//i["Received"] = mustIcon(widget.NewIcon(Received))
	//i["Search"] = mustIcon(widget.NewIcon(Search))
	//i["Send"] = mustIcon(widget.NewIcon(Send))
	//i["Sent"] = mustIcon(widget.NewIcon(Sent))
	//i["Settings"] = mustIcon(widget.NewIcon(Settings))
	//i["TxNumber"] = mustIcon(widget.NewIcon(TxNumber))
	//i["Unconfirmed"] = mustIcon(widget.NewIcon(Unconfirmed))
	return i
}

func mustIcon(ic *widget.Icon, err error) *widget.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}
