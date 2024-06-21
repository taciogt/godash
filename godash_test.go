package godash

import (
	"reflect"
	"strings"
	"testing"
)

type customStruct struct {
	int    int
	string string
}

type typeAlias int

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
			})
		}
	})

	t.Run("slices of a type alias", func(t *testing.T) {
		// TODO
	})
}
