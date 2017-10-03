package stringset

import (
	"fmt"
	"sort"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// API:
//
// New() Set
// NewFromSlice([]string) Set
// (s Set) String() string
// (s Set) IsEmpty() bool
// (s Set) Has(string) bool
// Subset(s1, s2 Set) bool
// Disjoint(s1, s2 Set) bool
// Equal(s1, s2 Set) bool
// (s Set) Add(string)
// Intersection(s1, s2 Set) Set
// Difference(s1, s2 Set) Set
// Union(s1, s2 Set) Set
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements.  For example {"a", "b"}.
// Format the empty set as {}.

const testVersion = 4

// Set is a custom list of strings.
type Set struct {
	elements []string
}

// New creates a blank Set.
func New() Set {
	return Set{}
}

// NewFromSlice creates a Set from a slice of strings.
func NewFromSlice(elements []string) Set {
	s := Set{}
	for _, e := range elements {
		if !s.Has(e) {
			s.elements = append(s.elements, e)
		}
	}
	sort.Strings(s.elements)
	return s
}

func (s Set) String() string {
	if len(s.elements) == 0 {
		return "{}"
	}
	return fmt.Sprintf("{\"%s\"}", strings.Join(s.elements, "\", \""))
}

// IsEmpty returns true if the Set has nothing in it.
func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

// Has returns true if the set contains string `t` as an element.
func (s Set) Has(t string) bool {
	return Include(s, t)
}

// Subset returns true if all the elements in `s1` are also in `s2`.
func Subset(s1, s2 Set) bool {
	matches := Filter(s1, func(t string) bool {
		return Include(s2, t)
	})
	return len(matches.elements) == len(s1.elements)
}

// Disjoint returns true if the two sets have nothing in common.
func Disjoint(s1, s2 Set) bool {
	shorterSet, longerSet := sortSetsByLen(s1, s2)
	hasNothingInCommon := true
	for _, e := range longerSet.elements {
		if Include(shorterSet, e) {
			hasNothingInCommon = false
		}
	}
	return hasNothingInCommon
}

// sortSetsByLen returns two structs: the shorter set by length, and then the
// longer set by length.
func sortSetsByLen(s1, s2 Set) (Set, Set) {
	longerSet, shorterSet := s1, s2
	if len(s1.elements) < len(s2.elements) {
		shorterSet, longerSet = s1, s2
	}
	return shorterSet, longerSet
}

// Equal returns true if all elements in each set are the same.
func Equal(s1, s2 Set) bool {
	if len(s1.elements) != len(s2.elements) {
		return false
	}
	for index, e := range s1.elements {
		if s2.elements[index] != e {
			return false
		}
	}
	return true
}

// Add appends string `t` to the set if it is not already there.
func (s *Set) Add(t string) {
	if !s.Has(t) {
		s.elements = append(s.elements, t)
		sort.Strings(s.elements)
	}
}

// Intersection returns a set of the elements that are in common between the
// two sets.
func Intersection(s1, s2 Set) Set {
	shorterSet, longerSet := sortSetsByLen(s1, s2)
	commonElements := []string{}
	for _, e := range longerSet.elements {
		if Include(shorterSet, e) {
			commonElements = append(commonElements, e)
		}
	}
	return NewFromSlice(commonElements)
}

// Difference returns a set of the items that are in s1 but not in s2.
func Difference(s1, s2 Set) Set {
	uncommonElements := []string{}
	for _, e := range s1.elements {
		if !Include(s2, e) {
			uncommonElements = append(uncommonElements, e)
		}
	}
	return NewFromSlice(uncommonElements)
}

// Union returns a unique set of all the items that are in both sets.
func Union(s1, s2 Set) Set {
	s3 := NewFromSlice(s1.elements)
	for _, e := range s2.elements {
		s3.Add(e)
	}
	return s3
}

// Index returns an int if Set `s` contains string `t`.
func Index(s Set, t string) int {
	for i, v := range s.elements {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true if Set `s` contains string `t`.
func Include(s Set, t string) bool {
	return Index(s, t) >= 0
}

// Filter iterates over Set `s` and returns the items for which `f` is true.
func Filter(s Set, f func(string) bool) Set {
	vsf := make([]string, 0)
	for _, v := range s.elements {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return NewFromSlice(vsf)
}
