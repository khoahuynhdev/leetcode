package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		x        int
		k        int
		expected int64
	}{
		{"example 1: two windows of size 3", []int{5, -2, 1, 3, 7, 3, 6, 4, -1}, 3, 2, 8},
		{"example 2: adjacent windows of size 2", []int{9, -2, -2, -2, 1, 5}, 2, 2, 3},
		{"edge case: already equal elements", []int{4, 4, 4, 4}, 2, 2, 0},
		{"edge case: whole array is one window", []int{1, 2, 3, 10}, 4, 1, 10},
		{"edge case: k*x equals n forces partition", []int{1, 5, 2, 8}, 2, 2, 10},
		{"edge case: single window picks cheapest", []int{1, 100, 3, 4, 50}, 2, 1, 1},
		{"edge case: large values no overflow", []int{-1000000, 1000000, -1000000, 1000000}, 2, 2, 4000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.nums, tt.x, tt.k)
			if got != tt.expected {
				t.Errorf("minOperations(%v, %d, %d) = %d, want %d", tt.nums, tt.x, tt.k, got, tt.expected)
			}
		})
	}
}
