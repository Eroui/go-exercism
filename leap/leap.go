// Package leap package provide utitilies for leap years
package leap

// IsLeapYear returns true if a given year is a leap year and false otherwise
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		} else {
			return true
		}
	} else {
		return false
	}
}
