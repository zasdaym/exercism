package wordsearch

import (
	"fmt"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	for _, word := range words {
		point := find(word, puzzle)
		if point[0][0] == -1 {
			return nil, fmt.Errorf("word not found")
		}
		result[word] = point
	}
	return result, nil
}

func find(word string, puzzle []string) [2][2]int {
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			offsets := [][2]int{
				{0, 1},
				{0, -1},
				{1, 0},
				{-1, 0},
				{1, 1},
				{1, -1},
				{-1, 1},
				{-1, -1},
			}
			for _, offset := range offsets {
				point := check(word, puzzle, i, j, offset)
				if point[0] == -1 {
					continue
				}
				return [2][2]int{{j, i}, point}
			}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}
}

func check(word string, puzzle []string, y, x int, offset [2]int) [2]int {
	if y < 0 || y >= len(puzzle) || x < 0 || x >= len(puzzle[y]) || puzzle[y][x] != word[0] {
		return [2]int{-1, -1}
	}
	if len(word) == 1 {
		return [2]int{x, y}
	}
	return check(word[1:], puzzle, y+offset[0], x+offset[1], offset)
}
