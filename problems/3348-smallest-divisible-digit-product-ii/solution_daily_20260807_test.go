package main

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		num      string
		divisor  int64
		expected string
	}{
		{"example 1: raise a digit and pad the tail", "1234", 256, "1488"},
		{"example 2: num already satisfies t", "12355", 50, "12355"},
		{"example 3: t has the prime factor 13", "11111", 26, "-1"},
		{"edge case: t has the prime factor 11", "11", 11, "-1"},
		{"edge case: t is 1 but num contains a zero", "10", 1, "11"},
		{"edge case: zero inside num with a factor still left", "1023", 2, "1112"},
		{"edge case: no same-length answer, must grow", "99", 8, "118"},
		{"edge case: large power of two needs four digits", "99", 1024, "2888"},
		{"edge case: a factor 5 can only come from the digit 5", "46", 5, "51"},
		{"edge case: longer than num and than the digit floor", "12", 100, "455"},
		{"edge case: keeps the longest matching prefix", "2222", 81, "2299"},
		{"edge case: factors of 7 force the digit 7", "11", 49, "77"},
		{"edge case: num already divisible, left untouched", "222222222", 2, "222222222"},
		{"edge case: t is 1 and num is already zero-free", "99", 1, "99"},
		{"edge case: trailing zero only needs the last digit bumped", "90", 1, "91"},
		{"edge case: all nines force a longer number", "99", 5, "115"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestNumber(tt.num, tt.divisor)
			if got != tt.expected {
				t.Errorf("smallestNumber(%q, %d) = %q, want %q", tt.num, tt.divisor, got, tt.expected)
			}
		})
	}
}

// TestSolutionAgainstBruteForce cross-checks the greedy construction against an
// exhaustive upward scan on small inputs.
func TestSolutionAgainstBruteForce(t *testing.T) {
	for start := 10; start <= 60; start++ {
		for divisor := int64(1); divisor <= 200; divisor++ {
			num := strconv.Itoa(start)
			want := bruteForce(start, divisor)
			if got := smallestNumber(num, divisor); got != want {
				t.Fatalf("smallestNumber(%q, %d) = %q, want %q", num, divisor, got, want)
			}
		}
	}
}

// bruteForce scans upward for the first zero-free number whose digit product is
// divisible by divisor. Divisors carrying a prime above 7 are impossible, so
// they short-circuit instead of running the (endless) scan.
func bruteForce(start int, divisor int64) string {
	rest := divisor
	for _, p := range []int64{2, 3, 5, 7} {
		for rest%p == 0 {
			rest /= p
		}
	}
	if rest != 1 {
		return "-1"
	}
	for v := start; ; v++ {
		product := int64(1)
		zeroFree := true
		for x := v; x > 0; x /= 10 {
			d := int64(x % 10)
			if d == 0 {
				zeroFree = false
				break
			}
			product *= d
		}
		if zeroFree && product%divisor == 0 {
			return strconv.Itoa(v)
		}
	}
}
