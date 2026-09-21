package main

import (
	"reflect"
	"testing"
)

func TestResultArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int64
	}{
		{
			name:     "example 1: nums=[1,2,3,4,5], k=3",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int64{9, 2, 4},
		},
		{
			name:     "example 2: powers of two with k=4, residue 3 unreachable",
			nums:     []int{1, 2, 4, 8, 16, 32},
			k:        4,
			expected: []int64{18, 1, 2, 0},
		},
		{
			name:     "example 3: nums=[1,1,2,1,1], k=2",
			nums:     []int{1, 1, 2, 1, 1},
			k:        2,
			expected: []int64{9, 6},
		},
		{
			name:     "edge case: single element, product below k",
			nums:     []int{7},
			k:        5,
			expected: []int64{0, 0, 1, 0, 0},
		},
		{
			name:     "edge case: k=1 collapses every product to residue 0",
			nums:     []int{9, 4, 7, 1000000000},
			k:        1,
			expected: []int64{10},
		},
		{
			name:     "edge case: every element divisible by k",
			nums:     []int{5, 10, 15},
			k:        5,
			expected: []int64{6, 0, 0, 0, 0},
		},
		{
			name:     "edge case: all ones never leave residue 1",
			nums:     []int{1, 1, 1, 1},
			k:        3,
			expected: []int64{0, 10, 0},
		},
		{
			name:     "edge case: large values reduced mod k",
			nums:     []int{1000000000, 999999999, 3},
			k:        4,
			expected: []int64{3, 1, 0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resultArray(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("resultArray(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}

// TestResultArrayLargeInputOverflow guards the 64-bit accumulation: with n = 1e5
// the total subarray count is 5,000,050,000, far beyond a 32-bit counter.
func TestResultArrayLargeInputOverflow(t *testing.T) {
	const n = 100000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	got := resultArray(nums, 1)
	want := []int64{int64(n) * int64(n+1) / 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("resultArray(1..%d, 1) = %v, want %v", n, got, want)
	}
}

// TestResultArrayMatchesBruteForce cross-checks the O(n*k) DP against a direct
// O(n^2) enumeration of every subarray on small inputs.
func TestResultArrayMatchesBruteForce(t *testing.T) {
	inputs := [][]int{
		{1},
		{2, 3},
		{4, 6, 9, 1},
		{5, 5, 5, 5, 5},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1000000000, 7, 13, 2, 1},
	}

	for _, nums := range inputs {
		for k := 1; k <= 5; k++ {
			got := resultArray(nums, k)
			want := bruteForceResultArray(nums, k)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("resultArray(%v, %d) = %v, want %v", nums, k, got, want)
			}
		}
	}
}

func bruteForceResultArray(nums []int, k int) []int64 {
	result := make([]int64, k)
	for i := range nums {
		product := 1 % k
		for j := i; j < len(nums); j++ {
			product = product * (nums[j] % k) % k
			result[product]++
		}
	}
	return result
}
