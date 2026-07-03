package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		online   []bool
		k        int64
		expected int64
	}{
		{
			name:     "example 1: two routes, cheaper one wins the bottleneck",
			edges:    [][]int{{0, 1, 5}, {1, 3, 10}, {0, 2, 3}, {2, 3, 4}},
			online:   []bool{true, true, true, true},
			k:        10,
			expected: 3,
		},
		{
			name:     "example 2: offline node forces a higher-bottleneck path",
			edges:    [][]int{{0, 1, 7}, {1, 4, 5}, {0, 2, 6}, {2, 3, 6}, {3, 4, 2}, {2, 4, 6}},
			online:   []bool{true, true, true, false, true},
			k:        12,
			expected: 6,
		},
		{
			name:     "edge case: no edges at all -> -1",
			edges:    [][]int{},
			online:   []bool{true, true},
			k:        100,
			expected: -1,
		},
		{
			name:     "edge case: budget too small for any path -> -1",
			edges:    [][]int{{0, 1, 5}, {1, 2, 5}},
			online:   []bool{true, true, true},
			k:        4,
			expected: -1,
		},
		{
			name:     "edge case: direct edge within budget",
			edges:    [][]int{{0, 1, 10}},
			online:   []bool{true, true},
			k:        10,
			expected: 10,
		},
		{
			name:     "edge case: direct edge just over budget -> -1",
			edges:    [][]int{{0, 1, 10}},
			online:   []bool{true, true},
			k:        9,
			expected: -1,
		},
		{
			name:     "edge case: zero-cost edge with k = 0",
			edges:    [][]int{{0, 1, 0}},
			online:   []bool{true, true},
			k:        0,
			expected: 0,
		},
		{
			name:     "edge case: only online detour is available",
			edges:    [][]int{{0, 1, 8}, {1, 3, 8}, {0, 2, 2}, {2, 3, 2}},
			online:   []bool{true, false, true, true},
			k:        20,
			expected: 2,
		},
		{
			name:     "edge case: source and sink disconnected -> -1",
			edges:    [][]int{{0, 1, 5}},
			online:   []bool{true, true, true},
			k:        1000,
			expected: -1,
		},
		{
			name:     "edge case: large costs need 64-bit arithmetic",
			edges:    [][]int{{0, 1, 1000000000}, {1, 2, 1000000000}},
			online:   []bool{true, true, true},
			k:        2000000000,
			expected: 1000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxPathScore(tt.edges, tt.online, tt.k)
			if result != tt.expected {
				t.Errorf("findMaxPathScore() = %v, want %v", result, tt.expected)
			}
		})
	}
}
