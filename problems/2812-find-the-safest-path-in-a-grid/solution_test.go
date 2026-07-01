package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "example 1: thieves block both corners",
			grid:     [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			expected: 0,
		},
		{
			name:     "example 2: single thief top-right",
			grid:     [][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}},
			expected: 2,
		},
		{
			name:     "example 3: thieves on opposite corners of 4x4",
			grid:     [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}},
			expected: 2,
		},
		{
			name:     "edge case: single cell that is a thief",
			grid:     [][]int{{1}},
			expected: 0,
		},
		{
			name:     "edge case: every cell is a thief",
			grid:     [][]int{{1, 1}, {1, 1}},
			expected: 0,
		},
		{
			name:     "edge case: start cell is a thief",
			grid:     [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			expected: 0,
		},
		{
			name:     "edge case: single thief in the center of 5x5",
			grid:     [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumSafenessFactor(tt.grid)
			if result != tt.expected {
				t.Errorf("maximumSafenessFactor(%v) = %d, want %d", tt.grid, result, tt.expected)
			}
		})
	}
}
