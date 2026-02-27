package main

import "testing"

func TestMakeLargestSpecial(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example 1: swap to get largest",
			input:    "11011000",
			expected: "11100100",
		},
		{
			name:     "example 2: single pair unchanged",
			input:    "10",
			expected: "10",
		},
		{
			name:     "edge case: flat structure of identical pairs",
			input:    "101010",
			expected: "101010",
		},
		{
			name:     "edge case: deeply nested",
			input:    "11001100",
			expected: "11001100",
		},
		{
			name:     "edge case: multiple top-level groups needing reorder",
			input:    "1010110010",
			expected: "1100101010",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := makeLargestSpecial(tt.input)
			if result != tt.expected {
				t.Errorf("makeLargestSpecial(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
