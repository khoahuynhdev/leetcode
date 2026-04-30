package main

import "testing"

func TestMaxPathScoreDaily20260430(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected int
	}{
		{
			name:     "example 1: 2x2 grid with budget 1",
			grid:     [][]int{{0, 1}, {2, 0}},
			k:        1,
			expected: 2,
		},
		{
			name:     "example 2: 2x2 grid where every path overspends",
			grid:     [][]int{{0, 1}, {1, 2}},
			k:        1,
			expected: -1,
		},
		{
			name:     "edge case: single cell grid with zero budget",
			grid:     [][]int{{0}},
			k:        0,
			expected: 0,
		},
		{
			name:     "edge case: zero budget but every path needs cost",
			grid:     [][]int{{0, 1}, {1, 0}},
			k:        0,
			expected: -1,
		},
		{
			name:     "all zeros grid: cost is always 0",
			grid:     [][]int{{0, 0, 0}, {0, 0, 0}},
			k:        0,
			expected: 0,
		},
		{
			name:     "budget allows the value-2 corner: prefer 2 over 1",
			grid:     [][]int{{0, 2}, {1, 2}},
			k:        2,
			expected: 4,
		},
		{
			name:     "budget exactly fits: choose path with 2 instead of 1",
			grid:     [][]int{{0, 1}, {1, 2}},
			k:        2,
			expected: 3,
		},
		{
			name:     "budget far larger than path length: take everything",
			grid:     [][]int{{0, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			k:        1000,
			expected: 8,
		},
		{
			name:     "tiny budget on rich grid: pick the single best non-zero cell",
			grid:     [][]int{{0, 2, 2}, {1, 0, 0}, {1, 1, 0}},
			k:        1,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxPathScoreDaily20260430(tt.grid, tt.k)
			if got != tt.expected {
				t.Errorf("maxPathScoreDaily20260430(%v, %d) = %d, want %d", tt.grid, tt.k, got, tt.expected)
			}
		})
	}
}
