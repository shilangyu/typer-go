package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/utils"
	"github.com/shilangyu/typeracer-go/ui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	utils.Check(err)
	defer g.Close()

	err = ui.CreateWelcome(g)
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
