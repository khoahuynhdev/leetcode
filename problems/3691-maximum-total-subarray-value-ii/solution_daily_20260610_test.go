package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{"example 1: pick two overlapping subarrays", []int{1, 3, 2}, 2, 4},
		{"example 2: three subarrays each value 4", []int{4, 2, 5, 1}, 3, 12},
		{"edge case: single element forces value 0", []int{7}, 1, 0},
		{"edge case: all equal elements yield 0", []int{5, 5, 5, 5}, 4, 0},
		{"edge case: k equals total subarrays sums all ranges", []int{1, 3, 2}, 6, 5},
		{"edge case: two elements single pick", []int{0, 1000000000}, 1, 1000000000},
		{"edge case: pick only the single largest range", []int{4, 2, 5, 1}, 1, 4},
		{"edge case: strictly increasing array", []int{1, 2, 3, 4}, 3, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTotalValue(tt.nums, tt.k)
			if got != tt.expected {
				t.Errorf("maxTotalValue(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
