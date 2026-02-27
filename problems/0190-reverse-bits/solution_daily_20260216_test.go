package main

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected uint32
	}{
		{
			name:     "example 1: mixed bits",
			input:    43261596,
			expected: 964176192,
		},
		{
			name:     "example 2: mostly ones",
			input:    2147483644,
			expected: 1073741822,
		},
		{
			name:     "edge case: zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "edge case: smallest even positive",
			input:    2,
			expected: 1073741824,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseBits(tt.input)
			if result != tt.expected {
				t.Errorf("reverseBits(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
