// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/common-nighthawk/go-figure"
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
	"github.com/shilangyu/typeracer-go/widgets"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	utils.Check(err)
	defer g.Close()

	w, h := g.Size()

	// main menu views
	sign := widgets.NewText("sign", figure.NewFigure("typeracer", "", false).String(), false, true, w/2, h/5)

	infoItems := utils.Center([]string{
		"Single player mode - test your typing skills offline!",
		"Multi player mode - battle against other typers",
		"Settings - change app settings",
		"Exit - exit the app",
	})
	info := widgets.NewText("info", infoItems[0], true, true, w/2, 3*h/4)

	menuItems := []string{"single player", "multi player", "settings", "exit"}
	menu := widgets.NewMenu("menu", utils.Center(menuItems), w/2, h/2, true, true, func(i int) {
		g.Update(info.ChangeText(infoItems[i]))
	})

	g.SetManager(sign, menu, info)

	err = menu.Init(g)
	utils.Check(err)

	err = keybindings(g)
	utils.Check(err)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
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
