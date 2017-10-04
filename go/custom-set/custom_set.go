package stringset

import (
	"fmt"
	"sort"
	"strings"
)

const testVersion = 4

// Set is a custom list of strings.
type Set []string

// New creates a blank Set.
func New() Set {
	return Set{}
}

// NewFromSlice creates a Set from a slice of strings.
func NewFromSlice(elements []string) Set {
	s := Set{}
	s.Add(elements...)
	sort.Strings(s)
	return s
}

func (s Set) String() string {
	quoted := []string{}
	for _, e := range s {
		quoted = append(quoted, fmt.Sprintf("%q", e))
	}
	return fmt.Sprintf("{%s}", strings.Join(quoted, ", "))
}

// IsEmpty returns true if the Set has nothing in it.
func (s Set) IsEmpty() bool {
	return len(s) == 0
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
	return len(matches) == len(s1)
}

// Disjoint returns true if the two sets have nothing in common.
func Disjoint(s1, s2 Set) bool {
	for _, e := range s1 {
		if Include(s2, e) {
			return false
		}
	}
	return true
}

// Equal returns true if all elements in each set are the same.
func Equal(s1, s2 Set) bool {
	return len(s1) == len(s2) && Subset(s1, s2)
}

// Add appends one or more `elements` to the set if they are not already there.
func (s *Set) Add(elements ...string) {
	for _, e := range elements {
		if !s.Has(e) {
			*s = append(*s, e)
			sort.Strings(*s)
		}
	}
}

// Intersection returns a set of the elements that are in common between the
// two sets.
func Intersection(s1, s2 Set) Set {
	common := Set{}
	for _, e := range s1 {
		if Include(s2, e) {
			common.Add(e)
		}
	}
	return common
}

// Difference returns a set of the items that are in s1 but not in s2.
func Difference(s1, s2 Set) Set {
	uncommon := Set{}
	for _, e := range s1 {
		if !Include(s2, e) {
			uncommon.Add(e)
		}
	}
	return uncommon
}

// Union returns a unique set of all the items that are in both sets.
func Union(s1, s2 Set) Set {
	for _, e := range s2 {
		s1.Add(e)
	}
	return s1
}

// Index returns an int if Set `s` contains string `t`.
func Index(s Set, t string) int {
	for i, v := range s {
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
	for _, v := range s {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return NewFromSlice(vsf)
}
