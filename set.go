package godash

import (
	"fmt"
	"slices"
	"strings"
)

type setElement interface {
	comparable
}

type Set[T setElement] map[T]struct{}

func NewSet[T setElement](elements ...T) Set[T] {
	m := make(map[T]struct{})
	for _, element := range elements {
		m[element] = struct{}{}
	}
	return m
}

// Add inserts the specified element into the set.
// If the element already exists in the set, no action is taken.
func (s *Set[T]) Add(element T) {
	(*s)[element] = struct{}{}
}

// Delete removes the specified element from the set.
// If the element doesn't exist in the set, no action is taken.
func (s *Set[T]) Delete(element T) {
	delete(*s, element)
}

// Has checks if the specified element exists in the set.
// It returns true if the element exists, otherwise it returns false.
func (s Set[T]) Has(element T) bool {
	_, ok := s[element]
	return ok
}

// Values returns a slice containing all the elements in the set.
// These elements won't be returned in any specific order.
func (s *Set[T]) Values() []T {
	result := make([]T, 0, len(*s))
	for key, _ := range *s {
		result = append(result, key)
	}
	return result
}

// String returns a string representation of the set.
// The elements in the set are joined by commas and surrounded by curly braces.
// The elements are sorted in ascending order before joining.
// The string representation that is returned is in the format: "set{element1, element2, ...}"
// The order of the elements guaranteed to make troubleshooting easier.
// The elements are converted to strings using the format "%v".
func (s Set[T]) String() string {
	elements := s.Values()
	elementsStr := make([]string, len(elements))
	for i, element := range elements {
		elementsStr[i] = fmt.Sprintf("%v", element)
	}
	slices.Sort(elementsStr)

	elementsStrJoined := strings.Join(elementsStr, ", ")
	return fmt.Sprintf("set{%s}", elementsStrJoined)
}
