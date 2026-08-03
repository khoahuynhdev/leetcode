package main

import "testing"

func TestStoneGameIII(t *testing.T) {
	allOnes := make([]int, 50000)
	for i := range allOnes {
		allOnes[i] = 1
	}

	allZeros := make([]int, 1000)

	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"example 1: Alice's best line concedes the 7", []int{1, 2, 3, 7}, "Bob"},
		{"example 2: negative tail forces taking all three", []int{1, 2, 3, -9}, "Alice"},
		{"example 3: best Alice can manage is a draw", []int{1, 2, 3, 6}, "Tie"},

		{"edge case: single positive stone", []int{5}, "Alice"},
		{"edge case: single negative stone", []int{-1}, "Bob"},
		{"edge case: single zero stone", []int{0}, "Tie"},
		{"edge case: two stones, Alice takes only the first", []int{1, -1}, "Alice"},
		{"edge case: exactly three stones, all negative", []int{-1, -2, -3}, "Tie"},
		{"edge case: all values zero", allZeros, "Tie"},
		{"edge case: seven equal stones", []int{1, 1, 1, 1, 1, 1, 1}, "Alice"},
		{"edge case: extreme values at both bounds", []int{1000, -1000, 1000, -1000}, "Alice"},
		// For all-ones the margin cycles 1,2,3,2,1,0 with period 6 in the
		// length, so it is a Tie exactly when n % 6 == 0. 50000 % 6 == 2.
		{"edge case: max length, all ones", allOnes, "Alice"},
		{"edge case: all ones, length divisible by 6", []int{1, 1, 1, 1, 1, 1}, "Tie"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Guard against the solution mutating its input.
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			got := stoneGameIII(tt.input)
			if got != tt.expected {
				t.Errorf("stoneGameIII(%v) = %q, want %q", original, got, tt.expected)
			}
			for i := range original {
				if tt.input[i] != original[i] {
					t.Fatalf("input was mutated at index %d: got %d, want %d", i, tt.input[i], original[i])
				}
			}
		})
	}
}

// TestStoneGameIIIAgainstBrute cross-checks the O(n)/O(1) DP against a plain
// exponential minimax on small inputs, so a sign error in the recurrence
// cannot slip past the hand-picked cases above.
func TestStoneGameIIIAgainstBrute(t *testing.T) {
	// Every array of length 1..5 over the values {-2,-1,0,1,2}.
	values := []int{-2, -1, 0, 1, 2}

	var walk func(cur []int)
	walk = func(cur []int) {
		if len(cur) > 0 {
			want := bruteStoneGameIII(cur)
			if got := stoneGameIII(cur); got != want {
				t.Fatalf("stoneGameIII(%v) = %q, brute force says %q", cur, got, want)
			}
		}
		if len(cur) == 5 {
			return
		}
		for _, v := range values {
			walk(append(cur, v))
		}
	}
	walk(make([]int, 0, 5))
}

// bruteStoneGameIII explores the full game tree without memoization.
func bruteStoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	// margin returns the mover's score minus the opponent's over stoneValue[i:].
	var margin func(i int) int
	margin = func(i int) int {
		if i >= n {
			return 0
		}
		best := 0
		sum := 0
		for k := 1; k <= 3 && i+k <= n; k++ {
			sum += stoneValue[i+k-1]
			if cand := sum - margin(i+k); k == 1 || cand > best {
				best = cand
			}
		}
		return best
	}

	switch d := margin(0); {
	case d > 0:
		return "Alice"
	case d < 0:
		return "Bob"
	default:
		return "Tie"
	}
}
