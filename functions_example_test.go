package godash_test

import (
	"fmt"
	"github.com/taciogt/godash"
	"strconv"
)

func ExampleMapperToMustMapper() {
	doubleToString := func(i int) (string, error) {
		return strconv.Itoa(i * 2), nil
	}
	input := []int{0, 1, 2, 3, 4}

	fmt.Println(godash.Map(input, doubleToString))
	// MustMap will panic if the function passed to godash.MapperToMustMapper returns an error
	fmt.Println(godash.MustMap(input, godash.MapperToMustMapper(doubleToString)))

	// Output:
	// [0 2 4 6 8] <nil>
	// [0 2 4 6 8]
}
