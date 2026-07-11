package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected int
	}{
		{
			name:     "example 1: all components complete",
			n:        6,
			edges:    [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}},
			expected: 3,
		},
		{
			name:     "example 2: one incomplete triangle path",
			n:        6,
			edges:    [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}},
			expected: 1,
		},
		{
			name:     "edge case: no edges, all isolated vertices",
			n:        4,
			edges:    [][]int{},
			expected: 4,
		},
		{
			name:     "edge case: single vertex",
			n:        1,
			edges:    [][]int{},
			expected: 1,
		},
		{
			name:     "edge case: fully connected graph is one complete component",
			n:        4,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			expected: 1,
		},
		{
			name:     "edge case: single connected but incomplete path",
			n:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			expected: 0,
		},
		{
			name:     "edge case: single edge pair is complete",
			n:        2,
			edges:    [][]int{{0, 1}},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countCompleteComponents(tt.n, tt.edges)
			if result != tt.expected {
				t.Errorf("countCompleteComponents(%d, %v) = %d, want %d",
					tt.n, tt.edges, result, tt.expected)
			}
		})
	}
}
