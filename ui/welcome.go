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
