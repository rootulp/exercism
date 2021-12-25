package ledger

import (
	"errors"
	"fmt"
	"math"
	"sort"
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

	output += header(locale)
	for _, entry := range sortEntries(entries) {
		output += formatEntry(locale, currency, entry)
	}
	return output, nil
}

func sortEntries(entries []Entry) (sorted []Entry) {
	sorted = make([]Entry, len(entries))
	copy(sorted, entries)
	sort.Slice(sorted, func(i int, j int) bool {
		a := entries[i]
		b := entries[j]
		if a.Date != b.Date {
			return a.Date < b.Date
		} else if a.Description != b.Description {
			return a.Description < b.Description
		} else {
			return a.Change < b.Change
		}
	})
	return sorted
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
	paddedChange := fmt.Sprintf("%03s", strconv.Itoa(absoluteValue(cents)))
	change += formatCurrencySymbol(locale, currency)
	change += formatNumber(locale, paddedChange)
	if isNegative(cents) {
		change = formatNegative(locale, change)
	}
	return change
}

func formatNumber(locale string, paddedChange string) (formatted string) {
	integralPart := formatIntegralPart(locale, paddedChange)
	decimalPart := paddedChange[len(paddedChange)-2:]
	return strings.Join([]string{integralPart, decimalPart}, locales[locale].DecimalPoint)
}

func formatCurrencySymbol(locale string, currency string) (formatted string) {
	if locale == "en-US" {
		return currencyToSymbol[currency]
	} else if locale == "nl-NL" {
		return currencyToSymbol[currency] + " "
	}
	panic("invalid locale")
}

func formatNegative(locale string, change string) (formatted string) {
	if locale == "en-US" {
		return fmt.Sprintf("(%s)", change)
	} else if locale == "nl-NL" {
		return fmt.Sprintf("%s-", change)
	}
	panic("invalid locale")
}

func formatIntegralPart(locale string, paddedChange string) (formatted string) {
	var parts []string
	remainder := paddedChange[:len(paddedChange)-2]
	for len(remainder) > 3 {
		parts = append([]string{remainder[len(remainder)-3:]}, parts...)
		remainder = remainder[:len(remainder)-3]
	}
	if len(remainder) > 0 {
		parts = append([]string{remainder}, parts...)
	}
	return strings.Join(parts, locales[locale].IntegralSeperator)
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
		seperatorA, seperatorB := entry.Date[4], entry.Date[7]
		if seperatorA != '-' || seperatorB != '-' {
			return false
		}
	}
	return true
}

func isNegative(cents int) bool {
	return cents < 0
}

func absoluteValue(n int) int {
	return int(math.Abs(float64(n)))
}
