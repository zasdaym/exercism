// Package clock is solution for problem Clock.
package clock

import (
	"fmt"
)

const day = 1440

// Clock represents a clock time.
type Clock struct {
	minutes int
}

// New returns new clock from given hours and minutes.
func New(hours, minutes int) Clock {
	return Clock{
		minutes: mod((hours*60 + minutes), day),
	}
}

// String returns string format of Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add returns Clock added with given minutes.
func (c Clock) Add(minutes int) Clock {
	return Clock{
		minutes: mod(int(c.minutes)+minutes, day),
	}
}

// Subtract returns Clock subtracted with given minutes.
func (c Clock) Subtract(minutes int) Clock {
	return Clock{
		minutes: mod(int(c.minutes)-minutes, day),
	}
}

func mod(x, m int) int {
	return (x%m + m) % m
}
