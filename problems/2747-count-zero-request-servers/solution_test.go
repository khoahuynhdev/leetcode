package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		logs     [][]int
		x        int
		queries  []int
		expected []int
	}{
		{
			name:     "example 1: windows shift past older requests",
			n:        3,
			logs:     [][]int{{1, 3}, {2, 6}, {1, 5}},
			x:        5,
			queries:  []int{10, 11},
			expected: []int{1, 2},
		},
		{
			name:     "example 2: all servers busy then one idle",
			n:        3,
			logs:     [][]int{{2, 4}, {2, 1}, {1, 2}, {3, 1}},
			x:        2,
			queries:  []int{3, 4},
			expected: []int{0, 1},
		},
		{
			name:     "edge case: unsorted and duplicate queries",
			n:        3,
			logs:     [][]int{{1, 3}, {2, 6}, {1, 5}},
			x:        5,
			queries:  []int{11, 10, 11},
			expected: []int{2, 1, 2},
		},
		{
			name:     "edge case: inclusive window boundaries",
			n:        2,
			logs:     [][]int{{1, 5}, {2, 10}},
			x:        5,
			queries:  []int{10, 9},
			expected: []int{0, 1},
		},
		{
			name:     "edge case: all logs outside window",
			n:        4,
			logs:     [][]int{{1, 1}, {2, 2}},
			x:        1,
			queries:  []int{100},
			expected: []int{4},
		},
		{
			name:     "edge case: same server multiple logs stays active",
			n:        2,
			logs:     [][]int{{1, 1}, {1, 4}, {1, 6}},
			x:        3,
			queries:  []int{4, 6, 10},
			expected: []int{1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countServers(tt.n, tt.logs, tt.x, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
