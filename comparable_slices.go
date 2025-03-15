package godash

// ComparableSlice is a generic slice type that restricts its elements to types that satisfy the comparable constraint.
// It extends the behavior of the standard Slice with some extra functionalities that don't require predicates,
// like the [Includes] and [IndexOf] methods.
type ComparableSlice[T comparable] struct {
	Slice[T]
}

func NewComparableSlice[T comparable](elems ...T) ComparableSlice[T] {
	return ComparableSlice[T]{NewSlice(elems...)}
}

// Includes checks whether the specified value exists within the given ComparableSlice.
// Includes reports whether the specified value is present in the slice. It returns true if any element equals the provided value, and false otherwise.
func Includes[T comparable, S ~[]T](s S, value T) bool {
	return Some(s, func(v T) bool {
		return v == value
	})
}

// Includes behaves exactly like [Includes] function, except it is called directly on the slice.
func (s ComparableSlice[T]) Includes(value T) bool {
	return Includes(s.Slice, value)
}

func IndexOf[T comparable](s ComparableSlice[T], value T) (int, bool) {
	return s.FindIndex(func(v T) bool {
		return v == value
	})
}

func (s ComparableSlice[T]) IndexOf(value T) (int, bool) {
	return IndexOf(s, value)
}
