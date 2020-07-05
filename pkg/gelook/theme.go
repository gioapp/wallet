// SPDX-License-Identifier: Unlicense OR MIT

package gelook

import (
	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"

	"github.com/gioapp/wallet/pkg/fonts"
)

type DuoUItheme struct {
	Shaper        text.Shaper
	TextSize      unit.Value
	Colors        map[string]string
	Fonts         map[string]text.Typeface
	Icons         map[string]*DuoUIicon
	scrollBarSize int
}

func init() {
	fonts.Register()
}
func NewDuoUItheme() *DuoUItheme {
	t := &DuoUItheme{
		Shaper: font.Default(),
	}
	t.Colors = NewDuoUIcolors()
	t.Fonts = NewDuoUIfonts()
	t.TextSize = unit.Sp(16)
	t.Icons = NewDuoUIicons()
	return t
}

func NewDuoUIfonts() (f map[string]text.Typeface) {
	f = make(map[string]text.Typeface)
	f["Primary"] = "bariol"
	f["Secondary"] = "plan9"
	f["Mono"] = "go"
	return f
}

func (t *DuoUItheme) ChangeLightDark() {
	light := t.Colors["Light"]
	dark := t.Colors["Dark"]
	lightGray := t.Colors["LightGrayIII"]
	darkGray := t.Colors["DarkGrayII"]
	t.Colors["Light"] = dark
	t.Colors["Dark"] = light
	t.Colors["LightGrayIII"] = darkGray
	t.Colors["DarkGrayII"] = lightGray
}
