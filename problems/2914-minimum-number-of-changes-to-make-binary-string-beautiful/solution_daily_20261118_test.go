package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: two mismatched pairs", "1001", 2},
		{"example 2: single mismatched pair", "10", 1},
		{"example 3: already beautiful", "0000", 0},
		{"edge case: minimal length already uniform", "11", 0},
		{"edge case: all pairs mismatched", "101010", 3},
		{"edge case: unaligned run of ones", "0110", 2},
		{"edge case: mixed uniform blocks", "110011", 0},
		{"edge case: one mismatch among uniform pairs", "110100", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minChanges(tt.s)
			if result != tt.expected {
				t.Errorf("minChanges(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
