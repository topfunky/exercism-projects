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

type Set struct {
	elements []string
}

func New() Set {
	return Set{}
}

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
	quotedElements := Map(s.elements, func(e string) string {
		return fmt.Sprintf("\"%s\"", e)
	})
	if len(s.elements) == 0 {
		return "{}"
	}
	return fmt.Sprintf("{%s}", strings.Join(quotedElements, ", "))
}

func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Set) Has(st string) bool {
	return Include(s.elements, st)
}

func Subset(s1, s2 Set) bool {
	return true
}

func Disjoint(s1, s2 Set) bool {
	return true
}

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

func (s *Set) Add(st string) {
	if !s.Has(st) {
		s.elements = append(s.elements, st)
		sort.Strings(s.elements)
	}
}

func Intersection(s1, s2 Set) Set {
	return Set{}
}

func Difference(s1, s2 Set) Set {
	return Set{}
}

func Union(s1, s2 Set) Set {
	s3 := NewFromSlice(s1.elements)
	for _, e := range s2.elements {
		s3.Add(e)
	}
	return s3
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}
