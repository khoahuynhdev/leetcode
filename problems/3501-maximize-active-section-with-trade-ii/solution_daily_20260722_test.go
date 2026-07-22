package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		queries  [][]int
		expected []int
	}{
		{
			name:     "example 1: no valid trade possible",
			s:        "01",
			queries:  [][]int{{0, 1}},
			expected: []int{1},
		},
		{
			name:     "example 2: single zero-run vs. two zero-runs",
			s:        "0100",
			queries:  [][]int{{0, 3}, {0, 2}, {1, 3}, {2, 3}},
			expected: []int{4, 3, 1, 1},
		},
		{
			name:     "example 3: clipped boundary zero-run",
			s:        "1000100",
			queries:  [][]int{{1, 5}, {0, 6}, {0, 4}},
			expected: []int{6, 7, 2},
		},
		{
			name:     "example 4: alternating pattern",
			s:        "01010",
			queries:  [][]int{{0, 3}, {1, 4}, {1, 3}},
			expected: []int{4, 4, 2},
		},
		{
			name:     "edge case: all ones (no zero-run, gain always 0)",
			s:        "1111",
			queries:  [][]int{{0, 3}, {1, 2}, {0, 0}},
			expected: []int{4, 4, 4},
		},
		{
			name:     "edge case: all zeros (single zero-run, no trade)",
			s:        "0000",
			queries:  [][]int{{0, 3}, {1, 2}},
			expected: []int{0, 0},
		},
		{
			name:     "edge case: single character",
			s:        "0",
			queries:  [][]int{{0, 0}},
			expected: []int{0},
		},
		{
			name:     "edge case: interior pair beats boundary pairs",
			s:        "010010010",
			queries:  [][]int{{0, 8}, {3, 6}},
			expected: []int{7, 6},
		},
		{
			name:     "edge case: three zero-runs, best is the wider adjacent pair",
			s:        "0100100",
			queries:  [][]int{{0, 6}},
			expected: []int{6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxActiveSectionsAfterTrade(tt.s, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maxActiveSectionsAfterTrade(%q, %v) = %v, want %v",
					tt.s, tt.queries, result, tt.expected)
			}
		})
	}
}
