// Solution for problem Leap.
package leap

// IsLeapYear check if a given integer is a leap year.
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
