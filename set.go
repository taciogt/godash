package godash

import (
	"fmt"
	"slices"
	"strings"
)

type setElement interface {
	comparable
}

// Set is a type that represents a set data structure.
// It is implemented using a map where the keys represent the elements of the set.
// The values of the map are of type `struct{}` and don't contain any meaningful data.
// The `Set` type can be used to store elements of any type that implements the `comparable` interface.
// The zero value of Set is an empty set.
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

// Size returns the number of elements in the set.
// It does not count duplicate elements multiple times.
func (s Set[T]) Size() int {
	return len(s)
}

// Intersection returns a new set that contains the common elements between the set and the other set.
// The resulting set will only contain elements that are present in both sets.
func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	result := make(Set[T])

	pivot, other := s, s2
	if pivot.Size() > other.Size() {
		pivot, other = other, pivot
	}

	for key := range s {
		if s2.Has(key) {
			result.Add(key)
		}
	}
	return result
}

// Union returns a new set that contains all the elements from the receiver set (s)
// and the other set (s2). The resulting set does not contain any duplicate elements
// and is returned as the result. The receiver set (s) and the other set (s2) are not modified.
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	result := NewSet[T]()

	for _, v := range s.Values() {
		result.Add(v)
	}
	for _, v := range s2.Values() {
		result.Add(v)
	}

	return result
}

// Difference returns a new set that contains the elements of the receiver set (s) but not in the other set (s2).
// The resulting set does not contain any duplicate elements and is returned as the result.
// The receiver set (s) is not modified.
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	result := NewSet[T]()

	for key := range s {
		if !s2.Has(key) {
			result.Add(key)
		}
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
