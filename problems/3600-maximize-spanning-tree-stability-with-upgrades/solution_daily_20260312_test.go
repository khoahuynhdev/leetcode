package main

import "testing"

func TestMaxStability(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		k        int
		expected int
	}{
		{
			name:     "example 1: must edge caps stability",
			n:        3,
			edges:    [][]int{{0, 1, 2, 1}, {1, 2, 3, 0}},
			k:        1,
			expected: 2,
		},
		{
			name:     "example 2: upgrade two edges",
			n:        3,
			edges:    [][]int{{0, 1, 4, 0}, {1, 2, 3, 0}, {0, 2, 1, 0}},
			k:        2,
			expected: 6,
		},
		{
			name:     "example 3: must edges form cycle",
			n:        3,
			edges:    [][]int{{0, 1, 1, 1}, {1, 2, 1, 1}, {2, 0, 1, 1}},
			k:        0,
			expected: -1,
		},
		{
			name:     "edge case: disconnected graph",
			n:        4,
			edges:    [][]int{{0, 1, 5, 0}, {2, 3, 5, 0}},
			k:        0,
			expected: -1,
		},
		{
			name:     "edge case: no upgrades needed, all must edges",
			n:        3,
			edges:    [][]int{{0, 1, 10, 1}, {1, 2, 7, 1}},
			k:        0,
			expected: 7,
		},
		{
			name:     "edge case: single edge graph",
			n:        2,
			edges:    [][]int{{0, 1, 5, 0}},
			k:        1,
			expected: 10,
		},
		{
			name:     "edge case: upgrade limited by k",
			n:        4,
			edges:    [][]int{{0, 1, 3, 0}, {1, 2, 3, 0}, {2, 3, 3, 0}},
			k:        1,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxStability(tt.n, tt.edges, tt.k)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
