package widgets

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
)

// Text is a widget that displays text
type Text struct {
	name   string
	Text   string
	Frame  bool
	Center bool
	x, y   int
	w, h   int
}

// NewText initializes the Text widget
// if frame is true a border is rendered
// if center is true x and y becomes the center not start
func NewText(name, text string, frame, center bool, x, y int) *Text {
	w, h := utils.StringDimensions(text)
	w++
	h++

	if center {
		x = x - w/2
		y = y - h/2
	}

	return &Text{name, text, frame, center, x, y, w, h}
}

// Name returns the widget name
func (w * Text) Name() string {
	return w.name
}

// Coord returns the x and y of the widget
func (w *Text) Coord() (int, int) {
	return w.x, w.y
}

// Size returns the width and height of the widget
func (w *Text) Size() (int, int) {
	return w.w, w.h
}

// Layout renders the Text widget
func (w *Text) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = w.Frame
		fmt.Fprint(v, w.Text)
	}

	return nil
}

// ChangeText changes the text within
func (w *Text) ChangeText(s string) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		v, err := g.View(w.name)
		if err != nil {
			return err
		}
		v.Clear()
		fmt.Fprint(v, s)

		return nil
	}
}
