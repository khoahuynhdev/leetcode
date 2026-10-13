package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{"example 1: three positives", []int{1, 2, 3}, 12},
		{"example 2: mixed signs pairs", []int{1, -1, 2, -2, 3, -3}, 16},
		{"edge case: single positive element", []int{5}, 25},
		{"edge case: single negative element", []int{-7}, 49},
		{"edge case: two elements with mixed signs", []int{3, -4}, 7},
		{"edge case: all zeros", []int{0, 0, 0}, 0},
		{"edge case: equal magnitudes even length cancel out", []int{2, -2, 2, -2}, 0},
		{"edge case: negative large magnitudes dominate", []int{-10, 1, 2, -9, 3}, 185},
		{"edge case: max values require int64", []int{40000, -40000, 40000}, 1600000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxAlternatingSum(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
