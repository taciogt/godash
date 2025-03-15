package godash

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type customStruct struct {
	int    int
	string string
}

type typeAlias int

func TestAt(t *testing.T) {
	t.Run("tests for At() method", func(t *testing.T) {
		tcs := []struct {
			name        string
			input       []int
			index       int
			expected    int
			shouldPanic bool
		}{
			{name: "Positive index", input: []int{1, 2, 3, 4, 5}, index: 2, expected: 3, shouldPanic: false},
			{name: "Negative index", input: []int{1, 2, 3, 4, 5}, index: -1, expected: 5, shouldPanic: false},
			{name: "Zero-length slice", input: []int{}, index: 0, shouldPanic: true},
			{name: "Index out of bounds", input: []int{1, 2, 3}, index: 5, shouldPanic: true},
			{name: "Negative index out of bounds", input: []int{1, 2, 3}, index: -4, shouldPanic: true},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				defer func() {
					if r := recover(); tc.shouldPanic == (r == nil) {
						t.Errorf("expectedIndex panic=%t but got: %v", tc.shouldPanic, r)
					}
				}()
				got := At(tc.input, tc.index)
				if got != tc.expected {
					t.Errorf("expectedIndex %v, got %v", tc.expected, got)
				}

				s := NewSlice(tc.input...)
				got = s.At(tc.index)
				if got != tc.expected {
					t.Errorf("expectedIndex %v, got %v", tc.expected, got)
				}
			})
		}
	})
}

func TestEvery(t *testing.T) {
	t.Run("integer slices", func(t *testing.T) {
		even := func(n int) bool { return n%2 == 0 }
		positive := func(n int) bool { return n > 0 }

		tests := []struct {
			name string
			s    Slice[int]
			p    func(int) bool
			want bool
		}{{
			name: "every even number",
			s:    Slice[int]{2, 4, 6, 8, 10},
			p:    even,
			want: true,
		}, {
			name: "not every even number",
			s:    Slice[int]{2, 4, 5, 8, 10},
			p:    even,
			want: false,
		}, {
			name: "every positive number",
			s:    Slice[int]{1, 2, 3, 4, 5},
			p:    positive,
			want: true,
		}, {
			name: "empty slice",
			s:    Slice[int]{},
			p:    even,
			want: true,
		}, {
			name: "single element",
			s:    []int{2},
			p:    even,
			want: true,
		}, {
			name: "negative elements",
			s:    []int{-2, -4, -6, -8, -10},
			p:    positive,
			want: false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Every(tt.s, tt.p); got != tt.want {
					t.Errorf("Every(%v) = %v, want %v", tt.s, got, tt.want)
				}

				if got := tt.s.Every(tt.p); got != tt.want {
					t.Errorf("s.Every(%v) = %v, want %v", tt.s, got, tt.want)
				}
			})
		}
	})

	t.Run("string slices", func(t *testing.T) {
		tests := []struct {
			name string
			s    []string
			p    func(string) bool
			want bool
		}{{
			name: "every string has length greater than 0",
			s:    []string{"abc", "de", "f"},
			p:    func(s string) bool { return len(s) > 0 },
			want: true,
		}, {
			name: "not every string has length greater than 3",
			s:    []string{"abc", "defg", "hijklm"},
			p:    func(s string) bool { return len(s) > 3 },
			want: false,
		}, {
			name: "no strings in slice",
			s:    []string{},
			p:    func(s string) bool { return len(s) > 0 },
			want: true,
		}, {
			name: "single element with length 3",
			s:    []string{"abc"},
			p:    func(s string) bool { return len(s) == 3 },
			want: true,
		}, {
			name: "single element with length less than 3",
			s:    []string{"ab"},
			p:    func(s string) bool { return len(s) == 3 },
			want: false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Every(tt.s, tt.p); got != tt.want {
					t.Errorf("Every(%v) = %v, want %v", tt.s, got, tt.want)
				}

				s := NewSlice(tt.s...)
				if got := s.Every(tt.p); got != tt.want {
					t.Errorf("s.Every(%v) = %v, want %v", tt.s, got, tt.want)
				}
			})
		}
	})

	t.Run("slices of a custom struct", func(t *testing.T) {
		tests := []struct {
			name string
			s    []customStruct
			p    func(customStruct) bool
			want bool
		}{{
			name: "every int field is greater than zero",
			s:    []customStruct{{int: 1}, {int: 2}},
			p:    func(cs customStruct) bool { return cs.int > 0 },
			want: true,
		}, {
			name: "not every int field is greater than zero",
			s:    []customStruct{{int: 1}, {int: 2}, {int: 0}},
			p:    func(cs customStruct) bool { return cs.int > 0 },
			want: false,
		}, {
			name: "every string len is greater than zero",
			s:    []customStruct{{string: "a"}, {string: "b"}, {string: "c"}, {string: "d"}, {string: "e"}},
			p:    func(cs customStruct) bool { return len(cs.string) > 0 },
			want: true,
		}, {
			name: "not every string len is greater than zero",
			s:    []customStruct{{string: "a"}, {string: "b"}, {string: "c"}, {string: "d"}, {}},
			p:    func(cs customStruct) bool { return len(cs.string) > 0 },
			want: false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Every(tt.s, tt.p); got != tt.want {
					t.Errorf("Every(%v) = %v, want %v", tt.s, got, tt.want)
				}

				s := NewSlice(tt.s...)
				if got := s.Every(tt.p); got != tt.want {
					t.Errorf("s.Every(%v) = %v, want %v", tt.s, got, tt.want)
				}
			})
		}
	})

	t.Run("slices of a type alias", func(t *testing.T) {
		greaterThanZero := func(i typeAlias) bool {
			return i > 0
		}

		tests := []struct {
			name string
			s    []typeAlias
			p    func(alias typeAlias) bool
			want bool
		}{{
			name: "every int field is greater than zero",
			s:    []typeAlias{1, 2, 3},
			p:    greaterThanZero,
			want: true,
		}, {
			name: "not every int field is greater than zero",
			s:    []typeAlias{-1, 0, -1},
			p:    greaterThanZero,
			want: false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Every(tt.s, tt.p); got != tt.want {
					t.Errorf("Every(%v) = %v, want %v", tt.s, got, tt.want)
				}
			})
		}
	})
}

func TestSome(t *testing.T) {
	t.Run("integer slices", func(t *testing.T) {
		greaterThanThree := func(value int) bool {
			return value > 3
		}
		isNegative := func(value int) bool {
			return value < 0
		}

		tests := []struct {
			name  string
			slice []int
			p     Predicate[int]
			want  bool
		}{{
			name:  "some elements greater than three",
			slice: []int{1, 2, 3, 4},
			p:     greaterThanThree,
			want:  true,
		}, {
			name:  "no elements greater than three",
			slice: []int{1, 2, 3},
			p:     greaterThanThree,
			want:  false,
		}, {
			name:  "some elements negative",
			slice: []int{-1, 2, 3},
			p:     isNegative,
			want:  true,
		}, {
			name:  "empty slice",
			slice: []int{},
			p:     isNegative,
			want:  false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Some(tt.slice, tt.p); got != tt.want {
					t.Errorf("Some(%v) = %v, want %v", tt.slice, got, tt.want)
				}

				s := NewSlice(tt.slice...)
				if got := s.Some(tt.p); got != tt.want {
					t.Errorf("s.Some(%v) = %v, want %v", tt.slice, got, tt.want)
				}
			})
		}
	})

	t.Run("string slices", func(t *testing.T) {
		containsA := func(value string) bool {
			return strings.Contains(value, "a")
		}

		tests := []struct {
			name  string
			slice []string
			p     Predicate[string]
			want  bool
		}{{
			name:  "some strings contain 'a'",
			slice: []string{"apple", "banana", "pear"},
			p:     containsA,
			want:  true,
		}, {
			name:  "no strings contain 'a'",
			slice: []string{"melon", "lemon", "kiwi"},
			p:     containsA,
			want:  false,
		}, {
			name:  "empty slice",
			slice: []string{},
			p:     containsA,
			want:  false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Some(tt.slice, tt.p); got != tt.want {
					t.Errorf("Some(%v) = %v, want %v", tt.slice, got, tt.want)
				}

				s := NewSlice(tt.slice...)
				if got := s.Some(tt.p); got != tt.want {
					t.Errorf("s.Some(%v) = %v, want %v", tt.slice, got, tt.want)
				}
			})
		}
	})

	t.Run("custom struct slices", func(t *testing.T) {
		hasPositiveInt := func(cs customStruct) bool {
			return cs.int > 0
		}

		tests := []struct {
			name  string
			slice []customStruct
			p     Predicate[customStruct]
			want  bool
		}{{
			name:  "some structs have positive int",
			slice: []customStruct{{int: -1}, {int: 2}, {int: 0}},
			p:     hasPositiveInt,
			want:  true,
		}, {
			name:  "no structs have positive int",
			slice: []customStruct{{int: -1}, {int: -2}, {int: 0}},
			p:     hasPositiveInt,
			want:  false,
		}, {
			name:  "empty slice",
			slice: []customStruct{},
			p:     hasPositiveInt,
			want:  false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Some(tt.slice, tt.p); got != tt.want {
					t.Errorf("Some(%v) = %v, want %v", tt.slice, got, tt.want)
				}

				s := NewSlice(tt.slice...)
				if got := s.Some(tt.p); got != tt.want {
					t.Errorf("s.Some(%v) = %v, want %v", tt.slice, got, tt.want)
				}
			})
		}
	})
}

func TestFill(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		value     int
		positions []int
		expected  []int
	}{{
		name:      "fill entire slice",
		input:     []int{1, 2, 3, 4},
		value:     0,
		positions: nil,
		expected:  []int{0, 0, 0, 0},
	}, {
		name:      "fill with lower boundary",
		input:     []int{1, 2, 3, 4},
		value:     99,
		positions: []int{1},
		expected:  []int{1, 99, 99, 99},
	}, {
		name:      "fill with lower and upper boundary",
		input:     []int{1, 2, 3, 4},
		value:     99,
		positions: []int{1, 2},
		expected:  []int{1, 99, 99, 4},
	}, {
		name:      "fill no effect with empty positions",
		input:     []int{1, 2, 3, 4},
		value:     5,
		positions: []int{},
		expected:  []int{5, 5, 5, 5},
	}, {
		name:      "fill empty slice",
		input:     []int{},
		value:     10,
		positions: nil,
		expected:  []int{},
	}, {
		name:      "fill out-of-bounds indices",
		input:     []int{1, 2, 3, 4},
		value:     5,
		positions: []int{-1, 5},
		expected:  []int{5, 5, 5, 5},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fill(tt.input, tt.value, tt.positions...)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Fill() = %v, want %v", got, tt.expected)
			}

			slice := NewSlice(tt.input...)
			gotSlice := slice.Fill(tt.value, tt.positions...)
			expectedSlice := NewSlice(tt.expected...)
			if !reflect.DeepEqual(gotSlice, expectedSlice) {
				t.Errorf("slice.Fill() = %v, want %v", gotSlice, expectedSlice)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Run("tests for slices of integers", func(t *testing.T) {
		greaterThanFive := func(value int) bool {
			return value > 5
		}

		isEven := func(value int) bool {
			return value%2 == 0
		}

		tests := []struct {
			name     string
			source   []int
			pred     Predicate[int]
			expected []int
		}{{
			name:     "empty slice",
			source:   []int{},
			pred:     greaterThanFive,
			expected: []int{},
		}, {
			name:     "all elements are greater than five",
			source:   []int{6, 7, 9, 11},
			pred:     greaterThanFive,
			expected: []int{6, 7, 9, 11},
		}, {
			name:     "non existent greater than five",
			source:   []int{1, 2, 3, 4, 5},
			pred:     greaterThanFive,
			expected: []int{},
		}, {
			name:     "partial greater than five",
			source:   []int{1, 6, 3, 7},
			pred:     greaterThanFive,
			expected: []int{6, 7},
		}, {
			name:     "all elements even",
			source:   []int{2, 4, 6, 8},
			pred:     isEven,
			expected: []int{2, 4, 6, 8},
		}, {
			name:     "no even elements",
			source:   []int{1, 3, 5, 7},
			pred:     isEven,
			expected: []int{},
		}, {
			name:     "some even elements",
			source:   []int{1, 2, 3, 4},
			pred:     isEven,
			expected: []int{2, 4},
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := Filter(tt.source, tt.pred)
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("got %v, want %v", result, tt.expected)
				}

				s := NewSlice(tt.source...)
				expectedSlice := NewSlice(tt.expected...)
				resultSlice := s.Filter(tt.pred)
				if !reflect.DeepEqual(resultSlice, expectedSlice) {
					t.Errorf("got %v, want %v", resultSlice, expectedSlice)
				}
			})
		}
	})

	t.Run("test for string slices", func(t *testing.T) {
		stringsWithLetterA := func(s string) bool {
			return strings.Contains(s, "a")
		}

		tests := []struct {
			name  string
			slice []string
			p     Predicate[string]
			want  []string
		}{{
			name:  "some strings with letter a",
			slice: []string{"a", "bacon", "no first letter"},
			p:     stringsWithLetterA,
			want:  []string{"a", "bacon"},
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := Filter(tt.slice, tt.p)
				if !reflect.DeepEqual(result, tt.want) {
					t.Errorf("got %v, want %v", result, tt.want)
				}

				s := NewSlice(tt.slice...)
				expectedSlice := NewSlice(tt.want...)
				resultSlice := s.Filter(tt.p)
				if !reflect.DeepEqual(resultSlice, expectedSlice) {
					t.Errorf("got %v, want %v", resultSlice, expectedSlice)
				}
			})
		}

	})

	t.Run("test for custom struct slices", func(t *testing.T) {
		tests := []struct {
			name string
			s    []customStruct
			p    func(customStruct) bool
			want []customStruct
		}{{
			name: "all elements",
			s:    []customStruct{{int: 1}, {int: 2}},
			p:    func(cs customStruct) bool { return true },
			want: []customStruct{{int: 1}, {int: 2}},
		}, {
			name: "only the middle one",
			s:    []customStruct{{int: 1}, {int: 2}, {int: 0}},
			p:    func(cs customStruct) bool { return cs.int == 2 },
			want: []customStruct{{int: 2}},
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Filter(tt.s, tt.p)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Filter(%v) = %v, want %v", tt.s, got, tt.want)
				}

				s := NewSlice(tt.s...)
				expectedSlice := NewSlice(tt.want...)
				resultSlice := s.Filter(tt.p)
				if !reflect.DeepEqual(resultSlice, expectedSlice) {
					t.Errorf("got %v, want %v", resultSlice, expectedSlice)
				}
			})
		}
	})

	t.Run("test for custom alias types", func(t *testing.T) {
		greaterThanZero := func(i typeAlias) bool {
			return i > 0
		}

		tests := []struct {
			name string
			s    []typeAlias
			p    func(i typeAlias) bool
			want []typeAlias
		}{{
			name: "all elements are greater than zero",
			s:    []typeAlias{1, 2, 3},
			p:    greaterThanZero,
			want: []typeAlias{1, 2, 3},
		}, {
			name: "only the last one is greater than zero",
			s:    []typeAlias{-1, 0, 1},
			p:    greaterThanZero,
			want: []typeAlias{1},
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Filter(tt.s, tt.p)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Filter(%v) = %v, want %v", tt.s, got, tt.want)
				}
			})
		}
	})
}

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
		s := []string{"a", "bc", "def", "1234"}
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

	t.Run("slices of a custom struct", func(t *testing.T) {
		// TODO: implement this test
	})

	t.Run("slices of a type alias", func(t *testing.T) {
		// TODO
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

				s := NewSlice(tt.arr...)
				gotSliceIndex, gotSliceBool := s.FindIndex(tt.predicate)
				if gotSliceIndex != tt.wantIndex || gotSliceBool != tt.wantBool {
					t.Errorf("FindIndex() = %v, %v; want %v, %v", gotSliceIndex, gotSliceBool, tt.wantIndex, tt.wantBool)
				}
			})
		}
	})

	t.Run("slices of a type alias", func(t *testing.T) {
		// TODO
	})
}

func TestFindLast(t *testing.T) {
	t.Run("integer slices", func(t *testing.T) {
		tests := []struct {
			name     string
			s        Slice[int]
			p        Predicate[int]
			expected int
			found    bool
		}{{
			name:     "find last even number",
			s:        Slice[int]{1, 2, 3, 4, 5},
			p:        func(x int) bool { return x%2 == 0 },
			expected: 4,
			found:    true,
		}, {
			name:     "empty slice",
			s:        Slice[int]{},
			p:        func(x int) bool { return x > 0 },
			expected: 0, // default value for int
			found:    false,
		}, {
			name:     "no match found",
			s:        Slice[int]{1, 3, 5},
			p:        func(x int) bool { return x%2 == 0 },
			expected: 0,
			found:    false,
		}, {
			name:     "only one element matches",
			s:        Slice[int]{1, 2, 3},
			p:        func(x int) bool { return x == 2 },
			expected: 2,
			found:    true,
		}, {
			name:     "last element matches",
			s:        Slice[int]{1, 2, 3, 4},
			p:        func(x int) bool { return x == 4 },
			expected: 4,
			found:    true,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, found := tt.s.FindLast(tt.p)
				if got != tt.expected || found != tt.found {
					t.Errorf("FindLast() = (%v, %v); want (%v, %v)", got, found, tt.expected, tt.found)
				}
			})
		}
	})

	t.Run("string slices", func(t *testing.T) {
		tests := []struct {
			name     string
			s        Slice[string]
			p        Predicate[string]
			expected string
			found    bool
		}{{
			name:     "find last string containing 'a'",
			s:        Slice[string]{"apple", "banana", "cherry"},
			p:        func(x string) bool { return strings.Contains(x, "a") },
			expected: "banana",
			found:    true,
		}, {
			name:     "no string matches predicate",
			s:        Slice[string]{"dog", "cat", "mouse"},
			p:        func(x string) bool { return strings.Contains(x, "z") },
			expected: "",
			found:    false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, found := tt.s.FindLast(tt.p)
				if got != tt.expected || found != tt.found {
					t.Errorf("FindLast() = (%v, %v); want (%v, %v)", got, found, tt.expected, tt.found)
				}
			})
		}
	})

	t.Run("custom struct slices", func(t *testing.T) {
		tests := []struct {
			name     string
			s        Slice[customStruct]
			p        Predicate[customStruct]
			expected customStruct
			found    bool
		}{{
			name: "find last struct with int field > 0",
			s: Slice[customStruct]{
				{int: 1, string: "a"},
				{int: 0, string: "b"},
				{int: 2, string: "c"},
			},
			p:        func(cs customStruct) bool { return cs.int > 0 },
			expected: customStruct{int: 2, string: "c"},
			found:    true,
		}, {
			name: "no match found",
			s: Slice[customStruct]{
				{int: -1, string: "x"},
				{int: -2, string: "y"},
			},
			p:        func(cs customStruct) bool { return cs.int > 0 },
			expected: customStruct{},
			found:    false,
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, found := tt.s.FindLast(tt.p)
				if got != tt.expected || found != tt.found {
					t.Errorf("FindLast() = (%v, %v); want (%v, %v)", got, found, tt.expected, tt.found)
				}
			})
		}
	})
}

func TestForEach(t *testing.T) {
	noOp := func(_ *testing.T, _ []int) {}
	traversedItems := make([][]int, 0)

	tests := []struct {
		name      string
		input     []int
		processFn func(i int, v int)
		expected  func(t *testing.T, input []int)
	}{{
		name:  "empty slice",
		input: []int{},
		processFn: func(i int, v int) {
			t.Error("function called for empty slice")
		},
		expected: noOp,
	}, {
		name:  "function doesn't change input",
		input: []int{5},
		processFn: func(i int, v int) {
			v *= 2
		},
		expected: func(t *testing.T, input []int) {
			if !reflect.DeepEqual(input, []int{5}) {
				t.Error("function should not change input slice")
			}
		},
	}, {
		name:  "function traverse through all input items",
		input: []int{1, 2, 3},
		processFn: func(i int, v int) {
			traversedItems = append(traversedItems, []int{i, v})
		},
		expected: func(t *testing.T, input []int) {
			expected := [][]int{{0, 1}, {1, 2}, {2, 3}}
			if !reflect.DeepEqual(traversedItems, expected) {
				t.Error("function should traverse through all items")
			}
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := make([]int, len(test.input))
			copy(output, test.input)
			ForEach(output, test.processFn)
			test.expected(t, output)
		})
	}
}

func TestMap(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected []string
		mapper   func(int) (string, error)
		err      error
	}{{
		name:     "simple",
		input:    []int{1, 2, 3},
		expected: []string{"2", "4", "6"},
		mapper: func(i int) (string, error) {
			return strconv.Itoa(i * 2), nil
		},
		err: nil,
	}, {
		name:     "empty input",
		input:    []int{},
		expected: []string{},
		mapper: func(i int) (string, error) {
			return strconv.Itoa(i), nil
		},
		err: nil,
	}, {
		name:     "mapper error",
		input:    []int{1, 2, 3},
		expected: nil,
		mapper: func(i int) (string, error) {
			return "", errors.New("error")
		},
		err: errors.New("error"),
	}}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Map(tc.input, tc.mapper)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expectedIndex %v, got %v", tc.expected, actual)
			}
			if !reflect.DeepEqual(tc.err, err) {
				t.Errorf("expectedIndex err %v, got %v", tc.err, err)
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name           string
		slice          Slice[int]
		expectedResult int
		expectedOk     bool
		expectedSlice  Slice[int]
	}{{
		name:           "Pop from non-empty slice",
		slice:          Slice[int]{1, 2, 3},
		expectedResult: 3,
		expectedOk:     true,
		expectedSlice:  Slice[int]{1, 2},
	}, {
		name:           "Pop from single-element slice",
		slice:          Slice[int]{10},
		expectedResult: 10,
		expectedOk:     true,
		expectedSlice:  Slice[int]{},
	}, {
		name:           "Pop from empty slice",
		slice:          Slice[int]{},
		expectedResult: 0, // Default value for int
		expectedOk:     false,
		expectedSlice:  Slice[int]{},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("standalone function", func(t *testing.T) {
				s := make([]int, len(tt.slice))
				copy(s, tt.slice)

				gotResult, gotOk := Pop(&s)
				if gotResult != tt.expectedResult || gotOk != tt.expectedOk {
					t.Errorf("Pop() = (%v, %v), want (%v, %v)", gotResult, gotOk, tt.expectedResult, tt.expectedOk)
				}

				if !reflect.DeepEqual(s, tt.expectedSlice.ToRaw()) {
					t.Errorf("resulting Slice = %v, want %v", s, tt.expectedSlice)
				}
			})

			t.Run("method on Slice", func(t *testing.T) {
				s := NewSlice[int](tt.slice...)
				gotResult, gotOk := s.Pop()

				if gotResult != tt.expectedResult || gotOk != tt.expectedOk {
					t.Errorf("Slice.Pop() = (%v, %v), want (%v, %v)", gotResult, gotOk, tt.expectedResult, tt.expectedOk)
				}

				if !reflect.DeepEqual(s, tt.expectedSlice) {
					t.Errorf("resulting Slice = %v, want %v", s, tt.expectedSlice)
				}
			})

		})
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		name           string
		initialSlice   []int
		valuesToAdd    []int
		expectedSlice  []int
		expectedLength int
	}{{
		name:           "push to empty slice",
		initialSlice:   []int{},
		valuesToAdd:    []int{1, 2, 3},
		expectedSlice:  []int{1, 2, 3},
		expectedLength: 3,
	}, {
		name:           "push single value",
		initialSlice:   []int{1, 2, 3},
		valuesToAdd:    []int{4},
		expectedSlice:  []int{1, 2, 3, 4},
		expectedLength: 4,
	}, {
		name:           "push multiple values",
		initialSlice:   []int{1},
		valuesToAdd:    []int{2, 3, 4},
		expectedSlice:  []int{1, 2, 3, 4},
		expectedLength: 4,
	}, {
		name:           "push empty values to slice",
		initialSlice:   []int{1, 2, 3},
		valuesToAdd:    []int{},
		expectedSlice:  []int{1, 2, 3},
		expectedLength: 3,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("standalone function", func(t *testing.T) {
				s := make([]int, len(tt.initialSlice))
				copy(s, tt.initialSlice)

				gotLength := Push(&s, tt.valuesToAdd...)
				if !reflect.DeepEqual(s, tt.expectedSlice) || gotLength != tt.expectedLength {
					t.Errorf("Push() got slice=%v length=%d, want slice=%v, length=%d",
						s, gotLength, tt.expectedSlice, tt.expectedLength)
				}
			})

			t.Run("method on Slice", func(t *testing.T) {
				s := NewSlice[int](tt.initialSlice...)
				gotLength := s.Push(tt.valuesToAdd...)
				if !reflect.DeepEqual(s.ToRaw(), tt.expectedSlice) || gotLength != tt.expectedLength {
					t.Errorf("Push() got slice=%v length=%d, want slice=%v length=%d",
						s, gotLength, tt.expectedSlice, tt.expectedLength)
				}
			})
		})
	}
}

func TestReduce(t *testing.T) {
	type inputs struct {
		slice        []int
		reducer      func(int, int) (int, error)
		initialValue int
	}
	type expected struct {
		result int
		error  error
	}

	expectedErr := errors.New("intentional reducer error")

	tests := []struct {
		name                string
		inputArray          []int
		reducerFunc         func(int, int) (int, error)
		initialValue        int
		expectedResult      int
		expectedErrorIsNill bool

		inputs   inputs
		expected expected
	}{{
		name: "EmptyInput",
		inputs: inputs{
			slice:        []int{},
			reducer:      func(acc, cur int) (int, error) { return acc + cur, nil },
			initialValue: 100,
		},
		expected: expected{
			result: 100,
			error:  nil,
		},
	}, {
		name: "SingleElementInput",
		inputs: inputs{
			slice:        []int{2},
			reducer:      func(acc, cur int) (int, error) { return acc * cur, nil },
			initialValue: 10,
		},
		expected: expected{
			result: 20,
			error:  nil,
		},
	}, {
		name: "MultipleElementInput",
		inputs: inputs{
			slice:        []int{2, 3, 4},
			reducer:      func(acc, cur int) (int, error) { return acc + cur, nil },
			initialValue: 10,
		},
		expected: expected{
			result: 19,
			error:  nil,
		},
	}, {
		name: "ErrorPropagation",
		inputs: inputs{
			slice:        []int{2, 3, 4},
			reducer:      func(acc, cur int) (int, error) { return 0, expectedErr },
			initialValue: 0,
		},
		expected: expected{
			result: 0,
			error:  expectedErr,
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Reduce(test.inputs.slice, test.inputs.reducer, test.inputs.initialValue)
			if result != test.expected.result {
				t.Errorf("expectedIndex %v, got %v", test.expected.result, result)
			}
			if !errors.Is(err, test.expected.error) {
				t.Errorf("expectedIndex error %v, but got %v", test.expected.error, err)
			}
		})
	}
}

func TestReduceRight(t *testing.T) {
	t.Run("int sum", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		reducer := func(acc int, curr int) (int, error) {
			return acc + curr, nil
		}
		got, err := ReduceRight(nums, reducer, 0)

		want := 10
		if got != want || err != nil {
			t.Errorf("unexpected result: got result=%d, err=%v, want result=%d, error=%v", got, err, want, nil)
		}

	})

	t.Run("string concatenation", func(t *testing.T) {
		words := []string{"foo", "bar", "baz"}
		reducer := func(acc string, curr string) (string, error) {
			return acc + curr, nil
		}
		got, err := ReduceRight(words, reducer, "")

		want := "bazbarfoo"
		if got != want || err != nil {
			t.Errorf("unexpected result: got result=%s, err=%v, want result=%s, error=%v", got, err, want, nil)
		}
	})

	t.Run("with error", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		customErr := errors.New("value 2 encountered")
		reducer := func(acc int, curr int) (int, error) {
			if curr == 2 {
				return acc, customErr
			}
			return acc + curr, nil
		}
		got, err := ReduceRight(nums, reducer, 0)

		want := 7
		if got != want || !errors.Is(err, customErr) {
			t.Errorf("unexpected result: got result=%d, err=%v, want result=%d, error=%v", got, err, want, customErr)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		nums := []int{}
		reducer := func(acc int, curr int) (int, error) {
			return acc + curr, nil
		}
		got, err := ReduceRight(nums, reducer, 42)

		want := 42
		if got != want || err != nil {
			t.Errorf("unexpected result: got result=%d, err=%v, want result=%d, error=%v", got, err, want, nil)
		}
	})

	t.Run("different input and output types", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		reducer := func(acc string, curr int) (string, error) {
			return acc + strconv.Itoa(curr), nil
		}
		got, err := ReduceRight(nums, reducer, "")

		want := "4321"
		if got != want || err != nil {
			t.Errorf("unexpected result: got result=%s, err=%v, want result=%s, error=%v", got, err, want, nil)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("integer slices", func(t *testing.T) {
		tests := []struct {
			name  string
			input []int
			want  []int
		}{{
			name:  "odd length slice",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		}, {
			name:  "even length slice",
			input: []int{1, 2, 3, 4},
			want:  []int{4, 3, 2, 1},
		}, {
			name:  "single element slice",
			input: []int{42},
			want:  []int{42},
		}, {
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		}, {
			name:  "negative integers",
			input: []int{-3, -2, -1, 0},
			want:  []int{0, -1, -2, -3},
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Run("standalone function", func(t *testing.T) {
					// Test the standalone function
					inputCopy := make([]int, len(tt.input))
					copy(inputCopy, tt.input)

					got := Reverse(inputCopy)
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("Reverse(%v) = %v, want %v", tt.input, got, tt.want)
					}
					if !reflect.DeepEqual(inputCopy, tt.want) {
						t.Errorf("original slice wasn't modified in-place. got=%v, want=%v", inputCopy, tt.want)
					}
				})

				t.Run("method on slice type", func(t *testing.T) {
					inputCopy := make([]int, len(tt.input))
					copy(inputCopy, tt.input)

					s := NewSlice(inputCopy...)
					gotSlice := s.Reverse()

					if !reflect.DeepEqual(gotSlice, Slice[int](tt.want)) {
						t.Errorf("s.Reverse() = %v, want %v", gotSlice, tt.want)
					}
					if !reflect.DeepEqual(s, Slice[int](tt.want)) {
						t.Errorf("original Slice wasn't modified in-place. Got %v, want %v", s, tt.want)
					}
				})
			})
		}
	})

	t.Run("string slices", func(t *testing.T) {
		tests := []struct {
			name  string
			input []string
			want  []string
		}{{
			name:  "words",
			input: []string{"hello", "world", "go", "reverse"},
			want:  []string{"reverse", "go", "world", "hello"},
		}, {
			name:  "empty strings present",
			input: []string{"", "abc", "", "def"},
			want:  []string{"def", "", "abc", ""},
		}, {
			name:  "single string",
			input: []string{"singleton"},
			want:  []string{"singleton"},
		}, {
			name:  "empty slice",
			input: []string{},
			want:  []string{},
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Run("standalone function", func(t *testing.T) {
					inputCopy := make([]string, len(tt.input))
					copy(inputCopy, tt.input)

					got := Reverse(inputCopy)
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("Reverse(%v) = %v, want %v", tt.input, got, tt.want)
					}
				})

				t.Run("method on slice type", func(t *testing.T) {
					inputCopy := make([]string, len(tt.input))
					copy(inputCopy, tt.input)

					s := NewSlice(inputCopy...)
					gotSlice := s.Reverse()

					if !reflect.DeepEqual(gotSlice, Slice[string](tt.want)) {
						t.Errorf("s.Reverse() = %v, want %v", gotSlice, tt.want)
					}
				})
			})
		}
	})
}

func TestToReversed(t *testing.T) {
	t.Run("integer slices", func(t *testing.T) {
		tests := []struct {
			name  string
			input []int
			want  []int
		}{{
			name:  "odd length slice",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		}, {
			name:  "even length slice",
			input: []int{1, 2, 3, 4},
			want:  []int{4, 3, 2, 1},
		}, {
			name:  "single element slice",
			input: []int{42},
			want:  []int{42},
		}, {
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		}, {
			name:  "negative integers",
			input: []int{-3, -2, -1, 0},
			want:  []int{0, -1, -2, -3},
		}}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Run("standalone function", func(t *testing.T) {
					// Test the standalone function
					inputCopy := make([]int, len(tt.input))
					copy(inputCopy, tt.input)

					got := ToReversed(inputCopy)
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("Reverse(%v) = %v, want %v", tt.input, got, tt.want)
					}
					if !reflect.DeepEqual(inputCopy, tt.input) {
						t.Errorf("original slice was modified in-place. got=%v, want=%v", inputCopy, tt.want)
					}
				})

				t.Run("method on slice type", func(t *testing.T) {
					inputCopy := make([]int, len(tt.input))
					copy(inputCopy, tt.input)

					s := NewSlice(inputCopy...)
					gotSlice := s.ToReversed()

					if !reflect.DeepEqual(gotSlice, Slice[int](tt.want)) {
						t.Errorf("s.Reverse() = %v, want %v", gotSlice, tt.want)
					}
					if !reflect.DeepEqual(s, Slice[int](tt.input)) {
						t.Errorf("original Slice wasn't modified in-place. Got %v, want %v", s, tt.want)
					}
				})
			})
		}
	})
}

func TestShift(t *testing.T) {
	tests := []struct {
		name           string
		initialSlice   []int
		expectedResult int
		expectedOk     bool
		expectedSlice  []int
	}{{
		name:           "Shift from non-empty slice",
		initialSlice:   []int{1, 2, 3},
		expectedResult: 1,
		expectedOk:     true,
		expectedSlice:  []int{2, 3},
	}, {
		name:           "Shift single-element slice",
		initialSlice:   []int{10},
		expectedResult: 10,
		expectedOk:     true,
		expectedSlice:  []int{},
	}, {
		name:           "Shift empty slice",
		initialSlice:   []int{},
		expectedResult: 0, // Default value for int
		expectedOk:     false,
		expectedSlice:  []int{},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("standalone function", func(t *testing.T) {
				s := make([]int, len(tt.initialSlice))
				copy(s, tt.initialSlice)

				gotResult, gotOk := Shift(&s)
				if gotResult != tt.expectedResult || gotOk != tt.expectedOk {
					t.Errorf("Shift() = (%v, %v), want (%v, %v)", gotResult, gotOk, tt.expectedResult, tt.expectedOk)
				}

				if !reflect.DeepEqual(s, tt.expectedSlice) {
					t.Errorf("resulting slice = %v, want %v", s, tt.expectedSlice)
				}
			})

			t.Run("method on Slice", func(t *testing.T) {
				s := NewSlice[int](tt.initialSlice...)
				gotResult, gotOk := s.Shift()

				if gotResult != tt.expectedResult || gotOk != tt.expectedOk {
					t.Errorf("Slice.Shift() = (%v, %v), want (%v, %v)", gotResult, gotOk, tt.expectedResult, tt.expectedOk)
				}

				if !reflect.DeepEqual(s.ToRaw(), tt.expectedSlice) {
					t.Errorf("resulting Slice = %v, want %v", s, tt.expectedSlice)
				}
			})
		})
	}
}

func TestUnshift(t *testing.T) {
	tests := []struct {
		name           string
		initialSlice   []int
		valuesToAdd    []int
		expectedSlice  []int
		expectedLength int
	}{{
		name:           "unshift to empty slice",
		initialSlice:   []int{},
		valuesToAdd:    []int{1, 2, 3},
		expectedSlice:  []int{1, 2, 3},
		expectedLength: 3,
	}, {
		name:           "unshift single value",
		initialSlice:   []int{4, 5, 6},
		valuesToAdd:    []int{1},
		expectedSlice:  []int{1, 4, 5, 6},
		expectedLength: 4,
	}, {
		name:           "unshift multiple values",
		initialSlice:   []int{3, 4},
		valuesToAdd:    []int{1, 2},
		expectedSlice:  []int{1, 2, 3, 4},
		expectedLength: 4,
	}, {
		name:           "unshift empty values to slice",
		initialSlice:   []int{7, 8, 9},
		valuesToAdd:    []int{},
		expectedSlice:  []int{7, 8, 9},
		expectedLength: 3,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("standalone function", func(t *testing.T) {
				s := make([]int, len(tt.initialSlice))
				copy(s, tt.initialSlice)

				gotLength := Unshift(&s, tt.valuesToAdd...)
				if !reflect.DeepEqual(s, tt.expectedSlice) || gotLength != tt.expectedLength {
					t.Errorf("Unshift() got slice=%v length=%d, want slice=%v length=%d",
						s, gotLength, tt.expectedSlice, tt.expectedLength)
				}
			})

			t.Run("method on Slice", func(t *testing.T) {
				s := NewSlice[int](tt.initialSlice...)
				gotLength := s.Unshift(tt.valuesToAdd...)
				if !reflect.DeepEqual(s.ToRaw(), tt.expectedSlice) || gotLength != tt.expectedLength {
					t.Errorf("Unshift() got slice=%v length=%d, want slice=%v length=%d",
						s, gotLength, tt.expectedSlice, tt.expectedLength)
				}
			})
		})
	}
}
