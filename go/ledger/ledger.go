package ledger

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	symbol, err := currencySymbol(currency)
	if err != nil {
		return "", err
	}

	header, err := generateHeader(locale)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.WriteString(header)

	if len(entries) == 0 {
		return builder.String(), nil
	}

	localEntries := make([]Entry, len(entries))
	copy(localEntries, entries)
	sort.Slice(localEntries, func(i, j int) bool {
		if localEntries[i].Date == localEntries[j].Date {
			return localEntries[i].Change < localEntries[j].Change
		}
		return localEntries[i].Date < localEntries[j].Date
	})

	for _, entry := range localEntries {
		entryDate, err := parseEntryDate(locale, entry.Date)
		if err != nil {
			return "", err
		}

		entryDescription := entry.Description
		if len(entryDescription) > 25 {
			entryDescription = entryDescription[:22] + "..."
		}

		entryChange, err := formatChange(locale, symbol, entry.Change)
		if err != nil {
			return "", err
		}

		line := fmt.Sprintf("%10s | %-25s | %13s\n", entryDate, entryDescription, entryChange)
		builder.WriteString(line)
	}
	return builder.String(), nil
}

func generateHeader(locale string) (string, error) {
	var date, description, change string
	switch locale {
	case "en-US":
		date = "Date"
		description = "Description"
		change = "Change"
	case "nl-NL":
		date = "Datum"
		description = "Omschrijving"
		change = "Verandering"
	default:
		return "", fmt.Errorf("unknown locale")
	}
	return date +
		strings.Repeat(" ", 10-len(date)) +
		" | " +
		description +
		strings.Repeat(" ", 25-len(description)) +
		" | " + change + "\n", nil
}

func parseEntryDate(locale string, s string) (string, error) {
	parsedTime, err := time.Parse("2006-01-02", s)
	if err != nil {
		return "", err
	}
	year, month, date := parsedTime.Date()
	var result string
	switch locale {
	case "en-US":
		result = fmt.Sprintf("%02d/%02d/%d", month, date, year)
	case "nl-NL":
		result = fmt.Sprintf("%02d-%02d-%d", date, month, year)
	}
	return result, nil
}

func currencySymbol(currency string) (string, error) {
	switch currency {
	case "USD":
		return "$", nil
	case "EUR":
		return "€", nil
	default:
		return "", fmt.Errorf("unknown currency")
	}
}

func formatChange(locale string, symbol string, change int) (string, error) {
	var (
		format            string
		thousandSeparator string
		decimalSeparator  string
	)

	absChange := change
	if absChange < 0 {
		absChange = -absChange
	}

	// Format based on locale
	switch locale {
	case "en-US":
		thousandSeparator = ","
		decimalSeparator = "."
		if change < 0 {
			format = "(%s%s)"
		} else {
			format = "%s%s "
		}
	case "nl-NL":
		thousandSeparator = "."
		decimalSeparator = ","
		if change < 0 {
			format = "%s %s-"
		} else {
			format = "%s %s "
		}
	default:
		return "", fmt.Errorf("unknown locale")
	}

	formattedChange := fmt.Sprintf(format, symbol, humanizeChange(absChange, thousandSeparator, decimalSeparator))
	return formattedChange, nil
}

func humanizeChange(n int, thousandSeparator, decimalSeparator string) string {
	decimalPart := n % 100
	integerPart := n / 100

	var parts []string
	var remainder int
	for integerPart > 999 {
		remainder = integerPart % 1000
		parts = append([]string{fmt.Sprintf("%03d", remainder)}, parts...)
		integerPart /= 1000
	}
	parts = append(parts, strconv.Itoa(integerPart))
	sort.Slice(parts, func(i, j int) bool {
		return j < i
	})

	return strings.Join(parts, thousandSeparator) + decimalSeparator + fmt.Sprintf("%02d", decimalPart)
}
