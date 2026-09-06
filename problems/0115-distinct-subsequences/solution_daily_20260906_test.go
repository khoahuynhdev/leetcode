package main

import "testing"

func TestNumDistinctDaily20260906(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{"example 1: rabbbit contains rabbit three ways", "rabbbit", "rabbit", 3},
		{"example 2: babgbag contains bag five ways", "babgbag", "bag", 5},
		{"edge case: target longer than source", "a", "aa", 0},
		{"edge case: empty target is matched once", "abc", "", 1},
		{"edge case: empty source and non-empty target", "", "a", 0},
		{"edge case: both strings empty", "", "", 1},
		{"edge case: source equals target", "abc", "abc", 1},
		{"edge case: no shared characters", "abc", "d", 0},
		{"edge case: repeated characters count as binomial", "aaaa", "aa", 6},
		{"edge case: single character repeated in source", "aaa", "a", 3},
		{"edge case: comparison is case sensitive", "aA", "A", 1},
		{"edge case: characters present but wrong order", "ba", "ab", 0},
		{"edge case: same characters in matching order", "ba", "ba", 1},
		{"edge case: interleaved matches", "aabb", "ab", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t2 *testing.T) {
			if got := numDistinct(tt.s, tt.t); got != tt.expected {
				t2.Errorf("numDistinct(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.expected)
			}
		})
	}
}

func TestNumDistinctLargeInputDaily20260906(t *testing.T) {
	// 1000 'a's choose 1 -- exercises the upper bound of the constraints and
	// confirms the rolling array is swept in the right direction.
	s := make([]byte, 1000)
	for i := range s {
		s[i] = 'a'
	}

	if got := numDistinct(string(s), "a"); got != 1000 {
		t.Errorf("numDistinct(1000 a's, %q) = %v, want %v", "a", got, 1000)
	}

	// C(1000, 2) = 499500.
	if got := numDistinct(string(s), "aa"); got != 499500 {
		t.Errorf("numDistinct(1000 a's, %q) = %v, want %v", "aa", got, 499500)
	}
}
