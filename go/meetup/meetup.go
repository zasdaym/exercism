// Package meetup is solution for problem Meetup.
package meetup

import (
	"time"
)

// WeekSchedule represents the week of the month when meetup scheduled.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns calendar day that matches the meetup restrictions.
func Day(week WeekSchedule, day time.Weekday, month time.Month, year int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.FixedZone("UTC", 0))
	var matches []int

	for date.Month() == month {
		if date.Weekday() == day {
			matches = append(matches, date.Day())
		}
		date = date.AddDate(0, 0, 1)
	}
	
	switch week {
	case First:
		return matches[0]
	case Second:
		return matches[1]
	case Third:
		return matches[2]
	case Fourth:
		return matches[3]
	case Last:
		return matches[len(matches)-1]
	case Teenth:
		for _, d := range matches {
			if d >= 13 {
				return d
			}
		}
	}

	return -1
}
