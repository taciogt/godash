package godash

// Predicate defines a function type that takes a single argument of type T and returns a boolean value.
type Predicate[T any] func(T) bool

// Mapper defines a function type that transforms an input of type TInput into an output of type TOutput,
// possibly returning an error.
type Mapper[TInput any, TOutput any] func(TInput) (TOutput, error)

// MustMapper defines a function type that transforms an input of type TInput into an output of type TOutput without errors.
type MustMapper[TInput any, TOutput any] func(TInput) TOutput

// MapperToMustMapper converts a Mapper function to a MustMapper by panicking if an error occurs during execution.
func MapperToMustMapper[TInput any, TOutput any](mapper Mapper[TInput, TOutput]) MustMapper[TInput, TOutput] {
	return func(value TInput) TOutput {
		result, err := mapper(value)
		if err != nil {
			panic(err)
		}
		return result
	}
}
