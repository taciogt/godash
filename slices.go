package godash

// Slice is a generic type representing a slice of elements with any type.
// The elements of the slice are of type T, where T can be any type.
type Slice[T any] []T

func NewSlice[T any](elems ...T) Slice[T] {
	return elems
}

// ToRaw converts the generic Slice[T] into a standard Go slice []T, maintaining all elements in their order.
func (s Slice[T]) ToRaw() []T {
	return s
}

// At retrieves the element at the specified index of the Slice.
// Negative indexes count backward from the end of the Slice.
func At[T any](s Slice[T], index int) T {
	if index > 0 {
		return s[index]
	} else {
		return s[len(s)+index]
	}
}

// At returns the element at the specified index within the slice.
// Negative indexes are supported to retrieve elements from the end of the slice.
func (s Slice[T]) At(index int) T {
	return At(s, index)
}

//func Exists[T any](s Slice[T], p Predicate[T]) bool {
//	for _, v := range s {
//		if !p(v) {
//			return false
//		}
//	}
//	return true
//}
//
//func (s Slice[T]) Exists(p Predicate[T]) bool {
//	return Every(s, p)
//}

// Every returns true if every element in the given slice satisfies the provided predicate function.
// Otherwise, it returns false.
func Every[T any](s Slice[T], p Predicate[T]) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

// Every behaves exactly as the [Every] function, except it is called directly on the slice to be checked.
func (s Slice[T]) Every(p Predicate[T]) bool {
	return Every(s, p)
}

// Some checks if any element in the slice satisfies the provided predicate function.
func Some[T any, S ~[]T](s S, p Predicate[T]) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}

// Some behaves exactly as the [Some] function, except it is called directly on the slice to be checked.
func (s Slice[T]) Some(p Predicate[T]) bool {
	return Some(s, p)
}

// Fill replaces elements of a slice with the specified value within the given range or entire slice
// if no range is provided.
func Fill[T any, S ~[]T](s S, value T, positions ...int) []T {
	newSlice := make([]T, len(s))
	copy(newSlice, s)

	lowerBound := 0
	if len(positions) > 0 {
		lowerBound = positions[0]
	}
	upperBound := len(s)
	if len(positions) > 1 {
		upperBound = positions[1]
	}

	for i := range newSlice {
		if (i >= lowerBound) && (i <= upperBound) {
			newSlice[i] = value
		}
	}
	return newSlice
}

// Fill replaces elements of the slice with the specified value within the given range or entire slice
// if no range is provided.
func (s Slice[T]) Fill(value T, positions ...int) Slice[T] {
	return Fill(s, value, positions...)
}

// Filter applies a predicate function to each element in the given slice and returns a new slice
// containing only the elements for which the predicate function returns true.
func Filter[T any, S ~[]T](s S, p Predicate[T]) []T {
	result := make([]T, 0)
	for _, v := range s {
		if p(v) {
			result = append(result, v)
		}
	}
	return result
}

// Filter behaves exactly as the [Filter] function, except it is called directly on the slice to be filtered.
func (s Slice[T]) Filter(p Predicate[T]) Slice[T] {
	return Filter(s, p)
}

// Find returns the first element in the given slice that satisfies the provided predicate function.
// If no element satisfies the predicate, the zero value of the element type is returned along with false.
func Find[T any, S ~[]T](s S, p Predicate[T]) (T, bool) {
	if i, ok := FindIndex(s, p); ok {
		return s[i], true
	}

	var zero T
	return zero, false
}

// Find behaves exactly as the [Find] function, except it is called directly on the slice to be searched.
func (s Slice[T]) Find(p Predicate[T]) (T, bool) {
	return Find(s, p)
}

// FindIndex returns the index of the first element in the given slice that satisfies
// the provided predicate function. If no element satisfies the predicate,
// -1 is returned along with false.
func FindIndex[T any, S ~[]T](s S, p Predicate[T]) (int, bool) {
	for i, v := range s {
		if p(v) {
			return i, true
		}
	}
	return -1, false
}

// FindIndex behaves exactly as the [FindIndex] function, except it is called directly on the slice to be searched.
func (s Slice[T]) FindIndex(p Predicate[T]) (int, bool) {
	return FindIndex(s, p)
}

// FindLast returns the last element in the given slice that satisfies the provided predicate function.
// If no element satisfies the predicate, the zero value of the element type is returned along with false.
func FindLast[T any, S ~[]T](s S, p Predicate[T]) (T, bool) {
	if i, ok := FindLastIndex(s, p); ok {
		return s[i], true
	}

	var zero T
	return zero, false
}

// FindLast behaves exactly as the [FindLast] function, except it is called directly on the slice to be searched.
func (s Slice[T]) FindLast(p Predicate[T]) (T, bool) {
	return FindLast(s, p)
}

// FindLastIndex returns the index of the last element in the given slice that satisfies
// the provided predicate function. If no element satisfies the predicate,
// -1 is returned along with false.
func FindLastIndex[T any, S ~[]T](s S, p Predicate[T]) (int, bool) {
	for i := len(s) - 1; i >= 0; i-- {
		if p(s[i]) {
			return i, true
		}
	}
	return -1, false
}

// FindLastIndex behaves exactly as the [FindLastIndex] function,
// except it is called directly on the slice to be searched.
func (s Slice[T]) FindLastIndex(p Predicate[T]) (int, bool) {
	return FindLastIndex(s, p)
}

// ForEach applies the provided function f to each element in the slice s.
// The function f should take an index i and a value v as arguments.
// The index i represents the position of the element in the slice s,
// and the value v represents the current element.
// Note that this function does not modify the elements in the slice s.
// Example usage:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	ForEach(numbers, func(i int, v int) {
//	    fmt.Printf("index: %d, value: %d\n", i, v)
//	})
func ForEach[T any, S ~[]T](s S, f func(i int, v T)) {
	for i, v := range s {
		f(i, v)
	}
}

// ForEach behaves exactly like [ForEach] function, except it is called directly on the slice.
func (s Slice[T]) ForEach(f func(i int, v T)) {
	ForEach(s, f)
}

// Map takes in a slice of input values and a mapper function, and applies the mapper function to each
// input value. It returns a new slice containing the mapped values. If any error occurs during the mapping
// process, the function aborts and returns nil along with the error. Otherwise, it returns the new slice
// of mapped values and a nil error.
func Map[TInput any, TOutput any, S ~[]TInput](s S, mapper func(TInput) (TOutput, error)) ([]TOutput, error) {
	result := make([]TOutput, len(s))
	for i, value := range s {
		mapped, err := mapper(value)
		if err != nil {
			return nil, err
		}
		result[i] = mapped
	}
	return result, nil
}

// Pop removes and returns the last element from the slice pointed to by `s`.
// If the slice is empty, it returns the zero value of type `T` and `false`.
// The function modifies the original slice by updating it with one less element.
func Pop[T any, S ~*[]T](s S) (T, bool) {
	length := len(*s)
	if length == 0 {
		var zero T
		return zero, false
	}

	lastElem := (*s)[length-1]
	*s = (*s)[:length-1]

	return lastElem, true
}

// Pop behaves exactly like [Pop] function, except it is called directly on the slice.
func (s *Slice[T]) Pop() (T, bool) {
	rawSlice := s.ToRaw()
	result, ok := Pop(&rawSlice)
	*s = NewSlice(rawSlice...)
	return result, ok
}

// Push appends the provided values to the slice and returns the new size of the slice.
// The function modifies the original slice in place.
func Push[T any, S ~*[]T](s S, value ...T) (length int) {
	*s = append(*s, value...)
	return len(*s)
}

// Push behaves exactly like [Push] function, except it is called directly on the slice.
func (s *Slice[T]) Push(value ...T) (length int) {
	rawSlice := s.ToRaw()
	length = Push(&rawSlice, value...)
	*s = NewSlice(rawSlice...)
	return length
}

// Reduce iterates over the elements in the slice and applies the reducer function to each element,
// accumulating the result in the initial value. It returns the final accumulated value and an error,
// if any occurred during the reduction process. The reducer function takes two arguments: the current
// accumulated value and the current element value, and returns the updated accumulated value and an
// error, if any occurred.
func Reduce[TIn any, TOut any, S ~[]TIn](s S, reducer func(acc TOut, curr TIn) (TOut, error), initialValue TOut) (TOut, error) {
	result := initialValue
	var err error
	for _, v := range s {
		result, err = reducer(result, v)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

// ReduceRight applies a reducer function to each element of the slice, starting from the right (end),
// resulting in a single output value. The function is called with an accumulator and each element
// from right to left.
func ReduceRight[TIn any, TOut any, S ~[]TIn](s S, reducer func(acc TOut, curr TIn) (TOut, error), initialValue TOut) (TOut, error) {
	acc := initialValue

	for i := len(s) - 1; i >= 0; i-- {
		var err error
		acc, err = reducer(acc, s[i])
		if err != nil {
			return acc, err
		}
	}

	return acc, nil
}

// Reverse reverses the elements of a slice in place.
// This function modifies the original slice and returns it.
func Reverse[T any, S ~[]T](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Reverse reverses the elements of the slice in place.
// This method modifies the original slice and returns the modified slice for chaining.
func (s Slice[T]) Reverse() Slice[T] {
	return Reverse(s)
}

// ToReversed returns a new slice with the elements of the input slice reversed,
// preserving the original slice unmodified.
func ToReversed[T any, S ~[]T](s S) []T {
	result := make([]T, len(s))
	l := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		result[i] = s[l-1-i]
	}
	return result
}

// ToReversed creates and returns a new slice with elements in reverse order,
// leaving the original slice unchanged.
func (s Slice[T]) ToReversed() Slice[T] {
	return ToReversed(s)
}

// Shift removes and returns the first element of the provided slice, modifying the original slice.
// If the slice is empty, it returns the zero value of type `T` and `false`.
func Shift[T any, S ~*[]T](s S) (T, bool) {
	length := len(*s)
	if length == 0 {
		var zero T
		return zero, false
	}
	firstElem := (*s)[0]
	*s = (*s)[1:]
	return firstElem, true
}

// Shift removes and returns the first element of the slice, updating the original slice.
// If the slice is empty, it returns the zero value of type `T` and `false`.
func (s *Slice[T]) Shift() (T, bool) {
	rawSlice := s.ToRaw()
	result, ok := Shift(&rawSlice)
	*s = NewSlice(rawSlice...)
	return result, ok
}

// Unshift prepends one or more values to the beginning of the provided slice pointer
// and returns the new length of the slice.
//
// Parameters:
//   - value: One or more values of type T to be added to the beginning of the slice.
//
// Returns:
//   - length: The new length of the slice after the values have been prepended.
func Unshift[T any, S ~*[]T](s S, value ...T) (length int) {
	*s = append(value, *s...)
	return len(*s)
}

// Unshift prepends one or more values to the beginning of the slice, updating the original slice.
// It returns the new length of the slice.
//
// Parameters:
//   - value: One or more values of type T to be added to the beginning of the slice.
//
// Returns:
//   - length: The new length of the slice after the values have been prepended.
func (s *Slice[T]) Unshift(value ...T) (length int) {
	rawSlice := s.ToRaw()
	length = Unshift(&rawSlice, value...)
	*s = NewSlice(rawSlice...)
	return length
}
