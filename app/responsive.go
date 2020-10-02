package gwallet

func (g *GioWallet) Resposnsivity() {
	g.setMode(g.UI.Context.Constraints.Max.X)
	r := make(map[string]interface{})

	r["MainLayout"] = "vflexb(start,f(1,_),r(_))"
	r["NavSize"] = 255
	r["Logo"] = g.UI.Theme.Icons["Logo"]
	//r["xxxx"] = xxx
	//r["xxxx"] = xxx
	//r["xxxx"] = xxx
	//r["xxxx"] = xxx
	g.UI.res.mod = r
	return
}

func (g *GioWallet) setMode(w int) {
	switch w {
	case w < 640:

	}

}
