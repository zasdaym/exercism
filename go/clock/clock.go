// Package clock is solution for problem Clock.
package clock

import (
	"fmt"
)

// Clock represents a clock time.
type Clock int

// New returns new clock from given hours and minutes.
func New(hours, minutes int) Clock {
	return Clock(mod((hours*60 + minutes), 1440))
}

// String returns string format of Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add returns Clock added with given minutes.
func (c Clock) Add(minutes int) Clock {
	return Clock(mod(int(c)+minutes, 1440))
}

// Subtract returns Clock subtracted with given minutes.
func (c Clock) Subtract(minutes int) Clock {
	return Clock(mod(int(c)-minutes, 1440))
}

func mod(x, m int) int {
	return (x%m + m) % m
}
