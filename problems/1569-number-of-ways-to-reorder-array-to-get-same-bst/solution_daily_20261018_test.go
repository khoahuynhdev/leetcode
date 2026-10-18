package main

import "testing"

func TestNumOfWaysDaily20261018(t *testing.T) {
	sorted := make([]int, 1000)
	for i := range sorted {
		sorted[i] = i + 1
	}

	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: root with one child on each side", []int{2, 1, 3}, 1},
		{"example 2: mixed left and right subtrees", []int{3, 4, 5, 1, 2}, 5},
		{"example 3: right-skewed chain has no alternatives", []int{1, 2, 3}, 0},
		{"edge case: single element", []int{1}, 0},
		{"edge case: two elements", []int{2, 1}, 0},
		{"edge case: balanced-ish tree of six nodes", []int{3, 1, 2, 5, 4, 6}, 19},
		{"edge case: perfect tree of seven nodes", []int{4, 2, 1, 3, 6, 5, 7}, 79},
		{"edge case: left-skewed chain", []int{5, 4, 3, 2, 1}, 0},
		{"edge case: max length sorted input", sorted, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numOfWays(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
