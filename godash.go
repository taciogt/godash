package godash

type Slice[T any] []T

//type Slice interface {
//}

func Find[T any, S ~[]T](s S, predicate func(v T) bool) (T, bool) {
	for _, v := range s {
		if predicate(v) {
			return v, true
		}
	}

	var zero T
	return zero, false
}
