package widgets

import (
	"github.com/jroimartin/gocui"
)

// Input is a widget allows for user input
type Input struct {
	name   string
	Text   string
	Frame  bool
	Center bool
	x, y   int
	w, h   int
}

// NewInput initializes the Input widget
// if frame is true a border is rendered
// if center is true x and y becomes the center not start
func NewInput(name string, frame, center bool, x, y int, w, h int) *Input {
	w--
	h--
	
	if center {
		x = x - w/2
		y = y - h/2
	}

	return &Input{name, "", frame, center, x, y, w, h}
}

// Name returns the widget name
func (w *Input) Name() string {
	return w.name
}

// Coord returns the x and y of the widget
func (w *Input) Coord() (int, int) {
	return w.x, w.y
}

// Size returns the width and height of the widget
func (w *Input) Size() (int, int) {
	return w.w, w.h
}

// Layout renders the Input widget
func (w *Input) Layout(g *gocui.Gui) error {
	g.Cursor = true

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Wrap = true

		v.Frame = w.Frame
		if _, err := g.SetCurrentView(w.name); err != nil {
			return err
		}
	}

	return nil
}

// // ChangeText changes the text within
// func (w *Text) ChangeText(s string) func(g *gocui.Gui) error {
// 	return func(g *gocui.Gui) error {
// 		v, err := g.View(w.name)
// 		if err != nil {
// 			return err
// 		}
// 		v.Clear()
// 		fmt.Fprint(v, s)

// 		return nil
// 	}
// }
