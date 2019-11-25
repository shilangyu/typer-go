package ui

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/rivo/tview"
)

// CreateWelcome creates welcome screen widgets
func CreateWelcome(app *tview.Application) error {
	signWi := tview.NewTextView()
	fmt.Fprint(signWi, figure.NewFigure("typer-go", "", false).String())
	menuWi := tview.NewList().
		AddItem("single player", "test your typing skills offline!", 'a', nil).
		AddItem("multi player", "battle against other typers", 'b', nil).
		AddItem("stats", "TO BE RELEASED", 'c', nil).
		AddItem("settings", "change app settings", 'd', func() {
			CreateSettings(app)
		}).
		AddItem("exit", "exit the app", 'e', func() {
			app.Stop()
		})
	// switch i {
	// case 0:
	// 	utils.Check(CreateSingleplayer(g))
	// case 1:
	// 	utils.Check(CreateMultiplayerSetup(g))
	// }

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(signWi, 0, 1, false).
		AddItem(menuWi, 0, 1, true)

	app.SetRoot(layout, true)
	return nil
	//return keybindings(g, nil)
}
