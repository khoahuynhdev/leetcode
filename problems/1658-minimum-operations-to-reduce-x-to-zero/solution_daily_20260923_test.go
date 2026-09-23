package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		x        int
		expected int
	}{
		{"example 1: suffix only, remove last two", []int{1, 1, 4, 2, 3}, 5, 2},
		{"example 2: x exceeds total sum", []int{5, 6, 7, 8, 9}, 4, -1},
		{"example 3: two from the left and three from the right", []int{3, 2, 20, 1, 1, 3}, 10, 5},
		{"edge case: x equals total sum, remove everything", []int{1, 1, 4, 2, 3}, 11, 5},
		{"edge case: target reachable in value but not as a subarray sum", []int{2, 2, 2}, 3, -1},
		{"edge case: single element matches x exactly", []int{5}, 5, 1},
		{"edge case: single element smaller than x", []int{5}, 3, -1},
		{"edge case: one operation off the left end", []int{5, 2, 3, 1, 1}, 5, 1},
		{"edge case: prefer the longer of two matching windows", []int{1, 1, 1, 1, 1}, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			result := minOperations(nums, tt.x)
			if result != tt.expected {
				t.Errorf("minOperations(%v, %d) = %v, want %v", tt.nums, tt.x, result, tt.expected)
			}
			for i := range nums {
				if nums[i] != tt.nums[i] {
					t.Errorf("minOperations mutated input: got %v, want %v", nums, tt.nums)
					break
				}
			}
		})
	}
}
