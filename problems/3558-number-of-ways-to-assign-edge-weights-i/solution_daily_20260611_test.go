package main

import "testing"

func TestAssignEdgeWeights(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		expected int
	}{
		{
			name:     "example 1: single edge, depth 1",
			edges:    [][]int{{1, 2}},
			expected: 1,
		},
		{
			name:     "example 2: max depth 2 with branching",
			edges:    [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}},
			expected: 2,
		},
		{
			name:     "edge case: star graph, all leaves at depth 1",
			edges:    [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}},
			expected: 1,
		},
		{
			name:     "edge case: straight chain of depth 5",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
			expected: 16, // 2^(5-1)
		},
		{
			name:     "edge case: unbalanced tree, deepest branch wins",
			edges:    [][]int{{1, 2}, {2, 3}, {1, 4}},
			expected: 2, // max depth 2 -> 2^(2-1)
		},
		{
			name:     "edge case: edges given out of order",
			edges:    [][]int{{3, 4}, {1, 3}, {1, 2}},
			expected: 2, // path 1->3->4 has depth 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := assignEdgeWeights(tt.edges)
			if result != tt.expected {
				t.Errorf("assignEdgeWeights(%v) = %d, want %d", tt.edges, result, tt.expected)
			}
		})
	}
}
