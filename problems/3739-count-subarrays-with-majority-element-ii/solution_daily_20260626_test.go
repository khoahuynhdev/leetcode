package main

import "testing"

func TestCountMajoritySubarrays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int64
	}{
		{"example 1: [1,2,2,3] target 2", []int{1, 2, 2, 3}, 2, 5},
		{"example 2: all ones target 1", []int{1, 1, 1, 1}, 1, 10},
		{"example 3: target absent", []int{1, 2, 3}, 4, 0},
		{"edge case: single matching element", []int{5}, 5, 1},
		{"edge case: single non-matching element", []int{5}, 7, 0},
		{"edge case: target never the majority", []int{2, 1, 1, 1, 2}, 2, 2},
		{"edge case: alternating values", []int{1, 2, 1, 2, 1}, 1, 6},
		{"edge case: two matching at the ends", []int{3, 9, 9, 3}, 9, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countMajoritySubarrays(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("countMajoritySubarrays(%v, %d) = %d, want %d",
					tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
