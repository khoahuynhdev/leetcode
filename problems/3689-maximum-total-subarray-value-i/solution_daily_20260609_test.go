package main

import "testing"

func TestMaxTotalValue(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{"example 1: [1,3,2] k=2", []int{1, 3, 2}, 2, 4},
		{"example 2: [4,2,5,1] k=3", []int{4, 2, 5, 1}, 3, 12},
		{"edge case: single element", []int{7}, 5, 0},
		{"edge case: all equal elements", []int{3, 3, 3, 3}, 4, 0},
		{"edge case: k=1 takes whole-array span", []int{10, 0, 5}, 1, 10},
		{"edge case: large product needs int64", []int{0, 1000000000}, 100000, 100000000000000},
		{"edge case: contains zero as min", []int{0, 2, 9, 4}, 3, 27},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxTotalValue(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("maxTotalValue(%v, %d) = %d, want %d", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
