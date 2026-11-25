package main

import "testing"

// bruteForce checks every length from 1 upward by comparing slices directly.
func bruteForce(nums []int) int {
	n := len(nums)
	for l := 1; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			unique := true
			for j := 0; j+l <= n && unique; j++ {
				if j == i {
					continue
				}
				same := true
				for k := 0; k < l; k++ {
					if nums[i+k] != nums[j+k] {
						same = false
						break
					}
				}
				if same {
					unique = false
				}
			}
			if unique {
				return l
			}
		}
	}
	return n
}

func TestSmallestUniqueSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: all equal elements", []int{3, 3, 3}, 3},
		{"example 2: a single element appears once", []int{2, 1, 2, 3, 3}, 1},
		{"example 3: length-2 window is unique", []int{1, 1, 2, 2, 1}, 2},
		{"edge case: single element", []int{7}, 1},
		{"edge case: two equal elements", []int{5, 5}, 2},
		{"edge case: all distinct", []int{1, 2, 3, 4}, 1},
		{"edge case: periodic pattern", []int{1, 2, 1, 2, 1, 2}, 4},
		{"edge case: max values", []int{100000, 100000, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := smallestUniqueSubarray(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSmallestUniqueSubarrayAgainstBruteForce(t *testing.T) {
	seed := uint32(12345)
	next := func() uint32 {
		seed = seed*1664525 + 1013904223
		return seed >> 16
	}
	for iter := 0; iter < 300; iter++ {
		n := int(next()%12) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = int(next()%3) + 1
		}
		got := smallestUniqueSubarray(nums)
		want := bruteForce(nums)
		if got != want {
			t.Fatalf("nums=%v: got %d, want %d", nums, got, want)
		}
	}
}
