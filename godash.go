package godash

type Slice[T any] []T

//type Slice interface {
//}

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
