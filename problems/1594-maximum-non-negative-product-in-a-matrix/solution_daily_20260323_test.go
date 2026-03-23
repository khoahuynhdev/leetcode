package main

import "testing"

func TestMaxProductPath(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "example 1: all negative paths",
			grid:     [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}},
			expected: -1,
		},
		{
			name:     "example 2: negatives cancel out",
			grid:     [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}},
			expected: 8,
		},
		{
			name:     "example 3: zero in path",
			grid:     [][]int{{1, 3}, {0, -4}},
			expected: 0,
		},
		{
			name:     "edge case: single cell positive",
			grid:     [][]int{{3}},
			expected: 3,
		},
		{
			name:     "edge case: single cell negative",
			grid:     [][]int{{-2}},
			expected: -1,
		},
		{
			name:     "edge case: single cell zero",
			grid:     [][]int{{0}},
			expected: 0,
		},
		{
			name:     "edge case: all positive",
			grid:     [][]int{{1, 2, 3}, {4, 1, 2}, {1, 3, 1}},
			expected: 12,
		},
		{
			name:     "edge case: 1x3 row with negative product",
			grid:     [][]int{{2, -1, 3}},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProductPath(tt.grid)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
