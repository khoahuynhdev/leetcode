package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{"example 1: single character is already minimal", "z", "z"},
		{"example 2: odd length, middle char is forced to 'b'", "babab", "abbba"},
		{"example 3: even length, all counts even", "daccad", "acddca"},
		{"edge case: two identical characters", "aa", "aa"},
		{"edge case: all characters identical, odd length", "aaaaa", "aaaaa"},
		{"edge case: all characters identical, even length", "aaaa", "aaaa"},
		{"edge case: odd-count letter is not the smallest", "cbabc", "bcacb"},
		{"edge case: input palindrome is not minimal", "bbaabb", "abbbba"},
		{"edge case: input already minimal is idempotent", "acddca", "acddca"},
		{"edge case: descending letters, even length", "zyxxyz", "xyzzyx"},
		{"edge case: three distinct letters, even length", "cbaabc", "abccba"},
		{"edge case: middle letter differs from the half's letters", "cabac", "acbca"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestPalindrome(tt.s)
			if got != tt.expected {
				t.Errorf("smallestPalindrome(%q) = %q, want %q", tt.s, got, tt.expected)
			}
		})
	}
}

// TestSolutionProperties checks structural invariants that must hold for any
// input: the result is a palindrome, and it is a permutation of the input.
func TestSolutionProperties(t *testing.T) {
	inputs := []string{
		"z",
		"babab",
		"daccad",
		"aa",
		"aaaaa",
		"cbabc",
		"bbaabb",
		"zyxxyz",
		"cabac",
		strings.Repeat("ab", 500) + "c" + strings.Repeat("ba", 500),
	}

	for _, in := range inputs {
		t.Run("input "+truncate(in), func(t *testing.T) {
			got := smallestPalindrome(in)

			if len(got) != len(in) {
				t.Fatalf("length = %d, want %d", len(got), len(in))
			}
			for i, j := 0, len(got)-1; i < j; i, j = i+1, j-1 {
				if got[i] != got[j] {
					t.Fatalf("result %q is not a palindrome (index %d vs %d)", truncate(got), i, j)
				}
			}

			var want, have [26]int
			for i := 0; i < len(in); i++ {
				want[in[i]-'a']++
				have[got[i]-'a']++
			}
			if want != have {
				t.Fatalf("result %q is not a permutation of the input", truncate(got))
			}
		})
	}
}

func truncate(s string) string {
	if len(s) <= 20 {
		return s
	}
	return s[:20] + "..."
}
