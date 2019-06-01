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
	return []string{"Background", "Text"}[e]
}
