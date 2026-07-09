package main

import (
	"reflect"
	"testing"
)

func TestPathExistenceQueries(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		nums     []int
		maxDiff  int
		queries  [][]int
		expected []bool
	}{
		{
			name:     "example 1: self path and disconnected nodes",
			n:        2,
			nums:     []int{1, 3},
			maxDiff:  1,
			queries:  [][]int{{0, 0}, {0, 1}},
			expected: []bool{true, false},
		},
		{
			name:     "example 2: path through intermediate node",
			n:        4,
			nums:     []int{2, 5, 6, 8},
			maxDiff:  2,
			queries:  [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
			expected: []bool{false, false, true, true},
		},
		{
			name:     "edge case: single node self query",
			n:        1,
			nums:     []int{5},
			maxDiff:  0,
			queries:  [][]int{{0, 0}},
			expected: []bool{true},
		},
		{
			name:     "edge case: all nodes connected in one component",
			n:        5,
			nums:     []int{1, 2, 3, 4, 5},
			maxDiff:  1,
			queries:  [][]int{{0, 4}, {1, 3}, {2, 2}},
			expected: []bool{true, true, true},
		},
		{
			name:     "edge case: maxDiff zero connects only equal duplicates",
			n:        4,
			nums:     []int{4, 4, 7, 7},
			maxDiff:  0,
			queries:  [][]int{{0, 1}, {2, 3}, {1, 2}, {0, 3}},
			expected: []bool{true, true, false, false},
		},
		{
			name:     "edge case: no edges, every node isolated",
			n:        3,
			nums:     []int{0, 10, 20},
			maxDiff:  5,
			queries:  [][]int{{0, 1}, {1, 2}, {0, 2}, {2, 2}},
			expected: []bool{false, false, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pathExistenceQueries(tt.n, tt.nums, tt.maxDiff, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("pathExistenceQueries(%d, %v, %d, %v) = %v, want %v",
					tt.n, tt.nums, tt.maxDiff, tt.queries, result, tt.expected)
			}
		})
	}
}
