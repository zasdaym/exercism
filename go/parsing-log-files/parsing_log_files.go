package parsinglogfiles

import "regexp"

var (
	validLine = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*$`)
	separator = regexp.MustCompile(`<[-=*~]*>`)
	password  = regexp.MustCompile(`\"[A-Za-z ]+\"`)
	eol       = regexp.MustCompile(`end-of-line[0-9]+`)
	user      = regexp.MustCompile(`User\s+([A-Za-z0-9]+)`)
)

func IsValidLine(text string) bool {
	return validLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	return separator.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var result int
	for _, line := range lines {
		if !IsValidLine(line) {
			continue
		}
		if password.MatchString(line) {
			result++
		}
	}
	return result
}

func RemoveEndOfLineText(text string) string {
	return eol.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var result []string
	for _, line := range lines {
		matches := user.FindStringSubmatch(line)
		if len(matches) > 0 {
			line = "[USR] " + matches[1] + " " + line
		}
		result = append(result, line)
	}
	return result
}
