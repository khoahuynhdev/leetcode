package main

import "testing"

func TestMinFlips(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"example 1: rotate and flip", "111000", 2},
		{"example 2: already alternating", "010", 0},
		{"example 3: single flip needed", "1110", 1},
		{"edge case: single character 0", "0", 0},
		{"edge case: single character 1", "1", 0},
		{"edge case: all ones", "1111", 2},
		{"edge case: all zeros", "00000", 2},
		{"edge case: two same chars", "00", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minFlips(tt.input)
			if result != tt.expected {
				t.Errorf("minFlips(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
