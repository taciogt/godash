package godash

type Mapper[TInput any, TOutput any] func(TInput) (TOutput, error)
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
