package godash

import (
	"errors"
	"testing"
)

func TestMapperToMustMapper(t *testing.T) {
	mapper := func(input int) (string, error) {
		if input < 0 {
			return "", errors.New("negative input is not allowed")
		}
		return string(rune(input)), nil
	}

	tests := []struct {
		name        string
		mapper      Mapper[int, string]
		input       int
		want        string
		expectPanic bool
	}{{
		name:        "successful mapping",
		mapper:      mapper,
		input:       65,
		want:        "A",
		expectPanic: false,
	}, {
		name:        "error in mapping",
		mapper:      mapper,
		input:       -1,
		expectPanic: true,
	}, {
		name:        "panic with nil mapper",
		mapper:      nil,
		expectPanic: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.expectPanic {
					t.Errorf("unexpected panic: %v", r)
				} else if r == nil && tt.expectPanic {
					t.Error("expected a panic but didn't occur")
				}
			}()

			if tt.mapper == nil {
				_ = MapperToMustMapper[int, string](tt.mapper)(tt.input)
				return
			}

			mustMapper := MapperToMustMapper(tt.mapper)
			got := mustMapper(tt.input)
			if got != tt.want && !tt.expectPanic {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
