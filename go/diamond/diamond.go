package diamond

import (
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("invalid char")
	}

	var (
		maxLetterIndex = int(char - 'A')
		length         = int(maxLetterIndex*2 + 1)
	)

	var builder strings.Builder

	for i := 0; i < length; i++ {
		if i == 0 || i == length-1 {
			// Handle the first and last row
			outerSpaces := strings.Repeat(" ", (length-1)/2)
			builder.WriteString(fmt.Sprintf("%s%c%s", outerSpaces, 'A', outerSpaces))
			if i == 0 && i != length-1 {
				builder.WriteString("\n")
			}
			continue
		}

		var letterIndex int
		if i <= length/2 {
			// Top half of the diamond
			letterIndex = i
		} else {
			// Bottom half of the diamond
			letterIndex = maxLetterIndex - (i - maxLetterIndex)
		}

		letter := 'A' + letterIndex
		innerSpacesCount := letterIndex*2 - 1
		outerSpacesCount := (length - innerSpacesCount - 2) / 2
		outerSpaces := strings.Repeat(" ", outerSpacesCount)
		innerSpaces := strings.Repeat(" ", innerSpacesCount)
		builder.WriteString(fmt.Sprintf("%s%c%s%c%s\n", outerSpaces, letter, innerSpaces, letter, outerSpaces))
	}

	return builder.String(), nil
}
