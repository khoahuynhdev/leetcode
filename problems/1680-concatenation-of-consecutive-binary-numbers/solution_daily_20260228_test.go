package main

import "testing"

func TestConcatenatedBinary(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"example 1: single number", 1, 1},
		{"example 2: first three numbers", 3, 27},
		{"example 3: twelve numbers with modulo", 12, 505379714},
		{"edge case: n=2 binary 110", 2, 6},
		{"power of two: n=4", 4, 220},
		{"large input: n=100000", 100000, 757631812},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := concatenatedBinary(tt.input)
			if result != tt.expected {
				t.Errorf("concatenatedBinary(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
