package markdown

import (
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) (html string) {
	markdown = handleBold(markdown)
	markdown = handleItalic(markdown)
	return wrapParagraph(markdown)
}

func wrapParagraph(s string) string {
	return "<p>" + s + "</p>"
}

func handleBold(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	return markdown
}

func handleItalic(markdown string) (result string) {
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	return markdown
}
