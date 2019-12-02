package ui

import (
	"github.com/rivo/tview"
	"github.com/shilangyu/typer-go/utils"
)

// CreateWelcome creates welcome screen widgets
func CreateWelcome(app *tview.Application) error {
	const welcomeSign = `
 _                                          
| |_ _   _ _ __   ___ _ __       __ _  ___  
| __| | | | '_ \ / _ \ '__|____ / _  |/ _ \ 
| |_| |_| | |_) |  __/ | |_____| (_| | (_) |
 \__|\__, | .__/ \___|_|        \__, |\___/ 
     |___/|_|                   |___/       
`
	signWi := tview.NewTextView().SetText(welcomeSign)
	menuWi := tview.NewList().
		AddItem("single player", "test your typing skills offline!", 0, func() {
			utils.Check(CreateSingleplayer(app))
		}).
		AddItem("multi player", "battle against other typers", 0, func() {
			utils.Check(CreateMultiplayerSetup(app))
		}).
		AddItem("stats", "TO BE RELEASED", 0, nil).
		AddItem("settings", "change app settings", 0, func() {
			utils.Check(CreateSettings(app))
		}).
		AddItem("exit", "exit the app", 0, func() {
			app.Stop()
		})

	signW, signH := utils.StringDimensions(welcomeSign)
	menuW, menuH := 32, menuWi.GetItemCount()*2
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(Center(signW, signH, signWi), 0, 1, false).
		AddItem(Center(menuW, menuH, menuWi), 0, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)

	app.SetRoot(layout, true)
	return nil
}
