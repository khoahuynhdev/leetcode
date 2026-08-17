package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestStoneGameVDaily20260817(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"example 1: mixed values, cut into [6,2,3] and [4,5,5]", []int{6, 2, 3, 4, 5, 5}, 18},
		{"example 2: all equal, odd length", []int{7, 7, 7, 7, 7, 7, 7}, 28},
		{"example 3: single stone, game already over", []int{4}, 0},
		{"edge case: empty input", []int{}, 0},
		{"edge case: two equal stones, Alice breaks the tie", []int{5, 5}, 5},
		{"edge case: two unequal stones, forced into the smaller", []int{10, 1}, 1},
		{"edge case: heavy head dominates every split", []int{100, 1, 1}, 3},
		{"edge case: heavy tail dominates every split", []int{1, 1, 100}, 3},
		{"edge case: strictly increasing", []int{1, 2, 3}, 4},
		{"edge case: all ones, even length", []int{1, 1, 1, 1}, 3},
		{"edge case: all ones, longer run", []int{1, 1, 1, 1, 1, 1}, 4},
		{"edge case: maximum stone values", []int{1000000, 1000000, 1000000}, 1000000},
		{"edge case: tie decides the outer split", []int{3, 1, 2, 2}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)
			got := stoneGameV(input)
			if got != tt.expected {
				t.Errorf("stoneGameV(%v) = %d, want %d", tt.input, got, tt.expected)
			}
			if !reflect.DeepEqual(input, tt.input) {
				t.Errorf("stoneGameV mutated its input: got %v, want %v", input, tt.input)
			}
		})
	}
}

// stoneGameVBrute is the direct O(n^3) recurrence, used only to cross-check the
// optimized O(n^2) solution on random inputs.
func stoneGameVBrute(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}
	prefix := make([]int, n+1)
	for i, v := range stoneValue {
		prefix[i+1] = prefix[i] + v
	}
	sum := func(i, j int) int { return prefix[j+1] - prefix[i] }

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			best := 0
			for k := i; k < j; k++ {
				left, right := sum(i, k), sum(k+1, j)
				var cur int
				switch {
				case left < right:
					cur = left + dp[i][k]
				case left > right:
					cur = right + dp[k+1][j]
				default:
					cur = left + max(dp[i][k], dp[k+1][j])
				}
				if cur > best {
					best = cur
				}
			}
			dp[i][j] = best
		}
	}
	return dp[0][n-1]
}

func TestStoneGameVDaily20260817AgainstBrute(t *testing.T) {
	rng := rand.New(rand.NewSource(1563))
	for trial := 0; trial < 400; trial++ {
		n := 1 + rng.Intn(9)
		// Small value ranges make exact ties frequent, which is the branch most
		// likely to break.
		maxVal := []int{1, 2, 3, 10}[rng.Intn(4)]
		stones := make([]int, n)
		for i := range stones {
			stones[i] = 1 + rng.Intn(maxVal)
		}
		want := stoneGameVBrute(stones)
		if got := stoneGameV(stones); got != want {
			t.Fatalf("stoneGameV(%v) = %d, brute force says %d", stones, got, want)
		}
	}
}
