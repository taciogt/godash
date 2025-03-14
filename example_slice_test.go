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

func ExampleReduce() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := func(acc, val int) (int, error) {
		return acc + val, nil
	}
	result, err := godash.Reduce(numbers, sum, 0)
	fmt.Println(result, err)
	// Output:
	// 15 <nil>
}

func ExampleReduceRight() {
	words := []string{"Go", "is", "fun"}
	concat := func(acc, val string) (string, error) {
		return acc + " " + val, nil
	}
	result, err := godash.ReduceRight(words, concat, "")
	fmt.Println(result, err)
	// Output:
	// fun is Go <nil>
}

func ExampleReverse() {
	// Reverse a slice of integers
	nums := []int{1, 2, 3, 4, 5}
	reversed := godash.Reverse(nums)
	fmt.Println("reversed integers:", reversed)
	fmt.Println("original slice (modified):", nums)

	// Output:
	// reversed integers: [5 4 3 2 1]
	// original slice (modified): [5 4 3 2 1]
}

func ExampleSlice_Reverse() {
	// Create and reverse a slice of integers using the method
	s := godash.NewSlice(1, 2, 3, 4, 5)
	s.Reverse()
	fmt.Println("reversed integers slice:", s)

	// Chaining with other methods
	s2 := godash.NewSlice(1, 2, 3, 4, 5, 6)
	isEven := func(n int) bool { return n%2 == 0 }

	// Chain operations: reverse the slice, then filter for even numbers
	result := s2.Reverse().Filter(isEven)
	fmt.Println("reversed and filtered:", result)

	// Output:
	// reversed integers slice: [5 4 3 2 1]
	// reversed and filtered: [6 4 2]
}

func ExampleToReversed() {
	nums := []int{1, 2, 3, 4, 5}
	reversed := godash.ToReversed(nums)
	fmt.Println("reversed slice:", reversed)
	fmt.Println("original slice (unchanged):", nums)

	// Output:
	// reversed slice: [5 4 3 2 1]
	// original slice (unchanged): [1 2 3 4 5]
}

func ExampleSlice_ToReversed() {
	s := godash.NewSlice(1, 2, 3, 4, 5)
	reversed := s.ToReversed()
	fmt.Println("reversed slice:", reversed)
	fmt.Println("original slice (unchanged):", s)

	// Output:
	// reversed slice: [5 4 3 2 1]
	// original slice (unchanged): [1 2 3 4 5]
}

func ExampleShift() {
	s := []int{1, 2, 3, 4, 5}
	element, ok := godash.Shift(&s)
	fmt.Println("non-empty slice")
	fmt.Println(element, ok)
	fmt.Println(s)

	emptySlice := []int{}
	element, ok = godash.Shift(&emptySlice)
	fmt.Println("empty slice")
	fmt.Println(element, ok)
	fmt.Println(emptySlice)

	// Output:
	// non-empty slice
	// 1 true
	// [2 3 4 5]
	// empty slice
	// 0 false
	// []
}

func ExampleSlice_Shift() {
	s := godash.NewSlice(1, 2, 3, 4, 5)
	element, ok := s.Shift()
	fmt.Println(element, ok)
	fmt.Println(s)
	// Output:
	// 1 true
	// [2 3 4 5]
}
