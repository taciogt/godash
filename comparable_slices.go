package godash

// ComparableSlice is a generic slice type that restricts its elements to types that satisfy the comparable constraint.
// It extends the behavior of the standard Slice with some extra functionality that don't require predicates,
// like the Includes method.
type ComparableSlice[T comparable] Slice[T]

func NewComparableSlice[T comparable](elems ...T) ComparableSlice[T] {
	return elems
}

// Includes checks whether the specified value exists within the given ComparableSlice.
// Returns true if found, false otherwise.
func Includes[T comparable](s ComparableSlice[T], value T) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

// Includes behaves exactly like [Includes] function, except it is called directly on the slice.
func (s ComparableSlice[T]) Includes(value T) bool {
	return Includes(s, value)
}

func IndexOf[T comparable](s ComparableSlice[T], value T) (int, bool) {
	return FindIndex(s, func(v T) bool {
		return v == value
	})
}

func (s ComparableSlice[T]) IndexOf(value T) (int, bool) {
	return IndexOf(s, value)
}
