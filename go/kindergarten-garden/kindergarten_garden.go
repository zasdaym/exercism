package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

type Garden struct {
	plantsByChild map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	result := &Garden{
		plantsByChild: make(map[string][]string, 0),
	}

	lines := strings.Split(diagram, "\n")
	if len(lines) != 3 {
		return nil, fmt.Errorf("invalid diagram")
	}
	if lines[0] != "" {
		return nil, fmt.Errorf("invalid diagram")
	}
	if len(lines[1]) != len(lines[2]) {
		return nil, fmt.Errorf("mismatched rows")
	}
	if len(lines[1])%2 == 1 {
		return nil, fmt.Errorf("odd number of plants in a row")
	}

	sortedChildren := append([]string(nil), children...)
	sort.Slice(sortedChildren, func(i, j int) bool {
		return sortedChildren[i] < sortedChildren[j]
	})

	for i := range sortedChildren {
		if _, ok := result.plantsByChild[sortedChildren[i]]; ok {
			return nil, fmt.Errorf("duplicate child name")
		}
		positions := [4][2]int{
			{1, i * 2},
			{1, i*2 + 1},
			{2, i * 2},
			{2, i*2 + 1},
		}
		var plants []string
		for _, position := range positions {
			x, y := position[1], position[0]
			plant := plantFromCode[lines[y][x]]
			if plant == "" {
				return nil, fmt.Errorf("invalid plant code")
			}
			plants = append(plants, plant)
		}
		result.plantsByChild[sortedChildren[i]] = plants
	}

	return result, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	result, ok := g.plantsByChild[child]
	return result, ok
}

var plantFromCode = map[byte]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}
