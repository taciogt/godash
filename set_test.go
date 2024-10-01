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

func TestSet_Has(t *testing.T) {
	tests := []struct {
		name     string
		set      Set[int]
		element  int
		expected bool
	}{{
		name:     "element exists",
		set:      NewSet[int](1, 2, 3, 4, 5),
		element:  3,
		expected: true,
	}, {
		name:     "element doesn't exist",
		set:      NewSet[int](1, 2, 3, 4, 5),
		element:  6,
		expected: false,
	}, {
		name:     "empty set",
		set:      NewSet[int](),
		element:  1,
		expected: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Has(tt.element); got != tt.expected {
				t.Errorf("Set.Has() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want int
	}{{
		name: "Empty set",
		s:    NewSet[int](),
		want: 0,
	}, {
		name: "Single element",
		s:    NewSet[int](1),
		want: 1,
	}, {
		name: "Multiple elements",
		s:    NewSet[int](1, 2, 3),
		want: 3,
	}, {
		name: "Duplicate elements",
		s:    NewSet[int](1, 1, 2, 2, 3, 3),
		want: 3,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Set.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	tests := []struct {
		name string
		s1   Set[int]
		s2   Set[int]
		want Set[int]
	}{
		{name: "EmptySets", s1: NewSet[int](), s2: NewSet[int](), want: NewSet[int]()},
		{name: "OneEmptySet", s1: NewSet(1, 2, 3), s2: NewSet[int](), want: NewSet[int]()},
		{name: "NoMatch", s1: NewSet(1, 2, 3), s2: NewSet(4, 5, 6), want: NewSet[int]()},
		{name: "PartialMatch", s1: NewSet(1, 2, 3), s2: NewSet(2, 3, 4), want: NewSet(2, 3)},
		{name: "FullMatch", s1: NewSet(1, 2, 3), s2: NewSet(1, 2, 3), want: NewSet(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s1.Intersection(tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name string
		set1 Set[int]
		set2 Set[int]
		want Set[int]
	}{
		{
			name: "union two normal sets",
			set1: NewSet(1, 2, 3),
			set2: NewSet(3, 4, 5),
			want: NewSet(1, 2, 3, 4, 5),
		}, {
			name: "union with empty set",
			set1: NewSet(1, 2, 3),
			set2: NewSet[int](),
			want: NewSet(1, 2, 3),
		}, {
			name: "union two empty sets",
			set1: NewSet[int](),
			set2: NewSet[int](),
			want: NewSet[int](),
		}, {
			name: "union with set with some duplicates",
			set1: NewSet(1, 2, 3),
			set2: NewSet(2, 2, 2),
			want: NewSet(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set1.Union(tt.set2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set[%T].Union() = %v, want %v", tt.set1, got, tt.want)
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
