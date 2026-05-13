package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		limit    int
		expected int
	}{
		{"example 1: change one element", []int{1, 2, 4, 3}, 4, 1},
		{"example 2: limit too small forces two moves", []int{1, 2, 2, 1}, 2, 2},
		{"example 3: already complementary", []int{1, 2, 1, 2}, 2, 0},
		{"edge case: n==2 single pair is always complementary", []int{1, 1}, 1, 0},
		{"edge case: n==2 with larger limit", []int{1, 5}, 10, 0},
		{"edge case: all equal, already complementary", []int{5, 5, 5, 5}, 5, 0},
		{"edge case: one move via a nested zero point", []int{1, 1, 2, 1}, 5, 1},
		{"edge case: extremes force two moves", []int{10, 1, 1, 10}, 10, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMoves(tt.nums, tt.limit)
			if got != tt.expected {
				t.Errorf("minMoves(%v, %d) = %d, want %d", tt.nums, tt.limit, got, tt.expected)
			}
		})
	}
}
