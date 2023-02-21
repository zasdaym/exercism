package minesweeper

import (
	"strings"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	result := make([]string, 0, len(board))
	for i := range board {
		var sb strings.Builder
		for j := range board[i] {
			// Mine itself
			if board[i][j] == '*' {
				sb.WriteRune('*')
				continue
			}

			// Check surrounding mines
			var mines int
			cur := &coordinate{y: i, x: j}
			for _, neighbor := range cur.neighbors(len(board), len(board[i])) {
				if board[neighbor.y][neighbor.x] == '*' {
					mines++
				}
			}

			// Insert number of mines or a space if empty
			char := byte('0' + mines)
			if mines == 0 {
				char = ' '
			}
			sb.WriteByte(char)
		}
		result = append(result, sb.String())
	}
	return result
}

// coordinate represents a coordinate in a minesweeper board.
type coordinate struct {
	y int
	x int
}

// neighbors returns all valid neighbors of a coordinate in a minesweeper board.
func (c *coordinate) neighbors(maxHeight, maxWidth int) []coordinate {
	var result []coordinate
	offsets := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, offset := range offsets {
		y, x := c.y+offset[0], c.x+offset[1]
		if y >= 0 && x >= 0 && y < maxHeight && x < maxWidth {
			result = append(result, coordinate{y, x})
		}
	}
	return result
}
