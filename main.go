// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/common-nighthawk/go-figure"
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/widgets"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Mouse = true

	w, h := g.Size()

	// main menu views
	text := widgets.NewText("sign", figure.NewFigure("typeracer", "", false).String(), true, true, w/2, h/5)
	g.SetManager(text)
	// g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	w, h := g.Size()

	xoff, yoff := 33, 7
	xoff, yoff = 7, 3
	if v, err := g.SetView("menu", w/2-xoff, h/2-yoff, w/2+xoff, h/2+yoff-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "single player")
		fmt.Fprintln(v, "multi player ")
		fmt.Fprintln(v, "  settings   ")
		fmt.Fprintln(v, "    exit     ")
	}

	xoff, yoff = 27, 1
	if v, err := g.SetView("info", w/2-xoff, h/2-yoff+20, w/2+xoff, h/2+yoff+20); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Single player mode - test your typing skills offline!")
	}
	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("menu", gocui.MouseLeft, gocui.ModNone, showInfo); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, downMenu); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, upMenu); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

var currMenu int

func downMenu(g *gocui.Gui, v *gocui.View) error {
	if currMenu != 3 {
		currMenu++
		return showInfo(g, nil)
	}
	return nil
}

func upMenu(g *gocui.Gui, v *gocui.View) error {
	if currMenu != 0 {
		currMenu--
		return showInfo(g, nil)
	}
	return nil
}

func showInfo(g *gocui.Gui, v *gocui.View) (err error) {
	if v != nil {
		_, currMenu = v.Cursor()
	}

	v, _ = g.SetCurrentView("menu")
	if err != nil {
		return err
	}
	v.SetCursor(0, currMenu)

	v, err = g.SetCurrentView("info")
	if err != nil {
		return err
	}

	desc := [...]string{
		"Single player mode - test your typing skills offline!",
		"Multi player mode - battle against other typers",
		"Settings - change app settings",
		"Exit - exit the app",
	}[currMenu]
	v.Clear()
	fmt.Fprintln(v, desc)
	return nil
}
