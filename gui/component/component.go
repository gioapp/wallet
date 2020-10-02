package component

import (
	"github.com/gioapp/gel/theme"
)

type DuoUIcomponent struct {
	Name    string
	Version string
	Theme   *theme.DuoUItheme
	M       interface{}
	V       func()
	C       func()
}
