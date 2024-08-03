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
