package igpay

import (
	"fmt"
	"strings"
)

const testVersion = 1

// PigLatin translates a phrase from English to a funny language ending in 'ay'.
// It returns the translated string.
func PigLatin(phrase string) (pl string) {
	words := strings.Split(phrase, " ")
	translations := make([]string, len(words))
	for index, word := range words {
		translations[index] = translate(word)
	}
	return strings.Join(translations, " ")
}

func translate(word string) (pl string) {
	switch {
	case startsWithVowel(word):
		pl = fmt.Sprintf("%say", word)
	case startsWithChars(word, []string{"xr", "ytt"}):
		pl = fmt.Sprintf("%say", word)
	case startsWithChars(word, []string{"squ", "thr", "sch"}):
		pl = transposeLeadingChars(word, 3)
	case startsWithChars(word, []string{"ch", "qu", "th"}):
		pl = transposeLeadingChars(word, 2)
	default:
		pl = transposeLeadingChars(word, 1)
	}
	return pl
}

func transposeLeadingChars(phrase string, numberOfChars int) string {
	first, rest := phrase[0:numberOfChars], phrase[numberOfChars:]
	return fmt.Sprintf("%s%say", rest, first)
}

func startsWithVowel(phrase string) bool {
	return strings.IndexAny(phrase, "aeiou") == 0
}

// startsWithChars examines a string to see if it starts with any of a slice of
// strings handed in the second argument.
// It returns true if so.
func startsWithChars(phrase string, pairs []string) bool {
	for _, pair := range pairs {
		if strings.Index(phrase, pair) == 0 {
			return true
		}
	}
	return false
}
