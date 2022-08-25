package expenses

import (
	"errors"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(records []Record, predicate func(Record) bool) []Record {
	result := make([]Record, 0, len(records))
	for _, r := range records {
		if predicate(r) {
			result = append(result, r)
		}
	}
	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day <= p.To && r.Day >= p.From
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(records []Record, p DaysPeriod) float64 {
	var sum float64
	filtered := Filter(records, ByDaysPeriod(p))
	for _, r := range filtered {
		sum += r.Amount
	}
	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(records []Record, p DaysPeriod, c string) (float64, error) {
	var sum float64

	inCategory := Filter(records, ByCategory(c))
	if len(inCategory) == 0 {
		return 0, errors.New("no records in the given category")
	}

	inPeriod := Filter(inCategory, ByDaysPeriod(p))
	for _, r := range inPeriod {
		sum += r.Amount
	}

	return sum, nil
}
