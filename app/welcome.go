package gwallet

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

var (
	apiAddressInput = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func (g *GioWallet) GetWelcome() {
	//f, err := g.sh.ID()
	//checkError(err)
	//fv, ft, err := g.sh.Version()
	//checkError(err)
	//g.Status = Status{
	//	Title: "string",
	//	//HostingSize: "uint",
	//	PeerId:  f.ID,
	//	Version: fv + " " + ft,
	//	//Gateway:   ,
	//	//Api:       "string",
	//	Addresses: f.Addresses,
	//	Pub:       f.PublicKey,
	//}
}

func (g *GioWallet) welcomeHeader() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], 10, 10, 10, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		helper.Fill(gtx, helper.HexARGB(g.UI.Theme.Colors["PanelBg"]))
		return D{}
	})
}

func (g *GioWallet) welcomeBody() []func(gtx C) D {
	return []func(gtx C) D{
		g.twoPanels(10, 0, g.welcomeLeft(), g.welcomeRight()),
	}
}

func (g *GioWallet) welcomeLeft() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflexb(start,r(_),r(_))",
			g.welcomeLeftTop(),
			g.welcomeLeftBottom())
	}
}

func (g *GioWallet) welcomeLeftTop() func(gtx C) D {
	var d func(gtx C) D
	var tt bool
	//if g.sh.IsUp() {

	if tt {
		d = func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "vflexb(start,r(inset(0dp0dp30dp0dp,_)),r(inset(0dp0dp80dp0dp,_)))",
				func(gtx C) D {
					title := theme.H3(g.UI.Theme, "Connected to IPFS")
					title.Alignment = text.Start
					title.Color = helper.HexARGB(g.UI.Theme.Colors["Success"])
					return title.Layout(gtx)
				},
				func(gtx C) D {
					title := theme.Body(g.UI.Theme, "Now, it’s time for you to explore your node. Head to Files page to manage and share your files, or explore the Merkle Forest of peer-hosted hash-linked data via IPLD explorer.\n\nYou can always come back to this address to change the IPFS node you're connected to.")
					title.Alignment = text.Start
					return title.Layout(gtx)

				})

		}
	} else {
		d = func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "vflexb(start,r(inset(0dp0dp30dp0dp,_)),r(inset(0dp0dp80dp0dp,_)),r(inset(0dp0dp30dp0dp,_)),r(inset(0dp0dp80dp0dp,_))r(inset(0dp0dp30dp0dp,_)))",
				func(gtx C) D {
					title := theme.H3(g.UI.Theme, "Is your IPFS daemon running?")
					title.Alignment = text.Start
					title.Color = helper.HexARGB(g.UI.Theme.Colors["Check"])
					return title.Layout(gtx)
				},
				func(gtx C) D {
					title := theme.Body(g.UI.Theme, "Failed to connect to the API.")
					title.Alignment = text.Start
					return title.Layout(gtx)
				},
				func(gtx C) D {
					title := theme.Body(g.UI.Theme, "Start an IPFS daemon in a terminal:")
					title.Alignment = text.Start
					return title.Layout(gtx)
				},
				func(gtx C) D {
					title := theme.Body(g.UI.Theme, "SHEL $ ipfs daemon Initializing daemon... API server listening on /ip4/127.0.0.1/tcp/5001")
					title.Alignment = text.Start
					return title.Layout(gtx)
				},
				func(gtx C) D {
					title := theme.Body(g.UI.Theme, "For more info on how to get started with IPFS you can read the guide.")
					title.Alignment = text.Start
					return title.Layout(gtx)
				})

		}
	}
	return d
}

func (g *GioWallet) welcomeLeftBottom() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflexb(start,r(inset(0dp0dp10dp0dp,_)),r(inset(0dp0dp10dp0dp,_)),r(inset(0dp0dp10dp0dp,_)),r(inset(0dp0dp0dp0dp,_)),r(inset(0dp0dp0dp0dp,_)))",
			func(gtx C) D {
				title := theme.H3(g.UI.Theme, "Is your API on a port other than 5001?")
				title.Alignment = text.Start
				title.Color = helper.HexARGB(g.UI.Theme.Colors["Check"])
				return title.Layout(gtx)
			},
			func(gtx C) D {
				title := theme.Body(g.UI.Theme, "If your IPFS node is configured with a custom API address, please set it here.")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				title := theme.Small(g.UI.Theme, "API ADDRESS")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			ContainerLayout(g.UI.Theme.Colors["White"], g.UI.Theme.Colors["White"], g.UI.Theme.Colors["White"], 8, 8, 8, func(gtx C) D {
				gtx.Constraints.Min.X = 430
				e := material.Editor(g.UI.Theme.T, apiAddressInput, "Api address")
				return e.Layout(gtx)
			}),
			func(gtx C) D {
				e := material.Button(g.UI.Theme.T, browseBtn, "Submit")
				e.Inset = layout.Inset{
					Top:    unit.Dp(4),
					Right:  unit.Dp(4),
					Bottom: unit.Dp(4),
					Left:   unit.Dp(4),
				}
				e.CornerRadius = unit.Dp(4)
				return e.Layout(gtx)
			})

	}
}

func (g *GioWallet) welcomeRight() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], g.UI.Theme.Colors["PanelBg"], 10, 10, 10, func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflexb(start,r(inset(0dp0dp0dp0dp,_)),r(inset(0dp0dp0dp0dp,_)))",
			func(gtx C) D {
				title := theme.H3(g.UI.Theme, "What is IPFS?")
				title.Alignment = text.Start
				return title.Layout(gtx)
			},
			func(gtx C) D {
				title := theme.Body(g.UI.Theme, "IPFS is a protocol that defines a content-addressed file system, coordinates content delivery and combines ideas from Kademlia, BitTorrent, Git and more.\n\nIPFS is a filesystem. It has directories and files and mountable filesystem via FUSE.\n\nIPFS is a web. Files are accessible via HTTP gateways like https://ipfs.io. Browsers can be extended to use the ipfs:// scheme directly, and hash-addressed content guarantees authenticity.\n\nIPFS is p2p. It supports worldwide peer-to-peer file transfers with a completely decentralized architecture and no central point of failure.\n\nIPFS is a CDN. Add a file to your local repository, and it's now available to the world with cache-friendly content-hash addressing and BitTorrent-like bandwidth distribution.")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	})
}
