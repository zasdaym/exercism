// Package matrix is solution for problem Matrix.
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix.
type Matrix [][]int

// New creates a new Matrix.
func New(in string) (Matrix, error) {
	matrix := make(Matrix, 0)
	lines := strings.Split(in, "\n")
	lastLen := 0
	for i, line := range lines {
		row := make([]int, 0)
		for _, s := range strings.Split(strings.TrimSpace(line), " ") {
			num, err := strconv.Atoi(s)
			if err != nil {
				return matrix, fmt.Errorf("error converting input to number")
			}
			row = append(row, num)
		}
		if i != 0 && len(row) != lastLen {
			return matrix, fmt.Errorf("inconsistent row length")
		}
		lastLen = len(row)
		matrix = append(matrix, row)
	}
	return matrix, nil
}

// Rows returns rows of a Matrix.
func (m Matrix) Rows() [][]int {
	var newRows [][]int
	for _, row := range m {
		newRow := make([]int, len(row))
		copy(newRow, row)
		newRows = append(newRows, newRow)
	}
	return newRows
}

// Cols returns columns of a Matrix.
func (m Matrix) Cols() [][]int {
	rowLen := len(m)
	colLen := len(m[0])

	var newRows [][]int
	for i := 0; i < colLen; i++ {
		newRow := make([]int, rowLen)
		newRows = append(newRows, newRow)
	}

	for i, row := range m {
		for j, num := range row {
			newRows[j][i] = num
		}
	}
	return newRows
}

// Set sets a matrix element on given row and col.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m) || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
