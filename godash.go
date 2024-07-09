/*Package godash provides a set of modular and generic functions to manipulate common data structures.
 */
package godash

type Slice[T any] []T

type Predicate[T any] func(T) bool

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

// Filter returns a new slice containing only the elements from the given slice that satisfy the provided predicate function.
func Filter[T any, S ~[]T](s S, p Predicate[T]) []T {
	result := make([]T, 0)
	for _, v := range s {
		if p(v) {
			result = append(result, v)
		}
	}
	return result
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
