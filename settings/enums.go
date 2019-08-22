package settings

// Highlight describes how the text should be highlighted
type Highlight int

const (
	// HighlightBackground says the text should have a background highlight
	HighlightBackground Highlight = iota
	// HighlightText says the text should have a text highlight
	HighlightText
)

func (e Highlight) String() string {
	return []string{"background", "text"}[e]
}

// ErrorDisplay describes how incorrectly typed letters should be displayed
type ErrorDisplay int

const (
	// ErrorDisplayTyped says the typed letter should be displayed
	ErrorDisplayTyped ErrorDisplay = iota
	// ErrorDisplayText says the text letter should be displayed
	ErrorDisplayText
)

func (e ErrorDisplay) String() string {
	return []string{"typed", "text"}[e]
}
