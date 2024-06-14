package godash

type Slice[T any] []T

//type Slice interface {
//}

// Find returns the first element in the given slice that satisfies the provided predicate function.
// If no element satisfies the predicate, the zero value of the element type is returned along with false.
func Find[T any, S ~[]T](s S, predicate func(v T) bool) (T, bool) {
	for _, v := range s {
		if predicate(v) {
			return v, true
		}
	}

	var zero T
	return zero, false
}
