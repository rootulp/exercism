package parsinglogfiles

import (
	"fmt"
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

func CountQuotedPasswords(lines []string) (count int) {
	for _, line := range lines {
		lowercase := strings.ToLower(line)
		regex, err := regexp.Compile(`".*password.*"`)
		if err != nil {
			panic(err)
		}
		if regex.Match([]byte(lowercase)) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex, err := regexp.Compile(`end-of-line\d*`)
	if err != nil {
		panic(err)
	}

	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) (result []string) {
	regex, err := regexp.Compile(`User\s+(\w*)`)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		if regex.MatchString(line) {
			match := regex.FindStringSubmatch(line)
			user := match[1]
			prefix := fmt.Sprintf(`[USR] %s`, user)
			line = prefix + " " + line
		}
		result = append(result, line)
	}
	return result
}
