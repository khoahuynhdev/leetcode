package main

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{"example 1: single zero with k=1", "110", 1, 1},
		{"example 2: two zeros with k=3", "0101", 3, 2},
		{"example 3: impossible k even z odd", "101", 2, -1},
		{"edge case: all ones already", "1111", 2, 0},
		{"edge case: all zeros k equals n", "000", 3, 1},
		{"edge case: k equals n mixed", "01", 2, -1},
		{"edge case: single zero", "0", 1, 1},
		{"edge case: all zeros k=1", "00", 1, 2},
		{"edge case: capacity constraint dominates", "1110", 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minOperations(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("minOperations(%q, %d) = %d, want %d", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
