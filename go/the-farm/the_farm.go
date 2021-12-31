// Package thefarm provides The Farm solution.
package thefarm

import (
	"errors"
	"fmt"
)

// SillyNephewError represents error caused by silly nephew when counting cows.
type SillyNephewError struct {
	cows int
}

func (s SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0, SillyNephewError{cows: cows}
	}
	amount, err := weightFodder.FodderAmount()
	if errors.Is(err, ErrScaleMalfunction) {
		amount *= 2
	} else if err != nil {
		return 0, err
	}
	if amount < 0 {
		return 0, errors.New("negative fodder")
	}
	return amount / float64(cows), nil
}
