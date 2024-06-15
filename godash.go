package godash

type Slice[T any] []T

// Every returns true if every element in the given slice satisfies the provided predicate function.
// Otherwise, it returns false.
func Every[T any](s Slice[T], p func(T) bool) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

// FindIndex returns the index of the first element in the given slice that satisfies
// the provided predicate function. If no element satisfies the predicate,
// -1 is returned along with false.
func FindIndex[T any, S ~[]T](s S, p func(T) bool) (int, bool) {
	for i, v := range s {
		if p(v) {
			return i, true
		}
	}
	return -1, false
}

// Find returns the first element in the given slice that satisfies the provided predicate function.
// If no element satisfies the predicate, the zero value of the element type is returned along with false.
func Find[T any, S ~[]T](s S, p func(T) bool) (T, bool) {
	if i, ok := FindIndex(s, p); ok {
		return s[i], true
	}

	var zero T
	return zero, false
}

// Filter returns a new slice containing only the elements from the given slice that satisfy the provided predicate function.
func Filter[T any, S ~[]T](s S, p func(T) bool) Slice[T] {
	result := make([]T, 0)

	for _, v := range s {
		if p(v) {
			result = append(result, v)
		}
	}

	return result
}
