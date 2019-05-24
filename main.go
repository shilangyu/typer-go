package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typeracer-go/ui"
	"github.com/shilangyu/typeracer-go/utils"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	utils.Check(err)
	defer g.Close()

	err = ui.CreateWelcome(g)
	utils.Check(err)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
