package main

import (
	"reflect"
	"testing"
)

func TestGetResultsDaily20260530(t *testing.T) {
	tests := []struct {
		name     string
		queries  [][]int
		expected []bool
	}{
		{
			name:     "example 1: obstacle at 2 limits the largest fitting block",
			queries:  [][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}},
			expected: []bool{false, true, true},
		},
		{
			name:     "example 2: inserting a second obstacle splits the gap",
			queries:  [][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}},
			expected: []bool{true, true, false},
		},
		{
			name:     "edge case: no obstacles, block fits exactly in [0, x]",
			queries:  [][]int{{2, 5, 5}},
			expected: []bool{true},
		},
		{
			name:     "edge case: no obstacles, block strictly larger than x",
			queries:  [][]int{{2, 5, 6}},
			expected: []bool{false},
		},
		{
			name:     "edge case: query x equals an obstacle position (touching allowed)",
			queries:  [][]int{{1, 3}, {1, 5}, {2, 3, 3}},
			expected: []bool{true},
		},
		{
			name:     "edge case: obstacle sits beyond the query range",
			queries:  [][]int{{1, 10}, {2, 5, 5}},
			expected: []bool{true},
		},
		{
			name:     "edge case: two obstacles create small gaps, trailing decides feasibility",
			queries:  [][]int{{1, 1}, {1, 3}, {2, 5, 2}, {2, 5, 3}},
			expected: []bool{true, false},
		},
		{
			name:     "edge case: obstacle at 1, query for size that only fits trailing",
			queries:  [][]int{{1, 1}, {2, 5, 4}, {2, 5, 5}},
			expected: []bool{true, false},
		},
		{
			name:     "edge case: dense obstacles still leave a touch-fit block of size 1",
			queries:  [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 3, 1}, {2, 3, 2}},
			expected: []bool{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getResults(tt.queries)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("getResults(%v) = %v, want %v", tt.queries, got, tt.expected)
			}
		})
	}
}
