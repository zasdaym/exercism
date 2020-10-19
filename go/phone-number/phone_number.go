package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

func Number(s string) (string, error) {
	cleaned := cleanNumber(s)
	var result strings.Builder
	for _, n := range cleaned {
		if !unicode.IsDigit(n) {
			return "", fmt.Errorf("invalid character")
		}
		result.WriteRune(n)
	}
	number := result.String()
	if len(number) > 11 || len(number) < 10 {
		return "", fmt.Errorf("invalid length")
	}
	if len(number) == 11 && number[0] != '1' {
		return "", fmt.Errorf("11 digit must be started with 1")
	}
	if len(number) == 11 {
		number = number[1:]
	}
	if number[0] == '0' || number[0] == '1' {
		return "", fmt.Errorf("area code can't be 0 or 1")
	}
	if number[3] == '0' || number[3] == '1' {
		return "", fmt.Errorf("exchange code can't be 0 or 1")
	}
	return number, nil
}

func cleanNumber(s string) string {
	replacer := strings.NewReplacer("(", "", ")", "", "-", "", ".", "", "+", "", " ", "")
	return replacer.Replace(s)
}

func AreaCode(s string) (string, error) {
	phone, err := Number(s)
	if err != nil {
		return "", err
	}
	return phone[:3], nil
}

func Format(s string) (string, error) {
	phone, err := Number(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%v) %v-%v", phone[:3], phone[3:6], phone[6:]), nil
}
