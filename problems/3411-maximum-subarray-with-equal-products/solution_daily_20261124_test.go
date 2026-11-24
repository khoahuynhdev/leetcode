package main

import "testing"

func TestMaxLengthDaily20261124(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: ones around a single 2", []int{1, 2, 1, 2, 1, 1, 1}, 5},
		{"example 2: pairwise coprime window", []int{2, 3, 4, 5, 6}, 3},
		{"example 3: mixed with ones", []int{1, 2, 3, 1, 4, 5, 1}, 5},
		{"edge case: two equal non-coprime values", []int{4, 4}, 2},
		{"edge case: all ones", []int{1, 1, 1, 1, 1}, 5},
		{"edge case: all same value greater than one", []int{2, 2, 2, 2}, 2},
		{"edge case: coprime run of primes", []int{2, 3, 5, 7}, 4},
		{"edge case: shared factors split the array", []int{1, 6, 1, 9, 1, 10, 1}, 5},
		{"edge case: pairwise coprime with ones", []int{1, 7, 1, 9, 1, 10, 1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.nums); got != tt.expected {
				t.Errorf("maxLength(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
