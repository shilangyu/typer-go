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

	textFrameWi := widgets.NewCollection("singleplayer-text", "", false, w/5+1, 0, 4*w/5, 5*h/6+1)

	words := strings.Split("Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls Cock and balls", " ")
	points := organiseText(words, 4*w/5-2)
	var textWis []*widgets.Text
	for i, p := range points {
		words[i] += " "
		textWis = append(textWis, widgets.NewText("singleplayer-text-"+strconv.Itoa(i), words[i], false, false, w/5+1+p.x, p.y))
	}

	var inputWi *widgets.Input
	inputWi = widgets.NewInput("singleplayer-input", true, false, w/5+1, h-h/6, w-w/5-1, h/6, func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
		if key == gocui.KeyEnter || len(v.Buffer()) == 0 && ch == 0 {
			return false
		}
		gocui.DefaultEditor.Edit(v, key, ch, mod)

		b := v.Buffer()[:len(v.Buffer())-1]

		ansiWord := wordsDiff(words[currWord], b)

		g.Update(textWis[currWord].ChangeText(ansiWord))

		if key == gocui.KeySpace && b == words[currWord] {
			currWord++
			g.Update(inputWi.ChangeText(""))
		}

		return false
	})

	var wis []gocui.Manager
	wis = append(wis, statsFrameWi)
	for _, stat := range statWis {
		wis = append(wis, stat)
	}
	wis = append(wis, textFrameWi)
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

// takes a slice of words and length of a line
// retuns xs and ys of the words on a plane
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

// adds ANSI colors to indicate diff
func wordsDiff(toColor, differ string) (ansiWord string) {
	for i := range differ {
		if i >= len(toColor) || differ[i] != toColor[i] {
			ansiWord += "\u001b[41m"
		} else {
			ansiWord += "\u001b[42m"
		}
		ansiWord += string(differ[i])
	}
	ansiWord += "\u001b[0m"

	if len(differ) < len(toColor) {
		ansiWord += toColor[len(differ):]
	}

	return
}
