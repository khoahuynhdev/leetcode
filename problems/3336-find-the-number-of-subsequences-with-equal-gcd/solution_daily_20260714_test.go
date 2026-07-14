package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: nums = [1,2,3,4]", []int{1, 2, 3, 4}, 10},
		{"example 2: nums = [10,20,30]", []int{10, 20, 30}, 2},
		{"example 3: nums = [1,1,1,1]", []int{1, 1, 1, 1}, 50},
		{"edge case: single element cannot form two subsequences", []int{5}, 0},
		{"edge case: two equal elements (ordered pairs)", []int{7, 7}, 2},
		{"edge case: two coprime elements have no equal-gcd pair", []int{2, 3}, 0},
		{"edge case: max value repeated", []int{200, 200, 200}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsequencePairCount(tt.nums)
			if result != tt.expected {
				t.Errorf("subsequencePairCount(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
