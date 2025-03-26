package matrix

import (
	"slices"
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix struct {
	rows [][]int
}

type Pair struct {
	y int
	x int
}

func New(s string) (*Matrix, error) {
	rows := make([][]int, 0)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		var row []int
		for _, unparsed := range strings.Fields(line) {
			num, err := strconv.Atoi(unparsed)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		if len(row) == 0 {
			continue
		}
		rows = append(rows, row)
	}
	return &Matrix{rows: rows}, nil
}

func (m *Matrix) Saddle() []Pair {
	var result []Pair

	for i := range m.rows {
		var candidates []int
		highest := slices.Max(m.rows[i])
		for j := range m.rows[i] {
			if m.rows[i][j] == highest {
				candidates = append(candidates, j)
			}
		}

		for _, candidate := range candidates {
			for ii := range m.rows {
				if m.rows[ii][candidate] < m.rows[i][candidate] {
					candidate = -1
					break
				}
			}
			if candidate != -1 {
				result = append(result, Pair{y: i + 1, x: candidate + 1})
			}
		}
	}

	return result
}
