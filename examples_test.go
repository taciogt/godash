package godash

import (
	"fmt"
	"strconv"
)

func Example() {
	s := []int{-2 - 1, 0, 1, 2, 3, 4}
	isGreaterThanZero := func(i int) bool { return i > 0 }

	filtered := Filter(s, isGreaterThanZero)
	fmt.Println(filtered)
	// Output:
	// [1 2 3 4]
}

func ExampleEvery() {
	isEven := func(i int) bool { return i%2 == 0 }

	allEvens := []int{-2, 0, 2, 4}
	someEvens := []int{0, 1, 2, 3, 4}

	fmt.Println(Every(allEvens, isEven))
	fmt.Println(Every(someEvens, isEven))
	// Output:
	// true
	// false
}

func ExampleMap() {
	doubleToString := func(i int) (string, error) {
		return strconv.Itoa(i * 2), nil
	}
	input := []int{0, 1, 2, 3, 4}

	fmt.Println(Map(input, doubleToString))
	// Output:
	// [0 2 4 6 8] <nil>
}
