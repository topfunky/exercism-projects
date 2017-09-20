// Package raindrops includes functions for generating funny strings from numbers.
package raindrops

import (
	"fmt"
	"strconv"
)

const testVersion = 3

// Convert turns a number into a funny string based on the number's factors.
// It returns a funny string based on the prime factors 3, 5, and 7.
func Convert(n int) string {
	pling, plang, plong := "", "", ""

	if n%3 == 0 {
		pling = "Pling"
	}
	if n%5 == 0 {
		plang = "Plang"
	}
	if n%7 == 0 {
		plong = "Plong"
	}

	if pling == "" && plang == "" && plong == "" {
		return strconv.Itoa(n)
	}

	return fmt.Sprintf("%v%v%v", pling, plang, plong)
}
