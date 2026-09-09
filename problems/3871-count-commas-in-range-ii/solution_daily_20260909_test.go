package main

import "testing"

func TestCountCommas(t *testing.T) {
	tests := []struct {
		name     string
		n        int64
		expected int64
	}{
		{"example 1: n = 1002, three numbers with one comma each", 1002, 3},
		{"example 2: n = 998, every number has fewer than four digits", 998, 0},
		{"edge case: minimum n = 1", 1, 0},
		{"edge case: n = 999, just below the first comma", 999, 0},
		{"edge case: n = 1000, the very first comma", 1000, 1},
		{"edge case: n = 9999, all four-digit numbers counted", 9999, 9000},
		{"edge case: n = 999999, just below the second comma", 999999, 999000},
		{"edge case: n = 1000000, exact power of one thousand adds one", 1000000, 999002},
		{"edge case: n = 1000002, two thresholds contribute", 1000002, 999006},
		{"edge case: n = 10^12, exact threshold deep in the range", 1000000000000, 2998998999004},
		{"edge case: n = 10^15 - 1, just below the maximum threshold", 999999999999999, 3998998998999000},
		{"edge case: maximum n = 10^15", 1000000000000000, 3998998998999005},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countCommas(tt.n)
			if result != tt.expected {
				t.Errorf("countCommas(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}

// TestCountCommasAgainstBruteForce cross-checks the closed form against a direct
// per-number count over a small range, where enumeration is still cheap.
func TestCountCommasAgainstBruteForce(t *testing.T) {
	digits := func(x int64) int {
		d := 0
		for ; x > 0; x /= 10 {
			d++
		}
		return d
	}

	var want int64
	for n := int64(1); n <= 200000; n++ {
		want += int64((digits(n) - 1) / 3)

		if got := countCommas(n); got != want {
			t.Fatalf("countCommas(%d) = %v, want %v", n, got, want)
		}
	}
}
