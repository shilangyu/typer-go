package ui

import (
	"github.com/rivo/tview"
	"github.com/shilangyu/typer-go/settings"
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
			10,
			nil,
			func(text string) {
				settings.I.TextsPath = text
				settings.Save()
			}).
		AddButton("OK", func() { CreateWelcome(app) })

	// infoItems := utils.Center([]string{
	// 	"How your text should be highlighted",
	// 	"What part should be displayed when you type an error",
	// 	"Path to your typer texts (separated with a new line)",
	// })

	app.SetRoot(settingsWi, true)
	return nil
}
