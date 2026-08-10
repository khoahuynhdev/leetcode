package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"example 1: n=1, Alice takes the only stone", 1, true},
		{"example 2: n=2, Alice must leave one stone for Bob", 2, false},
		{"example 3: n=4, perfect square cleared in one move", 4, true},
		{"edge case: n=3, take 1 leaving Bob the losing 2", 3, true},
		{"edge case: n=5, losing position", 5, false},
		{"edge case: n=7, losing position despite two options", 7, false},
		{"edge case: n=10, losing position", 10, false},
		{"edge case: n=17, losing position", 17, false},
		{"edge case: n=100, large perfect square", 100, true},
		{"edge case: n=99999, near upper constraint", 99999, true},
		{"edge case: n=100000, upper constraint", 100000, true},
		{"edge case: n=0, outside constraints, mover has no move", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := winnerSquareGame(tt.n); result != tt.expected {
				t.Errorf("winnerSquareGame(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}

// TestSolutionAgainstBruteForce cross-checks the DP against a plain,
// unmemoized minimax search over the small range where exponential search is
// still affordable. This guards against an inverted win/lose condition, which
// is the classic bug here and one the handful of examples would not catch.
func TestSolutionAgainstBruteForce(t *testing.T) {
	var bruteForce func(n int) bool
	bruteForce = func(n int) bool {
		for k := 1; k*k <= n; k++ {
			if !bruteForce(n - k*k) {
				return true
			}
		}
		return false
	}

	for n := 0; n <= 25; n++ {
		if got, want := winnerSquareGame(n), bruteForce(n); got != want {
			t.Errorf("winnerSquareGame(%d) = %v, brute force says %v", n, got, want)
		}
	}
}

func BenchmarkWinnerSquareGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		winnerSquareGame(100000)
	}
}
