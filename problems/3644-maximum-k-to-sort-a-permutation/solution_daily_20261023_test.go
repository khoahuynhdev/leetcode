package main

import "testing"

func TestSortPermutation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: swap 3 and 1 with k=1", []int{0, 3, 2, 1}, 1},
		{"example 2: swap 3 and 2 with k=2", []int{0, 1, 3, 2}, 2},
		{"example 3: fully reversed needs k=0", []int{3, 2, 1, 0}, 0},
		{"edge case: single element", []int{0}, 0},
		{"edge case: already sorted", []int{0, 1, 2, 3, 4}, 0},
		{"edge case: zero misplaced forces k=0", []int{1, 0, 2}, 0},
		{"edge case: swap 7 and 5 with k=5", []int{0, 1, 2, 3, 4, 7, 6, 5}, 5},
		{"edge case: cycle of 5, 7, 6 gives k=4", []int{0, 1, 2, 3, 4, 7, 5, 6}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortPermutation(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
