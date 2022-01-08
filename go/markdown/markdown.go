package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) (html string) {
	markdown = handleBold(markdown)
	markdown = handleItalic(markdown)
	if isHeader(markdown) {
		return handleHeader(markdown)
	}
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

func handleItalic(markdown string) string {
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	return markdown
}

func isHeader(line string) bool {
	return strings.HasPrefix(line, "#")
}

func handleHeader(line string) string {
	if strings.HasPrefix(line, "######") {
		return handleHeaderWithLevel(line, 6)
	}
	if strings.HasPrefix(line, "#####") {
		return handleHeaderWithLevel(line, 5)
	}
	if strings.HasPrefix(line, "####") {
		return handleHeaderWithLevel(line, 4)
	}
	if strings.HasPrefix(line, "###") {
		return handleHeaderWithLevel(line, 3)
	}
	if strings.HasPrefix(line, "##") {
		return handleHeaderWithLevel(line, 2)
	}
	if strings.HasPrefix(line, "#") {
		return handleHeaderWithLevel(line, 1)
	}
	panic("handleHeader called with a line that doesn't contain a header")
}

func handleHeaderWithLevel(markdown string, level int) string {
	return fmt.Sprintf("<h%d>%s</h%d>", level, markdown[level+1:], level)
}
