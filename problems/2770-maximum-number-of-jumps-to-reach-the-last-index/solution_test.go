package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"example 1: target 2 allows 3 jumps", []int{1, 3, 6, 4, 1, 2}, 2, 3},
		{"example 2: target 3 allows 5 single-step jumps", []int{1, 3, 6, 4, 1, 2}, 3, 5},
		{"example 3: target 0 makes last index unreachable", []int{1, 3, 6, 4, 1, 2}, 0, -1},
		{"edge case: two equal elements, direct jump", []int{0, 0}, 0, 1},
		{"edge case: two elements out of target range, unreachable", []int{0, 10}, 5, -1},
		{"edge case: all elements equal, jump one step at a time", []int{1, 1, 1, 1}, 0, 3},
		{"edge case: negative values within target", []int{-3, 0, -2, 5, -1}, 5, 3},
		{"edge case: must skip an unreachable middle index", []int{0, 100, 1, 2}, 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumJumps(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("maximumJumps(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
