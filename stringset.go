package exercism

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

import (
	"sort"
	"strings"
)

// Define the Set type here.
type Set struct {
	data []string
}

func New() Set {
	return Set{
		data: make([]string, 0),
	}
}

func NewFromSlice(l []string) Set {
	m := make(map[string]bool)
	list := []string{}

	for _, value := range l {
		if _, exists := m[value]; !exists {
			list = append(list, value)
			m[value] = true
		}
	}

	sort.Strings(list)
	return Set{
		data: list,
	}
}

func (s Set) String() string {
	result := strings.Join(s.data, "\", \"")
	if len(result) > 0 {
		return "{\"" + result + "\"}"
	} else {
		return "{}"
	}
}

func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Set) Has(elem string) bool {
	for _, val := range s.data {
		if elem == val {
			return true
		}
	}
	return false
}

func (s *Set) Add(elem string) {
	if s.Has(elem) {
		return
	} else {
		s.data = append(s.data, elem)
		sort.Strings(s.data)
	}
}

func Subset(s1, s2 Set) bool {
	for _, val := range s1.data {
		if s2.Has(val) == false {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for _, val := range s1.data {
		if s2.Has(val) == true {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.data) != len(s2.data) {
		return false
	}

	for idx, val := range s1.data {
		if val != s2.data[idx] {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	intersec := make([]string, len(s1.data))
	intIdx := 0
	for _, val := range s1.data {
		if s2.Has(val) {
			intersec[intIdx] = val
			intIdx++
		}
	}
	return NewFromSlice(intersec[:intIdx])
}

func Difference(s1, s2 Set) Set {
	values := make(map[string]bool)

	for _, val := range s2.data {
		values[val] = true
	}

	diff := make([]string, len(s1.data)+len(s2.data))
	diffIdx := 0
	for _, val := range s1.data {
		_, exists := values[val]
		if !exists {
			diff[diffIdx] = val
			diffIdx++
		}
	}
	return NewFromSlice(diff[:diffIdx])
}

func Union(s1, s2 Set) Set {
	len1 := len(s1.data)
	len2 := len(s2.data)
	list := make([]string, len1+len2)
	copy(list, s1.data)
	copy(list[len1:], s2.data)
	return NewFromSlice(list)
}
