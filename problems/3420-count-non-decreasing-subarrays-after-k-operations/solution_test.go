package main

import (
	"math/rand"
	"testing"
)

// bruteForce computes the answer in O(n^2) for verification.
func bruteForce(nums []int, k int) int64 {
	var res int64
	for l := 0; l < len(nums); l++ {
		cost, mx := 0, 0
		for r := l; r < len(nums); r++ {
			if nums[r] > mx {
				mx = nums[r]
			}
			cost += mx - nums[r]
			if cost > k {
				break
			}
			res++
		}
	}
	return res
}

func TestCountNonDecreasingSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{"example 1: only subarrays starting with 6,3,1 fail", []int{6, 3, 1, 2, 4, 4}, 7, 17},
		{"example 2: one long subarray plus short ones", []int{6, 3, 1, 3, 6}, 4, 12},
		{"edge case: single element", []int{5}, 1, 1},
		{"edge case: already non-decreasing", []int{1, 2, 3, 4}, 1, 10},
		{"edge case: strictly decreasing with small k", []int{5, 4, 3, 2, 1}, 1, 9},
		{"edge case: large values, only the full array exceeds k", []int{1000000000, 1, 1000000000, 1}, 1000000000, 9},
		{"edge case: all equal elements", []int{7, 7, 7}, 1, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countNonDecreasingSubarrays(tt.nums, tt.k)
			if got != tt.expected {
				t.Errorf("countNonDecreasingSubarrays(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}

func TestCountNonDecreasingSubarraysRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(3420))
	for iter := 0; iter < 500; iter++ {
		n := rng.Intn(12) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(10) + 1
		}
		k := rng.Intn(20) + 1
		if got, want := countNonDecreasingSubarrays(nums, k), bruteForce(nums, k); got != want {
			t.Fatalf("nums=%v k=%d: got %d, want %d", nums, k, got, want)
		}
	}
}
