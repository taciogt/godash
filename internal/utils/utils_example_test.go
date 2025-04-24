package utils_test

import (
	"fmt"
	"github.com/taciogt/godash/internal/utils"
)

func ExampleIsEven() {
	// Check if numbers are even
	fmt.Println(utils.IsEven(0))  // zero
	fmt.Println(utils.IsEven(2))  // positive even
	fmt.Println(utils.IsEven(3))  // positive odd
	fmt.Println(utils.IsEven(-4)) // negative even
	fmt.Println(utils.IsEven(-5)) // negative odd

	// Output:
	// true
	// true
	// false
	// true
	// false
}
