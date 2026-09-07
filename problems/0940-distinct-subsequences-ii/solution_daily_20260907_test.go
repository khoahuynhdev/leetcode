package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "example 1: all distinct characters, 2^3-1 subsequences",
			s:        "abc",
			expected: 7,
		},
		{
			name:     "example 2: repeated non-adjacent character forces dedup",
			s:        "aba",
			expected: 6,
		},
		{
			name:     "example 3: all identical characters collapse to length count",
			s:        "aaa",
			expected: 3,
		},
		{
			name:     "edge case: single character is the smallest legal input",
			s:        "a",
			expected: 1,
		},
		{
			name:     "edge case: empty string has no non-empty subsequences",
			s:        "",
			expected: 0,
		},
		{
			name:     "edge case: two identical characters, not 3",
			s:        "zz",
			expected: 2,
		},
		{
			name:     "edge case: interleaved repeats abab",
			s:        "abab",
			expected: 11,
		},
		{
			name:     "edge case: full alphabet gives exactly 2^26-1",
			s:        "abcdefghijklmnopqrstuvwxyz",
			expected: 67108863,
		},
		{
			name:     "edge case: alphabet plus a repeated 'a' gives 2^27-2",
			s:        "abcdefghijklmnopqrstuvwxyza",
			expected: 134217726,
		},
		{
			name:     "edge case: max length all identical characters",
			s:        strings.Repeat("a", 2000),
			expected: 2000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := distinctSubseqII(tt.s)
			if result != tt.expected {
				t.Errorf("distinctSubseqII(%q) = %v, want %v", truncate(tt.s), result, tt.expected)
			}
		})
	}
}

// TestSolutionMatchesBruteForce cross-checks the DP against an exhaustive
// set-based enumeration on every short string over a small alphabet.
func TestSolutionMatchesBruteForce(t *testing.T) {
	alphabet := []byte{'a', 'b', 'c'}

	var walk func(prefix []byte)
	walk = func(prefix []byte) {
		if len(prefix) > 0 {
			s := string(prefix)
			want := bruteForceDistinctSubseq(s)
			if got := distinctSubseqII(s); got != want {
				t.Errorf("distinctSubseqII(%q) = %v, want %v", s, got, want)
			}
		}
		if len(prefix) == 8 {
			return
		}
		for _, c := range alphabet {
			walk(append(prefix, c))
		}
	}
	walk(nil)
}

// TestSolutionLongInputStaysInModulus verifies that the count is reduced at
// every step: the true answer for a 2000-character string overflows any fixed
// width integer, so the result must remain a valid residue.
func TestSolutionLongInputStaysInModulus(t *testing.T) {
	s := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 77) // 2002 chars, trimmed below
	s = s[:2000]

	got := distinctSubseqII(s)
	if got < 0 || got >= 1_000_000_007 {
		t.Errorf("distinctSubseqII(<2000 chars>) = %v, want a value in [0, 1000000007)", got)
	}
	if got == 0 {
		t.Errorf("distinctSubseqII(<2000 chars>) = 0, which is almost certainly a bug rather than a genuine residue")
	}
}

// bruteForceDistinctSubseq enumerates every subsequence into a set. Only usable
// for very short strings (2^len subsets).
func bruteForceDistinctSubseq(s string) int {
	seen := make(map[string]struct{})
	for mask := 1; mask < 1<<len(s); mask++ {
		var sb strings.Builder
		for i := 0; i < len(s); i++ {
			if mask&(1<<i) != 0 {
				sb.WriteByte(s[i])
			}
		}
		seen[sb.String()] = struct{}{}
	}
	return len(seen)
}

// truncate keeps failure messages readable for very long inputs.
func truncate(s string) string {
	if len(s) <= 40 {
		return s
	}
	return s[:40] + "..."
}
