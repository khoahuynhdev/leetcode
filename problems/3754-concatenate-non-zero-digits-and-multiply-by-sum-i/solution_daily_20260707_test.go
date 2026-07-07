package main

import "testing"

func TestConcatenateNonZeroDigits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1: interior zeros, x=1234 sum=10", 10203004, 12340},
		{"example 2: trailing zeros, x=1 sum=1", 1000, 1},
		{"edge case: n=0 has no non-zero digits", 0, 0},
		{"edge case: single non-zero digit", 7, 49},
		{"edge case: max input 10^9 keeps only leading 1", 1000000000, 1},
		{"edge case: no zeros present", 123, 738},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := concatenateNonZeroDigits(tt.n)
			if result != tt.expected {
				t.Errorf("concatenateNonZeroDigits(%d) = %d, want %d", tt.n, result, tt.expected)
			}
		})
	}
}
