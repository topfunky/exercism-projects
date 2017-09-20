package acronym

import (
	"bytes"
	"strings"
	"unicode"
)

const testVersion = 3

// Abbreviate creates an acronym from a phrase.
// It returns a string that is an acronym of the phrase.
func Abbreviate(phrase string) string {
	var buffer bytes.Buffer

	var words = strings.FieldsFunc(phrase, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	for _, word := range words {
		letter := strings.ToUpper(string(word[0]))
		buffer.WriteString(letter)
	}

	return buffer.String()
}
