package main

import "testing"

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		name     string
		side     int
		points   [][]int
		k        int
		expected int
	}{
		{
			name:     "example 1: four corners of a 2x2 square",
			side:     2,
			points:   [][]int{{0, 2}, {2, 0}, {2, 2}, {0, 0}},
			k:        4,
			expected: 2,
		},
		{
			name:     "example 2: cluster forces minimum gap of 1",
			side:     2,
			points:   [][]int{{0, 0}, {1, 2}, {2, 0}, {2, 2}, {2, 1}},
			k:        4,
			expected: 1,
		},
		{
			name:     "example 3: pick five points along the boundary",
			side:     2,
			points:   [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 0}, {2, 2}, {2, 1}},
			k:        5,
			expected: 1,
		},
		{
			name:     "edge case: tiny 1x1 square has only the four corners",
			side:     1,
			points:   [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			k:        4,
			expected: 1,
		},
		{
			name:     "edge case: only the four corners of a large square",
			side:     10,
			points:   [][]int{{0, 0}, {10, 0}, {10, 10}, {0, 10}},
			k:        4,
			expected: 10,
		},
		{
			name:     "edge case: dense bottom edge with three other corners",
			side:     5,
			points:   [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {5, 5}, {0, 5}},
			k:        4,
			expected: 5,
		},
		{
			name:     "edge case: all 16 boundary lattice points on a 4x4 square, k=5",
			side:     4,
			points:   [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 3}, {0, 2}, {0, 1}},
			k:        5,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistance(tt.side, tt.points, tt.k)
			if got != tt.expected {
				t.Errorf("maxDistance(%d, %v, %d) = %d, want %d",
					tt.side, tt.points, tt.k, got, tt.expected)
			}
		})
	}
}
