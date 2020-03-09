// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	isLeap := false
	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0{
				isLeap = true
			}
		} else {
			isLeap = true
		}
	}
	return isLeap
}
