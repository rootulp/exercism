package ledger

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type Locale struct {
	Date              string
	Description       string
	Change            string
	DateSeperator     string
	DecimalPoint      string
	IntegralSeperator string
}

var locales = map[string]Locale{
	"en-US": {
		Date:              "Date",
		Description:       "Description",
		Change:            "Change",
		DateSeperator:     "/",
		DecimalPoint:      ".",
		IntegralSeperator: ",",
	},
	"nl-NL": {
		Date:              "Datum",
		Description:       "Omschrijving",
		Change:            "Verandering",
		DateSeperator:     "-",
		DecimalPoint:      ",",
		IntegralSeperator: ".",
	},
}

var currencyToSymbol = map[string]string{
	"USD": "$",
	"EUR": "€",
}

func FormatLedger(currency string, locale string, entries []Entry) (output string, e error) {
	if !isValidCurrency(currency) {
		return "", errors.New("invalid currency")
	}
	if !isValidLocale(locale) {
		return "", errors.New("invalid locale")
	}
	if !isValidDate(entries) {
		return "", errors.New("invalid date")
	}

	entriesCopy := append([]Entry{}, entries...)

	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}

	output += header(locale)
	for _, entry := range entriesCopy {
		output += formatEntry(locale, currency, entry)
	}
	return output, nil
}

func formatEntry(locale string, currency string, entry Entry) (formatted string) {
	date := formatDate(locale, entry.Date)
	description := formatDescription(entry.Description)
	change := formatChange(locale, currency, entry.Change)

	// This conditional is necessary because expected output is aligned
	// differently for negative vs. non-negative values
	if isNegative(entry.Change) {
		return fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, change)
	} else {
		return fmt.Sprintf("%-10s | %-25s | %12s\n", date, description, change)
	}
}

func header(locale string) (output string) {
	return fmt.Sprintf("%-10s | %-25s | %s\n", locales[locale].Date, locales[locale].Description, locales[locale].Change)
}

func formatDate(locale string, date string) string {
	year, month, day := date[0:4], date[5:7], date[8:10]

	if locale == "nl-NL" {
		return strings.Join([]string{day, month, year}, locales[locale].DateSeperator)
	} else if locale == "en-US" {
		return strings.Join([]string{month, day, year}, locales[locale].DateSeperator)
	}
	panic("invalid locale")
}

// formatDescription will ellipsize the description if it is longer than 25
// characters
func formatDescription(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	}
	return description
}

func formatChange(locale string, currency string, cents int) (change string) {
	isNegative := isNegative(cents)
	absoluteValueCents := int(math.Abs(float64(cents)))
	change += currencyToSymbol[currency]
	if locale == "nl-NL" {
		change += " "
		centsStr := fmt.Sprintf("%03s", strconv.Itoa(absoluteValueCents))
		change += formatIntegralPart(centsStr, locale)
		change += locales[locale].DecimalPoint
		change += centsStr[len(centsStr)-2:]
		if isNegative {
			// Append `-`
			change = fmt.Sprintf("%s-", change)
		}
	} else if locale == "en-US" {
		centsStr := fmt.Sprintf("%03s", strconv.Itoa(absoluteValueCents))
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			change += parts[i] + locales[locale].IntegralSeperator
		}
		change = change[:len(change)-1]
		change += locales[locale].DecimalPoint
		change += centsStr[len(centsStr)-2:]
		if isNegative {
			// Surround with parenthesis
			change = fmt.Sprintf("(%s)", change)
		}
	}
	return change
}

func formatIntegralPart(centsStr string, locale string) (formatted string) {
	integralRemainder := centsStr[:len(centsStr)-2]
	var integralParts []string
	for len(integralRemainder) > 3 {
		integralParts = append([]string{integralRemainder[len(integralRemainder)-3:]}, integralParts...)
		integralRemainder = integralRemainder[:len(integralRemainder)-3]
	}
	if len(integralRemainder) > 0 {
		integralParts = append([]string{integralRemainder}, integralParts...)
	}
	return strings.Join(integralParts, locales[locale].IntegralSeperator)
}

func isValidCurrency(currency string) bool {
	_, ok := currencyToSymbol[currency]
	return ok
}

func isValidLocale(locale string) bool {
	_, ok := locales[locale]
	return ok
}

func isValidDate(entries []Entry) bool {
	for _, entry := range entries {
		if len(entry.Date) != 10 {
			return false
		}
		_, d2, _, d4, _ := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
		if d2 != '-' || d4 != '-' {
			return false
		}
	}
	return true
}

func isNegative(cents int) bool {
	return cents < 0
}
