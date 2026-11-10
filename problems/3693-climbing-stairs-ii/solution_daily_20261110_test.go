package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		costs    []int
		expected int
	}{
		{"example 1: path 0->1->2->4", 4, []int{1, 2, 3, 4}, 13},
		{"example 2: path 0->2->4", 4, []int{5, 1, 6, 2}, 11},
		{"example 3: single jump of 3", 3, []int{9, 8, 3}, 12},
		{"edge case: single step", 1, []int{5}, 6},
		{"edge case: two steps, jump directly", 2, []int{100, 1}, 5},
		{"edge case: all ones prefers unit steps", 6, []int{1, 1, 1, 1, 1, 1}, 12},
		{"edge case: expensive middle steps skipped", 6, []int{10000, 10000, 1, 10000, 10000, 1}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := climbStairs(tt.n, tt.costs)
			if result != tt.expected {
				t.Errorf("climbStairs(%d, %v) = %d, want %d", tt.n, tt.costs, result, tt.expected)
			}
		})
	}
}
