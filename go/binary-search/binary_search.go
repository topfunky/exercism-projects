// Package binarysearch implements a classic search algorithm.
package binarysearch

import "fmt"

const testVersion = 1

// SearchInts looks for a key in a slice.
// It returns the index of the match with some caveats.
// - 0 if less than all values
// - Max index of slice of more than all values
// - Index of value one larger than key if between values
// - Lowest index of match if there are several matches in the slice.
func SearchInts(s []int, key int) int {
	return searchInRange(s, 0, len(s)-1, key)
}

func searchInRange(s []int, left, right, key int) int {
	mid := calculateMidpoint(left, right)
	if right >= left {
		switch {
		case s[mid] == key:
			return peekLeftwardForDuplicates(s, mid, key)
		case s[mid] > key:
			// Search leftward
			return searchInRange(s, left, mid-1, key)
		case s[mid] < key:
			// Search rightward
			return searchInRange(s, mid+1, right, key)
		}
	}
	return mid
}

// Message looks for a key in a slice.
// It returns a descriptive string according to the following pattern.
//
//   k found at beginning of slice.
//   k found at end of slice.
//   k found at index fx.
//   k < all values.
//   k > all n values.
//   k > lv at lx, < gv at gx.
//   slice has no values.
func Message(s []int, key int) string {
	index := SearchInts(s, key)
	// Cases where key is not in range of slice.
	switch {
	case len(s) == 0:
		return "slice has no values"
	case index >= len(s):
		return fmt.Sprintf("%d > all %d values", key, len(s))
	case index == 0 && key < s[index]:
		return fmt.Sprintf("%d < all values", key)
	}

	// Cases where key matches exactly.
	if s[index] == key {
		switch {
		case index == 0:
			return fmt.Sprintf("%d found at beginning of slice", key)
		case index == len(s)-1:
			return fmt.Sprintf("%d found at end of slice", key)
		case index > 0:
			return fmt.Sprintf("%d found at index %d", key, index)
		}
	}

	// Key is in range of slice but doesn't match exactly.
	return fmt.Sprintf("%d > %d at index %d, < %d at index %d", key, s[index-1], index-1, s[index], index)
}

func calculateMidpoint(left, right int) int {
	return left + (right-left)/2
}

func peekLeftwardForDuplicates(s []int, left, key int) int {
	if left >= 0 && s[left] == key {
		return peekLeftwardForDuplicates(s, left-1, key)
	}
	// If we get here, we went too far left and have to add 1.
	return left + 1
}
