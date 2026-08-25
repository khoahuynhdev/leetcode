package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"example 1: multiples 2,4,6,8 present so 10 is the first gap", []int{8, 2, 3, 4, 6}, 2, 10},
		{"example 2: no multiple of 5 below 5 so k itself is missing", []int{1, 4, 7, 10, 15}, 5, 5},
		{"edge case: k = 1 degenerates to smallest missing positive", []int{1, 2, 3}, 1, 4},
		{"edge case: k larger than every element", []int{1, 2, 3}, 7, 7},
		{"edge case: duplicates collapse to a single blocker", []int{2, 2, 2}, 2, 4},
		{"edge case: contiguous run of multiples", []int{2, 4, 6, 8, 10}, 2, 12},
		{"edge case: answer exceeds the input value range", []int{100}, 100, 200},
		{"edge case: single element that is not a multiple", []int{7}, 3, 3},
		{"edge case: k = 1 with a gap in the middle", []int{1, 2, 4, 5}, 1, 3},
		{"edge case: non-multiples are ignored as blockers", []int{3, 9, 12, 6}, 6, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := smallestMissingMultiple(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("smallestMissingMultiple(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
