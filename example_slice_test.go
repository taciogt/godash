package godash_test

import (
	"fmt"
	"github.com/taciogt/godash"
)

func ExampleEvery() {
	isEven := func(i int) bool { return i%2 == 0 }

	allEvens := []int{-2, 0, 2, 4}
	fmt.Println(godash.Every(allEvens, isEven))

	someEvens := []int{0, 1, 2, 3, 4}
	fmt.Println(godash.Every(someEvens, isEven))
	// Output:
	// true
	// false
}

func ExampleSlice_Every() {
	isEven := func(i int) bool { return i%2 == 0 }

	someEvens := godash.NewSlice(0, 1, 2, 3, 4)
	fmt.Println(someEvens.Every(isEven))
	// Output:
	// false
}

func ExampleFilter() {
	s := []int{-3, -2, -1, 0, 1, 2, 3}
	isGreaterThanZero := func(n int) bool {
		return n > 0
	}

	filtered := godash.Filter(s, isGreaterThanZero)
	fmt.Println(filtered)
	// Output:
	// [1 2 3]
}

func ExampleSlice_Filter() {
	isGreaterThanZero := func(n int) bool {
		return n > 0
	}

	s := godash.NewSlice(-3, -2, -1, 0, 1, 2, 3)
	fmt.Println(s.Filter(isGreaterThanZero))
	// Output:
	// [1 2 3]
}

func ExampleFindIndex() {
	s := []string{"a", "ab", "abc", "abcd"}
	findTrigram := func(s string) bool {
		return len(s) == 3
	}

	idx, ok := godash.FindIndex(s, findTrigram)
	fmt.Println(idx, ok)
	// Output:
	// 2 true
}

func ExampleSlice_FindIndex() {
	s := godash.NewSlice("a", "ab", "abc", "abcd")
	findTrigram := func(s string) bool {
		return len(s) == 3
	}

	idx, ok := s.FindIndex(findTrigram)
	fmt.Println(idx, ok)
	// Output:
	// 2 true
}
