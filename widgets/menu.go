package widgets

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
)

// Menu is a widget that creates a vertical clickable menu
type Menu struct {
	name     string
	Items    []string
	x, y     int
	w, h     int
	Center   bool
	Arrows   bool
	OnChange func(i int)
	OnSubmit func(i int)
	currItem int
}

// NewMenu initializes the Menu widget
// if arrows is true menu will be controlled through arrows as well
func NewMenu(name string, items []string, x, y int, center, arrows bool, onChange, onSubmit func(i int)) *Menu {
	w, h := utils.StringDimensions(strings.Join(items, "\n"))
	w++
	h++

	if center {
		x = x - w/2
		y = y - h/2
	}

	return &Menu{name, items, x, y, w, h, center, arrows, onChange, onSubmit, 0}
}

// Name returns the widget name
func (w * Menu) Name() string {
	return w.name
}

// Coord returns the x and y of the widget
func (w *Menu) Coord() (int, int) {
	return w.x, w.y
}

// Size returns the width and height of the widget
func (w *Menu) Size() (int, int) {
	return w.w, w.h
}

// Init initializes the gocui side of things
func (w *Menu) Init(g *gocui.Gui) error {
	g.Mouse = true

	if err := g.SetKeybinding(w.name, gocui.MouseLeft, gocui.ModNone, w.onMouse); err != nil {
		return err
	}

	if w.Arrows {
		if err := g.SetKeybinding(w.name, gocui.KeyArrowDown, gocui.ModNone, w.onArrow(1)); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyArrowUp, gocui.ModNone, w.onArrow(-1)); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			w.OnSubmit(w.currItem)
			return nil
		}); err != nil {
			return err
		}
	}

	return nil
}

// handles keystroke events
func (w *Menu) onArrow(change int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		w.currItem += change
		if w.currItem == -1 {
			w.currItem++
		} else if w.currItem == len(w.Items) {
			w.currItem--
		} else {
			v, err := g.View(w.name)
			if err != nil {
				return err
			}
			v.SetCursor(0, w.currItem)

			if w.OnChange != nil {
				w.OnChange(w.currItem)
			}
		}

		return nil
	}
}

// handles mouse event
func (w *Menu) onMouse(g *gocui.Gui, v *gocui.View) error {
	_, currItem := v.Cursor()
	if currItem != w.currItem {
		w.currItem = currItem
		if w.OnChange != nil {
			w.OnChange(w.currItem)
		}
	} else {
		w.OnSubmit(currItem)
	}
	return nil
}

// Layout renders the Menu widget
func (w *Menu) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.SetCurrentView(w.name)

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		for _, text := range w.Items {
			fmt.Fprintln(v, text)
		}
	}

	return nil
}
