package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"example 1: both endpoints unique, larger one wins", []int{3, 9, 2, 1, 7}, 3, 7},
		{"example 2: last element repeats, only first qualifies", []int{3, 9, 7, 2, 1, 7}, 4, 3},
		{"example 3: k=1 with a duplicated value", []int{0, 0}, 1, -1},
		{"edge case: single element, k equals n equals 1", []int{5}, 1, 5},
		{"edge case: k equals n covers every value", []int{2, 2, 3}, 3, 3},
		{"edge case: k=1 picks largest value seen exactly once", []int{4, 1, 9, 1}, 1, 9},
		{"edge case: same value at both endpoints", []int{5, 3, 5}, 2, -1},
		{"edge case: no almost missing integer", []int{1, 2, 3, 1}, 2, -1},
		{"edge case: first endpoint is the larger candidate", []int{9, 2, 3, 1}, 2, 9},
		{"edge case: zeros are valid answers, not sentinels", []int{0, 1, 1}, 2, 0},
		{"edge case: strictly increasing, last endpoint wins", []int{1, 2, 3, 4, 5}, 3, 5},
		{"edge case: endpoint repeats in the interior", []int{0, 1, 0}, 2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestInteger(tt.nums, tt.k); got != tt.expected {
				t.Errorf("largestInteger(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}

// bruteForceLargestInteger tallies, for every value, how many size-k windows
// contain it, then returns the largest value with a tally of exactly one. It is
// the O(n*k) definition-driven reference used to validate the O(n) case
// analysis in solution.go.
func bruteForceLargestInteger(nums []int, k int) int {
	n := len(nums)
	windows := make(map[int]int, n)

	for s := 0; s+k <= n; s++ {
		seen := make(map[int]bool, k)
		for _, v := range nums[s : s+k] {
			seen[v] = true
		}
		for v := range seen {
			windows[v]++
		}
	}

	best := -1
	for v, c := range windows {
		if c == 1 && v > best {
			best = v
		}
	}
	return best
}

// TestSolutionAgainstBruteForce exhaustively compares the closed-form solution
// with the reference over every array of length 1..4 drawn from {0,1,2}, for
// every legal k.
func TestSolutionAgainstBruteForce(t *testing.T) {
	const alphabet = 3

	for n := 1; n <= 4; n++ {
		total := 1
		for i := 0; i < n; i++ {
			total *= alphabet
		}

		for code := 0; code < total; code++ {
			nums := make([]int, n)
			for i, rest := 0, code; i < n; i, rest = i+1, rest/alphabet {
				nums[i] = rest % alphabet
			}

			for k := 1; k <= n; k++ {
				name := fmt.Sprintf("nums=%v k=%d", nums, k)
				t.Run(name, func(t *testing.T) {
					want := bruteForceLargestInteger(nums, k)
					if got := largestInteger(nums, k); got != want {
						t.Errorf("largestInteger(%v, %d) = %v, want %v", nums, k, got, want)
					}
				})
			}
		}
	}
}
