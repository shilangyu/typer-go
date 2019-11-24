package main

import (
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/shilangyu/typer-go/ui"
	"github.com/shilangyu/typer-go/utils"
)

func main() {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	// tview.Styles.PrimaryTextColor = tcell.ColorBlack

	app := tview.NewApplication()
	defer app.Stop()

	app.SetBeforeDrawFunc(func(s tcell.Screen) bool {
		s.Clear()
		return false
	})

	err := ui.CreateWelcome(app)
	utils.Check(err)

	if err := app.Run(); err != nil {
		log.Panicln(err)
	}
}
