package main

import "testing"

func TestFurthestDistanceFromOrigin(t *testing.T) {
	tests := []struct {
		name     string
		moves    string
		expected int
	}{
		{"example 1: mixed with equal L and R", "L_RL__R", 3},
		{"example 2: more L than R with underscores", "_R__LL_", 5},
		{"example 3: all underscores", "_______", 7},
		{"edge case: single L", "L", 1},
		{"edge case: single R", "R", 1},
		{"edge case: single underscore", "_", 1},
		{"edge case: all L no underscores", "LLLL", 4},
		{"edge case: all R no underscores", "RRRR", 4},
		{"edge case: balanced L and R no underscores", "LRLR", 0},
		{"edge case: balanced L and R with underscores", "LR__", 2},
		{"edge case: more R than L with underscores", "R_R_L", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := furthestDistanceFromOrigin(tt.moves)
			if got != tt.expected {
				t.Errorf("furthestDistanceFromOrigin(%q) = %d, want %d", tt.moves, got, tt.expected)
			}
		})
	}
}
