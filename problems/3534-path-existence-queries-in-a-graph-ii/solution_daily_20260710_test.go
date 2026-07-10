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
		expected []int
	}{
		{
			name:     "example 1: two direct-edge queries",
			n:        5,
			nums:     []int{1, 8, 3, 4, 2},
			maxDiff:  3,
			queries:  [][]int{{0, 3}, {2, 4}},
			expected: []int{1, 1},
		},
		{
			name:     "example 2: mixed distances and a disconnected pair",
			n:        5,
			nums:     []int{5, 3, 1, 9, 10},
			maxDiff:  2,
			queries:  [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}},
			expected: []int{1, 2, -1, 1},
		},
		{
			name:     "example 3: fully disconnected graph with a self query",
			n:        3,
			nums:     []int{3, 6, 1},
			maxDiff:  1,
			queries:  [][]int{{0, 0}, {0, 1}, {1, 2}},
			expected: []int{0, -1, -1},
		},
		{
			name:     "edge case: single node, only a self query",
			n:        1,
			nums:     []int{7},
			maxDiff:  5,
			queries:  [][]int{{0, 0}},
			expected: []int{0},
		},
		{
			name:     "edge case: maxDiff zero connects only equal values",
			n:        4,
			nums:     []int{2, 2, 3, 2},
			maxDiff:  0,
			queries:  [][]int{{0, 1}, {0, 3}, {0, 2}, {2, 2}},
			expected: []int{1, 1, -1, 0},
		},
		{
			name:     "edge case: long chain needs many hops",
			n:        5,
			nums:     []int{1, 2, 3, 4, 5},
			maxDiff:  1,
			queries:  [][]int{{0, 4}, {0, 3}, {1, 4}, {2, 2}},
			expected: []int{4, 3, 3, 0},
		},
		{
			name:     "edge case: everything within reach in one hop",
			n:        4,
			nums:     []int{10, 11, 12, 13},
			maxDiff:  100,
			queries:  [][]int{{0, 3}, {1, 2}, {3, 0}},
			expected: []int{1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathExistenceQueries(tt.n, tt.nums, tt.maxDiff, tt.queries)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("pathExistenceQueries(%d, %v, %d, %v) = %v, want %v",
					tt.n, tt.nums, tt.maxDiff, tt.queries, got, tt.expected)
			}
		})
	}
}
