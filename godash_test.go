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
