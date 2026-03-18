package main

import "testing"

func TestCountSubmatrices(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected int
	}{
		{
			name:     "example 1: 2x3 grid k=18",
			grid:     [][]int{{7, 6, 3}, {6, 6, 1}},
			k:        18,
			expected: 4,
		},
		{
			name:     "example 2: 3x3 grid k=20",
			grid:     [][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}},
			k:        20,
			expected: 6,
		},
		{
			name:     "edge case: single cell within k",
			grid:     [][]int{{5}},
			k:        5,
			expected: 1,
		},
		{
			name:     "edge case: single cell exceeds k",
			grid:     [][]int{{10}},
			k:        3,
			expected: 0,
		},
		{
			name:     "edge case: all zeros",
			grid:     [][]int{{0, 0}, {0, 0}},
			k:        0,
			expected: 4,
		},
		{
			name:     "edge case: single row",
			grid:     [][]int{{1, 2, 3, 4}},
			k:        6,
			expected: 3,
		},
		{
			name:     "edge case: single column",
			grid:     [][]int{{1}, {2}, {3}, {4}},
			k:        6,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Deep copy grid since the solution modifies it in-place
			grid := make([][]int, len(tt.grid))
			for i := range tt.grid {
				grid[i] = make([]int, len(tt.grid[i]))
				copy(grid[i], tt.grid[i])
			}
			result := countSubmatrices(grid, tt.k)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
