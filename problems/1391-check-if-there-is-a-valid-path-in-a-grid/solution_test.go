package main

import "testing"

func TestHasValidPath(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected bool
	}{
		{
			name:     "example 1: zig-zag through 2x3 grid reaches end",
			grid:     [][]int{{2, 4, 3}, {6, 5, 2}},
			expected: true,
		},
		{
			name:     "example 2: vertical street at (0,1) does not accept the right opening from (0,0)",
			grid:     [][]int{{1, 2, 1}, {1, 2, 1}},
			expected: false,
		},
		{
			name:     "example 3: chain breaks at (0,2) where street 2 does not accept a left opening",
			grid:     [][]int{{1, 1, 2}},
			expected: false,
		},
		{
			name:     "edge case: 1x1 grid with arbitrary street, start equals destination",
			grid:     [][]int{{1}},
			expected: true,
		},
		{
			name:     "edge case: 1x1 grid with street type 5, start equals destination",
			grid:     [][]int{{5}},
			expected: true,
		},
		{
			name:     "edge case: single row 1x3 with curve-then-straight-then-curve forms a valid path",
			grid:     [][]int{{4, 1, 3}},
			expected: true,
		},
		{
			name:     "edge case: 2x2 with two valid paths via right-then-down and down-then-right",
			grid:     [][]int{{4, 3}, {6, 5}},
			expected: true,
		},
		{
			name:     "edge case: single column 3x1 of vertical streets connects top to bottom",
			grid:     [][]int{{2}, {2}, {2}},
			expected: true,
		},
		{
			name:     "edge case: single column 2x1 with horizontal street at top is disconnected",
			grid:     [][]int{{1}, {2}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasValidPath(tt.grid)
			if got != tt.expected {
				t.Errorf("hasValidPath(%v) = %v, want %v", tt.grid, got, tt.expected)
			}
		})
	}
}
