// Package booking provides Booking up for Beauty solution.
package booking

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	s, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	return !t.After(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		t.Weekday(),
		t.Month(),
		t.Day(),
		t.Year(),
		t.Hour(),
		t.Minute(),
	)
}

// annivDateWithoutYear represents anniversary date without the year.
const annivDateWithoutYear = "9-15 00:00:00"

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	annivDate := fmt.Sprintf("%d-%s", time.Now().Year(), annivDateWithoutYear)
	t, err := time.Parse("2006-1-2 15:04:05", annivDate)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
