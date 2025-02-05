package godash

// ComparableSlice is a generic slice type that restricts its elements to types that satisfy the comparable constraint.
// It extends the behavior of the standard Slice with some extra functionality that don't require predicates,
// like the Includes method.
type ComparableSlice[T comparable] Slice[T]

func NewComparableSlice[T comparable](elems ...T) ComparableSlice[T] {
	return elems
}

// At behaves as the method [Slice.At]
func (s ComparableSlice[T]) At(index int) T {
	return At[T](Slice[T](s), index)
}

// Every behaves as the method [Slice.Every]
func (s ComparableSlice[T]) Every(p Predicate[T]) bool {
	return Every(Slice[T](s), p)
}

// Fill behaves as the method [Slice.Fill]
func (s ComparableSlice[T]) Fill(value T, positions ...int) ComparableSlice[T] {
	return Fill(Slice[T](s), value, positions...)
}

// Filter behaves as the method [Slice.Filter]
func (s ComparableSlice[T]) Filter(p Predicate[T]) ComparableSlice[T] {
	return Filter(Slice[T](s), p)
}

// Find behaves as the method [Slice.Find]
func (s ComparableSlice[T]) Find(p Predicate[T]) (T, bool) {
	return Find(Slice[T](s), p)
}

// FindIndex behaves as the method [Slice.FindIndex]
func (s ComparableSlice[T]) FindIndex(p Predicate[T]) (int, bool) {
	return FindIndex(Slice[T](s), p)
}

// FindLast behaves as the method [Slice.FindLast]
func (s ComparableSlice[T]) FindLast(p Predicate[T]) (T, bool) {
	return FindLast(Slice[T](s), p)
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
