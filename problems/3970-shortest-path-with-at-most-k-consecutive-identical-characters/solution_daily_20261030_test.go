package main

import "testing"

func TestShortestPath(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		labels   string
		k        int
		expected int
	}{
		{"example 1: k=1 forces direct edge", 3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 3}}, "aab", 1, 3},
		{"example 2: k=2 allows cheaper path", 3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 3}}, "aab", 2, 2},
		{"example 3: only path has run of 3", 3, [][]int{{0, 1, 1}, {1, 2, 1}}, "aaa", 2, -1},
		{"edge case: single node", 1, [][]int{}, "a", 1, 0},
		{"edge case: no edges", 2, [][]int{}, "ab", 1, -1},
		{"edge case: direction matters", 2, [][]int{{1, 0, 5}}, "ab", 3, -1},
		{"edge case: detour through other label resets streak", 4,
			// 0->1 is "aa" (invalid for k=1); 0->2->1->3 is "abab"
			[][]int{{0, 1, 1}, {1, 3, 1}, {0, 2, 5}, {2, 1, 5}},
			"aabb", 1, 11},
		{"edge case: cheaper state with longer run must not block valid path", 5,
			// node 3 reached cheaply via 0->3 with run 2 ("aa"), but 3->4 then needs run 1;
			// the costlier 0->2->3 ("aba") reaches node 3 with run 1
			[][]int{{0, 3, 1}, {0, 2, 10}, {2, 3, 1}, {3, 4, 1}},
			"aabaa", 2, 12},
		{"edge case: all identical labels within k", 3, [][]int{{0, 1, 2}, {1, 2, 3}}, "zzz", 3, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestPath(tt.n, tt.edges, tt.labels, tt.k)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
