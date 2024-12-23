/*
Package godash provides a set of modular and generic functions to manipulate common data structures.
*/
package godash

type Predicate[T any] func(T) bool
