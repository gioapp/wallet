package rcd

func (r *RcVar) SaveDaemonCfg() {
	//marshalled, _ := js.Marshal(r.Settings.Daemon.Config)
	//config, _ := pod.EmptyConfig()
	//if err := js.Unmarshal(marshalled, config); err != nil {
	//}
	//config2.Configure(r.cx, r.cx.AppContext.Command.Name)
	//save.Pod(config)
}

//
//func settings(cx *conte.Xt) *model.DuoUIsettings {
//	settings := &model.DuoUIsettings{
//		Abbrevation: "DUO",
//		Tabs: &model.DuoUIconfTabs{
//			Current:  "wallet",
//			TabsList: make(map[string]*widget.Clickable),
//		},
//		Daemon: &model.DaemonConfig{
//			Config: cx.ConfigMap,
//			Schema: pod.GetConfigSchema(cx.Config, cx.ConfigMap),
//		},
//	}
//	// Settings tabs
//	settingsFields := make(map[string]interface{})
//	for _, group := range settings.Daemon.Schema.Groups {
//		settings.Tabs.TabsList[group.Legend] = new(widget.Clickable)
//		for _, field := range group.Fields {
//			switch field.Type {
//			case "stringSlice":
//				settingsFields[field.Model] = &widget.Editor{
//					SingleLine: false,
//				}
//				switch field.InputType {
//				case "text":
//					if settings.Daemon.Config[field.Model] != nil {
//						var text string
//						for _, str := range *settings.Daemon.
//							Config[field.Model].(*cli.StringSlice) {
//							text = text + str + "\n"
//						}
//						(settingsFields[field.Model]).(*widget.Editor).SetText(text)
//					}
//				}
//			case "input":
//				settingsFields[field.Model] = &widget.Editor{
//					SingleLine: true,
//				}
//				if settings.Daemon.Config[field.Model] != nil {
//					switch field.InputType {
//					case "text":
//						(settingsFields[field.Model]).(*widget.Editor).SetText(
//							fmt.Sprint(*settings.Daemon.Config[field.Model].(*string)))
//					case "number":
//						(settingsFields[field.Model]).(*widget.Editor).SetText(
//							fmt.Sprint(*settings.Daemon.Config[field.Model].(*int)))
//					case "decimal":
//						(settingsFields[field.Model]).(*widget.Editor).SetText(
//							fmt.Sprintf("%0.f", *settings.Daemon.Config[field.
//								Model].(*float64)))
//					case "time":
//						(settingsFields[field.Model]).(*widget.Editor).SetText(
//							fmt.Sprint(*settings.
//								Daemon.Config[field.Model].(*time.Duration)))
//					case "password":
//						(settingsFields[field.Model]).(*widget.Editor).SetText(
//							fmt.Sprint(*settings.Daemon.Config[field.Model].(*string)))
//					}
//				}
//			case "switch":
//				settingsFields[field.Model] = new(gel.CheckBox)
//				(settingsFields[field.Model]).(*gel.CheckBox).SetChecked(
//					*settings.Daemon.Config[field.Model].(*bool))
//			case "radio":
//				settingsFields[field.Model] = new(gel.Enum)
//			default:
//				settingsFields[field.Model] = new(widget.Clickable)
//			}
//		}
//	}
//	settings.Daemon.Widgets = settingsFields
//	return settings
//}
