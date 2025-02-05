package godash

import (
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

	t.Run(t.Name(), func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.slice.At(tt.index)
				if got != tt.expected {
					t.Errorf("At() = %v, want %v", got, tt.expected)
				}
			})
		}
	})
}

func TestIncludes(t *testing.T) {
	tests := []struct {
		name     string
		slice    ComparableSlice[int]
		search   int
		expected bool
	}{{
		name:     "element exists in slice",
		slice:    ComparableSlice[int]{1, 2, 3, 4, 5},
		search:   3,
		expected: true,
	}, {
		name:     "element does not exist in slice",
		slice:    ComparableSlice[int]{1, 2, 3, 4, 5},
		search:   6,
		expected: false,
	}, {
		name:     "empty slice",
		slice:    ComparableSlice[int]{},
		search:   1,
		expected: false,
	}, {
		name:     "slice with duplicates contains the element",
		slice:    ComparableSlice[int]{1, 2, 2, 3, 3},
		search:   2,
		expected: true,
	}, {
		name:     "slice with negative numbers contains the element",
		slice:    ComparableSlice[int]{-1, -2, -3, -4},
		search:   -3,
		expected: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.slice.Includes(tt.search)
			if got != tt.expected {
				t.Errorf("Includes() = %v, want %v", got, tt.expected)
			}
		})
	}

	t.Run("string slices", func(t *testing.T) {
		slice := ComparableSlice[string]{"apple", "banana", "cherry"}
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
		slice := ComparableSlice[customStruct]{
			{id: 1, name: "Alice"},
			{id: 2, name: "Bob"},
		}
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
		slice:          ComparableSlice[int]{1, 2, 3, 4, 5},
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
		slice:          ComparableSlice[int]{7, 8, 9},
		search:         7,
		expectedIndex:  0,
		expectedExists: true,
	}, {
		name:           "last element in slice",
		slice:          ComparableSlice[int]{10, 20, 30},
		search:         30,
		expectedIndex:  2,
		expectedExists: true,
	}, {
		name:           "slice with duplicates",
		slice:          ComparableSlice[int]{1, 2, 2, 3, 3},
		search:         2,
		expectedIndex:  1,
		expectedExists: true,
	}, {
		name:           "slice with negative numbers",
		slice:          ComparableSlice[int]{-1, -2, -3, -4},
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
		slice := ComparableSlice[string]{"apple", "banana", "cherry"}
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
		slice := ComparableSlice[customStruct]{
			{id: 1, name: "Alice"},
			{id: 2, name: "Bob"},
			{id: 3, name: "Charlie"},
		}
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
