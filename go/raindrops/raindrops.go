// Package raindrops includes functions for generating funny strings from numbers.
package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 3

// Convert turns a number into a funny string based on the number's factors.
// It returns a funny string based on the prime factors 3, 5, and 7.
func Convert(n int) string {
	var buffer bytes.Buffer

	if n%3 == 0 {
		buffer.WriteString("Pling")
	}
	if n%5 == 0 {
		buffer.WriteString("Plang")
	}
	if n%7 == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() == 0 {
		return strconv.Itoa(n)
	}

	return buffer.String()
}
