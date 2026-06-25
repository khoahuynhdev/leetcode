package main

import "testing"

func TestCountSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{"example 1: [1,2,2,3] target 2", []int{1, 2, 2, 3}, 2, 5},
		{"example 2: all elements equal target", []int{1, 1, 1, 1}, 1, 10},
		{"example 3: target absent from array", []int{1, 2, 3}, 4, 0},
		{"edge case: single matching element", []int{5}, 5, 1},
		{"edge case: single non-matching element", []int{5}, 7, 0},
		{"edge case: target never the majority", []int{2, 1, 1, 1}, 1, 8},
		{"edge case: alternating values", []int{1, 2, 1, 2, 1}, 1, 6},
		{"edge case: large values within constraints", []int{1000000000, 1000000000, 1}, 1000000000, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countSubarrays(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("countSubarrays(%v, %d) = %d, want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
