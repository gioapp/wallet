package gwallet

import "gioui.org/widget"

var (
	dummyStatus   = Status{}
	dummyBalances = Balances{
		available: "14924.222",
		pending:   "155.454",
		immature:  "15.1114",
		total:     "9155.05454",
	}
	dummyLatestTxs = []Tx{
		Tx{
			Id:      "1ec0bcc5b8847f038b1d8e87cb7932f08dc91c090cc764761b98627f359866c4",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "3173d5f88fef47b1392ddc08d3450ea8a49e3c8573186bc858441a95cc20e6eb",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "fd82b2f547ceee39ddc5f6267eb585b54b1f951615c6dca2cfcc7c4a5b345b4c",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "770cd48980eb6d61558c477556462cf99561ca1fa301f31a446149ee28c246a1",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "a81a7dadd1683df9366936916b19709943268336a3d92edcb187337032a5bad8",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
	}

	dummyTxs = []Tx{
		Tx{
			Id:      "1ec0bcc5b8847f038b1d8e87cb7932f08dc91c090cc764761b98627f359866c4",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "3173d5f88fef47b1392ddc08d3450ea8a49e3c8573186bc858441a95cc20e6eb",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "fd82b2f547ceee39ddc5f6267eb585b54b1f951615c6dca2cfcc7c4a5b345b4c",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "770cd48980eb6d61558c477556462cf99561ca1fa301f31a446149ee28c246a1",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "a81a7dadd1683df9366936916b19709943268336a3d92edcb187337032a5bad8",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "1ec0bcc5b8847f038b1d8e87cb7932f08dc91c090cc764761b98627f359866c4",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "3173d5f88fef47b1392ddc08d3450ea8a49e3c8573186bc858441a95cc20e6eb",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "fd82b2f547ceee39ddc5f6267eb585b54b1f951615c6dca2cfcc7c4a5b345b4c",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "770cd48980eb6d61558c477556462cf99561ca1fa301f31a446149ee28c246a1",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "1ec0bcc5b8847f038b1d8e87cb7932f08dc91c090cc764761b98627f359866c4",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "3173d5f88fef47b1392ddc08d3450ea8a49e3c8573186bc858441a95cc20e6eb",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "fd82b2f547ceee39ddc5f6267eb585b54b1f951615c6dca2cfcc7c4a5b345b4c",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "770cd48980eb6d61558c477556462cf99561ca1fa301f31a446149ee28c246a1",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "a81a7dadd1683df9366936916b19709943268336a3d92edcb187337032a5bad8",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "1ec0bcc5b8847f038b1d8e87cb7932f08dc91c090cc764761b98627f359866c4",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "3173d5f88fef47b1392ddc08d3450ea8a49e3c8573186bc858441a95cc20e6eb",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "fd82b2f547ceee39ddc5f6267eb585b54b1f951615c6dca2cfcc7c4a5b345b4c",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "770cd48980eb6d61558c477556462cf99561ca1fa301f31a446149ee28c246a1",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "1ec0bcc5b8847f038b1d8e87cb7932f08dc91c090cc764761b98627f359866c4",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "3173d5f88fef47b1392ddc08d3450ea8a49e3c8573186bc858441a95cc20e6eb",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "fd82b2f547ceee39ddc5f6267eb585b54b1f951615c6dca2cfcc7c4a5b345b4c",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "770cd48980eb6d61558c477556462cf99561ca1fa301f31a446149ee28c246a1",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "a81a7dadd1683df9366936916b19709943268336a3d92edcb187337032a5bad8",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "1ec0bcc5b8847f038b1d8e87cb7932f08dc91c090cc764761b98627f359866c4",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "3173d5f88fef47b1392ddc08d3450ea8a49e3c8573186bc858441a95cc20e6eb",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "fd82b2f547ceee39ddc5f6267eb585b54b1f951615c6dca2cfcc7c4a5b345b4c",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
		Tx{
			Id:      "770cd48980eb6d61558c477556462cf99561ca1fa301f31a446149ee28c246a1",
			Time:    "1562750685",
			Address: "mhmE6N5fg4g1KX9omfYtEkVfaMc1xzMSo6",
			Type:    "mined",
			Amount:  "1.265",
			Btn:     new(widget.Clickable),
		},
	}
)
