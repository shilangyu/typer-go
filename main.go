package main

import (
	"log"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/typer-go/settings"
	"github.com/shilangyu/typer-go/ui"
	"github.com/shilangyu/typer-go/utils"
)

func main() {
	settings.H()
	time.Sleep(10 * time.Second)
	g, err := gocui.NewGui(gocui.OutputNormal)
	utils.Check(err)
	defer g.Close()

	err = ui.CreateWelcome(g)
	utils.Check(err)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
