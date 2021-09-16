// Package clock provides Clock solution.
package clock

import (
	"fmt"
)

// Clock represents a clock that handles times without dates.
type Clock struct {
	hour, minute int
}

// New creates a new Clock.
func New(hour, minute int) Clock {
	c := Clock{
		hour:   hour,
		minute: minute,
	}
	return c.normalize()
}

const (
	hoursInADay    = 24
	minutesInADay  = 1440
	minutesInAHour = 60
)

// normalize normalizes hour and minute of Clock.
func (c Clock) normalize() Clock {
	totalMinutes := (c.hour*minutesInAHour + c.minute) % minutesInADay
	if totalMinutes < 0 {
		totalMinutes = minutesInADay + totalMinutes
	}
	c.hour = totalMinutes / 60
	c.minute = totalMinutes % 60
	return c
}

// String returns string representation of Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds hour and minute into Clock.
func (c Clock) Add(minute int) Clock {
	c.minute += minute
	return c.normalize()
}

// Subtract subtracts hour and minute from Clock.
func (c Clock) Subtract(minute int) Clock {
	c.minute -= minute
	return c.normalize()
}
