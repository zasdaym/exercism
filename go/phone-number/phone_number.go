package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	var sb strings.Builder
	var pos int
	for _, c := range phoneNumber {
		if !unicode.IsNumber(c) {
			continue
		}
		number := c - '0'
		// Ignore country code
		if pos == 0 && number == 1 {
			continue
		}
		if (pos == 0 || pos == 3) && number < 2 {
			return "", fmt.Errorf("invalid number")
		}
		sb.WriteRune(c)
		pos++
	}
	result := sb.String()

	if len(result) != 10 {
		return "", fmt.Errorf("invalid phone number length")
	}
	return result, nil
}

func AreaCode(phoneNumber string) (string, error) {
	parsed, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return parsed[:3], nil
}

func Format(phoneNumber string) (string, error) {
	parsed, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", parsed[:3], parsed[3:6], parsed[6:]), nil
}
