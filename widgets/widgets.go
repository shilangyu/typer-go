package widgets

import "github.com/jroimartin/gocui"

// Widget describes the base of a widget
type Widget interface {
	gocui.Manager
	Name() string
	Coord() (int, int)
	Size() (int, int)
}
