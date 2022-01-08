package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) (html string) {
	markdown = handleBold(markdown)
	markdown = handleItalic(markdown)
	lines := strings.Split(markdown, "\n")
	htmlLines := []string{}
	for _, line := range lines {
		htmlLines = append(htmlLines, render(line))
	}
	result := strings.Join(htmlLines, "")
	return handleUnorderedList(result)
}

func render(line string) (html string) {
	if isHeader(line) {
		return handleHeader(line)
	}
	if isListItem(line) {
		return handleListItem(line)
	}
	return handleParagraph(line)
}

func isListItem(line string) bool {
	return strings.HasPrefix(line, "*")
}

func isHeader(line string) bool {
	return strings.HasPrefix(line, "#")
}

func handleListItem(line string) string {
	return fmt.Sprintf("<li>%s</li>", line[2:])
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

func handleParagraph(line string) string {
	return fmt.Sprintf("<p>%s</p>", line)
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

func handleUnorderedList(line string) string {
	// HACKHACK this assumes there is only one consecutive unordered list in the input
	line = strings.Replace(line, "<li>", "<ul><li>", 1)
	line = replaceLast(line, "</li>", "</li></ul>", 1)
	return line

}

func replaceLast(line string, search string, replacement string, occurences int) string {
	i := strings.LastIndex(line, search)
	if i == -1 {
		return line
	}
	return line[:i] + strings.Replace(line[i:], search, replacement, 1)
}
