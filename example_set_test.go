package godash_test

import (
	"fmt"
	"github.com/taciogt/godash"
)

func ExampleSet_Add() {
	s := godash.NewSet(1, 2, 3, 4)

	s.Add(5)
	fmt.Println(s)

	s.Add(2)
	fmt.Println(s)

	// Output:
	// set{1, 2, 3, 4, 5}
	// set{1, 2, 3, 4, 5}
}

func ExampleSet_Delete() {
	s := godash.NewSet(1, 2, 3, 4)

	s.Delete(5)
	fmt.Println(s)

	s.Delete(2)
	fmt.Println(s)

	// Output:
	// set{1, 2, 3, 4}
	// set{1, 3, 4}
}

func ExampleSet_Has() {
	s := godash.NewSet(1, 2, 3, 4)

	fmt.Println("Has(1) =", s.Has(1))
	fmt.Println("Has(5) =", s.Has(5))

	// Output:
	// Has(1) = true
	// Has(5) = false
}
