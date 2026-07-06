package main

import "testing"

func TestRemoveCoveredIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name:      "example 1: [3,6] covered by [2,8]",
			intervals: [][]int{{1, 4}, {3, 6}, {2, 8}},
			expected:  2,
		},
		{
			name:      "example 2: [2,3] covered by [1,4]",
			intervals: [][]int{{1, 4}, {2, 3}},
			expected:  1,
		},
		{
			name:      "edge case: single interval is never covered",
			intervals: [][]int{{5, 9}},
			expected:  1,
		},
		{
			name:      "edge case: equal starts, shorter is covered",
			intervals: [][]int{{1, 4}, {1, 10}},
			expected:  1,
		},
		{
			name:      "edge case: nested chain sharing a start",
			intervals: [][]int{{1, 2}, {1, 4}, {1, 6}, {1, 8}},
			expected:  1,
		},
		{
			name:      "edge case: disjoint intervals, none covered",
			intervals: [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected:  3,
		},
		{
			name:      "edge case: shared end with later start is covered",
			intervals: [][]int{{2, 8}, {3, 8}},
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeCoveredIntervals(tt.intervals)
			if result != tt.expected {
				t.Errorf("removeCoveredIntervals(%v) = %d, want %d", tt.intervals, result, tt.expected)
			}
		})
	}
}
