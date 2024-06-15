package godash

import (
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("an integer slice and a predicate that finds something", func(t *testing.T) {
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		predicate := func(n int) bool { return n > 0 && n%2 == 0 }

		result, ok := Find(s, predicate)
		if result != 2 || !ok {
			t.Errorf("Find(%v, predicate)=%v, %t, want 2, true", s, result, ok)
		}
	})

	t.Run("a string slice and a predicate that finds something", func(t *testing.T) {
		s := []string{"a", "bc", "def", "hijk"}
		predicate := func(s string) bool { return len(s) == 3 }

		result, ok := Find(s, predicate)
		if result != "def" || !ok {
			t.Errorf("Find(%v, predicate)=%v, %t, want def, true", s, result, ok)
		}
	})

	t.Run("a predicate that doesn't find an anything on the slice", func(t *testing.T) {
		s := []string{"a", "bc", "def", "hijk"}
		predicate := func(s string) bool { return len(s) == 10 }

		result, ok := Find(s, predicate)
		if result != "" || ok {
			t.Errorf("Find(%v, predicate)=%v, %t, want def, true", s, result, ok)
		}
	})
}

func TestFindIndex(t *testing.T) {
	t.Run("integer slices", func(t *testing.T) {
		cases := []struct {
			name      string
			arr       []int
			predicate func(int) bool
			wantIndex int
			wantBool  bool
		}{
			{"all positive", []int{1, 2, 3}, func(n int) bool { return n > 0 }, 0, true},
			{"all negative", []int{-1, -2, -3}, func(n int) bool { return n > 0 }, -1, false},
			{"empty", []int{}, func(n int) bool { return n > 0 }, -1, false},
			{"no match", []int{1, 2, 3}, func(n int) bool { return n > 3 }, -1, false},
			{"first match", []int{1, 2, 3}, func(n int) bool { return n == 1 }, 0, true},
			{"last match", []int{1, 2, 3}, func(n int) bool { return n == 3 }, 2, true},
			{"middle match", []int{1, 2, 3}, func(n int) bool { return n == 2 }, 1, true},
		}

		for _, tt := range cases {
			t.Run(tt.name, func(t *testing.T) {
				gotIndex, gotBool := FindIndex(tt.arr, tt.predicate)
				if gotIndex != tt.wantIndex || gotBool != tt.wantBool {
					t.Errorf("FindIndex() = %v, %v; want %v, %v", gotIndex, gotBool, tt.wantIndex, tt.wantBool)
				}
			})
		}

	})

}
