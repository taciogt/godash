package godash

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func TestAdd(t *testing.T) {
	type test struct {
		name     string
		input    []string
		element  string
		expected []string
	}

	tests := []test{{
		name:     "Add Without Conflict",
		input:    []string{"Apple", "Banana", "Carrot"},
		element:  "Zucchini",
		expected: []string{"Apple", "Banana", "Carrot", "Zucchini"},
	}, {
		name:     "Add With Conflict",
		input:    []string{"Apple", "Banana", "Carrot"},
		element:  "Apple",
		expected: []string{"Apple", "Banana", "Carrot"},
	}, {
		name:     "Add to Empty Set",
		input:    []string{},
		element:  "Apple",
		expected: []string{"Apple"},
	}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			set := NewSet(tc.input...)
			set.Add(tc.element)

			result := set.Values()
			slices.Sort(result)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("got %q, want %q", result, tc.expected)
			}
		})
	}
}

func TestSet_Delete(t *testing.T) {
	tests := []struct {
		name    string
		initial Set[int]
		arg     int
		want    Set[int]
	}{{
		name:    "Delete existing element",
		initial: NewSet[int](1, 2, 3),
		arg:     2,
		want:    NewSet[int](1, 3),
	}, {
		name:    "Delete non-existing element",
		initial: NewSet[int](1, 2, 3),
		arg:     4,
		want:    NewSet[int](1, 2, 3),
	}, {
		name:    "Delete from empty set",
		initial: NewSet[int](),
		arg:     2,
		want:    NewSet[int](),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.initial.Delete(tt.arg)
			if !reflect.DeepEqual(tt.initial, tt.want) {
				t.Errorf("Set.Delete() = %v, want %v", tt.initial, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type test struct {
		name     string
		input    Set[int]
		expected []int
	}

	tests := []test{{
		name:     "Empty set",
		input:    NewSet[int](),
		expected: []int{},
	}, {
		name:     "Single value",
		input:    NewSet(1),
		expected: []int{1},
	}, {
		name:     "Multiple values",
		input:    NewSet(1, 2, 3),
		expected: []int{1, 2, 3},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Values()
			sort.Ints(got)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
