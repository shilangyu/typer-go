// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/common-nighthawk/go-figure"
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
	"github.com/shilangyu/typeracer-go/widgets"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	w, h := g.Size()

	// main menu views
	sign := widgets.NewText("sign", figure.NewFigure("typeracer", "", false).String(), true, true, w/2, h/5)
	menuItems := []string{"single player", "multi player", "settings", "exit"}
	menu := widgets.NewMenu("menu", utils.Center(menuItems), w/2, h/2, true, true, nil)
	g.SetManager(sign, menu)

	err = menu.Init(g)
	if err != nil {
		log.Panicln(err)
	}
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
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

var currMenu int

func showInfo(g *gocui.Gui, v *gocui.View) (err error) {
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
