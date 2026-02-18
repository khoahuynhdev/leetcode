package main

import "testing"

func TestHasAlternatingBits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"example 1: n=5 (101)", 5, true},
		{"example 2: n=7 (111)", 7, false},
		{"example 3: n=11 (1011)", 11, false},
		{"edge case: n=1 (single bit)", 1, true},
		{"edge case: n=2 (10)", 2, true},
		{"edge case: n=3 (11)", 3, false},
		{"alternating: n=10 (1010)", 10, true},
		{"alternating: n=21 (10101)", 21, true},
		{"non-alternating: n=4 (100)", 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasAlternatingBits(tt.input)
			if result != tt.expected {
				t.Errorf("hasAlternatingBits(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
