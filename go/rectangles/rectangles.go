package rectangles

import (
	"fmt"
	"slices"
	"sort"
)

type Point struct {
	y, x int
}

func (p Point) String() string {
	return fmt.Sprintf("(y: %d, x: %d)", p.y, p.x)
}

func Count(diagram []string) int {
	var corners []Point
	for y, line := range diagram {
		for x, char := range line {
			if char == '+' {
				corners = append(corners, Point{y: y, x: x})
			}
		}
	}

	rectangles := make(map[string]bool)
	for i := 0; i < len(corners); i++ {
		for j := i + 1; j < len(corners); j++ {
			if corners[i].x == corners[j].x || corners[i].y == corners[j].y {
				continue
			}

			var (
				a = corners[i]
				b = corners[j]
				c = Point{y: a.y, x: b.x}
				d = Point{y: b.y, x: a.x}
			)

			if !slices.Contains(corners, c) || !slices.Contains(corners, d) {
				continue
			}

			points := []Point{a, b, c, d}
			sort.Slice(points, func(i, j int) bool {
				if points[i].y == points[j].y {
					return points[i].x < points[j].x
				}
				return points[i].y < points[j].y
			})

			hashKey := fmt.Sprintf("%s-%s-%s-%s", points[0], points[1], points[2], points[3])
			if rectangles[hashKey] {
				continue
			}

			if validRectangle(diagram, points[0], points[1], points[2], points[3]) {
				rectangles[hashKey] = true
			}
		}
	}

	return len(rectangles)
}

func validRectangle(diagram []string, topLeft, topRight, bottomLeft, bottomRight Point) bool {
	// Top
	for x := topLeft.x + 1; x < topRight.x; x++ {
		if diagram[topLeft.y][x] != '-' && diagram[topLeft.y][x] != '+' {
			return false
		}
	}

	// Bottom
	for x := bottomLeft.x + 1; x < bottomRight.x; x++ {
		if diagram[bottomLeft.y][x] != '-' && diagram[bottomLeft.y][x] != '+' {
			return false
		}
	}

	// Left
	for y := topLeft.y + 1; y < bottomLeft.y; y++ {
		if diagram[y][topLeft.x] != '|' && diagram[y][topLeft.x] != '+' {
			return false
		}
	}

	// Right
	for y := topRight.y + 1; y < bottomRight.y; y++ {
		if diagram[y][topRight.x] != '|' && diagram[y][topRight.x] != '+' {
			return false
		}
	}

	return true
}
