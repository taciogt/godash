package godash

// Slice is a generic type representing a slice of elements of any type.
// The elements of the slice are of type T, where T can be any type.
type Slice[T any] []T

func NewSlice[T any](elems ...T) Slice[T] {
	return elems
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
