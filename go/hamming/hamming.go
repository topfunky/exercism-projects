// Package hamming contains functions for reporting on DNA sequences.
package hamming

import "errors"

const testVersion = 6

// Distance calculates the Hamming distance between two strings, representing
// DNA strands. The distance is an integer equal to the number of characters
// that differ between the two strings.
//
// It returns an integer for the distance and an optional error (the distance
// will be -1 if there is an error).
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA Sequences must be the same length, but are not")
	}

	score := 0
	runesB := []rune(b)
	for index, character := range a {
		if character != runesB[index] {
			score++
		}
	}
	return score, nil // errors.New("Something went wrong")
}
