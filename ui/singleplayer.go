package ui

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/shilangyu/typer-go/game"
	"github.com/shilangyu/typer-go/settings"
)

// CreateSingleplayer creates singleplayer screen widgets
func CreateSingleplayer(app *tview.Application) error {
	text, err := game.ChooseText()
	if err != nil {
		return err
	}
	state := game.NewState(text)

	statsWis := [...]*tview.TextView{
		tview.NewTextView().SetText("wpm: 0"),
		tview.NewTextView().SetText("time: 0s"),
	}

	pages := tview.NewPages().
		AddPage("modal", tview.NewModal().
			SetText("Play again?").
			SetBackgroundColor(tcell.ColorDefault).
			AddButtons([]string{"yes", "exit"}).
			SetDoneFunc(func(index int, label string) {
				switch index {
				case 0:
					CreateSingleplayer(app)
				case 1:
					CreateWelcome(app)
				}
			}), false, false)

	var textWis []*tview.TextView
	for _, word := range state.Words {
		textWis = append(textWis, tview.NewTextView().SetText(word).SetDynamicColors(true))
	}

	currInput := ""
	inputWi := tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorDefault)
	inputWi.
		SetChangedFunc(func(text string) {
			if state.StartTime.IsZero() {
				state.Start()
				go func() {
					ticker := time.NewTicker(100 * time.Millisecond)
					for range ticker.C {
						if state.CurrWord == len(state.Words) {
							ticker.Stop()
							return
						}
						app.QueueUpdateDraw(func() {
							statsWis[0].SetText(fmt.Sprintf("wpm: %.0f", state.Wpm()))
							statsWis[1].SetText(fmt.Sprintf("time: %.02fs", time.Since(state.StartTime).Seconds()))
						})
					}
				}()
			}

			if len(currInput) < len(text) {
				if len(text) > len(state.Words[state.CurrWord]) || state.Words[state.CurrWord][len(text)-1] != text[len(text)-1] {
					state.IncError()
				}
			}

			app.QueueUpdateDraw(func(i int) func() {
				return func() {
					textWis[i].SetText(paintDiff(state.Words[i], text))
				}
			}(state.CurrWord))

			if text == state.Words[state.CurrWord] {
				state.NextWord()
				if state.CurrWord == len(state.Words) {
					state.End()
					pages.ShowPage("modal")
				} else {
					inputWi.SetText("")
				}
			}

			currInput = text
		})

	layout := tview.NewFlex()
	statsFrame := tview.NewFlex().SetDirection(tview.FlexRow)
	statsFrame.SetBorder(true).SetBorderPadding(1, 1, 1, 1).SetTitle("STATS")
	for _, statsWi := range statsWis {
		statsFrame.AddItem(statsWi, 1, 1, false)
	}
	layout.AddItem(statsFrame, 0, 1, false)

	secondColumn := tview.NewFlex().SetDirection(tview.FlexRow)
	textsLayout := tview.NewFlex()
	for _, textWi := range textWis {
		textsLayout.AddItem(textWi, len(textWi.GetText(true)), 1, false)
	}
	textsLayout.SetBorder(true)
	secondColumn.AddItem(textsLayout, 0, 3, false)
	inputWi.SetBorder(true)
	secondColumn.AddItem(inputWi, 0, 1, true)
	layout.AddItem(secondColumn, 0, 3, true)

	pages.AddPage("game", layout, true, true).SendToBack("game")
	app.SetRoot(pages, true)

	keybindings(app, CreateWelcome)
	return nil
}

// paintDiff returns an tview-painted string displaying the difference
func paintDiff(toColor string, differ string) (colorText string) {
	var h string
	if settings.I.Highlight == settings.HighlightBackground {
		h = ":"
	}

	for i := range differ {
		if i >= len(toColor) || differ[i] != toColor[i] {
			colorText += "[" + h + "red]"
		} else {
			colorText += "[" + h + "green]"
		}

		switch settings.I.ErrorDisplay {
		case settings.ErrorDisplayText:
			colorText += string(differ[i])
		case settings.ErrorDisplayTyped:
			if i < len(toColor) {
				colorText += string(toColor[i])
			}
		}
	}
	colorText += "[-:-:-]"

	if len(differ) < len(toColor) {
		colorText += toColor[len(differ):]
	}

	return
}
