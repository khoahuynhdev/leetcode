package main

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		x        int
		expected int
	}{
		{
			name:     "example 1: 2x2 grid even step makes all equal to 4",
			grid:     [][]int{{2, 4}, {6, 8}},
			x:        2,
			expected: 4,
		},
		{
			name:     "example 2: 2x2 grid step 1 makes all equal to 3",
			grid:     [][]int{{1, 5}, {2, 3}},
			x:        1,
			expected: 5,
		},
		{
			name:     "example 3: residues differ so impossible",
			grid:     [][]int{{1, 2}, {3, 4}},
			x:        2,
			expected: -1,
		},
		{
			name:     "edge case: 1x1 grid is already uni-value",
			grid:     [][]int{{7}},
			x:        3,
			expected: 0,
		},
		{
			name:     "edge case: all elements already equal",
			grid:     [][]int{{5, 5, 5}, {5, 5, 5}},
			x:        4,
			expected: 0,
		},
		{
			name:     "edge case: single row with x=1 collapses to median",
			grid:     [][]int{{1, 2, 3, 4, 5}},
			x:        1,
			expected: 6,
		},
		{
			name:     "edge case: large step infeasible across rows",
			grid:     [][]int{{10, 20}, {30, 41}},
			x:        10,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.grid, tt.x)
			if got != tt.expected {
				t.Errorf("minOperations(%v, %d) = %d, want %d", tt.grid, tt.x, got, tt.expected)
			}
		})
	}
}
