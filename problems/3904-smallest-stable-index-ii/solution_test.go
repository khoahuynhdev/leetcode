package main

import "testing"

func TestSmallestStableIndex(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"example 1: answer only at the last index", []int{5, 0, 1, 4}, 3, 3},
		{"example 2: strictly decreasing, no stable index", []int{3, 2, 1}, 1, -1},
		{"example 3: single element with k = 0", []int{0}, 0, 0},
		{"edge case: single large element", []int{1000000000}, 0, 0},
		{"edge case: all elements equal", []int{7, 7, 7}, 0, 0},
		{"edge case: non-decreasing array is stable everywhere", []int{1, 2, 3, 4}, 0, 0},
		{"edge case: k = 0 never satisfied", []int{2, 1}, 0, -1},
		{"edge case: k large enough for index 0", []int{10, 1, 5}, 100, 0},
		{"edge case: stable index followed by unstable ones", []int{1, 3, 2}, 0, 0},
		{"edge case: valley shape stable at index 0", []int{2, 1, 3}, 1, 0},
		{"edge case: max value range just below k", []int{1000000000, 0}, 999999999, -1},
		{"edge case: max value range exactly k", []int{1000000000, 0}, 1000000000, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestStableIndex(tt.nums, tt.k)
			if got != tt.expected {
				t.Errorf("smallestStableIndex(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}

// bruteForceSmallestStableIndex is the O(n^2) definition-following reference
// implementation, used to cross-check the optimized solution.
func bruteForceSmallestStableIndex(nums []int, k int) int {
	for i := range nums {
		maxLeft := nums[0]
		for j := 0; j <= i; j++ {
			maxLeft = max(maxLeft, nums[j])
		}
		minRight := nums[i]
		for j := i; j < len(nums); j++ {
			minRight = min(minRight, nums[j])
		}
		if maxLeft-minRight <= k {
			return i
		}
	}
	return -1
}

func TestSmallestStableIndexAgainstBruteForce(t *testing.T) {
	inputs := [][]int{
		{5, 0, 1, 4},
		{3, 2, 1},
		{0},
		{1, 3, 2},
		{2, 1, 3},
		{4, 4, 0, 9, 1, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{0, 1, 2, 3, 4, 5},
		{6, 2, 6, 2, 6, 2},
	}

	for _, nums := range inputs {
		for k := 0; k <= 10; k++ {
			want := bruteForceSmallestStableIndex(nums, k)
			got := smallestStableIndex(nums, k)
			if got != want {
				t.Errorf("smallestStableIndex(%v, %d) = %d, want %d", nums, k, got, want)
			}
		}
	}
}
