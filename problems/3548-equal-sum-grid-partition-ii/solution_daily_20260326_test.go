package main

import "testing"

func TestCanPartitionGrid(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected bool
	}{
		{
			name:     "example 1: equal horizontal split",
			grid:     [][]int{{1, 4}, {2, 3}},
			expected: true,
		},
		{
			name:     "example 2: vertical cut discount one cell",
			grid:     [][]int{{1, 2}, {3, 4}},
			expected: true,
		},
		{
			name:     "example 3: discount breaks connectivity",
			grid:     [][]int{{1, 2, 4}, {2, 3, 5}},
			expected: false,
		},
		{
			name:     "example 4: no valid cut exists",
			grid:     [][]int{{4, 1, 8}, {3, 2, 6}},
			expected: false,
		},
		{
			name:     "edge case: 1x2 grid equal values",
			grid:     [][]int{{5, 5}},
			expected: true,
		},
		{
			name:     "edge case: 1x2 grid unequal values",
			grid:     [][]int{{3, 7}},
			expected: false,
		},
		{
			name:     "edge case: 2x1 grid equal values",
			grid:     [][]int{{4}, {4}},
			expected: true,
		},
		{
			name:     "edge case: single column discount endpoint",
			grid:     [][]int{{3}, {1}, {2}},
			expected: true,
		},
		{
			name:     "edge case: single row discount endpoint",
			grid:     [][]int{{1, 2, 3}},
			expected: true,
		},
		{
			name:     "edge case: large equal partition no discount needed",
			grid:     [][]int{{1, 1, 1}, {1, 1, 1}},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canPartitionGrid(tt.grid)
			if result != tt.expected {
				t.Errorf("canPartitionGrid(%v) = %v, want %v", tt.grid, result, tt.expected)
			}
		})
	}
}
