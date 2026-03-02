package main

import "testing"

func TestMinSwaps(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "example 1: swap needed to clear upper triangle",
			grid:     [][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}},
			expected: 3,
		},
		{
			name:     "example 2: impossible identical rows",
			grid:     [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}},
			expected: -1,
		},
		{
			name:     "example 3: already valid",
			grid:     [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}},
			expected: 0,
		},
		{
			name:     "edge case: 1x1 grid",
			grid:     [][]int{{1}},
			expected: 0,
		},
		{
			name:     "edge case: all zeros",
			grid:     [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			expected: 0,
		},
		{
			name:     "edge case: reverse order needs maximum swaps",
			grid:     [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 0}},
			expected: 3,
		},
		{
			name:     "edge case: 2x2 impossible",
			grid:     [][]int{{1, 1}, {1, 1}},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSwaps(tt.grid)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
