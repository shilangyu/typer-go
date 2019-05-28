package ui

import (
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
)

// CreateSingleplayer creates welcome screen widgets
func CreateSingleplayer(g *gocui.Gui) error {
	var currWord int

	w, h := g.Size()

	statsFrameWi := widgets.NewCollection("singleplayer-stats", "STATS", false, 0, 0, w/5, h)

	statWis := []*widgets.Text{
		widgets.NewText("singleplayer-stats-wpm", "wpm: 0", false, false, 2, 1),
		widgets.NewText("singleplayer-stats-time", "time: 0", false, false, 2, 2),
	}

	words := strings.Split("Cock and balls", " ")
	points := organiseText(words, 4*w/5)
	var textWis []*widgets.Text
	for i, p := range points {
		words[i] += " "
		textWis = append(textWis, widgets.NewText("singleplayer-text-"+strconv.Itoa(i), words[i], false, false, w/5+1+p.x, p.y))
	}

	var inputWi *widgets.Input
	inputWi = widgets.NewInput("singleplayer-input", true, false, w/5+1, h-h/6, w-w/5-1, h/6, func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
		b := v.Buffer()[:len(v.Buffer())-1]

		if key == gocui.KeySpace && b == words[currWord] {
			currWord++
			g.Update(inputWi.ChangeText(""))
		} else {
			ansiWord := ""
			for i := range b {
				if i >= len(words[currWord]) || b[i] != words[currWord][i] {
					ansiWord += "\u001b[31m"
				} else {
					ansiWord += "\u001b[32m"
				}
				ansiWord += string(b[i])
			}
			ansiWord += "\u001b[0m"

			if len(b) < len(words[currWord]) {
				ansiWord += words[currWord][len(b):]
			}

			g.Update(textWis[currWord].ChangeText(ansiWord))
		}
	})

	var wis []gocui.Manager
	wis = append(wis, statsFrameWi)
	for _, stat := range statWis {
		wis = append(wis, stat)
	}
	for _, text := range textWis {
		wis = append(wis, text)
	}
	wis = append(wis, inputWi)

	g.SetManager(wis...)

	if err := keybindings(g); err != nil {
		return err
	}

	return nil
}

func organiseText(words []string, lineLength int) (points []struct{ x, y int }) {
	x, y := 0, 0

	for _, word := range words {
		if x+len(word) > lineLength {
			y++
			x = 0
		}
		points = append(points, struct{ x, y int }{x, y})
		x += len(word) + 1
	}

	return
}
