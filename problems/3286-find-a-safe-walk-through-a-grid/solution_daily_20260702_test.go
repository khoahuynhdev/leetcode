package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		health   int
		expected bool
	}{
		{
			name:     "example 1: reachable with health 1",
			grid:     [][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}},
			health:   1,
			expected: true,
		},
		{
			name:     "example 2: needs 4 health but only has 3",
			grid:     [][]int{{0, 1, 1, 0, 0, 0}, {1, 0, 1, 0, 0, 0}, {0, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}},
			health:   3,
			expected: false,
		},
		{
			name:     "example 3: must pass through safe center cell",
			grid:     [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			health:   5,
			expected: true,
		},
		{
			name:     "edge case: all-safe grid succeeds with minimal health",
			grid:     [][]int{{0, 0}, {0, 0}},
			health:   1,
			expected: true,
		},
		{
			name:     "edge case: unsafe start and end drains exactly to zero",
			grid:     [][]int{{1, 0}, {0, 1}},
			health:   2,
			expected: false,
		},
		{
			name:     "edge case: unsafe start and end survivable with extra health",
			grid:     [][]int{{1, 0}, {0, 1}},
			health:   3,
			expected: true,
		},
		{
			name:     "edge case: single row all safe",
			grid:     [][]int{{0, 0, 0, 0}},
			health:   1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSafeWalk(tt.grid, tt.health)
			if result != tt.expected {
				t.Errorf("findSafeWalk(%v, %d) = %v, want %v", tt.grid, tt.health, result, tt.expected)
			}
		})
	}
}
