package godash_test

import (
	"fmt"
	"github.com/taciogt/godash"
)

func ExampleComparableSlice_Includes() {
	slice := godash.NewComparableSlice(1, 2, 3)
	fmt.Println(slice.Includes(2))
	fmt.Println(slice.Includes(4))

	// Output:
	// true
	// false
}
