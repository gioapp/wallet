package gwallet

import (
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/lyt"
	"github.com/gioapp/wallet/pkg/theme"
)

func (g *GioWallet) GetExplore() {
	//f, err := g.sh.ID()
	//checkError(err)
	//fv, ft, err := g.sh.Version()
	//checkError(err)
	//g.Status = Status{
	//	Title: "string",
	//	HostingSize: "uint",
	//PeerId:  f.ID,
	//Version: fv + " " + ft,
	//Gateway:   ,
	//Api:       "string",
	//Addresses: f.Addresses,
	//Pub:       f.PublicKey,
	//}
}

func (g *GioWallet) exploreHeader() func(gtx C) D {
	return func(gtx C) D {
		//return ContainerLayout(g.UI.Theme.Colors["PanelBg"], 10, 10, 10, 10, func(gtx C) D {
		//
		//})
		return D{}
	}
}

func (g *GioWallet) exploreBody() []func(gtx C) D {
	return []func(gtx C) D{
		g.twoPanels(10, 0, g.exploreLeft(), g.exploreRight()),
	}
}

//func (g *GioWallet) exploreBody() []func(gtx C) D {
//	retu//		ContainerLayout(g.UI.Theme.Colors["Silver"], 1, 1, 1, 1, ContainerLayout(g.UI.Theme.Colors["PanelBg"], 30, 30, 30, 30, func(gtx C) D {
//			gtx.Constraints.Min.X = gtx.Constraints.Max.X
//			return lyt.Format(gtx, "vflexb(middle,f(0.6,inset(5dp0dp5dp0dp,_)),f(0.4, inset(5dp0dp5dp0dp,_))))",
//				func(gtx C) D {
//					gtx.Constraints.Max.Y = gtx.Constraints.Min.Y
//					return g.6UI.logoIpld.Layout(gtx)
//				},
//				func(gtx C) D {
//					title := theme.Body(g.UI.Theme, "IPLD is the data model of the content-addressable web. It allows us to treat all hash-linked data structures as subsets of a unified information space, unifying all data models that link data with hashes as instances of IPLD.\n\nContent addressing through hashes has become a widely-used means of connecting data in distributed systems, from the blockchains that run your favorite cryptocurrencies, to the commits that back your code, to the web’s content at large. Yet, whilst all of these tools rely on some common primitives, their specific underlying data structures are not interoperable.\n\nEnter IPLD: a single namespace for all hash-inspired protocols. Through IPLD, links can be traversed across protocols, allowing you to explore data regardless of the underlying protocol.")
//					title.Alignment = text.Start
//					return title.Layout(gtx)
//				})
//		})))

//}
//}
func (g *GioWallet) exploreLeft() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "vflexb(start,r(inset(0dp0dp30dp0dp,_)),r(inset(0dp0dp80dp0dp,_))))",
			func(gtx C) D {
				title := theme.H3(g.UI.Theme, "Explore the Merkle Forest")
				title.Alignment = text.Start
				title.Color = helper.HexARGB(g.UI.Theme.Colors["Success"])
				return title.Layout(gtx)
			},
			func(gtx C) D {
				title := theme.Body(g.UI.Theme, "Paste a CID into the box to fetch the IPLD node it addresses, or choose a featured dataset.")
				title.Alignment = text.Start
				return title.Layout(gtx)
			})
	}
}

func (g *GioWallet) exploreRight() func(gtx C) D {
	return ContainerLayout(g.UI.Theme.Colors["PanelBg"], 10, 10, 10, 10, func(gtx C) D {
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
