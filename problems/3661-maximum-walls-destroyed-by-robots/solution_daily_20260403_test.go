package main

import "testing"

func TestMaxWallsDestroyed(t *testing.T) {
	tests := []struct {
		name     string
		robots   []int
		distance []int
		walls    []int
		expected int
	}{
		{
			name:     "example 1: single robot fires left",
			robots:   []int{4},
			distance: []int{3},
			walls:    []int{1, 10},
			expected: 1,
		},
		{
			name:     "example 2: two robots cover all walls",
			robots:   []int{10, 2},
			distance: []int{5, 1},
			walls:    []int{5, 2, 7},
			expected: 3,
		},
		{
			name:     "example 3: bullet blocked by adjacent robot",
			robots:   []int{1, 2},
			distance: []int{100, 1},
			walls:    []int{10},
			expected: 0,
		},
		{
			name:     "edge case: wall at robot position always destroyed",
			robots:   []int{5},
			distance: []int{1},
			walls:    []int{5},
			expected: 1,
		},
		{
			name:     "edge case: no walls reachable",
			robots:   []int{5, 10},
			distance: []int{1, 1},
			walls:    []int{100},
			expected: 0,
		},
		{
			name:     "edge case: single robot fires right",
			robots:   []int{1},
			distance: []int{5},
			walls:    []int{3, 4, 6, 100},
			expected: 3,
		},
		{
			name:     "gap walls reachable by both robots",
			robots:   []int{1, 10},
			distance: []int{5, 5},
			walls:    []int{3, 5, 7},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxWallsDestroyed(tt.robots, tt.distance, tt.walls)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
