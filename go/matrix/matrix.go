// Package matrix providers Matrix solution.
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix of numbers.
type Matrix struct {
	rows [][]int
}

// New creates a new Matrix.
func New(s string) (*Matrix, error) {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return nil, nil
	}

	rows := make([][]int, len(lines))
	var rowLen int
	for i, line := range lines {
		chars := strings.Split(strings.TrimSpace(line), " ")
		if i == 0 {
			rowLen = len(chars)
		}

		if len(chars) != rowLen {
			return nil, fmt.Errorf("uneven row length between row %v and %v", lines[0], line)
		}

		row := make([]int, rowLen)
		for j, char := range chars {
			if char == "" {
				continue
			}

			num, err := strconv.Atoi(char)
			if err != nil {
				return nil, fmt.Errorf("invalid element: %w", err)
			}

			row[j] = num
		}
		rows[i] = row
	}

	m := &Matrix{
		rows: rows,
	}
	return m, nil
}

// Rows returns all rows of the matrix.
func (m *Matrix) Rows() [][]int {
	if m == nil {
		return nil
	}

	rows := make([][]int, len(m.rows))
	for i := range m.rows {
		row := make([]int, len(m.rows[i]))
		copy(row, m.rows[i])
		rows[i] = row
	}

	return rows
}

// Cols returns all rows of the matrix.
func (m *Matrix) Cols() [][]int {
	if m == nil {
		return nil
	}

	cols := make([][]int, len(m.rows[0]))
	for i := range m.rows[0] {
		col := make([]int, len(m.rows))
		for j := range m.rows {
			col[j] = m.rows[j][i]
		}
		cols[i] = col
	}

	return cols
}

// Set sets element at row & col in the matrix to val. Return false if the given coordinate is invalid.
func (m *Matrix) Set(row, col, val int) bool {
	if m == nil {
		return false
	}

	if row < 0 || row >= len(m.rows) {
		return false
	}

	if col < 0 || col >= len(m.rows[0]) {
		return false
	}

	m.rows[row][col] = val
	return true
}
