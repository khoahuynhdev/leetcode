package main

import "testing"

func TestCanPartitionGrid(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected bool
	}{
		{
			name:     "example 1: horizontal cut splits evenly",
			grid:     [][]int{{1, 4}, {2, 3}},
			expected: true,
		},
		{
			name:     "example 2: no valid cut exists",
			grid:     [][]int{{1, 3}, {2, 4}},
			expected: false,
		},
		{
			name:     "edge case: single row vertical cut works",
			grid:     [][]int{{3, 3}},
			expected: true,
		},
		{
			name:     "edge case: single column horizontal cut works",
			grid:     [][]int{{5}, {5}},
			expected: true,
		},
		{
			name:     "edge case: odd total sum",
			grid:     [][]int{{1, 2}, {3, 5}},
			expected: false,
		},
		{
			name:     "vertical cut on larger grid",
			grid:     [][]int{{1, 1, 2}, {1, 1, 2}},
			expected: true,
		},
		{
			name:     "all elements equal 3x2 horizontal cut",
			grid:     [][]int{{2, 2}, {2, 2}, {2, 2}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canPartitionGrid(tt.grid)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
