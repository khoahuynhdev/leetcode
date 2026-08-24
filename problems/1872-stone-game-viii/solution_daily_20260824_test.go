package main

import "testing"

func TestStoneGameVIII(t *testing.T) {
	tests := []struct {
		name     string
		stones   []int
		expected int
	}{
		{
			name:     "example 1: alice stops at index 3, forcing bob to sweep",
			stones:   []int{-1, 2, -3, 4, -5},
			expected: 5,
		},
		{
			name:     "example 2: alice takes everything in one move",
			stones:   []int{7, -6, 5, 10, 5, -2, -6},
			expected: 13,
		},
		{
			name:     "example 3: minimum length, alice has only one legal move",
			stones:   []int{-10, -12},
			expected: -22,
		},
		{
			name:     "edge case: n=2 all positive, alice must take both",
			stones:   []int{1, 2},
			expected: 3,
		},
		{
			name:     "edge case: all negative, alice defers the loss to bob",
			stones:   []int{-1, -2, -3},
			expected: 3,
		},
		{
			name:     "edge case: all positive, sweeping everything is optimal",
			stones:   []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "edge case: all zeros, difference is zero",
			stones:   []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "edge case: large negative in the middle limits the damage",
			stones:   []int{5, -100, 1},
			expected: -1,
		},
		{
			name:     "edge case: n=2 with mixed signs",
			stones:   []int{2, -1},
			expected: 1,
		},
		{
			name:     "edge case: boundary values at both extremes cancel out",
			stones:   []int{-10000, 10000, -10000, 10000},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.stones))
			copy(input, tt.stones)

			got := stoneGameVIII(input)
			if got != tt.expected {
				t.Errorf("stoneGameVIII(%v) = %d, want %d", tt.stones, got, tt.expected)
			}
		})
	}
}

// TestStoneGameVIIIMatchesBruteForce cross-checks the O(n) recurrence against a
// direct O(n^2) minimax over the same state space on small random-ish inputs.
func TestStoneGameVIIIMatchesBruteForce(t *testing.T) {
	inputs := [][]int{
		{-1, 2, -3, 4, -5},
		{7, -6, 5, 10, 5, -2, -6},
		{-10, -12},
		{3, -1, -1, -1, 8, -9, 2},
		{0, 0, 1, -1, 5},
		{-4, -4, -4, -4, -4, -4},
		{9, 1, 1, 1, 1, 1},
		{1, -2, 3, -4, 5, -6, 7, -8},
	}

	for _, stones := range inputs {
		got := stoneGameVIII(append([]int(nil), stones...))
		want := bruteForceStoneGameVIII(stones)
		if got != want {
			t.Errorf("stoneGameVIII(%v) = %d, brute force says %d", stones, got, want)
		}
	}
}

// bruteForceStoneGameVIII evaluates dp[i] = max over j >= i of (p[j]-dp[j+1])
// without the O(1) collapse, so it exercises the same model by a different path.
func bruteForceStoneGameVIII(stones []int) int {
	n := len(stones)
	prefix := make([]int, n)
	sum := 0
	for i, v := range stones {
		sum += v
		prefix[i] = sum
	}

	dp := make([]int, n)
	dp[n-1] = prefix[n-1]
	for i := n - 2; i >= 1; i-- {
		best := prefix[i] - dp[i+1]
		for j := i + 1; j < n; j++ {
			if cand := prefix[j] - dpAt(dp, j+1); cand > best {
				best = cand
			}
		}
		dp[i] = best
	}
	return dp[1]
}

func dpAt(dp []int, i int) int {
	if i >= len(dp) {
		return 0
	}
	return dp[i]
}
