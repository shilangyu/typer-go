package utils

import "strings"

// StringDimensions returns the width and height of a string
// where w = longest line, h = amount of lines
func StringDimensions(s string) (w, h int) {
	text := strings.Split(s, "\n")
	h = len(text)

	for _, line := range text {
		if len(line) > w {
			w = len(line)
		}
	}

	return
}
