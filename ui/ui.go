package ui

import (
	"github.com/jroimartin/gocui"
)

func keybindings(g *gocui.Gui, goBack func(g *gocui.Gui) error) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if goBack != nil {
		if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			return goBack(g)
		}); err != nil {
			return err
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
