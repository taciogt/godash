package utils

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "zero is even",
			input:    0,
			expected: true,
		},
		{
			name:     "positive even number",
			input:    2,
			expected: true,
		},
		{
			name:     "positive odd number",
			input:    3,
			expected: false,
		},
		{
			name:     "negative even number",
			input:    -4,
			expected: true,
		},
		{
			name:     "negative odd number",
			input:    -5,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
