// Package leap contains functions for working with years
// in the Gregorian calendar.
package leap

const testVersion = 3

/*
IsLeapYear takes a year determines if it is a leap year according to
the Gregorian calendar.

It returns a boolean if the year is a leap year.
*/
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	}
	return false
}
