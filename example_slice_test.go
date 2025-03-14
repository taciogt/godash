package godash_test

import (
	"fmt"
	"github.com/taciogt/godash"
	"strings"
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

func ExampleSlice_Fill() {
	s := godash.NewSlice(1, 2, 3, 4, 5)
	fmt.Println(s.Fill(9))
	fmt.Println(s.Fill(9, 2))
	fmt.Println(s.Fill(9, 1, 3))
	// Output:
	// [9 9 9 9 9]
	// [1 2 9 9 9]
	// [1 9 9 9 5]
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

func ExampleFind() {
	s := []string{"A", "Ab", "aBc", "ab-cd", "efg"}
	isLowerCase := func(s string) bool {
		return strings.ToLower(s) == s
	}

	result, ok := godash.Find(s, isLowerCase)
	fmt.Println(result, ok)
	// Output:
	// ab-cd true
}

func ExampleSlice_Find() {
	s := godash.NewSlice("A", "Ab", "aBc", "abc")
	isLowerCase := func(s string) bool {
		return strings.ToLower(s) == s
	}

	result, ok := s.Find(isLowerCase)
	fmt.Println(result, ok)
	// Output:
	// abc true
}

func ExamplePop() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(godash.Pop(&s))
	fmt.Println(s)
	// Output:
	// 5 true
	// [1 2 3 4]
}

func ExampleSlice_Pop() {
	s := godash.NewSlice(1, 2, 3, 4, 5)
	fmt.Println(s.Pop())
	fmt.Println(s)
	// Output:
	// 5 true
	// [1 2 3 4]
}

func ExamplePush() {
	s := []int{1, 2, 3}
	fmt.Println(godash.Push(&s, 4, 5))
	fmt.Println(s)
	// Output:
	// 5
	// [1 2 3 4 5]
}

func ExampleSlice_Push() {
	s := godash.NewSlice(1, 2, 3, 4, 5)
	fmt.Println(s.Push(6, 7, 8))
	fmt.Println(s)
	// Output:
	// 8
	// [1 2 3 4 5 6 7 8]
}
