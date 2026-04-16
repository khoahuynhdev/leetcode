package main

import (
	"reflect"
	"testing"
)

func TestSolveQueries(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  []int
		expected []int
	}{
		{
			name:     "example 1: mixed values with wrap-around",
			nums:     []int{1, 3, 1, 4, 1, 3, 2},
			queries:  []int{0, 3, 5},
			expected: []int{2, -1, 3},
		},
		{
			name:     "example 2: all unique values return -1",
			nums:     []int{1, 2, 3, 4},
			queries:  []int{0, 1, 2, 3},
			expected: []int{-1, -1, -1, -1},
		},
		{
			name:     "edge case: single element has no twin",
			nums:     []int{7},
			queries:  []int{0},
			expected: []int{-1},
		},
		{
			name:     "edge case: two identical elements",
			nums:     []int{5, 5},
			queries:  []int{0, 1},
			expected: []int{1, 1},
		},
		{
			name:     "edge case: all identical elements",
			nums:     []int{9, 9, 9, 9},
			queries:  []int{0, 1, 2, 3},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "edge case: wrap-around is the shorter path",
			nums:     []int{2, 1, 1, 1, 1, 2},
			queries:  []int{0, 5},
			expected: []int{1, 1},
		},
		{
			name:     "edge case: repeated queries on same index",
			nums:     []int{4, 2, 4, 2, 4},
			queries:  []int{0, 0, 2, 4, 4},
			expected: []int{1, 1, 2, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveQueries(tt.nums, tt.queries)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("solveQueries(%v, %v) = %v, want %v", tt.nums, tt.queries, got, tt.expected)
			}
		})
	}
}
