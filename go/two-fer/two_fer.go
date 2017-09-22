// Package twofer provides functions for playing a text-based kids game.
package twofer

import "fmt"

// ShareWith formats the name argument into a text game.
// It returns a string that includes the name or the default "you".
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
