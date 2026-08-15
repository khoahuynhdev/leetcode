package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: total XOR is zero, drop one element", []int{1, 2, 3}, 2},
		{"example 2: total XOR is already non-zero", []int{2, 3, 4}, 3},
		{"edge case: single zero has no valid subsequence", []int{0}, 0},
		{"edge case: all zeros", []int{0, 0, 0}, 0},
		{"edge case: single non-zero element", []int{5}, 1},
		{"edge case: identical pair cancels out", []int{7, 7}, 1},
		{"edge case: zeros padding a non-zero total", []int{0, 0, 5}, 3},
		{"edge case: zeros present but total XOR is zero", []int{1, 1, 0}, 2},
		{"edge case: max constraint values", []int{1000000000, 1000000000}, 1},
		{"edge case: four elements cancelling in pairs", []int{4, 4, 9, 9}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestSubsequence(tt.nums)
			if result != tt.expected {
				t.Errorf("longestSubsequence(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
