package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"example 1: merge bits into 15 and 6", []int{2, 6, 5, 8}, 2, 261},
		{"example 2: no operations needed", []int{4, 5, 4, 7}, 3, 90},
		{"edge case: single element", []int{1}, 1, 1},
		{"edge case: disjoint bits merge into one number", []int{1, 2, 4}, 1, 49},
		{"edge case: k equals n with disjoint bits leaves zeros", []int{1, 2, 4}, 3, 49},
		{"edge case: identical values", []int{1, 1, 1}, 2, 2},
		{"edge case: large values require modulo", []int{1_000_000_000, 1_000_000_000}, 2, 98},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSum(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
