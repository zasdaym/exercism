package railfence

import "strings"

func Encode(message string, rails int) string {
	if rails == 1 {
		return message
	}

	rows := make([]string, rails)
	direction := 1
	row := 0

	for _, char := range message {
		rows[row] += string(char)
		row += direction

		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}

	return strings.Join(rows, "")
}

func Decode(message string, rails int) string {
	if rails == 1 || len(message) == 0 {
		return message
	}

	// Create the rail pattern to determine position of characters
	pattern := make([][]bool, rails)
	for i := range pattern {
		pattern[i] = make([]bool, len(message))
	}

	// Mark positions where characters will be placed
	direction := 1
	row := 0
	for col := 0; col < len(message); col++ {
		pattern[row][col] = true
		row += direction

		if row == 0 || row == rails-1 {
			direction *= -1
		}
	}

	// Fill the rail pattern with characters from the encoded message
	idx := 0
	decoded := make([]byte, len(message))
	for r := 0; r < rails; r++ {
		for c := 0; c < len(message); c++ {
			if pattern[r][c] && idx < len(message) {
				decoded[c] = message[idx]
				idx++
			}
		}
	}

	return string(decoded)
}

// WECRLTEERDSOEEFEAOCAIVDEN
// W . . . E . . . C . . . R . . . L . . . T . . . E
// . E . R . D . S . O . E . E . F . E . A . O . C .
// . . A . . . I . . . V . . . D . . . E . . . N . .
