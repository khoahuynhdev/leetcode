package main

import (
	"math/rand"
	"testing"
)

func bruteSumCounts(nums []int) int {
	total := 0
	for i := range nums {
		seen := make(map[int]bool)
		for j := i; j < len(nums); j++ {
			seen[nums[j]] = true
			total = (total + len(seen)*len(seen)) % mod
		}
	}
	return total
}

func TestSumCounts(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: repeated value at both ends", []int{1, 2, 1}, 15},
		{"example 2: all identical pair", []int{2, 2}, 3},
		{"edge case: single element", []int{5}, 1},
		{"edge case: all distinct", []int{1, 2, 3}, 20},
		{"edge case: all identical", []int{7, 7, 7, 7}, 10},
		{"edge case: alternating values", []int{1, 2, 1, 2}, 1 + 1 + 1 + 1 + 4 + 4 + 4 + 4 + 4 + 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumCounts(tt.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSumCountsRandomAgainstBrute(t *testing.T) {
	rng := rand.New(rand.NewSource(2916))
	for iter := 0; iter < 200; iter++ {
		n := rng.Intn(30) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(6) + 1
		}
		if got, want := sumCounts(nums), bruteSumCounts(nums); got != want {
			t.Fatalf("nums=%v: got %v, want %v", nums, got, want)
		}
	}
}

func TestSumCountsLargeInput(t *testing.T) {
	nums := make([]int, 100000)
	for i := range nums {
		nums[i] = i + 1
	}
	// All distinct: sum over subarray lengths L of (n-L+1)*L^2, taken mod 1e9+7.
	n := int64(len(nums))
	var want int64
	for L := int64(1); L <= n; L++ {
		want = (want + (n-L+1)%mod*(L*L%mod)) % mod
	}
	if got := sumCounts(nums); int64(got) != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
