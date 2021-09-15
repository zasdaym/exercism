// Package logs provides Logs, Logs, Logs! solution.
package logs

import (
	"fmt"
	"strings"
)

const logSeparator = ":"

// Message extracts the message from the provided log line.
func Message(line string) string {
	parts := strings.Split(line, logSeparator)
	if len(parts) < 2 {
		return ""
	}
	return strings.TrimSpace(parts[1])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	msg := Message(line)
	n := 0
	for range msg {
		n++
	}
	return n
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	parts := strings.Split(line, logSeparator)
	lvl := strings.ToLower(strings.Trim(parts[0], "[]"))
	return lvl
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	msg := Message(line)
	lvl := LogLevel(line)
	newLog := fmt.Sprintf("%s (%s)", msg, lvl)
	return newLog
}
