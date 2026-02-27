package main

import "testing"

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"example 1: grouped pairs", "00110011", 6},
		{"example 2: alternating", "10101", 4},
		{"edge case: single character", "0", 0},
		{"edge case: all same characters", "1111", 0},
		{"edge case: two different characters", "01", 1},
		{"edge case: two groups equal length", "000111", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countBinarySubstrings(tt.input)
			if result != tt.expected {
				t.Errorf("countBinarySubstrings(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
