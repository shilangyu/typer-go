package widgets

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
)

// Text is a widget that displays text
type Text struct {
	Name   string
	Text   string
	Frame  bool
	Center bool
	X, Y   int
	W, H   int
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

// Layout renders the Text widget
func (w *Text) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.Name, w.X, w.Y, w.X+w.W, w.Y+w.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = w.Frame
		fmt.Fprint(v, w.Text)
	}

	return nil
}
