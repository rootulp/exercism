package ledger

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func isValidCurrency(currency string) bool {
	return currency == "USD" || currency == "EUR"
}

func isValidLocale(locale string) bool {
	return locale == "en-US" || locale == "nl-NL"
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

func formatDescription(description string) string {
	if len(description) > 25 {
		return description[:22] + "..."
	}
	return description + strings.Repeat(" ", 25-len(description))
}
func formatDate(locale string, date string) string {
	year, month, day := date[0:4], date[5:7], date[8:10]

	if locale == "nl-NL" {
		return fmt.Sprintf("%v-%v-%v", day, month, year)
	} else if locale == "en-US" {
		return fmt.Sprintf("%v/%v/%v", month, day, year)
	}
	panic("invalid locale")
}

func header(locale string) (output string) {
	if locale == "nl-NL" {
		return "Datum" +
			strings.Repeat(" ", 10-len("Datum")) +
			" | " +
			"Omschrijving" +
			strings.Repeat(" ", 25-len("Omschrijving")) +
			" | " + "Verandering" + "\n"
	} else if locale == "en-US" {
		return "Date" +
			strings.Repeat(" ", 10-len("Date")) +
			" | " +
			"Description" +
			strings.Repeat(" ", 25-len("Description")) +
			" | " + "Change" + "\n"
	}
	panic("invalid locale")
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
	co := make(chan struct {
		i int
		s string
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			de := formatDescription(entry.Description)
			d := formatDate(locale, entry.Date)
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				}
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
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
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				}
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
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
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			}
			co <- struct {
				i int
				s string
			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-len(a)) + a + "\n"}
		}(i, et)
	}
	temp := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		temp[v.i] = v.s
	}
	for i := 0; i < len(entriesCopy); i++ {
		output += temp[i]
	}
	return output, nil
}
