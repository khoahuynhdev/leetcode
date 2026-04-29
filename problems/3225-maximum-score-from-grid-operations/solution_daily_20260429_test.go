package main

import "testing"

func TestMaximumScoreDaily20260429(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int64
	}{
		{
			name: "example 1: single 5 carries column 0 plus 3+3 from middle columns",
			grid: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 3, 0, 0},
				{0, 1, 0, 0, 0},
				{5, 0, 0, 3, 0},
				{0, 0, 0, 0, 2},
			},
			expected: 11,
		},
		{
			name: "example 2: weave heights to expose the high-value white cells",
			grid: [][]int{
				{10, 9, 0, 0, 15},
				{7, 1, 0, 8, 0},
				{5, 20, 0, 11, 0},
				{0, 0, 0, 1, 2},
				{8, 12, 1, 10, 3},
			},
			expected: 94,
		},
		{
			name:     "edge case: single cell grid has no horizontal neighbor",
			grid:     [][]int{{42}},
			expected: 0,
		},
		{
			name: "edge case: all zeros, any heights yield score 0",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "edge case: 2x2 grid, fully blacken column 0 to expose the right column",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: 6,
		},
		{
			name: "edge case: corners on a 3x3 - blacken middle column to expose both sides",
			grid: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
			expected: 4,
		},
		{
			name: "edge case: large value reachable only via right neighbor",
			grid: [][]int{
				{100, 0},
				{0, 0},
			},
			expected: 100,
		},
		{
			name: "edge case: 3x3 with non-trivial optimum requiring two operations",
			grid: [][]int{
				{0, 0, 0},
				{5, 0, 7},
				{0, 0, 0},
			},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumScore(tt.grid)
			if got != tt.expected {
				t.Errorf("maximumScore(%v) = %d, want %d", tt.grid, got, tt.expected)
			}
		})
	}
}
