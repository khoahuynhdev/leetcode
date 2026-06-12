package main

import (
	"reflect"
	"testing"
)

func TestAssignEdgeWeights(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		queries  [][]int
		expected []int
	}{
		{
			name:     "example 1: single edge, self and adjacent queries",
			edges:    [][]int{{1, 2}},
			queries:  [][]int{{1, 1}, {1, 2}},
			expected: []int{0, 1},
		},
		{
			name:     "example 2: branching tree with varied path lengths",
			edges:    [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}},
			queries:  [][]int{{1, 4}, {3, 4}, {2, 5}},
			expected: []int{2, 1, 4},
		},
		{
			name:     "edge case: every query is a self-pair (distance 0)",
			edges:    [][]int{{1, 2}, {2, 3}},
			queries:  [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected: []int{0, 0, 0},
		},
		{
			name:     "edge case: straight-line (path) tree, distances 1..4",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			queries:  [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}},
			expected: []int{1, 2, 4, 8}, // 2^0, 2^1, 2^2, 2^3
		},
		{
			name:     "edge case: LCA is an interior node, not the root",
			edges:    [][]int{{1, 2}, {2, 3}, {2, 4}},
			queries:  [][]int{{3, 4}, {1, 3}},
			expected: []int{2, 2}, // dist(3,4)=2 via node 2; dist(1,3)=2 via node 1
		},
		{
			name:     "edge case: longer path crossing the root",
			edges:    [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}},
			queries:  [][]int{{2, 4}},
			expected: []int{4}, // path 2-1-3-4 has 3 edges -> 2^2 = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := assignEdgeWeights(tt.edges, tt.queries)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("assignEdgeWeights(%v, %v) = %v, want %v",
					tt.edges, tt.queries, got, tt.expected)
			}
		})
	}
}
