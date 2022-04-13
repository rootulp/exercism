package parsinglogfiles

import (
	"regexp"
	"strings"
)

var validPrefixes = []string{
	"[TRC]",
	"[DBG]",
	"[INF]",
	"[WRN]",
	"[ERR]",
	"[FTL]",
}

func IsValidLine(text string) bool {
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	regex, err := regexp.Compile(`<[-*~=]*>`)
	if err != nil {
		panic(err)
	}

	tabSeperated := regex.ReplaceAllString(text, "\t")
	return strings.Split(tabSeperated, "\t")
}

func CountQuotedPasswords(lines []string) int {
	panic("Please implement the CountQuotedPasswords function")
}

func RemoveEndOfLineText(text string) string {
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	panic("Please implement the TagWithUserName function")
}
