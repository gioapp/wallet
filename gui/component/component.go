package component

import (
	"github.com/gioapp/wallet/pkg/gelook"
)

type DuoUIcomponent struct {
	Name    string
	Version string
	Theme   *gelook.DuoUItheme
	M       interface{}
	V       func()
	C       func()
}
