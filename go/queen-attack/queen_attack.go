package queenattack

import (
	"fmt"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	white, err := parseQueen(whitePosition)
	if err != nil {
		return false, err
	}
	black, err := parseQueen(blackPosition)
	if err != nil {
		return false, err
	}
	if white == black {
		return false, fmt.Errorf("same queen")
	}
	return white.canAttack(black), nil
}

type queen struct {
	x int
	y int
}

func parseQueen(s string) (queen, error) {
	if len(s) != 2 {
		return queen{}, fmt.Errorf("invalid cell format")
	}
	x := int(s[0]) - 'a'
	y := int(s[1]) - '1'
	if x < 0 || x > 7 || y < 0 || y > 7 {
		return queen{}, fmt.Errorf("cell out of bounds")
	}
	return queen{x, y}, nil
}

func (q queen) canAttack(other queen) bool {
	return q.x == other.x || q.y == other.y || abs(q.x-other.x) == abs(q.y-other.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
