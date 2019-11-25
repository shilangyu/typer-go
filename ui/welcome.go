package ui

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/rivo/tview"
	"github.com/shilangyu/typer-go/utils"
)

// CreateWelcome creates welcome screen widgets
func CreateWelcome(app *tview.Application) error {
	signWi := tview.NewTextView().SetText(figure.NewFigure("typer-go", "", false).String())
	menuWi := tview.NewList().
		AddItem("single player", "test your typing skills offline!", 'a', func() {
			err := CreateSingleplayer(app)
			utils.Check(err)
		}).
		AddItem("multi player", "battle against other typers", 'b', nil).
		AddItem("stats", "TO BE RELEASED", 'c', nil).
		AddItem("settings", "change app settings", 'd', func() {
			err := CreateSettings(app)
			utils.Check(err)
		}).
		AddItem("exit", "exit the app", 'e', func() {
			app.Stop()
		})
	// switch i {
	// case 1:
	// 	utils.Check(CreateMultiplayerSetup(g))
	// }

	layout := tview.NewGrid().
		SetRows(10, 10, 0, 1).
		SetColumns(10, 5, 0, 5, 10).
		AddItem(signWi, 1, 1, 1, 3, 0, 0, false).
		AddItem(menuWi, 2, 2, 1, 1, 0, 0, true)

	app.SetRoot(layout, true)
	return nil
}
