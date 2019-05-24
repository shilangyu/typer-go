package widgets

import (
	"github.com/jroimartin/gocui"
)

// Collection is a widget that groups other widgets
type Collection struct {
	name   string
	Title  string
	Center bool
	x, y   int
	w, h   int
}

// NewCollection initializes the Collection widget
// if center is true x and y becomes the center not start
func NewCollection(name, title string, center bool, x, y int, w, h int) *Collection {
	w--
	h--

	if center {
		x = x - w/2
		y = y - h/2
	}

	return &Collection{name, title, center, x, y, w, h}
}

// Name returns the widget name
func (w *Collection) Name() string {
	return w.name
}

// Coord returns the x and y of the widget
func (w *Collection) Coord() (int, int) {
	return w.x, w.y
}

// Size returns the width and height of the widget
func (w *Collection) Size() (int, int) {
	return w.w, w.h
}

// Layout renders the Collection widget
func (w *Collection) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = w.Title
	}

	return nil
}
