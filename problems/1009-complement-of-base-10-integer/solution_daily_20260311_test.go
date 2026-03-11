package main

import "testing"

func TestBitwiseComplement(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"example 1: n=5 complement is 2", 5, 2},
		{"example 2: n=7 all ones complement is 0", 7, 0},
		{"example 3: n=10 complement is 5", 10, 5},
		{"edge case: n=0 complement is 1", 0, 1},
		{"edge case: n=1 complement is 0", 1, 0},
		{"power of two: n=4 (100) complement is 3 (011)", 4, 3},
		{"large value: n=999999999", 999999999, 73741824},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := bitwiseComplement(tt.input)
			if result != tt.expected {
				t.Errorf("bitwiseComplement(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
