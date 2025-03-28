package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	var (
		builder     strings.Builder
		ongoingList bool
	)

	lines := strings.Split(markdown, "\n")
	for _, line := range lines {
		var h int
		for h < len(line) && line[h] == '#' {
			h++
		}
		if h > 0 && h < 7 {
			line = line[h+1:]
			builder.WriteString(fmt.Sprintf("<h%d>%s</h%d>", h, line, h))
			continue
		}

		if strings.HasPrefix(line, "*") {
			if !ongoingList {
				builder.WriteString("<ul>")
				ongoingList = true
			}
			line = fmt.Sprintf("<li>%s</li>", line[2:])
			builder.WriteString(line)
			continue
		}

		if ongoingList {
			builder.WriteString("</ul>")
			ongoingList = false
		}

		builder.WriteString(fmt.Sprintf("<p>%s</p>", line))
	}

	if ongoingList {
		builder.WriteString("</ul>")
	}

	return builder.String()
}
