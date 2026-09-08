package main

import (
	"strconv"
	"testing"
)

func TestCountCommas(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"example 1: n = 1002, the three numbers 1,000 through 1,002", 1002, 3},
		{"example 2: n = 998, every number has fewer than four digits", 998, 0},
		{"edge case: n = 1, smallest allowed input", 1, 0},
		{"edge case: n = 999, just below the first comma", 999, 0},
		{"edge case: n = 1000, exactly at the first comma", 1000, 1},
		{"edge case: n = 1001, two commas total", 1001, 2},
		{"edge case: n = 9999, every four-digit number counted", 9999, 9000},
		{"edge case: n = 100000, upper bound of the constraints", 100000, 99001},
		{"edge case: n = 1000000, first number with two commas", 1000000, 999002},
		{"edge case: n = 1000001, past the second comma boundary", 1000001, 999004},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCommas(tt.n); got != tt.expected {
				t.Errorf("countCommas(%d) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}

// TestCountCommasMatchesBruteForce cross-checks the closed form against a
// reference that formats each number and counts separators directly. The
// running total is carried forward so every prefix [1, n] is verified in a
// single linear sweep.
func TestCountCommasMatchesBruteForce(t *testing.T) {
	want := 0
	for n := 1; n <= 200000; n++ {
		want += (len(strconv.Itoa(n)) - 1) / 3
		if got := countCommas(n); got != want {
			t.Fatalf("countCommas(%d) = %v, want %v", n, got, want)
		}
	}
}
