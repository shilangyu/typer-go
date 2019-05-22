package widgets

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
)

// Menu is a widget that creates a vertical clickable menu
type Menu struct {
	Name     string
	Items    []string
	X, Y     int
	W, H     int
	Center   bool
	Arrows   bool
	OnChange func(i int)
	OnSubmit func(i int, g *gocui.Gui, v *gocui.View) error
	currItem int
}

// NewMenu initializes the Menu widget
// if arrows is true menu will be controlled through arrows as well
func NewMenu(name string, items []string, x, y int, center, arrows bool, onChange func(i int), onSubmit func(i int, g *gocui.Gui, v *gocui.View) error) *Menu {
	w, h := utils.StringDimensions(strings.Join(items, "\n"))
	w++
	h++

	if center {
		x = x - w/2
		y = y - h/2
	}

	return &Menu{name, items, x, y, w, h, center, arrows, onChange, onSubmit, 0}
}

// Init initializes the gocui side of things
func (w *Menu) Init(g *gocui.Gui) error {
	g.Mouse = true

	if err := g.SetKeybinding(w.Name, gocui.MouseLeft, gocui.ModNone, w.onMouse); err != nil {
		return err
	}

	if w.Arrows {
		if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, w.onArrow(1)); err != nil {
			return err
		}
		if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, w.onArrow(-1)); err != nil {
			return err
		}
		if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			return w.OnSubmit(w.currItem, g, v)
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
			v, err := g.View(w.Name)
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
		return w.OnSubmit(currItem, g, v)
	}
	return nil
}

// Layout renders the Menu widget
func (w *Menu) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.Name, w.X, w.Y, w.X+w.W, w.Y+w.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		for _, text := range w.Items {
			fmt.Fprintln(v, text)
		}
	}

	return nil
}
