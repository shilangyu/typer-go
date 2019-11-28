package ui

import (
	"github.com/rivo/tview"
	"github.com/shilangyu/typer-go/settings"
	"github.com/shilangyu/typer-go/utils"
)

// CreateSettings creates a screen with settings
func CreateSettings(app *tview.Application) error {
	settingsWi := tview.NewForm().
		AddDropDown(
			"highlight",
			[]string{settings.HighlightBackground.String(), settings.HighlightText.String()},
			int(settings.I.Highlight),
			func(option string, index int) {
				settings.I.Highlight = settings.Highlight(index)
				settings.Save()
			}).
		AddDropDown(
			"error display",
			[]string{settings.ErrorDisplayText.String(), settings.ErrorDisplayTyped.String()},
			int(settings.I.ErrorDisplay),
			func(option string, index int) {
				settings.I.ErrorDisplay = settings.ErrorDisplay(index)
				settings.Save()
			}).
		AddInputField(
			"texts path",
			settings.I.TextsPath,
			20,
			nil,
			func(text string) {
				settings.I.TextsPath = text
				settings.Save()
			}).
		AddButton("DONE", func() { utils.Check(CreateWelcome(app)) })

	layout := tview.NewFlex().AddItem(Center(34, 10, settingsWi), 0, 1, true)

	app.SetRoot(layout, true)
	keybindings(app, CreateWelcome)
	return nil
}
