package main

import (
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		k        int
		expected int64
	}{
		{"example 1: coins [3,6,9], k=3 (6 and 9 are redundant multiples of 3)", []int{3, 6, 9}, 3, 9},
		{"example 2: coins [5,2], k=7 (overlap at multiples of 10)", []int{5, 2}, 7, 12},
		{"edge case: single coin, k=1", []int{2}, 1, 2},
		{"edge case: coin of 1 makes every integer reachable", []int{1}, 1, 1},
		{"edge case: coin of 1 with maximum k", []int{1}, 2000000000, 2000000000},
		{"edge case: 1 absorbs every other coin", []int{1, 2, 3}, 1000000000, 1000000000},
		{"three coins with pairwise overlaps: [2,3,5], k=10", []int{2, 3, 5}, 10, 14},
		{"non-coprime pair: [4,6], k=5 (12 counted once)", []int{4, 6}, 5, 16},
		{"unsorted input is handled: [9,3,6], k=5", []int{9, 3, 6}, 5, 15},
		{"coprime pair, k below first overlap: [5,2], k=1", []int{5, 2}, 1, 2},
		{"large k on a single denomination", []int{25}, 2000000000, 50000000000},
		{"large k with many coins", []int{2, 3, 5, 7, 11, 13, 17, 19, 23}, 1000000000, 1195583314},
		{"maximum coin count, all values 11..25", []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}, 100, 171},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKthSmallest(tt.coins, tt.k)
			if got != tt.expected {
				t.Errorf("findKthSmallest(%v, %d) = %d, want %d", tt.coins, tt.k, got, tt.expected)
			}
		})
	}
}

// TestSolutionAgainstBruteForce cross-checks the inclusion-exclusion binary
// search against a straightforward sieve for small inputs, where enumerating
// every reachable amount is affordable.
func TestSolutionAgainstBruteForce(t *testing.T) {
	cases := [][]int{
		{2},
		{5, 2},
		{3, 6, 9},
		{4, 6},
		{2, 3, 5},
		{6, 10, 15},
		{7, 11, 13},
		{2, 3, 4, 5, 6, 7},
		{25, 24, 23, 22, 21},
	}

	const limit = 400
	for _, coins := range cases {
		reachable := bruteForceReachable(coins, limit)
		for k := 1; k <= len(reachable); k++ {
			got := findKthSmallest(coins, k)
			if got != reachable[k-1] {
				t.Fatalf("findKthSmallest(%v, %d) = %d, want %d", coins, k, got, reachable[k-1])
			}
		}
	}
}

// bruteForceReachable lists, in ascending order, every value in [1, limit] that
// is a multiple of at least one coin.
func bruteForceReachable(coins []int, limit int) []int64 {
	seen := make(map[int64]bool)
	for _, c := range coins {
		for v := int64(c); v <= int64(limit); v += int64(c) {
			seen[v] = true
		}
	}
	values := make([]int64, 0, len(seen))
	for v := range seen {
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return values
}
