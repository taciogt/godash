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

func ExampleMap() {
	doubleToString := func(i int) (string, error) {
		return strconv.Itoa(i * 2), nil
	}
	input := []int{0, 1, 2, 3, 4}

	fmt.Println(Map(input, doubleToString))
	// Output:
	// [0 2 4 6 8] <nil>
}

func ExampleMustMap() {
	doubleToString := func(i int) string {
		return strconv.Itoa(i * 2)
	}
	input := []int{0, 1, 2, 3, 4}

	fmt.Println(MustMap(input, doubleToString))
	// Output:
	// [0 2 4 6 8]
}

func ExampleReduce() {
	sum := func(acc int, curr int) (int, error) {
		return acc + curr, nil
	}
	input := []int{1, 2, 3, 4}

	fmt.Println(Reduce(input, sum, 0))
	// Output:
	// 10 <nil>
}
