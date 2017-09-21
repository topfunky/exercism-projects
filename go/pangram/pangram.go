package pangram

import (
	"strings"
	"unicode"
)

const testVersion = 2

// IsPangram determines if a phrase contains all letters in the alphabet.
// It returns true if so.
func IsPangram(phrase string) bool {
	phrase = strings.ToLower(phrase)
	var letterTracker = make(map[rune]int)

	for _, letter := range phrase {
		if unicode.IsLetter(letter) {
			letterTracker[letter] = 1
		}
	}

	sum := 0
	for _, v := range letterTracker {
		sum += v
	}

	return sum == 26
}
