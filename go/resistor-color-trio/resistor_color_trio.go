package resistorcolortrio

import (
	"fmt"
	"math"
)

var codeByColor = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	var sum int
	for i, color := range colors[:3] {
		if i == 2 {
			sum *= int(math.Pow10(codeByColor[color]))
			continue
		}
		sum *= 10
		sum += codeByColor[color]
	}
	length := len(fmt.Sprint(sum))
	switch {
	case length > 9:
		return fmt.Sprintf("%d gigaohms", sum/1_000_000_000)
	case length > 6:
		return fmt.Sprintf("%d megaohms", sum/1_000_000)
	case length > 3:
		return fmt.Sprintf("%d kiloohms", sum/1_000)
	default:
		return fmt.Sprintf("%d ohms", sum)
	}
}
