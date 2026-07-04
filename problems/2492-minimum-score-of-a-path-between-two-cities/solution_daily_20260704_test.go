package main

import "testing"

func TestMinScore(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		roads    [][]int
		expected int
	}{
		{
			name:     "example 1: single component, min edge is 5",
			n:        4,
			roads:    [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}},
			expected: 5,
		},
		{
			name:     "example 2: answer edge is a detour off the direct path",
			n:        4,
			roads:    [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}},
			expected: 2,
		},
		{
			name:     "edge case: single road directly connects 1 and n",
			n:        2,
			roads:    [][]int{{1, 2, 42}},
			expected: 42,
		},
		{
			name:     "edge case: cheap road in another component is ignored",
			n:        5,
			roads:    [][]int{{1, 2, 8}, {2, 5, 6}, {3, 4, 1}},
			expected: 6,
		},
		{
			name:     "edge case: minimum comes from the last relaxed edge",
			n:        5,
			roads:    [][]int{{1, 2, 10}, {2, 3, 9}, {3, 4, 8}, {4, 5, 3}},
			expected: 3,
		},
		{
			name:     "edge case: duplicate-weight edges, smallest wins",
			n:        3,
			roads:    [][]int{{1, 2, 5}, {2, 3, 5}, {1, 3, 4}},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minScore(tt.n, tt.roads)
			if result != tt.expected {
				t.Errorf("minScore(%d, %v) = %d, want %d", tt.n, tt.roads, result, tt.expected)
			}
		})
	}
}
