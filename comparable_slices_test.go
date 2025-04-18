package godash

import (
	"reflect"
	"slices"
	"testing"
)

func TestComparableSlice_At(t *testing.T) {
	type test struct {
		name     string
		slice    ComparableSlice[int]
		index    int
		expected int
	}
	tests := []test{{
		name:     "find element at an existing positive index",
		slice:    NewComparableSlice(1, 2, 3, 4, 5),
		index:    1,
		expected: 2,
	}, {
		name:     "find element at an existing negative index",
		slice:    NewComparableSlice(1, 2, 3, 4, 5),
		index:    -1,
		expected: 5,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.slice.At(tt.index)
			if got != tt.expected {
				t.Errorf("At() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestComparableSlice_Every(t *testing.T) {
	type test struct {
		name      string
		slice     ComparableSlice[int]
		predicate func(int) bool
		expected  bool
	}
	isPositive := func(n int) bool {
		return n >= 0
	}

	tests := []test{{
		name:      "every element is positive",
		slice:     NewComparableSlice(1, 2, 3, 4, 5),
		predicate: isPositive,
		expected:  true,
	}, {
		name:      "not every element is positive",
		slice:     NewComparableSlice(-1, 0, 1, 2, 3, 4),
		predicate: isPositive,
		expected:  false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.slice.Every(tt.predicate)
			if got != tt.expected {
				t.Errorf("Every() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestComparableSlice_Fill(t *testing.T) {
	tests := []struct {
		name          string
		slice         ComparableSlice[int]
		value         int
		positions     []int
		expectedSlice ComparableSlice[int]
	}{{
		name:          "fill entire slice with a value",
		slice:         NewComparableSlice(1, 2, 3, 4),
		value:         5,
		positions:     nil,
		expectedSlice: NewComparableSlice(5, 5, 5, 5),
	}, {
		name:          "fill slice within range",
		slice:         NewComparableSlice(1, 2, 3, 4, 5),
		value:         9,
		positions:     []int{1, 3},
		expectedSlice: NewComparableSlice(1, 9, 9, 9, 5),
	}, {
		name:          "fill with single position (treat it as start index)",
		slice:         NewComparableSlice(10, 20, 30, 40),
		value:         7,
		positions:     []int{2},
		expectedSlice: NewComparableSlice(10, 20, 7, 7),
	}, {
		name:          "fill with range exceeding slice boundaries",
		slice:         NewComparableSlice(1, 2, 3),
		value:         0,
		positions:     []int{1, 5},
		expectedSlice: NewComparableSlice(1, 0, 0),
	}, {
		name:          "fill entire slice for empty positions parameter",
		slice:         NewComparableSlice(6, 7, 8),
		value:         3,
		positions:     []int{},
		expectedSlice: NewComparableSlice(3, 3, 3),
	}, {
		name:          "fill an empty slice",
		slice:         NewComparableSlice[int](),
		value:         9,
		positions:     nil,
		expectedSlice: NewComparableSlice[int](),
	}, {
		name:          "negative range boundaries are ignored",
		slice:         NewComparableSlice(1, 2, 3),
		value:         5,
		positions:     []int{-3, -1},
		expectedSlice: NewComparableSlice(1, 2, 3), // no change
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlice := tt.slice.Fill(tt.value, tt.positions...)
			if !slices.Equal(gotSlice, tt.expectedSlice.Slice) {
				t.Errorf("Fill() = %v, want %v", gotSlice, tt.expectedSlice)
			}
		})
	}
}

func TestComparableSlice_Filter(t *testing.T) {
	type test struct {
		name      string
		slice     ComparableSlice[int]
		predicate func(int) bool
		expected  ComparableSlice[int]
	}
	isEven := func(n int) bool {
		return n%2 == 0
	}
	tests := []test{{
		name:      "filter positive numbers",
		slice:     NewComparableSlice(1, 2, 3, 4, 5),
		predicate: isEven,
		expected:  NewComparableSlice(2, 4),
	}, {
		name:      "filter returns empty slice",
		slice:     NewComparableSlice(1, 3, 5),
		predicate: isEven,
		expected:  NewComparableSlice[int](),
	}, {
		name:      "filter on empty slice",
		slice:     NewComparableSlice[int](),
		predicate: isEven,
		expected:  NewComparableSlice[int](),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.slice.Filter(tt.predicate)
			if !slices.Equal(got, tt.expected.Slice) {
				t.Errorf("Filter() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestComparableSlice_Find(t *testing.T) {
	type test struct {
		name          string
		slice         ComparableSlice[int]
		predicate     func(int) bool
		expectedValue int
		expectedOk    bool
	}
	isEven := func(n int) bool {
		return n%2 == 0
	}
	isOdd := func(n int) bool {
		return n%2 != 0
	}

	tests := []test{{
		name:          "find even number",
		slice:         NewComparableSlice(1, 3, 5, 4, 6, 8),
		predicate:     isEven,
		expectedValue: 4,
		expectedOk:    true,
	}, {
		name:          "do not find even number",
		slice:         NewComparableSlice(1, 3, 5),
		predicate:     isEven,
		expectedValue: 0,
		expectedOk:    false,
	}, {
		name:          "find odd number",
		slice:         NewComparableSlice(1, 2, 3, 4, 5),
		predicate:     isOdd,
		expectedValue: 1,
		expectedOk:    true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := tt.slice.Find(tt.predicate)
			if gotValue != tt.expectedValue || gotOk != tt.expectedOk {
				t.Errorf("Find() = %v, %v, want %v, %v", gotValue, gotOk, tt.expectedValue, tt.expectedOk)
			}
		})
	}
}

func TestComparableSlice_FindIndex(t *testing.T) {
	type test struct {
		name          string
		slice         ComparableSlice[int]
		predicate     func(int) bool
		expectedIndex int
		expectedOk    bool
	}

	isEven := func(n int) bool {
		return n%2 == 0
	}
	isOdd := func(n int) bool {
		return n%2 != 0
	}
	tests := []test{{
		name:          "find even number",
		slice:         NewComparableSlice(1, 3, 5, 4, 6, 8),
		predicate:     isEven,
		expectedIndex: 3,
		expectedOk:    true,
	}, {
		name:          "do not find even number",
		slice:         NewComparableSlice(1, 3, 5),
		predicate:     isEven,
		expectedIndex: -1,
		expectedOk:    false,
	}, {
		name:          "find odd number",
		slice:         NewComparableSlice(1, 2, 3, 4, 5),
		predicate:     isOdd,
		expectedIndex: 0,
		expectedOk:    true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotOk := tt.slice.FindIndex(tt.predicate)
			if gotIndex != tt.expectedIndex || gotOk != tt.expectedOk {
				t.Errorf("FindIndex() = %v, %v, want %v, %v", gotIndex, gotOk, tt.expectedIndex, tt.expectedOk)
			}
		})
	}
}

func TestComparableSlice_FindLast(t *testing.T) {
	type test struct {
		name          string
		slice         ComparableSlice[int]
		predicate     func(int) bool
		expectedValue int
		expectedOk    bool
	}
	isEven := func(n int) bool {
		return n%2 == 0
	}
	tests := []test{{
		name:          "find even number",
		slice:         NewComparableSlice(1, 3, 5, 4, 6, 8),
		predicate:     isEven,
		expectedValue: 8,
		expectedOk:    true,
	}, {
		name:          "do not find even number",
		slice:         NewComparableSlice(1, 3, 5),
		predicate:     isEven,
		expectedValue: 0,
		expectedOk:    false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := tt.slice.FindLast(tt.predicate)
			if gotValue != tt.expectedValue || gotOk != tt.expectedOk {
				t.Errorf("FindLast() = %v, %v, want %v, %v", gotValue, gotOk, tt.expectedValue, tt.expectedOk)
			}
		})
	}
}

func TestComparableSlice_FindLastIndex(t *testing.T) {
	type test struct {
		name          string
		slice         ComparableSlice[int]
		predicate     func(int) bool
		expectedIndex int
		expectedOk    bool
	}
	isEven := func(n int) bool {
		return n%2 == 0
	}

	tests := []test{{
		name:          "find even number",
		slice:         NewComparableSlice(1, 3, 5, 4, 6, 8),
		predicate:     isEven,
		expectedIndex: 5,
		expectedOk:    true,
	}, {
		name:          "do not find even number",
		slice:         NewComparableSlice(1, 3, 5),
		predicate:     isEven,
		expectedIndex: -1,
		expectedOk:    false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotOk := tt.slice.FindLastIndex(tt.predicate)
			if gotIndex != tt.expectedIndex || gotOk != tt.expectedOk {
				t.Errorf("FindLastIndex() = %v, %v, want %v, %v", gotIndex, gotOk, tt.expectedIndex, tt.expectedOk)
			}
		})
	}
}

func TestComparableSlice_ForEach(t *testing.T) {
	t.Run("empty slice does not call the for each function", func(t *testing.T) {
		NewComparableSlice[int]().ForEach(func(i int, v int) {
			t.Error("function called for empty slice")
		})
	})

	t.Run("function does not change input slice", func(t *testing.T) {
		s := NewComparableSlice(1, 2, 3)
		s.ForEach(func(i int, v int) {
			v *= 2
		})
		if !slices.Equal(s.Slice, NewComparableSlice(1, 2, 3).Slice) {
			t.Error("function should not change input slice")
		}
	})

	t.Run("function traverse through all elements", func(t *testing.T) {
		s := NewComparableSlice(1, 2, 3)

		traversedItems := make([][]int, 0)
		s.ForEach(func(i int, v int) {
			traversedItems = append(traversedItems, []int{i, v})
		})

		expected := [][]int{{0, 1}, {1, 2}, {2, 3}}
		if !reflect.DeepEqual(traversedItems, expected) {
			t.Error("function should traverse through all items")
		}
	})
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		search   int
		expected bool
	}{{
		name:     "element exists in slice",
		slice:    []int{1, 2, 3, 4, 5},
		search:   3,
		expected: true,
	}, {
		name:     "element does not exist in slice",
		slice:    []int{1, 2, 3, 4, 5},
		search:   6,
		expected: false,
	}, {
		name:     "empty slice",
		slice:    []int{},
		search:   1,
		expected: false,
	}, {
		name:     "nil slice",
		slice:    nil,
		search:   1,
		expected: false,
	}, {
		name:     "slice with duplicates contains the element",
		slice:    []int{1, 2, 2, 3, 3},
		search:   2,
		expected: true,
	}, {
		name:     "slice with negative numbers contains the element",
		slice:    []int{-1, -2, -3, -4},
		search:   -3,
		expected: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("standalone function", func(t *testing.T) {
				got := Includes(tt.slice, tt.search)
				if got != tt.expected {
					t.Errorf("Includes() = %v, want %v", got, tt.expected)
				}
			})

			t.Run("receiver method", func(t *testing.T) {
				slice := NewComparableSlice(tt.slice...)
				got := slice.Includes(tt.search)
				if got != tt.expected {
					t.Errorf("Includes() = %v, want %v", got, tt.expected)
				}
			})
		})
	}

	t.Run("string slices", func(t *testing.T) {
		slice := NewComparableSlice[string]("apple", "banana", "cherry")
		if !slice.Includes("banana") {
			t.Errorf("Includes() = false, want true for 'banana'")
		}
		if slice.Includes("grape") {
			t.Errorf("Includes() = true, want false for 'grape'")
		}
	})

	t.Run("slices of custom structs", func(t *testing.T) {
		type customStruct struct {
			id   int
			name string
		}
		slice := NewComparableSlice[customStruct](
			customStruct{id: 1, name: "Alice"},
			customStruct{id: 2, name: "Bob"},
		)
		search := customStruct{id: 1, name: "Alice"}
		if !slice.Includes(search) {
			t.Errorf("Includes() = false, want true for %v", search)
		}
		searchNotExist := customStruct{id: 3, name: "Charlie"}
		if slice.Includes(searchNotExist) {
			t.Errorf("Includes() = true, want false for %v", searchNotExist)
		}
	})
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name           string
		slice          ComparableSlice[int]
		search         int
		expectedIndex  int
		expectedExists bool
	}{{
		name:           "element exists in slice",
		slice:          NewComparableSlice(1, 2, 3, 4, 5),
		search:         3,
		expectedIndex:  2,
		expectedExists: true,
	}, {
		name:           "element does not exist in slice",
		slice:          NewComparableSlice(1, 2, 3, 4, 5),
		search:         6,
		expectedIndex:  -1,
		expectedExists: false,
	}, {
		name:           "empty slice",
		slice:          ComparableSlice[int]{},
		search:         1,
		expectedIndex:  -1,
		expectedExists: false,
	}, {
		name:           "first element in slice",
		slice:          NewComparableSlice(7, 8, 9),
		search:         7,
		expectedIndex:  0,
		expectedExists: true,
	}, {
		name:           "last element in slice",
		slice:          NewComparableSlice(10, 20, 30),
		search:         30,
		expectedIndex:  2,
		expectedExists: true,
	}, {
		name:           "slice with duplicates",
		slice:          NewComparableSlice(1, 2, 2, 3, 3),
		search:         2,
		expectedIndex:  1,
		expectedExists: true,
	}, {
		name:           "slice with negative numbers",
		slice:          NewComparableSlice(-1, -2, -3, -4),
		search:         -3,
		expectedIndex:  2,
		expectedExists: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotIndex, gotExists := tt.slice.IndexOf(tt.search)
			if gotIndex != tt.expectedIndex || gotExists != tt.expectedExists {
				t.Errorf("IndexOf() = %v, %v, want %v, %v", gotIndex, gotExists, tt.expectedIndex, tt.expectedExists)
			}
		})
	}

	t.Run("string slices", func(t *testing.T) {
		slice := NewComparableSlice[string]("apple", "banana", "cherry")
		index, exists := slice.IndexOf("banana")
		if index != 1 || !exists {
			t.Errorf("IndexOf() = %v, %v, want 1, true for 'banana'", index, exists)
		}

		index, exists = slice.IndexOf("grape")
		if index != -1 || exists {
			t.Errorf("IndexOf() = %v, %v want 0, false for 'grape'", index, exists)
		}
	})

	t.Run("slices of custom structs", func(t *testing.T) {
		type customStruct struct {
			id   int
			name string
		}
		slice := NewComparableSlice[customStruct](
			customStruct{id: 1, name: "Alice"},
			customStruct{id: 2, name: "Bob"},
			customStruct{id: 3, name: "Charlie"},
		)
		search := customStruct{id: 2, name: "Bob"}
		index, exists := slice.IndexOf(search)
		if index != 1 || !exists {
			t.Errorf("IndexOf() = %v, %v, want 1, true for %v", index, exists, search)
		}

		searchNotExist := customStruct{id: 4, name: "David"}
		index, exists = slice.IndexOf(searchNotExist)
		if index != -1 || exists {
			t.Errorf("IndexOf() = %v, %v, want -1 for %v", index, exists, searchNotExist)
		}
	})
}

func TestComparableSlice_Pop(t *testing.T) {
	tests := []struct {
		name           string
		slice          []int
		expectedResult int
		expectedOk     bool
		expectedSlice  []int
	}{{
		name:           "Pop from non-empty slice",
		slice:          []int{1, 2, 3},
		expectedResult: 3,
		expectedOk:     true,
		expectedSlice:  []int{1, 2},
	}, {
		name:           "Pop from single-element slice",
		slice:          []int{10},
		expectedResult: 10,
		expectedOk:     true,
		expectedSlice:  []int{},
	}, {
		name:           "Pop from empty slice",
		slice:          []int{},
		expectedResult: 0, // Default value for int
		expectedOk:     false,
		expectedSlice:  []int{},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewComparableSlice(tt.slice...)

			gotResult, gotOk := s.Pop()
			if gotResult != tt.expectedResult || gotOk != tt.expectedOk {
				t.Errorf("Slice.Pop() = (%v, %v), want (%v, %v)", gotResult, gotOk, tt.expectedResult, tt.expectedOk)
			}

			if !reflect.DeepEqual(s.ToRaw(), tt.expectedSlice) {
				t.Errorf("resulting %v, want %v", s.ToRaw(), tt.expectedSlice)
			}
		})
	}
}

func TestComparableSlice_Push(t *testing.T) {
	tests := []struct {
		name           string
		slice          []int
		value          int
		expectedResult []int
		expectedLength int
	}{{
		name:           "Push to empty slice",
		slice:          []int{},
		value:          1,
		expectedResult: []int{1},
		expectedLength: 1,
	}, {
		name:           "Push to non-empty slice",
		slice:          []int{1, 2, 3},
		value:          4,
		expectedResult: []int{1, 2, 3, 4},
		expectedLength: 4,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewComparableSlice(tt.slice...)
			gotLength := s.Push(tt.value)
			if !reflect.DeepEqual(s.ToRaw(), tt.expectedResult) || gotLength != len(tt.expectedResult) {
				t.Errorf("Slice.Push() got slice=%v length=%d, want slice=%v length=%d",
					s.ToRaw(), gotLength, tt.expectedResult, tt.expectedLength)
			}
		})
	}
}
